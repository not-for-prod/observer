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
	name       string
	skipFrames int
	attributes map[string]any
}

// Name sets prospan name.
// If Name is not invoked before Start prospan name is set
// automatically using runtime.Caller().
func Name(name string) Builder {
	return Builder{
		name:       name,
		attributes: make(map[string]any, attributesPrealloc),
		skipFrames: initialSkipFrames,
	}
}

func WithAttribute(key string, value any) *Builder {
	return &Builder{
		name:       "",
		attributes: map[string]any{key: value},
		skipFrames: initialSkipFrames,
	}
}

func (b *Builder) Start(ctx context.Context) (context.Context, ProSpan) {
	funcName := autoname.GetRuntimeFunc(b.skipFrames)
	ctx, span := otel.Tracer(b.name).Start(ctx, funcName)
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
