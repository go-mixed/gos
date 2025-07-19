# Gos

Golang/[XGo](https://xgo.dev/) interpreter. Base on [ixgo v0.52.0](https://github.com/goplus/ixgo)

## ✨ Feature highlights
- Run the Golang **WITHOUT** [Golang compiler](https://go.dev/dl/)(150MB+)
- Only 18MB after built and [UPX -9](https://github.com/upx/upx)

- [X] Golang v1.24
- [x] XGo script file (*.xgo, *.gop)
- [x] Golang file
- [x] Golang project
- [X] Golang project in an archive file of `*.tar.gz`, `*.tar.xz`, ...
- [X] Support shebang line, like `#!/usr/bin/bash`
- [X] **Generic Function**
- [ ] `main.go` in the subdirectory
- [ ] `*.asm` file support

## TOC

- Run golang files
  - [Usage](#-run-golang-files)
  - [Single file mode](#single-file-mode)
  - [Project mode](#project-mode)
  - [Archive mode](#archive-mode)
- [Execute code](#execute-code)
- [REPL](#repl)
- [Run as shell file](#run-as-shell-file-like-sh)

## ⌛ Run Golang files 

```
gos <PATH>
  [-V | --debug]
  -- <arguments>
```

Run a [XGo script](https://xgo.dev/), or a Golang project

|                 | Type      | Default       |                                                                     |
|-----------------|-----------|---------------|---------------------------------------------------------------------|
| \<PATH>         | String    |               | File of Golang+ script, "*.xgo". <br/>Directory of Golang project. |
| -V<br/>--debug  | Boolean   | false         | Print the debug information.                                        |
| -- \<arguments> |           |               | arguments for script.<br/>Be read `os.Args` in the script.          |

<details>
  <summary>Advanced options</summary>

- `--vendor`: The path of Golang dependency packages.

  `gos run . --vendor=/path/to/vendor`

- `-I | --import <NAME=PATH>`: The package to be imported. 
  
  `gos run . -I mathex=/path/to/mathx -I json2=/path/to/json2`

- `-p | --plugin <path>`: (Only for linux)Load the "*.so" of golang plugin
  
  `gos run . --plugin /path/to/plugin1.so --plugin /path/to/plugin2.so`

  See https://github.com/go-mixed/gops_plugins

</details>

### Single file mode

Run a file with `*.xgo`、`*.go`

> `.go` must be `package main` and includes `func main()`
>
> See [examples/example2/1.go](examples/example2/1.xgo)、[examples/example2/2.go](examples/example2/2.go)

Run
```bash
gos /path/to/file.xgo
```

Run in the working directory
```bash
cd /path/to
gos file.xgo
```

With arguments
```bash
gos file.xgo -- --abc 123 --def
```



### Project mode

#### 1. Simple and flattened project 

- Must be `package main` and One `func main()` in the working directory
- The `.xgo` file implicitly contains a `func main()`, which is why only one `.xgo` file is allowed.


```
/path/to/
 - func.go
 - func.xgo
 - main.go
```

>  See [examples/example3](examples/example3)

Run

```bash
gos /path/to/examples/example3
```

Run in the working directory
```bash
cd /path/to/examples/example3
gos .
```

With arguments
```bash
gos . -- --abc 123 --def
```

#### 2. Project with submodules, or 3<sup>rd</sup> party modules.

- No allowed `*.xgo`
- `go.mod` MUST be in the working directory
- `vendor/modules.txt` MUST be in the working directory, if you need 3rd-party modules

```
/path/to
  - main.go
  - func/
    - func.go
  - go.mod
  - vensor/  <--- if you need 3rd-party modules
    - modules.txt
```

> See [examples/example1/](examples/example1)

### Archive mode

A packaging of project 

Supported archive format. When it runs, it'll actually be extract to `examples/__FILE_NAME__`

- tar.gz
- tar.bzip2
- tar.xz
- zip
- tar

> See [examples/example5.tar.gz](examples/example5.tar.gz)

Run an archive
```bash
gos examples/example5.tar.gz
```

With arguments
```bash
gos examples2/2.go -- --abc 123 --def
```

## ⚡ Execute code
```
gos
  [-s | --script <code>] 
  [--debug] 
  -- <script arguments>
```
Execute script code from **StdIn** or the argument of "--script"

|                         | Type   | Default |                                                            |
|-------------------------|--------|---------|------------------------------------------------------------|
| -V<br/>--debug          |        | false   | Print the debug information.                               |
| -s<br/>--script \<code> | String |         | The Golang/XGo script as string                           |
| -- \<arguments>         |        |         | arguments for script.<br/>Be read `os.Args` in the script. |

### Example

#### Code from StdIn

```bash
gos < example2/1.xgo
```

```bash
cat example2/1.xgo | gos
```

```bash
echo "i := 1+2; println(i)" | gos
printf "i := 1+2 \n println(i)" | gos
```

#### Code in argument "-s"

```
$ gos -s "i := 1+2; println(i)"
```

> Use `;`(semicolons) instead of carriage returns

## REPL
```
gos repl
```
A [XGo](https://xgo.dev/) Read Eval Print Loop

Online: [https://repl.xgo.dev/](https://repl.xgo.dev/)

![](docs/repl_examples.png)

## Run as shell file, like "*.sh"

Shebang line:

script.sh
```
///usr/bin/true; exec /usr/bin/gos -- "$0" "$@"

import "os"
import "fmt"

fmt.Printf("%v", os.Args)
```

> The first two lines are important

Run
```
$ chmod +x ./script.sh

$ ./scrpit.sh --argument1 --argument2
$ sh ./script.sh --argument1 --argument2
```

Print
```
[./script.sh.xgo --argument1 --argument2]
```

# Development

## Install dependencies

```shell
go install github.com/goplus/ixgo/cmd/qexp@latest
```

## Build build-in scripts

```shell
qexp -outdir pkgs -filename go_export github.com/inconshreveable/mousetrap github.com/spf13/pflag github.com/spf13/cobra go.uber.org/multierr gopkg.in/yaml.v3 github.com/pkg/errors

go build
```
