package main

import (
	"log"

	"github.com/Akagi201/utilgo/conflag"
	flags "github.com/jessevdk/go-flags"
)

var optCli struct {
	Conf string `long:"conf" default:"" json:"-" description:"config file"`
}

var optConf struct {
	LogLevel string `long:"log_level" default:"info" description:"log level"`
}

func main() {
	parserCli := flags.NewParser(&optCli, flags.Default|flags.IgnoreUnknown)
	parserConf := flags.NewParser(&optConf, flags.Default|flags.IgnoreUnknown)

	parserCli.Parse()

	if optCli.Conf != "" {
		conflag.LongHyphen = true
		conflag.BoolValue = false
		args, err := conflag.ArgsFrom(optCli.Conf)
		if err != nil {
			panic(err)
		}

		parserConf.ParseArgs(args)
	}

	parserConf.Parse()

	log.Printf("opts: %+v", optConf)
}
