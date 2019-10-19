# Intro

Based on [this](https://github.com/timtadh/lexmachine).

# Dependencies

```shell
go get -u golang.org/x/tools/cmd/goyacc
go install golang.org/x/tools/cmd/goyacc
goyacc lang/lang.y
```

# Running

Make sure that `$GOPATH` is set, that `$GOPATH/bin` is in `$PATH` and that this repository is reacable directly under `$GOPATH/src`. Then run:

```shell
make
```

# References

- [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc)
- [lexmachine](https://godoc.org/github.com/timtadh/lexmachine)
- [How to write a compiler in Go: a quick guide](https://www.freecodecamp.org/news/write-a-compiler-in-go-quick-guide-30d2f33ac6e0/)
- [Make a Lexer with Go](https://tylersommer.com/make-a-lexer-with-go)
- [gocc](https://github.com/goccmack/gocc)
- [GopherCon 2018 - How to Write a Parser in Go](https://about.sourcegraph.com/go/gophercon-2018-how-to-write-a-parser-in-go)
- [Lex YACC: Debugging](https://www.tldp.org/HOWTO/Lex-YACC-HOWTO-7.html)

