package framework

import "fmt"

func log(format string, a ...interface{}) {
	fmt.Printf("[FRMWRK] %s\n", fmt.Sprintf(format, a...))
}

func logerr(err error) {
	fmt.Printf("[FRMWRK] [ERROR] %s\n", err.Error())
}
