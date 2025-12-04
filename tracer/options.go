package tracer

import "github.com/not-for-prod/observer/git"

type options struct {
	host           string
	serviceName    string
	serviceVersion string
}

func newOptions(opts ...Option) *options {
	o := defaultOptions()

	for _, opt := range opts {
		opt.apply(&o)
	}

	return &o
}

// defaultOptions by default we use git info to fill service name and version.
func defaultOptions() options {
	gitData := git.GetCommitInfo()

	return options{
		host:           "localhost:4317",
		serviceName:    gitData.Project,
		serviceVersion: gitData.String(),
	}
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithHost(url string) Option {
	return optionFunc(
		func(o *options) {
			o.host = url
		},
	)
}

func WithServiceName(serviceName string) Option {
	return optionFunc(
		func(o *options) {
			o.serviceName = serviceName
		},
	)
}

func WithServiceVersion(serviceVersion string) Option {
	return optionFunc(
		func(o *options) {
			o.serviceVersion = serviceVersion
		},
	)
}
