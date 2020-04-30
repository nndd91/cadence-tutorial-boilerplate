# Readme


## Running the Cadence Service

1. `cd cadenceservice`
2. `docker-compose up`
3. Register the `simple-domain` with `docker run --network=host --rm ubercadence/cli:master --do simple-domain domain register --rd 10`

## Worker
Make sure go is installed in system.

Edit `worker/workflows/sample1.go` as you like then navigate back to the project root folder. 

* Worker
    1. `make worker`
    2. `./bins/worker`


## To start workflow
```
docker run --network=host --rm ubercadence/cli:master --domain simple-domain workflow start --tl "sampleGroup" --wt SampleWorkflow --et 600 --dt 600 -w SampleID
```
