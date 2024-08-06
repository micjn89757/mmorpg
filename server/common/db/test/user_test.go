package test

import (
	"server/common/db"
	"testing"
)

func TestGetUsername(t *testing.T) {
	user := db.GetUserByAccount("111")
	t.Log(user.Password)
}