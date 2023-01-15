package entity

import (
	"io"
	"net/http"
	"time"
)

func NewHttpResponseBuilder(statusCode int, contentType string, reader io.ReadCloser) *HttpResponseBuilder {
	return &HttpResponseBuilder{
		response: NewHttpResponse(statusCode, contentType, reader),
	}
}

type HttpResponseBuilder struct {
	response *HttpResponse
}

func (b *HttpResponseBuilder) Timeout(timeout *time.Duration) *HttpResponseBuilder {
	b.response.Timeout = *timeout
	return b
}

func (b *HttpResponseBuilder) HeaderValue(key string, value string) *HttpResponseBuilder {
	b.response.Header.Set(key, value)
	return b
}

func (b *HttpResponseBuilder) Reader(reader io.ReadCloser) *HttpResponseBuilder {
	b.response.Reader = reader
	return b
}

func NewHttpResponse(statusCode int, contentType string, reader io.ReadCloser) *HttpResponse {
	return &HttpResponse{
		StatusCode: statusCode,
		Reader:     reader,
		Header: http.Header{
			"content-type": []string{contentType},
		},
	}
}

type HttpResponse struct {
	StatusCode int
	Reader     io.ReadCloser
	Header     http.Header
	Timeout    time.Duration
}
