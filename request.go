package vndb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func encode(v any) ([]byte, *Error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, &Error{msg: fmt.Sprintf("encode request: %s", err)}
	}
	return b, nil
}

func decode[T any](r io.ReadCloser) (*T, *Error) {
	res := new(T)
	dec := json.NewDecoder(r)
	if err := dec.Decode(res); err != nil {
		return nil, &Error{msg: fmt.Sprintf("decode response: %s", err)}
	}
	return res, nil
}

func (v *Vndb) get(ctx context.Context, path string, query url.Values) (io.ReadCloser, *Error) {
	return v.do(ctx, path, http.MethodGet, query, nil)
}

func (v *Vndb) post(ctx context.Context, path string, body []byte) (io.ReadCloser, *Error) {
	return v.do(ctx, path, http.MethodPost, url.Values{}, body)
}

func (v *Vndb) patch(ctx context.Context, path string, body []byte) *Error {
	_, err := v.do(ctx, path, http.MethodPatch, url.Values{}, body)
	return err
}

func (v *Vndb) delete(ctx context.Context, path string) *Error {
	_, err := v.do(ctx, path, http.MethodDelete, url.Values{}, nil)
	return err
}

func (v *Vndb) do(ctx context.Context, path string, method string, query url.Values, body []byte) (io.ReadCloser, *Error) {
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		v.endpoint+path+"?"+query.Encode(),
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, &Error{msg: fmt.Sprintf("create request: %s", err)}
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")
	if v.token != "" {
		req.Header.Set("Authorization", "token "+v.token)
	}

	resp, err := v.client.Do(req)
	if err != nil {
		return nil, &Error{msg: fmt.Sprintf("do request: %s", err)}
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		_ = resp.Body.Close()
		return nil, &Error{
			apiError: &apiError{
				StatusCode: resp.StatusCode,
				Reason:     resp.Status,
			},
		}
	}

	return resp.Body, nil
}
