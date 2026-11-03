package folcon

import (
	"log"
	"os"
	"reflect"
	"runtime"
)

var Logger = log.New(os.Stdout, "[folcon-debug] ", 0)

// ANSI Color Constants for Terminal
const (
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
)

func colorizeMethod(method string) string {
	switch method {
	case "GET":
		return colorGreen + method + colorReset
	case "POST":
		return colorBlue + method + colorReset
	case "PUT":
		return colorYellow + method + colorReset
	case "DELETE":
		return colorRed + method + colorReset
	case "PATCH":
		return colorCyan + method + colorReset
	default:
		return colorGray + method + colorReset
	}
}

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func debugPrintRoute(httpMethod, absolutePath string, handlers HandlersChain) {
	nuHandlers := len(handlers)
	handlerName := nameOfFunction(handlers[nuHandlers-1])

	coloredMethod := colorizeMethod(httpMethod)

	LogPrint("%-15s %-25s --> %s (%d handlers)\n", coloredMethod, absolutePath, handlerName, nuHandlers)
}

func LogPrint(format string, values ...interface{}) {
	Logger.Printf(format, values...)
}
