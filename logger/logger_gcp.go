package logger

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/tukaianirban/sdk.go.common/bruce"

// 	"cloud.google.com/go/logging"
// 	"google.golang.org/api/option"
// )

// type gcpLogger struct {
// 	Client        *logging.Client
// 	LoggerClient  *logging.Logger
// 	LogName       string
// 	chIngressLogs chan LogMessage
// }

// var gcpLoggerObj *gcpLogger

// func getGCPLogger(ctx context.Context) (*gcpLogger, error) {

// 	gcpCredsFilename, err := bruce.GetString("provider_gcp_credentialsFile")
// 	if err != nil {
// 		return nil, err
// 	}

// 	// _, filename, _, _ := runtime.Caller(0)
// 	// dir := path.Join(path.Dir(filename), "../..")
// 	// os.Chdir(dir)

// 	// log.Printf("dir=%s", dir)
// 	opt := option.WithCredentialsFile(gcpCredsFilename)

// 	projectID, err := bruce.GetString("provider_gcp_projectID")
// 	if err != nil {
// 		return nil, err
// 	}

// 	client, err := logging.NewClient(ctx, projectID, opt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	logName, err := bruce.GetString("logs_gcp_logName")
// 	if err != nil {
// 		return nil, err
// 	}

// 	gcpLoggerObj = &gcpLogger{
// 		Client:        client,
// 		LoggerClient:  client.Logger(logName),
// 		LogName:       logName,
// 		chIngressLogs: make(chan LogMessage, 50),
// 	}

// 	go gcpLoggerObj.workerIngestLogs(ctx)

// 	return gcpLoggerObj, nil
// }

// func (self *gcpLogger) Log(name string, sev Sevlevel) {

// 	self.chIngressLogs <- LogMessage{
// 		Name:     name,
// 		Severity: sev,
// 	}
// }

// func (self *gcpLogger) workerIngestLogs(ctx context.Context) {

// 	defer func() {
// 		self.Client.Close()
// 	}()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Println("gcp logger worker received signal to finish")
// 			return

// 		case msg := <-self.chIngressLogs:
// 			self.LoggerClient.StandardLogger(getSeverityMapForGCP(msg.Severity)).Print(getGCPLogString(msg.Name))
// 		}
// 	}
// }

// func getGCPLogString(content string) string {

// 	return fmt.Sprintf("%s : %s", time.Now().Format(FORMAT_DATETIME_UNIVERSAL), content)
// }

// func getSeverityMapForGCP(sev Sevlevel) logging.Severity {

// 	switch sev {
// 	case Debug:
// 		return logging.Debug
// 	case Info:
// 		return logging.Info
// 	case Notice:
// 		return logging.Notice
// 	case Warning:
// 		return logging.Warning
// 	case Error:
// 		return logging.Error
// 	case Critical:
// 		return logging.Critical
// 	case Alert:
// 		return logging.Alert
// 	case Emergency:
// 		return logging.Emergency
// 	case Default:
// 		fallthrough
// 	default:
// 		return logging.Default
// 	}
// }
