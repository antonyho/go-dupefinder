update-deps:
	dep init
	dep ensure -update

test:
	go test

