package shim

import (
	"github.com/PGSSoft/terraform-provider-mssql/internal/provider"
	tf "github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProvider() tf.Provider {
	return provider.New(provider.VersionDev)()
}
