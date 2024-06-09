# logging utility

Purpose of this module is to enable applications to store/send logs in different environments like local file, in Google Cloud Platform, Amazon AWS, etc.

### Type and Level of logs supported

The log module supports 6 levels of logs:

- LEVEL_DEBUG - recommended to be used to denote logs that provide information useful for debugging issues with application / business logic.
- LEVEL_INFO - recommended to denote logs that provide information (only) to the viewer.
- LEVEL_WARNING - recommended to denote logs that inform of a possible error condition to occur if a certain state(s) are not fixed.
- LEVEL_ERROR - recommended for logs that denote a break in functionality, state or logic of an application
- LEVEL_ALERT - recommended for logs with errors that can potentially lead to degradation of application, business availability
- LEVEL_FATAL - recommended for logs denoting a complete failure of a application sub-system, business processing logic, or the entire application.

### Tracing

The log module could be running either with trace set / unset. Tracing can/should be enabled when init'ing the log module and remains relevant irrespective of the types/levels of log messages bwing created.

Tracing mode should be switched on when there is a need to identify the file and method name from where a log line originates. This must be set when init'ing the module.

```
log.Init(true)
```

Example log lines with tracing enabled:

```
08-06-2024 23:30:05 [W] logger_test.go:TestPrint:10 this should be a warning mesage with string=hello world Int=666
08-06-2024 23:30:05 [E] logger_test.go:TestPrint:11 error flagged with string=hello world Int=666
08-06-2024 23:30:05 [A] logger_test.go:TestPrint:12 this is an alert for string=hello world Int=666
08-06-2024 23:30:05 [D] logger_test.go:TestPrint:13 this log contains debugging
```

Note: To change the setting of `tracing`, the log module needs to be initialized once more, and thereby inferring the need for an application restart.

### Supported Methods

// todo
