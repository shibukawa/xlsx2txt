build:
	GOPATH=`pwd` go build xlsx2txt

test:
	GOPATH=`pwd` go test xlsx2txt

vet:
	GOPATH=`pwd` go tool vet xlsx2txt

fmt:
	GOPATH=`pwd` go fmt xlsx2txt

get:
	GOPATH=`pwd` go get xlsx2txt
