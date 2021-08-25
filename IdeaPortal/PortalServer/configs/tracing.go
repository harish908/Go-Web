package configs

import (
	"PortalServer/internal/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
)

type Tracing struct {
	tracer opentracing.Tracer
}

var trace *Tracing

func InitTracer() {
	trace = new(Tracing)
	tracer, closer := tracing.Init(viper.GetString("Trace_Name"))
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	trace.tracer = tracer
}

func GetTracer() opentracing.Tracer {
	return trace.tracer
}
