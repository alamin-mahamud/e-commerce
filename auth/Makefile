BINARY=e-commerce
TESTS=go test $(go list ./... | grep -v /vendor/) -cover

build:
	${TESTS}
	go build -o ${BINARY} $(go list ./... | grep -v /vendor/) 

install:
	${TESTS}
	go build -o ${BINARY} $(go list ./... | grep -v /vendor/)

unittest:
	go test -short $(go list ./... | grep -v /vendor/)


clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install unittest

