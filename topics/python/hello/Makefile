TARGETS = \
	test \


all: ${TARGETS}
	PYTHONPATH=. ./test

clean:
	rm ${TARGETS}

deps:
	go get github.com/sbinet/go-python


test: test.go
	go build -x test.go

