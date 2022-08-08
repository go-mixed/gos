# iGo+

## igop run [Path] [-V | --debug] [--vendor path] -- [arguments...]

- [Path]: a file of go+ script, or a path golang folder
- [-V | --debug]: print the debug information
- [--vendor]: the path of golang packages, You can generate via `go mod vendor`  
  - Optional
  - Default: [Path]/vendor

### examples
see folder "examples"

```
$ igop run examples/ --vendor examples/vendor
```

## igop repl
A go+ read-eval-print loop