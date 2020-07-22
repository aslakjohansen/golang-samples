TARGETS = \
	test.wasm \
	wasm_exec.js \
	test \


all: ${TARGETS}

clean:
	rm ${TARGETS}

deps:
	go get github.com/shurcooL/goexec


wasm_exec.js:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

test.wasm: test.go
	GOOS=js GOARCH=wasm go build -o test.wasm

test: test.go
	go build -x test.go
	./test

serve:
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

