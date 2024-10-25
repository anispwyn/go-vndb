package vndb

import "context"

func (v *Vndb) databaseQuering(ctx context.Context, path string, request Request) (*Response, *Error) {
	req, err := encode(request)
	if err != nil {
		return nil, err
	}
	resp, err := v.post(ctx, path, req)
	if err != nil {
		return nil, err
	}
	defer resp.Close()
	response, err := decode[Response](resp)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (v *Vndb) Vn(ctx context.Context, request Request) (*Response, *Error) {
	return v.databaseQuering(ctx, path.vn, request)
}

func (v *Vndb) Release(ctx context.Context, request Request) (*Response, *Error) {
	return v.databaseQuering(ctx, path.release, request)
}

func (v *Vndb) Producer(ctx context.Context, request Request) (*Response, *Error) {
	return v.databaseQuering(ctx, path.producer, request)
}

func (v *Vndb) Character(ctx context.Context, request Request) (*Response, *Error) {
	return v.databaseQuering(ctx, path.character, request)
}

func (v *Vndb) Staff(ctx context.Context, request Request) (*Response, *Error) {
	return v.databaseQuering(ctx, path.staff, request)
}

func (v *Vndb) Tag(ctx context.Context, request Request) (*Response, *Error) {
	return v.databaseQuering(ctx, path.tag, request)
}

func (v *Vndb) Trait(ctx context.Context, request Request) (*Response, *Error) {
	return v.databaseQuering(ctx, path.trait, request)
}
