package vndb_test

import (
	"context"
	"os"
	"testing"

	"github.com/tuihub/go-vndb"
)

func TestVndb_PostUList(t *testing.T) {
	res, err := v.PostUList(context.Background(), vndb.Request{
		User:    "u2",
		Fields:  "id, vote, vn.title",
		Filters: []string{"label", "=", "7"},
		Sort:    "vote",
		Reverse: true,
		Results: 10,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

func TestVndb_GetUListLabels(t *testing.T) {
	res, err := v.GetUListLabels(context.Background(), "u1", []string{"count"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}

func TestVndb_PatchUList(t *testing.T) {
	token := os.Getenv(TokenKey)
	if token == "" {
		t.Skip("token is empty")
	}
	err := v.PatchUList(context.Background(), "v17", vndb.PatchUList{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestVndb_PatchRList(t *testing.T) {
	token := os.Getenv(TokenKey)
	if token == "" {
		t.Skip("token is empty")
	}
	err := v.PatchRList(context.Background(), "r12", 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVndb_DeleteUList(t *testing.T) {
	token := os.Getenv(TokenKey)
	if token == "" {
		t.Skip("token is empty")
	}
	err := v.DeleteUList(context.Background(), "v17")
	if err != nil {
		t.Fatal(err)
	}
}

func TestVndb_DeleteRList(t *testing.T) {
	token := os.Getenv(TokenKey)
	if token == "" {
		t.Skip("token is empty")
	}
	err := v.DeleteRList(context.Background(), "r12")
	if err != nil {
		t.Fatal(err)
	}
}
