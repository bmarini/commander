package commander

import (
	"flag"
	"fmt"
	"os"
)

type Command interface {
	DefineFlags(*flag.FlagSet) interface{}
	Run(cmdConfig interface{})
}

type SubCommand interface {
	Name() string
	DefineFlags(*flag.FlagSet) interface{}
	Run(cmdConfig interface{}, subCmdConfig interface{})
}

type Config interface{}

type Parser struct {
	cmd Command
	cfg Config
	fs  *flag.FlagSet
}

type CLI struct {
	cmd     Command
	subCmds map[string]SubCommand
}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) addCommand(cmd Command) {
	c.cmd = cmd
}

func (c *CLI) addSubCommand(sub SubCommand) {
	c.subCmds[sub.Name()] = sub
}

func (c *CLI) hasCommand() bool {
	return c.cmd != nil
}

func (c *CLI) hasSubCommand() bool {
	return len(c.subCmds) > 0
}

func (c CLI) Run() {
	if c.hasCommand() {
		cp := &Parser{cmd: c.cmd, fs: flag.NewFlagSet(os.Args[0], flag.ExitOnError)}
		cp.cfg = c.cmd.DefineFlags(cp.fs)
		cp.fs.Parse(os.Args)
		cp.cmd.Run(cp.cfg)
	}

	// if len(os.Args) < 2 {
	// 	c.PrintUsage()
	// 	os.Exit(2)
	// }
}

func (c CLI) PrintUsage() {
	// c.FlagSet.Usage()
	fmt.Printf("Usage Instructions\n")
}
