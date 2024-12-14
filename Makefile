lib/worker.so:
	go build -o lib/worker.so -buildmode=c-shared lib/worker.go

clean:
	rm lib/worker.h lib/worker.so

