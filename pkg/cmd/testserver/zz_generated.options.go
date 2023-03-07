// Code generated by github.com/ecordell/optgen. DO NOT EDIT.
package testserver

import util "github.com/authzed/spicedb/pkg/cmd/util"

type ConfigOption func(c *Config)

// NewConfigWithOptions creates a new Config with the passed in options set
func NewConfigWithOptions(opts ...ConfigOption) *Config {
	c := &Config{}
	for _, o := range opts {
		o(c)
	}
	return c
}

// ToOption returns a new ConfigOption that sets the values from the passed in Config
func (c *Config) ToOption() ConfigOption {
	return func(to *Config) {
		to.GRPCServer = c.GRPCServer
		to.ReadOnlyGRPCServer = c.ReadOnlyGRPCServer
		to.HTTPGateway = c.HTTPGateway
		to.ReadOnlyHTTPGateway = c.ReadOnlyHTTPGateway
		to.LoadConfigs = c.LoadConfigs
		to.MaximumUpdatesPerWrite = c.MaximumUpdatesPerWrite
		to.MaximumPreconditionCount = c.MaximumPreconditionCount
		to.MaxCaveatContextSize = c.MaxCaveatContextSize
	}
}

// ConfigWithOptions configures an existing Config with the passed in options set
func ConfigWithOptions(c *Config, opts ...ConfigOption) *Config {
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithGRPCServer returns an option that can set GRPCServer on a Config
func WithGRPCServer(gRPCServer util.GRPCServerConfig) ConfigOption {
	return func(c *Config) {
		c.GRPCServer = gRPCServer
	}
}

// WithReadOnlyGRPCServer returns an option that can set ReadOnlyGRPCServer on a Config
func WithReadOnlyGRPCServer(readOnlyGRPCServer util.GRPCServerConfig) ConfigOption {
	return func(c *Config) {
		c.ReadOnlyGRPCServer = readOnlyGRPCServer
	}
}

// WithHTTPGateway returns an option that can set HTTPGateway on a Config
func WithHTTPGateway(hTTPGateway util.HTTPServerConfig) ConfigOption {
	return func(c *Config) {
		c.HTTPGateway = hTTPGateway
	}
}

// WithReadOnlyHTTPGateway returns an option that can set ReadOnlyHTTPGateway on a Config
func WithReadOnlyHTTPGateway(readOnlyHTTPGateway util.HTTPServerConfig) ConfigOption {
	return func(c *Config) {
		c.ReadOnlyHTTPGateway = readOnlyHTTPGateway
	}
}

// WithLoadConfigs returns an option that can append LoadConfigss to Config.LoadConfigs
func WithLoadConfigs(loadConfigs string) ConfigOption {
	return func(c *Config) {
		c.LoadConfigs = append(c.LoadConfigs, loadConfigs)
	}
}

// SetLoadConfigs returns an option that can set LoadConfigs on a Config
func SetLoadConfigs(loadConfigs []string) ConfigOption {
	return func(c *Config) {
		c.LoadConfigs = loadConfigs
	}
}

// WithMaximumUpdatesPerWrite returns an option that can set MaximumUpdatesPerWrite on a Config
func WithMaximumUpdatesPerWrite(maximumUpdatesPerWrite uint16) ConfigOption {
	return func(c *Config) {
		c.MaximumUpdatesPerWrite = maximumUpdatesPerWrite
	}
}

// WithMaximumPreconditionCount returns an option that can set MaximumPreconditionCount on a Config
func WithMaximumPreconditionCount(maximumPreconditionCount uint16) ConfigOption {
	return func(c *Config) {
		c.MaximumPreconditionCount = maximumPreconditionCount
	}
}

// WithMaxCaveatContextSize returns an option that can set MaxCaveatContextSize on a Config
func WithMaxCaveatContextSize(maxCaveatContextSize int) ConfigOption {
	return func(c *Config) {
		c.MaxCaveatContextSize = maxCaveatContextSize
	}
}
