package fasthttp

import (
	"time"
)

type Request struct {
	url         string
	method      string
	body        []byte
	contentType string
	headers     map[string]string
	queries     map[string]string
	pathParams  map[string]string
	timeout     time.Duration
	host        string
}

func Url(url string) *Request {
	return &Request{url: url}
}

func (r *Request) Url(url string) *Request {
	r.url = url
	return r
}

func Method(method string) *Request {
	return &Request{method: method}
}

func (r *Request) Method(method string) *Request {
	r.method = method
	return r
}

func Body(body []byte) *Request {
	return &Request{body: body}
}

func (r *Request) Body(body []byte) *Request {
	r.body = body
	return r
}

func ContentType(contentType string) *Request {
	return &Request{contentType: contentType}
}

func (r *Request) ContentType(contentType string) *Request {
	r.contentType = contentType
	return r
}

func PathParams(params map[string]string) *Request {
	r := &Request{
		pathParams: make(map[string]string),
	}
	r.pathParams = params
	return r
}

func (r *Request) PathParams(params map[string]string) *Request {
	if len(r.headers) == 0 {
		r.pathParams = make(map[string]string)
	}
	r.pathParams = params
	return r
}

func Queries(queries map[string]string) *Request {
	r := &Request{
		queries: make(map[string]string),
	}
	r.queries = queries
	return r
}

func (r *Request) Queries(queries map[string]string) *Request {
	if len(r.headers) == 0 {
		r.queries = make(map[string]string)
	}
	r.queries = queries
	return r
}

func Headers(headers map[string]string) *Request {
	r := &Request{
		headers: make(map[string]string),
	}
	r.headers = headers
	return r
}

func (r *Request) Headers(headers map[string]string) *Request {
	if len(r.headers) == 0 {
		r.headers = make(map[string]string)
	}
	r.headers = headers
	return r
}

func Timeout(timeout time.Duration) *Request {
	return &Request{timeout: timeout}
}

func (r *Request) Timeout(timeout time.Duration) *Request {
	r.timeout = timeout
	return r
}

func Host(host string) *Request {
	return &Request{host: host}
}

func (r *Request) Host(host string) *Request {
	r.host = host
	return r
}
