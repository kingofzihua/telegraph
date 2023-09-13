package app

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Command struct {
	usage    string
	desc     string
	options  CliOptions
	commands []*Command
	runFunc  RunCommandFunc
}

// CommandOption 定义用于初始化命令结构的可选参数。
type CommandOption func(*Command)

func WithCommandOptions(opt CliOptions) CommandOption {
	return func(c *Command) {
		c.options = opt
	}
}

// RunCommandFunc 定义应用程序的命令启动回调函数。
type RunCommandFunc func(args []string) error

// WithCommandRunFunc 用于设置应用程序的命令启动回调函数选项
func WithCommandRunFunc(run RunCommandFunc) CommandOption {
	return func(c *Command) {
		c.runFunc = run
	}
}

func (c *Command) cobraCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   c.usage,
		Short: c.desc,
	}
	cmd.SetOut(os.Stdout)
	cmd.Flags().SortFlags = false

	if len(c.commands) > 0 {
		for _, command := range c.commands {
			cmd.AddCommand(command.cobraCommand())
		}
	}

	if c.runFunc != nil {
		cmd.Run = c.runCommand
	}

	if c.options != nil {
		for _, f := range c.options.Flags().FlagSets {
			cmd.Flags().AddFlagSet(f)
		}
	}

	addHelpCommandFlag(c.usage, cmd.Flags())

	return cmd
}

func (c *Command) runCommand(cmd *cobra.Command, args []string) {
	if c.runFunc != nil {
		if err := c.runFunc(args); err != nil {
			fmt.Printf("%v %v\n", color.RedString("Error:"), err)
			os.Exit(1)
		}
	}
}

// FormatBaseName is formatted as an executable file name under different
// operating systems according to the given name.
func FormatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}

	return basename
}
