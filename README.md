# Logger Package

The `logger` package is a simple yet powerful logging library for Go applications. It supports different logging levels and multiple output options, making it flexible for various use cases.

## Features

- **Log Levels**: Supports `DEBUG`, `INFO`, `WARN`, `ERROR`, and `FATAL` levels.
- **Multiple Output Options**: Console, file, or both.
- **Automatic Log File Naming**: If no file is specified, the log file name defaults to the executable name.
- **Custom Log File Support**: Set a custom log file if needed.
- **Graceful Log File Closing**: Ensures proper resource management.

## Installation

To use the `logger` package in your project, install it using:

```bash
go get github.com/AfazTech/logger/v2@latest
```

## Usage

### Basic Setup

```go
package main

import (
	"github.com/AfazTech/logger/v2"
)

func main() {
	// Optional: Set output mode (default is CONSOLE_ONLY)
	logger.SetOutput(logger.CONSOLE_AND_FILE)

	logger.Info("Application started")
	logger.Debug("Debugging information")
	logger.Warn("Warning message")
	logger.Error("An error occurred")
}
```

### Log Levels

- `logger.Debug(messages ...)` – Debugging details
- `logger.Info(messages ...)` – General information
- `logger.Warn(messages ...)` – Warnings that may need attention
- `logger.Error(messages ...)` – Errors that should be handled
- `logger.Fatal(messages ...)` – Logs error and exits the program

Example:

```go
logger.Debug("This is a debug message")
logger.Info("User logged in successfully")
logger.Warn("Disk space running low")
logger.Error("Failed to connect to database")
logger.Fatal("Critical failure - shutting down")
```

### Formatted Logging

You can use formatted logging functions:

```go
logger.Infof("User %s logged in", "admin")
logger.Errorf("Error code: %d", 500)
```

### Output Options

You can control where logs are written:

- `logger.CONSOLE_ONLY` – Logs only to the console (default)
- `logger.FILE_ONLY` – Logs only to a file
- `logger.CONSOLE_AND_FILE` – Logs to both console and file

Example:

```go
logger.SetOutput(logger.FILE_ONLY)
```

### Custom Log File

By default, the log file is named after the executable. To set a custom file:

```go
logger.SetLogFile("my_logs.txt")
```

### Closing Log File

Call this when your application is shutting down:

```go
logger.CloseLogFile()
```

## Optional Functions

- `logger.SetOutput(option int)` – Change log output (`CONSOLE_ONLY`, `FILE_ONLY`, `CONSOLE_AND_FILE`). Default: `CONSOLE_ONLY`.
- `logger.SetLogFile(filename string)` – Set a custom log file. Default: executable name.
- `logger.CloseLogFile()` – Closes the log file (recommended before exit).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

