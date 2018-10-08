package tracing

import (
	"fmt"
	"io"
	"os"
	"time"

	alog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
	opentracing "github.com/opentracing/opentracing-go"
	config "github.com/uber/jaeger-client-go/config"
)

func NewTracer(serviceName, host string) (opentracing.Tracer, io.Closer) {
	sampler := &config.SamplerConfig{
		Type:  "const",
		Param: 1, // This reports 100%. Need to let user choose. Functional Options?
	}
	// NewCompositeReporter is used internally by library which ncludes logger reporter
	// created from logger automatically
	reporter := &config.ReporterConfig{
		LogSpans:            true,
		BufferFlushInterval: 1 * time.Second,
		LocalAgentHostPort:  host,
	}

	cfg := &config.Configuration{
		Sampler:  sampler,
		Reporter: reporter,
	}

	log := &alog.Logger{
		Handler: text.New(os.Stdout),
		Level:   alog.DebugLevel,
	}

	tracer, closer, err := cfg.New(serviceName, config.Logger(log))
	if err != nil {
		panic(fmt.Sprintf("cannot init Jaeger client error: %v\n", err))
	}
	return tracer, closer
}
