package vndb

import "net/http"

const UserAgent = "go-vndb/0.1"

type Vndb struct {
	client   *http.Client
	endpoint string
	token    string
}

type Option func(*Vndb)

func New(options ...Option) *Vndb {
	v := &Vndb{
		client:   http.DefaultClient,
		endpoint: "https://api.vndb.org/kana",
	}
	for _, o := range options {
		o(v)
	}
	return v
}

var (
	UseSandBox Option = func(v *Vndb) {
		v.endpoint = "https://beta.vndb.org/api/kana"
	}
)

func WithClient(client *http.Client) Option {
	return func(v *Vndb) {
		v.client = client
	}
}

func WithEndpoint(endpoint string) Option {
	return func(v *Vndb) {
		v.endpoint = endpoint
	}
}

func WithToken(token string) Option {
	return func(v *Vndb) {
		v.token = token
	}
}

func (v *Vndb) SetToken(token string) {
	v.token = token
}
