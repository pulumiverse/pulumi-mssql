// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "mssql:index/azureadServicePrincipal:AzureadServicePrincipal":
		r = &AzureadServicePrincipal{}
	case "mssql:index/azureadUser:AzureadUser":
		r = &AzureadUser{}
	case "mssql:index/database:Database":
		r = &Database{}
	case "mssql:index/databasePermission:DatabasePermission":
		r = &DatabasePermission{}
	case "mssql:index/databaseRole:DatabaseRole":
		r = &DatabaseRole{}
	case "mssql:index/databaseRoleMember:DatabaseRoleMember":
		r = &DatabaseRoleMember{}
	case "mssql:index/schema:Schema":
		r = &Schema{}
	case "mssql:index/schemaPermission:SchemaPermission":
		r = &SchemaPermission{}
	case "mssql:index/script:Script":
		r = &Script{}
	case "mssql:index/serverPermission:ServerPermission":
		r = &ServerPermission{}
	case "mssql:index/serverRole:ServerRole":
		r = &ServerRole{}
	case "mssql:index/serverRoleMember:ServerRoleMember":
		r = &ServerRoleMember{}
	case "mssql:index/sqlLogin:SqlLogin":
		r = &SqlLogin{}
	case "mssql:index/sqlUser:SqlUser":
		r = &SqlUser{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:mssql" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"mssql",
		"index/azureadServicePrincipal",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/azureadUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/databasePermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/databaseRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/databaseRoleMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/schema",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/schemaPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/script",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/serverPermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/serverRole",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/serverRoleMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/sqlLogin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"mssql",
		"index/sqlUser",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"mssql",
		&pkg{version},
	)
}
