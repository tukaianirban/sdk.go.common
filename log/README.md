# logging utility

Purpose of this module is to enable applications to store/send logs in different environments like local file, in Google Cloud Platform, Amazon AWS, etc.

### Modes of operation

The logger module can be used in different modes:

- MODE_DEFAULT: default logging of the passed in log message. No arg to be passed in the Init() call.

- MODE_LOCALFILE: logs are stored into a local file where the application is running.

  - Init() must be called with the path of the (log) storage file.
  - If the file exists, then logs are appended to it.
  - Application must have write permissions to the (log) storage file.

- MODE_GCP: logs are sent to a GCP service.

  - Init() must be called with the path of the GCP project config file.
  - Details of the implementation and AWS service yet to be decided.

- MODE_AWS: logs are sent to a AWS service.
  - Init() must be called with the path of the AWS project config file.
  - Details of the implementation and AWS service is yet to be decided.

### Some details

Calling fatal[f|ln]() will result in the log being pushed to the respective provider and then a call to os.Exit(1)

### Desired (Sample) logging outputs

Print:
DD-MM-YYYY HH:mm:ss [<SEV>] <... log message ...>
Debug:
DD-MM-YYYY HH:mm:ss.msec [<SEV>] <... log message ...>
Trace:
DD-MM-YYYY HH:mm:ss [<SEV>] <filaname:functionname:linenumber <... log message ...>

### Upcoming features

- Implementation of MODE_GCP
- Implementation of MODE_AWS
- Double-down on flushing out errors on graceful/ungraceful application shutdown.
