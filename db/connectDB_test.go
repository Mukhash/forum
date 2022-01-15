package db

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	db, _ := ConnectBD()

	if db != nil {
		t.Errorf("Failed at connecting to DB")
	}
}
