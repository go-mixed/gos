# iGo+

Golang/[Go+](https://goplus.org/) interpreter. Base on [igop v0.8.6](https://github.com/goplus/igop)

## Supported

- Run a [Go+ script](https://goplus.org/)
- Run a Golang file
- Run a Golang project (3rd party modules require a vendor directory)
- Run a Golang project in the archive file of `*.tar.gz`, `*.tar.xz`, ...
- Go+ [Read-Eval-Print-Loop](https://repl.goplus.org/)

## Run Golang 

```
igop run [Path] [-V | --debug] [--vendor path] -- [arguments...]
```

run a [Go+ script](https://goplus.org/), or run a Golang project

|                   | Values | Default       |                                                                                                    |
|-------------------|--------|---------------|----------------------------------------------------------------------------------------------------|
| [Path]            | string |               | File of golang+ script, "*.gop" <br/>Directory of golang, `package main` `func main(){}` in a file |
| -V<br/>--debug    |        | false         | print the debug information                                                                        |
| --vendor          | string | [Path]/vendor | Path of golang dependency packages.<br/>Generate by `go mod vendor`                                |
| -- [arguments...] |        |               | Executing arguments of golang source/go+<br/>you can read the arguments in the source              |

### Archive mode

supported archive format

- tar.gz
- tar.bzip2
- tar.xz
- zip
- tar

> Unless you specify an absolute path that mean path on the OS, `--vendor` would be a relative path in archive


### Examples
see  "example1"、"example2"

Golang project
```
$ igop run example1/
# as same as
$ igop run example1/ --vendor example1/vendor
```

Golang project in the archive
```
$ igop run examples1/example1.tar.gz --vendor vendor
```

> `--vendor` mean the vendor path is `examples1/example1.tar.gz/vendor`

Go+
```
$ igop run examples2/1.gop
```

Executing arguments
```
$ igop run examples2/2.gop -- --abc 123 --def
```

### Complex project
If the Golang project contains submodules, or 3rd party modules

MUST include these files
```
- go.mod
- go.sum
- vendor             <--- if you need 3rd-party modules
  - modules.txt
```

#### - go.mod

Init the project to generate `go.mod` at first
```
$ go mod init your project-name
```

#### - go.sum

Run this to generate `go.sum` after `go get`
```
$ go mod tidy
```

#### - `vendor` directory

Run this to create the vendor directory, `igop` need pre-reading the `vendor/modules.txt` to load the modules
```
$ go mod vendor
```

## REPL
```
igop repl
```
A [go+](https://goplus.org/) Read Eval Print Loop

online: [https://repl.goplus.org/](https://repl.goplus.org/)

