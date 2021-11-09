package tracing

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// StartSpanFromRequest extracts the parent span context from the inbound HTTP request
// and starts a new child span if there is a parent span.
func StartSpanFromRequest(r *http.Request) opentracing.Span {
	tracer := GetTracer()
	spanCtx, _ := Extract(tracer, r)
	return tracer.StartSpan("HTTP "+r.Method+":"+r.URL.Path, ext.RPCServerOption(spanCtx))
}
