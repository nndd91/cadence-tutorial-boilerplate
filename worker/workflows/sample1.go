package workflows

import (
	"context"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"time"
)

// ApplicationName is the task list for this sample
const TaskListName = "sampleGroup"

// This is registration process where you register all your workflows
// and activity function handlers.
func init() {
	workflow.RegisterWithOptions(SampleWorkflow, workflow.RegisterOptions{
		Name: "SampleWorkflow",
	})
	activity.RegisterWithOptions(SampleActivity, activity.RegisterOptions{
		Name: "SampleActivity",
	})
}

var activityOptions = workflow.ActivityOptions{
	ScheduleToStartTimeout: time.Minute,
	StartToCloseTimeout:    time.Minute,
	HeartbeatTimeout:       time.Second * 20,
}

func SampleActivity(ctx context.Context) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("sample activity started")
	return "Great!", nil
}

func SampleWorkflow(ctx workflow.Context) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	logger := workflow.GetLogger(ctx)
	logger.Info("Sample workflow started")
	var activityResult string
	err := workflow.ExecuteActivity(ctx, SampleActivity).Get(ctx, &activityResult)
	if err != nil {
		return "", err
	}

	return activityResult, nil
}
