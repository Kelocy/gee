package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// Print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	// Declare a string builder object for building strings
	var str strings.Builder
	// Append initial information
	str.WriteString(message + "\nTraceback:")
	// Only include n valid information
	for _, pc := range pcs[:n] {
		// Retrive the function object corresponding to pc (program counter)
		fn := runtime.FuncForPC(pc)
		// Get the file path and line number
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
