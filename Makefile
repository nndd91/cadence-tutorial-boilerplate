.PHONY: worker

default: bins

worker:
	go build -i -o bins/worker app/worker/main.go

bins: worker

clean:
	rm -rf bins

