package logger

import (
	"context"
	"errors"
	"log"

	"github.com/tukaianirban/sdk.go.common/bruce"
)

// todo : move this package to sdk.go.common with support for logging locally or with a cloud provider

type Logger interface {
	Log(name string, sev Sevlevel)
}

type Sevlevel uint8

const (
	Default Sevlevel = iota
	Debug
	Info
	Notice
	Warning
	Error
	Critical
	Alert
	Emergency
)

const FORMAT_DATETIME_UNIVERSAL string = "02-01-2006 15:04:05"

type LogMessage struct {
	Name     string
	Severity Sevlevel
}

// var chIngressLogs = make(chan LogMessage, 50)

// func Log(name string, sev Sevlevel) {

// 	chIngressLogs <- LogMessage{
// 		Name:     name,
// 		Severity: sev,
// 	}
// }

var loggerObj Logger

func Init(ctx context.Context) error {

	storeConfig, err := bruce.GetString("logs_store")
	if err != nil {
		return err
	}

	switch storeConfig {
	case "local":
		loggerObj, err = getLocalLogger(ctx)
		if err != nil {
			return err
		} else {
			log.Printf("log provider type=%s init done", storeConfig)
			return nil
		}

	case "gcp":
		loggerObj, err = getGCPLogger(ctx)
		if err != nil {
			return err
		} else {
			log.Printf("log provider type=%s init done", storeConfig)
			return nil
		}

	default:
		return errors.New("unknown logs provider configured")
	}
}

func Log(name string, sev Sevlevel) error {

	if loggerObj == nil {
		return errors.New("logger not initialized")
	}

	loggerObj.Log(name, sev)
	return nil
}

// func Init(ctx context.Context) error {

// 	outputLogFile, err := bruce.GetString("output_logfile")
// 	if err != nil {
// 		log.Fatalf("failed to open log output file, reason=%s", err.Error())
// 		return err
// 	}

// 	// open the log output file
// 	f, err := os.OpenFile(outputLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
// 	if err != nil {
// 		log.Fatalf("error opening log output file, reason=%s", err.Error())
// 		return err
// 	}

// 	// start the worker that continually consumes ingress logs
// 	go workerIngestLogs(ctx, f)

// 	return nil
// }

// func workerIngestLogs(ctx context.Context, fileObj *os.File) {

// 	w := bufio.NewWriter(fileObj)

// 	// w.WriteString(fmt.Sprintf("Date: %s \n", time.Now().Format(FORMAT_DATETIME_UNIVERSAL)))

// 	defer func() {
// 		w.Flush()
// 		fileObj.Close()
// 	}()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Println("logger worker received signal to finish")
// 			return

// 		case msg := <-chIngressLogs:
// 			w.WriteString(getLogString(msg.Name, msg.Severity))
// 			w.Flush()
// 		}
// 	}
// }
