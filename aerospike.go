package aerospike

import (
	"fmt"
	"net"
	"strconv"

	as "github.com/aerospike/aerospike-client-go"
	gzviper "github.com/gozix/viper/v2"
	"github.com/sarulabs/di/v2"
	"github.com/spf13/viper"
)

type (
	// Bundle implements the glue.Bundle interface.
	Bundle struct{}

	// Client is type alias of aerospike_client_go.Client
	Client as.Client
)

// BundleName is default definition name.
const BundleName = "aerospike.cluster"

// NewBundle create bundle instance.
func NewBundle() *Bundle {
	return new(Bundle)
}

// Name implements the glue.Bundle interface.
func (b *Bundle) Name() string {
	return BundleName
}

// Build implements the glue.Bundle interface.
func (b *Bundle) Build(builder *di.Builder) error {
	return builder.Add(
		di.Def{
			Name: BundleName,
			Build: func(ctn di.Container) (_ interface{}, err error) {
				var config *viper.Viper
				if err = ctn.Fill(gzviper.BundleName, &config); err != nil {
					return nil, err
				}

				var suffix = fmt.Sprintf("%s.", "aerospike_cluster")

				var (
					nodes = config.GetStringSlice(suffix + "nodes")
					hosts = make([]*as.Host, 0, len(nodes))
				)
				for _, node := range nodes {
					var host, portString, err = net.SplitHostPort(node)
					if err != nil {
						return nil, err
					}
					var port int
					if port, err = strconv.Atoi(portString); err != nil {
						return nil, err
					}
					hosts = append(hosts, as.NewHost(host, port))
				}

				var policy = as.NewClientPolicy()
				if config.IsSet(suffix + "idle_timeout") {
					policy.IdleTimeout = config.GetDuration(suffix + "idle_timeout")
				}
				if config.IsSet(suffix + "timeout") {
					policy.Timeout = config.GetDuration(suffix + "timeout")
				}
				if config.IsSet(suffix + "login_timeout") {
					policy.LoginTimeout = config.GetDuration(suffix + "login_timeout")
				}

				var client *as.Client
				client, err = as.NewClientWithPolicyAndHost(policy, hosts...)

				if err != nil {
					return nil, err
				}

				return client, nil
			},
			Close: func(obj interface{}) error {
				obj.(*as.Client).Close()
				return nil
			},
		},
	)
}

// DependsOn implements the glue.DependsOn interface.
func (b *Bundle) DependsOn() []string {
	return []string{gzviper.BundleName}
}
