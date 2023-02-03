package logger

import (
	"context"
	"log"
	"os"
)

type defaultLogger struct {
	logger *log.Logger
}

func getDefaultLogger(ctx context.Context) (*defaultLogger, error) {

	return &defaultLogger{
		logger: log.New(os.Stdout, "", os.O_APPEND|os.O_CREATE|os.O_WRONLY),
	}, nil
}

func (self *defaultLogger) Log(logline string, sev Sevlevel) {

	self.logger.Printf("%s: %s", SevLevelStrings[sev], logline)
}
