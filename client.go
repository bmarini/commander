package main

import (
	"flag"
	"fmt"

	"github.com/bmarini/cli/commander"
)

func main() {
	cli := commander.NewCLI()
	cmd := Command{}
	cli.AddCommand(cmd)
	cli.Run()
}

type ClientConfig struct {
	verbose bool
}

type Command struct{}

func (c Command) DefineFlags(f *flag.FlagSet) interface{} {
	var cfg ClientConfig
	f.BoolVar(&cfg.verbose, "verbose", false, "set to true for more verbose output")
	return cfg
}

func (c Command) Run(cfg interface{}) {
	fmt.Printf("%v", cfg)
}
