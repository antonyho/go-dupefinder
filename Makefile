build-deps:
	dep init

update-deps:
	dep ensure -update

test:
	go test

