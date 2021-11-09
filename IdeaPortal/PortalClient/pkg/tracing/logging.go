package tracing

import (
	"github.com/opentracing/opentracing-go"
	ottag "github.com/opentracing/opentracing-go/ext"
	otlog "github.com/opentracing/opentracing-go/log"
)

func LogURL(span opentracing.Span, path string, method string) {
	// set tags
	// ottag.SpanKindRPCClient.Set(span)
	ottag.HTTPUrl.Set(span, path)
	ottag.HTTPMethod.Set(span, method)
}

func LogError(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogFields(otlog.Error(err))
}
