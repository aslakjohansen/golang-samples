# Dependencies

go get -u golang.org/x/tools/cmd/goyacc
go install golang.org/x/tools/cmd/goyacc
goyacc lang/lang.y

