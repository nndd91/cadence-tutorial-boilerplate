module github.com/nndd91/cadence-tutorial-boilerplate

go 1.13

require (
	github.com/spf13/viper v1.6.3
	github.com/uber-go/tally v3.3.15+incompatible
	go.uber.org/cadence v0.11.2
	go.uber.org/yarpc v1.45.0
	go.uber.org/zap v1.15.0
)

replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02
