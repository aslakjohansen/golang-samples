# The Trick

The trick is to make sure that golang thinks that you are adhering to their decision that every codebase **MUST** be located under `$GOPATH`.

By doing something like the following I am no longer restricted by this annoying rule:

```shell
$ ln -s ~/vcs/git/golang-samples $GOPATH/src/golang-samples
```

