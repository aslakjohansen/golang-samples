TARGETS = \
	test \


all: ${TARGETS}

clean:
	rm ${TARGETS}


lang/y.go: lang/lang.y
	cd lang ; goyacc lang.y

test: test.go lang/y.go
#	go generate -x -v
	go run test.go

