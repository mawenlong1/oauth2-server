package log_test

import (
	"mwl/oauth2-server/log"
	"testing"
)

func TestLogger(t *testing.T) {
	log.INFO.Println("INFO message")
	log.WARNING.Println("WARNING message")
	log.ERROR.Println("ERROR message")
	log.FATAL.Println("FATAL message")
}
