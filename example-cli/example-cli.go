package main

import (
	"io"

	"github.com/pkg/errors"
	command "github.com/ungerik/go-command"
	logging "github.com/ungerik/go-logging"
	structflag "github.com/ungerik/go-structflag"

	"github.com/domonda/GoMeetupExample/go/logger"
)

var config struct {
	Production bool `usage:"Production disables test and debug functionalities"`
}

var (
	Debug = false
	Log   = logging.Vars{
		LoggerVar: &logger.Logger,
		DebugVar:  &Debug,
	}
)

func defaultCommand() {
	structflag.PrintUsage()
}

func main() {
	err := logger.SetFile("domonda.log")
	if err != nil {
		// Log is still working, just not writing to a file
		Log.UnresolvedErrorf(err, "Could not create log file")
		return
	}

	dispatcher := command.NewStringArgsDispatcher()
	dispatcher.MustAddDefaultCommand("Prints the command line options", defaultCommand, &command.WithoutArgs, command.Println)

	structflag.AppName = "domonda"
	structflag.PrintUsageIntro = func(output io.Writer) {
		dispatcher.PrintCommandsUsageIntro("domonda", output)
	}
	args := structflag.LoadFileIfExistsAndMustParseCommandLine("config.json", &config)

	cmd, err := dispatcher.DispatchCombined(args)
	switch {
	case errors.Cause(err) == command.ErrNotFound:
		Log.Printf("Command not found: '%s'\n\nAvailable commands:\n", cmd)
		dispatcher.PrintCommands("domonda")

	case err != nil:
		Log.Printf("Error while executing command %s: %+v\n", cmd, err)

	case err == nil:
		Log.Printf("Finished command %s", cmd)
	}
}
