package commands

import (
	logging "github.com/ungerik/go-logging"

	"github.com/domonda/GoMeetupExample/go/logger"
)

var (
	Debug = false
	Log   = logging.Vars{
		LoggerVar: &logger.Logger,
		DebugVar:  &Debug,
	}
)
