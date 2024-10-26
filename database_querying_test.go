package vndb_test

import (
	"context"
	"testing"

	"github.com/anispwyn/go-vndb"
)

func TestVndb_Vn(t *testing.T) {
	res, err := v.Vn(context.Background(), vndb.Request{
		Filters: []string{"id", "=", "v17"},
		Fields:  "title, image.url",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)
}
