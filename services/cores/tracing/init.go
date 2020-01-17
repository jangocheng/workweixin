package tracing

import (
	"io"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func Init(service string) (opentracing.Tracer, io.Closer) {
	agentHost := os.Getenv("JAEGER_AGENT_HOST")
	agentPort := os.Getenv("JAEGER_AGENT_PORT")
	if agentHost == "" {
		log.Fatal("jaeger agent host should not be null")
	}
	if agentPort == "" {
		agentPort = "6831"
	}
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: agentHost + ":" + agentPort,
		},
		ServiceName: service,
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("can not init Jaeger: #%v", err)
	}
	return tracer, closer
}
