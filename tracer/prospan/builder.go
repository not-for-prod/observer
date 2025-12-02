package prospan

import (
	"context"

	"github.com/not-for-prod/observer/logger"
	"github.com/not-for-prod/observer/tracer/autoname"
	"go.opentelemetry.io/otel"
)

const (
	initialSkipFrames  = 2
	attributesPrealloc = 2
)

// Builder дает добавть в спан какие то промежуточнык данные
type Builder struct {
	tracerName string
	spanName   string
	skipFrames int
	attributes map[string]any
}

// WithTracerName sets prospan name.
// If WithTracerName is not invoked before Start prospan name is set
// automatically using runtime.Caller().
func WithTracerName(name string) Builder {
	return Builder{
		tracerName: name,
		attributes: make(map[string]any, attributesPrealloc),
		skipFrames: initialSkipFrames,
	}
}

func WithSpanName(name string) Builder {
	return Builder{
		tracerName: name,
		attributes: make(map[string]any, attributesPrealloc),
		skipFrames: initialSkipFrames,
	}
}

func WithAttribute(key string, value any) *Builder {
	return &Builder{
		tracerName: "",
		attributes: map[string]any{key: value},
		skipFrames: initialSkipFrames,
	}
}

func WithRequest(request any) *Builder {
	key := "request"

	return &Builder{
		tracerName: "",
		attributes: map[string]any{key: request},
		skipFrames: initialSkipFrames,
	}
}

func (b *Builder) Start(ctx context.Context) (context.Context, ProSpan) {
	spanName := b.spanName

	if spanName == "" {
		spanName = autoname.GetRuntimeFunc(b.skipFrames)
	}

	ctx, span := otel.Tracer(b.tracerName).Start(ctx, spanName)
	ctx, l := logger.Instance().With(ctx, "trace", span.SpanContext().TraceID().String())

	for key, val := range b.attributes {
		setAttr(span, key, val)
	}

	return ctx, ProSpan{
		span:   span,
		logger: l,
	}
}

func (b *Builder) WithAttribute(key string, value any) *Builder {
	b.attributes[key] = value
	return b
}

func (b *Builder) WithSpanName(name string) *Builder {
	b.spanName = name
	return b
}
