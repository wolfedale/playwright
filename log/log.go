package log

import "github.com/fatih/color"

// LogSimple - prints white message
func LogSimple(format string, args ...interface{}) {
	color.White(format, args...)
}

// LogSkip - print yellow message
func LogSkip(format string, args ...interface{}) {
	color.Yellow(format, args...)
}

// LogError - prints red error message
func LogError(format string, args ...interface{}) {
	color.Red(format, args...)
}

// LogSuccess - prints green success message
func LogSuccess(format string, args ...interface{}) {
	color.Green(format, args...)
}
