package log

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

// Log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// log levels
const (
	InfoLevel  = iota // default 0
	ErrorLevel        // default 1
	Disabled          // default 2
)

// SetLevel controls log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		// Set logger output as standard output
		logger.SetOutput(os.Stdout)
	}

	if ErrorLevel < level {
		// Logger output disabled
		errorLog.SetOutput(io.Discard)
	}

	if InfoLevel < level {
		// Logger output disabled
		infoLog.SetOutput(io.Discard)
	}
}
