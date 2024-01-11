package main

import (
	"fmt"
	"gee"
	"net/http"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.Default()

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Gee\n")
	})

	// Index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"Gee"}
		c.String(http.StatusOK, names[100])
	})

	r.RUN(":9999")
}
