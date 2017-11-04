default:
	go build -o finddupe

install:
	go build -o ${GOPATH}/bin/finddupe github.com/antonyho/go-dupefinder

build-deps:
	dep init

update-deps:
	dep ensure -update

test:
	go test

