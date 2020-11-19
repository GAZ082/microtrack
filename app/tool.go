package microtrack

import (
	"log"
	"runtime/debug"
)

//CheckError two modes, info or panic.
func checkError(err error, mode string) bool {
	if err != nil {
		s := string(debug.Stack())
		if mode == "info" {
			log.Printf("INFO: (%v) %v", s, err)
		} else {
			log.Panicf("PANIC: (%v) %v", s, err)
		}
		return true
	}
	return false
}
