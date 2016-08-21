# Sigterm

Short example of how to handle a <kbd>Ctrl+C</kbd> interrupt.

```bash
$ go build
$ ./sigterm
1 mississippi...
2 mississippi...
3 mississippi...
4 mississippi...
5 mississippi...
^CBye!
```

**Note:** When using `go run` instead of building the binary the
`os.Exit` call will not have any effect and the `go run` command will
always have the exit status 1.
