package tracings

import (
	"io"
	
	"github.com/opentracing/opentracing-go"
	"github.com/vnotes/workweixin/services/cores/tracing"
)

var (
	Tracer opentracing.Tracer
	closer io.Closer
)

func InitTracing(service string) {
	Tracer, closer = tracing.Init(service)
	opentracing.SetGlobalTracer(Tracer)
}

func CloseTracer() error {
	return closer.Close()
}
