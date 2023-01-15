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
	b.response.Headers.Set(key, value)
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
		Headers: http.Header{
			"content-type": []string{contentType},
		},
	}
}

type HttpResponse struct {
	StatusCode int
	Reader     io.ReadCloser
	Headers    http.Header
	Timeout    time.Duration
}
