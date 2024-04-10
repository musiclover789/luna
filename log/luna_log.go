package luna_log

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	logger  *log.Logger
	logFile *os.File
)

// init is called automatically when the package is imported to initialize the logger.
// 当导入包时会自动调用 init 函数来初始化日志器。
func init() {
	if err := InitLogger(true); err != nil {
		panic(err)
	}
}

// InitLogger initializes the logger.
// 初始化日志器。
// Parameters:
//   - logToConsole: Whether to output logs to the console.
//
// Return:
//   - error: An error message. Returns nil if initialization is successful.
//
// Parameters:
//   - logToConsole: 是否将日志输出到控制台。
//
// Return:
//   - error: 错误信息，如果初始化成功则返回 nil。
func InitLogger(logToConsole bool) error {
	// Create the log file.
	logPath := filepath.Join(".", "app.log")
	var err error
	logFile, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// Output logs to both the file and the console.
	if logToConsole {
		logger = log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)
	} else {
		logger = log.New(logFile, "", log.LstdFlags)
	}

	return nil
}

// CloseLogger closes the logger.
// 关闭日志器。
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}

// Logf outputs formatted logs.
// 输出格式化日志。
// Parameters:
//   - format: The format string.
//   - v: The parameter list.
//
// Parameters:
//   - format: 格式化字符串。
//   - v: 参数列表。
func Logf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

// Log outputs logs.
// 输出日志。
// Parameters:
//   - v: The parameter list.
//
// Parameters:
//   - v: 参数列表。
func Log(v ...interface{}) {
	logger.Println(v...)
}

// LogError outputs error logs.
// 输出错误日志。
// Parameters:
//   - v: The parameter list.
//
// Parameters:
//   - v: 参数列表。
func LogError(v ...interface{}) {
	logger.SetPrefix("[ERROR] ")
	logger.SetOutput(os.Stderr)
	logger.Println(v...)
	logger.SetPrefix("")
	logger.SetOutput(logFile)
}

// LogFatal outputs fatal error logs and exits the program.
// 输出致命错误日志并退出程序。
// Parameters:
//   - v: The parameter list.
//
// Parameters:
//   - v: 参数列表。
func LogFatal(v ...interface{}) {
	logger.SetPrefix("[FATAL] ")
	logger.SetOutput(os.Stderr)
	logger.Fatalln(v...)
	logger.SetPrefix("")
	logger.SetOutput(logFile)
	defer CloseLogger()
}
