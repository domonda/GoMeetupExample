package logger

import (
	"os"

	logging "github.com/ungerik/go-logging"
)

const (
	MaxLogFileSize = 1024 * 1024 * 10
)

var (
	formatter = logging.NewTimeFormatter("2006-01-02 15:04:05.000", true)
	term      = logging.NewColorTerm(os.Stdout, os.Stdout, os.Stderr, formatter)

	Logger logging.Logger = logging.Tee{term}
)

func SetFile(filename string) error {
	file, err := logging.NewFile(filename, 0640, MaxLogFileSize, true, formatter)
	if err != nil {
		return err
	}
	Logger = logging.Tee{term, file}
	return nil
}
