# Logger Package

The Logger package is a very simple and flexible logging library for Go applications. It allows you to log messages at different levels (DEBUG, INFO, WARN, ERROR, FATAL) and supports multiple output options (console only, file only, or both).

## Features

- Log messages at various levels: DEBUG, INFO, WARN, ERROR, and FATAL.
- Choose output options: log to console, log to a file, or log to both.
- Simple and easy to use.

## Installation

To use the Logger package, you can clone the repository or use it as a module in your Go project.

```bash
go get github.com/imafaz/logger
```

## Usage

### Initialization

To initialize the logger, use the `InitLogger` function. You need to provide the log file name and the desired output option.

```go
package main

import (
	"github.com/imafaz/logger"
)

func main() {
	Logger.InitLogger("app.log", Logger.CONSOLE_AND_FILE)

	Logger.Log(Logger.INFO, "Application started")
	Logger.Log(Logger.DEBUG, "This is a debug message")
	Logger.Log(Logger.ERROR, "An error occurred")
}
```

### Log Levels

You can log messages at different levels:

- `Logger.DEBUG`: For debug messages.
- `Logger.INFO`: For informational messages.
- `Logger.WARN`: For warning messages.
- `Logger.ERROR`: For error messages.
- `Logger.FATAL`: For fatal messages that will terminate the application.

### Output Options

You can choose from the following output options:

- `Logger.CONSOLE_ONLY`: Log messages only to the console.
- `Logger.FILE_ONLY`: Log messages only to a file.
- `Logger.CONSOLE_AND_FILE`: Log messages to both the console and a file.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.