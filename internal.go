package goconfluence

import "fmt"

// DebugFlag is the global debugging variable
var DebugFlag = false

// EnableDebug enables debug output
func EnableDebug(enable bool) {
	if enable {
		DebugFlag = true
	}
}

// Debug outputs debug messages
func Debug(msg interface{}) {
	if DebugFlag {
		fmt.Printf("%+v\n", msg)
	}
}
