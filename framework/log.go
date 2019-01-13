package framework

import "fmt"

func log(msg string) {
	fmt.Printf("[FRMWRK] %s\n", msg)
}

func logerr(err error) {
	fmt.Printf("[FRMWRK] [ERROR] %s\n", err.Error())
}
