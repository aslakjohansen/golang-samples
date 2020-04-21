TARGETS = \
	test \


all: ${TARGETS}

clean:
	rm ${TARGETS}


update:
	go get github.com/eclipse/paho.mqtt.golang

test: test.go
	go run test.go

