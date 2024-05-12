package fasthttp

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type Settings struct {
	CacheSeconds int
	Proxy        string
	Mtls         MutualTLS
}

type Client struct {
	client *fasthttp.Client
}

type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string]string
}

func Init(settings ...Settings) Client {
	client := new(fasthttp.Client)
	cacheSeconds := 30
	if len(settings) > 0 {
		s := settings[0]

		if t := s.CacheSeconds; t != 0 {
			cacheSeconds = t
		}

		if p := s.Proxy; p != "" {
			client.Dial = fasthttpproxy.FasthttpHTTPDialer(p)

		}

		if mtls := s.Mtls; len(mtls.certFile) > 0 && len(mtls.keyFile) > 0 {
			cert, _ := tls.X509KeyPair(mtls.certFile, mtls.keyFile)

			tlsConfig := &tls.Config{
				Certificates: []tls.Certificate{cert},
			}

			client.TLSConfig = tlsConfig
		}
	}

	client.MaxIdleConnDuration = time.Duration(cacheSeconds) * time.Second

	return Client{
		client: client,
	}

}
func (c Client) Send(r *Request) (response Response, erro error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(r.method)

	if r.contentType != "" {
		req.Header.SetContentType(r.contentType)
	}

	req.Header.DisableNormalizing()
	if len(r.headers) > 0 {
		for k, v := range r.pathParams {
			r.url = strings.Replace(r.url, ":"+k, v, -1)
		}
	}

	if len(r.queries) > 0 {
		queryArgs := fasthttp.AcquireArgs()
		defer fasthttp.ReleaseArgs(queryArgs)

		for k, v := range r.queries {
			queryArgs.Add(k, v)
		}
		req.SetRequestURI(fmt.Sprintf("%s?%s", r.url, queryArgs.String()))
	} else {
		req.SetRequestURI(r.url)
	}

	if r.host != "" {
		req.UseHostHeader = true
		req.Header.SetHost(r.host)
	}

	if len(r.body) > 0 {
		req.SetBody(r.body)
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if r.timeout > 0 {
		erro = c.client.DoTimeout(req, resp, r.timeout)
	} else {
		erro = c.client.Do(req, resp)
	}

	if erro == nil {
		response.StatusCode, response.Body, response.Headers = resp.StatusCode(), resp.Body(), fillResponseHeaders(&resp.Header)
	}

	return
}

func fillResponseHeaders(resHeader *fasthttp.ResponseHeader) map[string]string {
	headersArray := make(map[string]string)

	resHeader.VisitAll(func(k, v []byte) {
		headersArray[string(k)] = string(v)
	})

	return headersArray
}
