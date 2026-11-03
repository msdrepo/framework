package folcon

import (
	"log"
	"net/http"
	"runtime/debug"

	res "github.com/msdrepo/framework/responses"
)

// ErrorHandler recovers from panics and returns a structured ErrorResponse JSON
func ErrorHandler() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the stack trace for developer debugging
				log.Printf("[Folcon Panic Recovery] %v\n%s", err, debug.Stack())

				// Send standard ErrorResponse JSON
				c.JSON(http.StatusInternalServerError, res.ErrorResponse{
					Status:     "error",
					StatusCode: http.StatusInternalServerError,
					Message:    "Internal Server Error",
					Data:       nil,
				})
			}
		}()

		// Proceed to the next handler/controller
		c.Next()
	}
}
