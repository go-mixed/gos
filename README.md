# iGo+

Golang/[Go+](https://goplus.org/) interpreter. Base on [igop v0.9.6](https://github.com/goplus/igop)

## ✨ Feature highlights
- Run the Golang project **WITHOUT** [Golang compiler](https://go.dev/dl/)(142MB)
- Only 25MB after built and [UPX](https://github.com/upx/upx)
- Real time & Run：
  - Go+ script file, or a Golang file
  - Golang project (3<sup>rd</sup> party modules require a vendor directory)
  - a Golang project in the archive file of `*.tar.gz`, `*.tar.xz`, ...
- Support shebang line, like `#!/usr/bin/bash`
- Go+ Read-Eval-Print-Loop
- **Go1.18, Go1.19 generics**

## TOC

- igop run: run golang files
  - [Usage](#-run-golang-files)
  - [Single file mode](#single-file-mode)
  - [Project mode](#project-mode)
    - [Without submodules, 3<sup>rd</sup> party modules](#without-submodules-3rd-party-modules)
    - [With submodules, 3<sup>rd</sup> party modules](#with-submodules-or-3rd-party-modules)
  - [Archive mode](#archive-mode)
- igop exec: execute script code
  - [Usage](#-execute-script-code)
- igop repl: [REPL](#repl)
- bash ./script.sh: [run as shell file](#run-as-shell-file-like-sh)

## ⌛ Run Golang files 

```
igop run <PATH> 
  [-V | --debug] 
  [--vendor <path>] 
  [-I | --import NAME=PATH] 
  [-p | --plugin <path>] 
  -- <arguments>
```

Run a [Go+ script](https://goplus.org/), or a Golang project

|                         | Type      | Default       |                                                                                                  |
|-------------------------|-----------|---------------|--------------------------------------------------------------------------------------------------|
| <PATH>                  | String    |               | File of Golang+ script, "*.gop". <br/>Directory of Golang project.                               |
| -V<br/>--debug          | Boolean   | false         | Print the debug information.                                                                     |
| --vendor                | String    | <PATH>/vendor | The path of Golang dependency packages.<br/>Generate by `go mod vendor`.                         |
| -I<br/>--import         | NAME=PATH |               | The package to be imported. `-I PACKAGE_NAME=PATH -I PACKAGE_NAME2=PATH2`.                       |
| -p<br/>--plugin \<path> | Array     |               | (Only for linux)Load the "*.so" of golang plugin, See https://github.com/fly-studio/igop_plugins |
| -- \<arguments>         |           |               | arguments for script.<br/>Be read `os.Args` in the script.                                       |

### Single file mode

#### Run a file with `*.gop`

```
$ igop run /path/to/file.gop 

# or
$ cd /path/to
$ igop run file.gop
```
The difference is that they have different **Working directories**

Similarly hereinafter.

> See [examples/example2/1.go](examples/example2/1.gop)

#### Run a file with `*.go`

Golang file, Must include `package main` and `func main()`

```
package main

func main() {

}
```

> See [examples/example2/2.go](examples/example2/2.go)

```
$ igop run /path/to/file.go 
```

### Project mode

#### Without submodules, 3<sup>rd</sup> party modules

- The package MUST be `package main` in all `.go` files 
- Included `func main()` in anyone `.go` file, or ONLY one `.gop` file
- No sub-folder. even if the sub-folders are present, they will be ignored
- Support Go/Go+ hybrid programming

```
/path/to/
 - func.go
 - func.gop
 - main.go
```

> ONLY allowed one `func main(){}`, Whether in the *.gop or *.go, See [examples/example3](examples/example3)

Run

```
$ cd /path/to
$ igop run .

# or
$ igop run /path/to
```

#### With submodules, or 3<sup>rd</sup> party modules 

If the Golang project contains submodules, or 3<sup>rd</sup> party modules.

Not allowed *.gop(Temporary)

```
/path/to
  - main.go
  - func/
    - func.go
```

MUST include these files
```
/path/to/
  - go.mod
  - vendor/             <--- if you need 3rd-party modules
    - modules.txt
```

> See [examples/example1/](examples/example1)

#### - go.mod

Init the project to generate `go.mod` at first

```
$ go mod init your project-name
```

#### - `vendor` directory

Run this command to create the vendor directory, 
`igop` need to load the modules from `vendor/modules.txt`

```
$ go mod vendor
```

### Archive mode

This is a convenient way that a packaging of project mode

Support archive format. 

- tar.gz
- tar.bzip2
- tar.xz
- zip
- tar

> See [examples/example1.tar.gz](examples/example5.tar.gz)

When it's running, it's actually being extract to `example1/__example1.tar.gz__`

```
$ igop run examples1/example1.tar.gz --vendor vendor
```

#### `vendor`, `import` path

Unless you specify an absolute path that mean path on the OS, `--vendor` would be a relative path in archive.

Same with the PATH of `--import NAME=PATH`.

> the argument of `--vendor vendor` mean the vendor path is `examples1/__example1.tar.gz__/vendor`

Use the vendor with absolute path in OS

```
$ igop run examples1/example1.tar.gz --vendor /path/to/vendor
```

### Examples

> See  "example1/"、"example2/"、"example3/"

Project mode without modules
```
$ igop run example3/

$ cd example3/
$ igop run .
```


Project mode with modules
```
$ igop run example1/

$ igop run example1/ --vendor example1/vendor

$ cd example1/
$ igop run . --vendor vendor
```



Single mode
```
$ igop run examples2/1.gop
```

Script arguments
```
$ igop run examples2/2.go -- --abc 123 --def
```

## ⚡ Execute script code
```
igop exec 
  [-s | --script <code>] 
  [--debug] 
  -- <script arguments>
```
Execute script code from **StdIn** or the argument of "--script"

|                         | Type   | Default |                                                            |
|-------------------------|--------|---------|------------------------------------------------------------|
| -V<br/>--debug          |        | false   | Print the debug information.                               |
| -s<br/>--script \<code> | String |         | The Golang/Go+ script as string                            |
| -- \<arguments>         |        |         | arguments for script.<br/>Be read `os.Args` in the script. |

### Example

#### Execute script from StdIn

```
$ igop exec < example2/1.gop

$ cat example2/1.gop | igop exec

$ echo "i := 1+2; println(i)" | igop exec
$ printf "i := 1+2 \n println(i)" | igop exec
```

#### Execute script from argument

```
$ igop exec -s "i := 1+2; println(i)"
```

> Use `;`(semicolons) instead of `\n`(carriage returns)

## REPL
```
igop repl
```
A [Go+](https://goplus.org/) Read Eval Print Loop

Online: [https://repl.goplus.org/](https://repl.goplus.org/)

![](docs/repl_examples.png)

## Run as shell file, like "*.sh"

Shebang line:

script.sh
```
///usr/bin/true; exec /usr/bin/igop run -- "$0" "$@"

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
[./script.sh.gop --argument1 --argument2]
```
