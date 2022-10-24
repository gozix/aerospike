// Copyright 2021 Stanislav Kustov. All rights reserved.
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package aerospike

import (
	"fmt"
	"net"
	"strconv"

	"github.com/gozix/di"
	"github.com/gozix/glue/v3"
	gzViper "github.com/gozix/viper/v3"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/spf13/viper"
)

// Bundle implements glue.Bundle interface.
type Bundle struct{}

// BundleName is default definition name.
const BundleName = "aerospike.cluster"

// Bundle implements glue.Bundle interface.
var _ glue.Bundle = (*Bundle)(nil)

// NewBundle create bundle instance.
func NewBundle() *Bundle {
	return new(Bundle)
}

// Name implements the glue.Bundle interface.
func (b *Bundle) Name() string {
	return BundleName
}

// Build implements the glue.Bundle interface.
func (b *Bundle) Build(builder di.Builder) error {
	return builder.Provide(b.provideAerospikeClient)
}

// DependsOn implements the glue.DependsOn interface.
func (b *Bundle) DependsOn() []string {
	return []string{
		gzViper.BundleName,
	}
}

func (b *Bundle) provideAerospikeClient(cfg *viper.Viper) (*as.Client, func() error, error) {
	var (
		suffix = fmt.Sprintf("%s.", "aerospike_cluster")
		nodes  = cfg.GetStringSlice(suffix + "nodes")
		hosts  = make([]*as.Host, 0, len(nodes))
	)

	for _, node := range nodes {
		var host, portString, err = net.SplitHostPort(node)
		if err != nil {
			return nil, nil, err
		}

		var port int
		if port, err = strconv.Atoi(portString); err != nil {
			return nil, nil, err
		}

		hosts = append(hosts, as.NewHost(host, port))
	}

	var policy = as.NewClientPolicy()
	if cfg.IsSet(suffix + "idle_timeout") {
		policy.IdleTimeout = cfg.GetDuration(suffix + "idle_timeout")
	}

	if cfg.IsSet(suffix + "timeout") {
		policy.Timeout = cfg.GetDuration(suffix + "timeout")
	}

	if cfg.IsSet(suffix + "login_timeout") {
		policy.LoginTimeout = cfg.GetDuration(suffix + "login_timeout")
	}

	var client, err = as.NewClientWithPolicyAndHost(policy, hosts...)
	if err != nil {
		return nil, nil, err
	}

	var closer = func() error {
		client.Close()
		return nil
	}

	return client, closer, nil
}
