# Readme


## Running the Cadence Service

1. `cd cadenceservice`
2. `docker-compose up`
3. Register the `simple-domain` with `docker run --network=host --rm ubercadence/cli:master --do simple-domain domain register --rd 10`

## Worker

Navigate back to the project root folder. Make sure go is installed in system.

* Worker
    1. `make worker`
    2. `./bins/worker`

## Endpoints

1. To start workflow
   * POST request to `http://localhost:3030/api/start-hello-world`
   * Note down the workflow id so you can use it to signal 

2. To signal workflow
    * Copy the workflow id from the previous response or check Cadence UI for the WorkflowID
    * POST Request to `http://localhost:3030/api/signal-hello-world?workflowId=<workflowId>&age=25`
    * Make sure to replace <workflowId> with the id retrieved from cadence web ui