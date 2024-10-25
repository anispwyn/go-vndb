package vndb_test

import (
	"context"
	"os"
	"testing"
)

func TestVndb_Stats(t *testing.T) {
	stats, err := v.Stats(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", stats)
}

func TestVndb_User(t *testing.T) {
	users, err := v.User(context.Background(), []string{"test"}, []string{"lengthvotes"})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", users)
}

func TestVndb_AuthInfo(t *testing.T) {
	token := os.Getenv(TokenKey)
	if token == "" {
		t.Skip("token is empty")
	}
	v.SetToken(token)
	info, err := v.AuthInfo(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", info)
}
