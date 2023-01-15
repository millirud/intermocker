package entity

import (
	"io"
	"time"
)

func NewHttpResponseBuilder(statusCode int, contentType string, contentReaderFn ContentReaderFn) *HttpResponseBuilder {
	return &HttpResponseBuilder{
		response: NewHttpResponse(statusCode, contentType, contentReaderFn),
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
	b.response.Header[key] = value
	return b
}

func (b *HttpResponseBuilder) New() *HttpResponse {
	return b.response
}

func NewHttpResponse(statusCode int, contentType string, contentReaderFn ContentReaderFn) *HttpResponse {
	return &HttpResponse{
		StatusCode:      statusCode,
		Header:          make(map[string]string, 0),
		ContentType:     contentType,
		ContentReaderFn: contentReaderFn,
	}
}

type ContentReaderFn = func() (reader io.Reader, length int64, contentType string, err error)

type HttpResponse struct {
	StatusCode      int
	ContentReaderFn ContentReaderFn
	ContentType     string
	Header          map[string]string
	Timeout         time.Duration
}
