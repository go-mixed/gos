# iGo+



## igop run [Path] [-V | --debug] [--vendor path] -- [arguments...]

run a [golang+ script](https://goplus.org/), or run golang source

- [Path]: 
  - a file of golang+ script, "*.gop"
  - a path of golang directory, `package main` `func main(){}` in the directory
- [-V | --debug]: print the debug information
- [--vendor]: the path of golang packages, You can generate via `go mod vendor`  
  - Optional
  - Default: [Path]/vendor

### examples
see  "example1"、"example2"

```
$ igop run example1/
# as same as
$ igop run example1/ --vendor example1/vendor
```

```
$ igop run examples2/1.gop
```

## igop repl
A [go+](https://goplus.org/) Read Eval Print Loop