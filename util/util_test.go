package util_test

import (
	"oauth2-server/log"
	"testing"

	"github.com/pborman/uuid"
)

func TestUUID(t *testing.T) {
	log.INFO.Println(uuid.New())
}
