# conflag

conflag is a hubrid configuration library, which combines standard go flags/go-flags with JSON/TOML/YAML config files.

## Features

- [x] Combine command-line flag and configuration file.
- [x] Specify configration section.
- [x] Specify list parameters.
- [x] Support TOML configuration file.
- [x] Support JSON configuration file.
- [x] Support YAML configuration file.

## Install

```sh
go get github.com/Akagi201/utilgo/conflag
```

## Example

```go
// define your flags.
var procs int
flag.IntVar(&procs, "procs", runtime.NumCPU(), "GOMAXPROCS")

// set flags from configuration before parse command-line flags.
args, err := conflag.ArgsFrom("/path/to/config.toml");
if err != nil {
    panic(err)
}
flag.CommandLine.Parse(args)

// parse command-line flags.
flag.Parse()
```

and you create `/path/to/config.toml`

```toml
procs = 2
```

and run your app without option, `procs` flag will be set in `2` that is defined at configration file.

## Priority of flag

A priority of flag is

`command-line flag` > `configration file` > `flag default value`

In the above case,

| run                         | procs                              |
| --------------------------- | ---------------------------------- |
| myapp -procs 3              | 3                                  |
| myapp (with config-file)    | 2                                  |
| myapp (without config-file) | runtime.NumCPU() (default of flag) |

## Position

You can specify `positions` arguments to `ArgsFrom` function.

```toml
[options]
flag = "value"

[other settings]
hoge = "fuga"
```

```go
// parse configration only under the options section.
conflag.ArgsFrom("/path/to/config.toml", "options")
```

## List

You can use list for multiple parameters.
The following toml makes `-flag value1 -flag value2` arguments.

```toml
flag = [ "value1", "value2" ]
```

## `go-flags`

If you use [go-flags](https://github.com/jessevdk/go-flags) package, you can specify options like the following.

```go
import (
        flags "github.com/jessevdk/go-flags"
)

parser := flags.NewParser(&opts, flags.Default)

conflag.LongHyphen = true
conflag.BoolValue = false
args, err := conflag.ArgsFrom("/path/to/config.toml");
if err != nil {
        panic(err)
}
parser.ParseArgs(args)
```