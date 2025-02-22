/*
Copyright 2023 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package placement

import "github.com/dapr/dapr/tests/integration/framework/process/exec"

// Option is a function that configures the process.
type Option func(*options)

// options contains the options for running Placement in integration tests.
type options struct {
	execOpts []exec.Option

	id                  string
	port                int
	healthzPort         int
	metricsPort         int
	initialCluster      string
	initialClusterPorts []int
}

func WithExecOptions(execOptions ...exec.Option) Option {
	return func(o *options) {
		o.execOpts = execOptions
	}
}

func WithPort(port int) Option {
	return func(o *options) {
		o.port = port
	}
}

func WithID(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func WithHealthzPort(port int) Option {
	return func(o *options) {
		o.healthzPort = port
	}
}

func WithMetricsPort(port int) Option {
	return func(o *options) {
		o.metricsPort = port
	}
}

func WithInitialCluster(initialCluster string) Option {
	return func(o *options) {
		o.initialCluster = initialCluster
	}
}

func WithInitialClusterPorts(ports []int) Option {
	return func(o *options) {
		o.initialClusterPorts = ports
	}
}
