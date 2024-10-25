package vndb

import (
	"context"
	"net/url"
	"strings"
)

func (v *Vndb) PostUList(ctx context.Context, request Request) (*Response, *Error) {
	return v.databaseQuering(ctx, path.uList, request)
}

func (v *Vndb) GetUListLabels(ctx context.Context, userID string, fields []string) (*UListLabels, *Error) {
	query := url.Values{}
	query.Add("user", userID)
	if len(fields) > 0 {
		query.Add("fields", strings.Join(fields, ","))
	}
	resp, err := v.get(ctx, path.uListLabels, query)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	labels, err := decode[UListLabels](resp)
	if err != nil {
		return nil, err
	}
	return labels, nil
}

func (v *Vndb) PatchUList(ctx context.Context, id string, uList PatchUList) *Error {
	req, err := encode(uList)
	if err != nil {
		return err
	}
	return v.patch(ctx, path.uList+"/"+id, req)
}

func (v *Vndb) PatchRList(ctx context.Context, id string, status int) *Error {
	req, err := encode(map[string]int{"status": status})
	if err != nil {
		return err
	}
	return v.patch(ctx, path.rList+"/"+id, req)
}

func (v *Vndb) DeleteUList(ctx context.Context, id string) *Error {
	return v.delete(ctx, path.uList+"/"+id)
}

func (v *Vndb) DeleteRList(ctx context.Context, id string) *Error {
	return v.delete(ctx, path.rList+"/"+id)
}
