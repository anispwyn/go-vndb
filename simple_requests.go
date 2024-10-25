package vndb

import (
	"context"
	"net/url"
	"strings"
)

func (v *Vndb) Stats(ctx context.Context) (*Stats, *Error) {
	resp, err := v.get(ctx, path.stats, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	stats, err := decode[Stats](resp)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func (v *Vndb) User(ctx context.Context, userIDOrName []string, fields []string) (map[string]User, *Error) {
	query := url.Values{}
	for _, q := range userIDOrName {
		query.Add("q", q)
	}
	if len(fields) > 0 {
		query.Add("fields", strings.Join(fields, ","))
	}
	resp, err := v.get(ctx, path.user, query)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	users, err := decode[map[string]User](resp)
	if err != nil {
		return nil, err
	}
	return *users, nil
}

func (v *Vndb) AuthInfo(ctx context.Context) (*AuthInfo, *Error) {
	resp, err := v.get(ctx, path.authInfo, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	info, err := decode[AuthInfo](resp)
	if err != nil {
		return nil, err
	}
	return info, nil
}
