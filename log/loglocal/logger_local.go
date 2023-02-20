package logger

// import (
// 	"bufio"
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/tukaianirban/sdk.go.common/bruce"
// )

// type localLogger struct {
// 	FileHandler   *os.File
// 	chIngressLogs chan LogMessage
// }

// var localLoggerObj *localLogger

// func getLocalLogger(ctx context.Context) (*localLogger, error) {

// 	outputLogFile, err := bruce.GetString("logs_local_logfile")
// 	if err != nil {
// 		log.Fatalf("failed to open log output file, reason=%s", err.Error())
// 		return nil, err
// 	}

// 	// open the log output file
// 	f, err := os.OpenFile(outputLogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
// 	if err != nil {
// 		log.Fatalf("error opening log output file, reason=%s", err.Error())
// 		return nil, err
// 	}

// 	localLoggerObj = &localLogger{
// 		FileHandler:   f,
// 		chIngressLogs: make(chan LogMessage, 50),
// 	}

// 	go localLoggerObj.workerIngestLogs(ctx)

// 	return localLoggerObj, nil
// }

// func (self *localLogger) Log(name string, sev Sevlevel) {

// 	self.chIngressLogs <- LogMessage{
// 		Name:     name,
// 		Severity: sev,
// 	}
// }

// func (self *localLogger) workerIngestLogs(ctx context.Context) {

// 	w := bufio.NewWriter(self.FileHandler)

// 	// w.WriteString(fmt.Sprintf("Date: %s \n", time.Now().Format(FORMAT_DATETIME_UNIVERSAL)))

// 	defer func() {
// 		w.Flush()
// 		self.FileHandler.Close()
// 	}()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Println("local logger worker received signal to finish")
// 			return

// 		case msg := <-self.chIngressLogs:
// 			w.WriteString(getLocalLogString(msg.Name, msg.Severity))
// 			w.Flush()
// 		}
// 	}
// }

// //
// // function that converts the raw log string to a formatted one
// //
// func getLocalLogString(content string, sev Sevlevel) string {

// 	return fmt.Sprintf("%s : %s : %s", time.Now().Format(FORMAT_DATETIME_UNIVERSAL), getSeverityMapForLocal(sev), content)
// }

// func getSeverityMapForLocal(sev Sevlevel) string {
// 	switch sev {
// 	case Debug:
// 		return "DEBUG"
// 	case Info:
// 		return "INFO"
// 	case Notice:
// 		return "NOTICE"
// 	case Warning:
// 		return "WARNING"
// 	case Error:
// 		return "ERROR"
// 	case Critical:
// 		return "CRITICAL"
// 	case Alert:
// 		return "ALERT"
// 	case Emergency:
// 		return "EMERGENCY"
// 	case Default:
// 		fallthrough
// 	default:
// 		return "DEFAULT"
// 	}
// }
