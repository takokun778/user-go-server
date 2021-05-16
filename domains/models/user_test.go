package models

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	result := NewUser("大園玲")
	t.Log(result)
}
