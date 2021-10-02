package configs

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

type Tracing struct{
	tracer opentracing.Tracer
	closer io.Closer
}
var trace *Tracing

// Init returns an instance of Jaeger Tracer.
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type: "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func InitTracer(){
	trace = new(Tracing)
	tracer, closer := Init(viper.GetString("Trace_Name"))
	opentracing.SetGlobalTracer(tracer)
	trace.tracer = tracer
	trace.closer = closer
}

func GetTracer() opentracing.Tracer{
	return trace.tracer
}

func CloseTracer(){
	defer trace.closer.Close()
}