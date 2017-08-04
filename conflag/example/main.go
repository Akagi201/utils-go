package main

import (
	"github.com/Akagi201/utilgo/conflag"
	flags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var opts struct {
	Conf     string `long:"conf" description:"config file"`
	LogLevel string `long:"log_level" default:"info" description:"log level"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default|flags.IgnoreUnknown)

	parser.Parse()

	if opts.Conf != "" {
		conflag.LongHyphen = true
		conflag.BoolValue = false
		args, err := conflag.ArgsFrom(opts.Conf)
		if err != nil {
			panic(err)
		}

		parser.ParseArgs(args)
	}

	log.Infof("opts: %+v", opts)
}
