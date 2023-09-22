// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mssql

import (
	_ "embed"
	"path/filepath"
	"strings"

	shimprovider "github.com/PGSSoft/terraform-provider-mssql/shim"
	"github.com/ettle/strcase"
	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumiverse/pulumi-mssql/provider/pkg/version"
)

//go:embed cmd/pulumi-resource-mssql/bridge-metadata.json
var bridgeMetadata []byte

// all of the token components used below.
const (
	// modules:
	mainMod = "index" // the mssql module
)

func convertName(name string) string {
	idx := strings.Index(name, "_")
	contract.Assertf(idx > 0 && idx < len(name)-1, "Invalid snake case name %s", name)
	name = name[idx+1:]
	contract.Assertf(len(name) > 0, "Invalid snake case name %s", name)
	return strcase.ToPascal(name)
}

func makeDataSource(mod string, name string) tokens.ModuleMember {
	name = convertName(name)
	return tfbridge.MakeDataSource("mssql", mod, "get"+name)
}

func makeResource(mod string, res string) tokens.Type {
	return tfbridge.MakeResource("mssql", mod, convertName(res))
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pf.ShimProvider(shimprovider.NewProvider())
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "mssql",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Microsoft SQL Server",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/pulumiverse/pulumi-mssql/main/docs/mssql.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/pulumiverse/pulumi-mssql",
		Description:       "A Pulumi Provider for Microsoft SQL Server and Azure SQL",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords: []string{
			"pulumi",
			"mssql",
			"category/database",
		},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/pulumiverse/pulumi-mssql",
		Repository: "https://github.com/pulumiverse/pulumi-mssql",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "PGSSoft",
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
		Config: map[string]*tfbridge.SchemaInfo{
			"hostname": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"MSSQL_HOSTNAME"},
				},
			},
			"port": {
				Default: &tfbridge.DefaultInfo{
					Value:   1433,
					EnvVars: []string{"MSSQL_PORT"},
				},
			},
			"sql_auth": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"username": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"MSSQL_USERNAME"},
						},
						MarkAsOptional: tfbridge.True(),
					},
					"password": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"MSSQL_PASSWORD"},
						},
						Secret:         tfbridge.True(),
						MarkAsOptional: tfbridge.True(),
					},
				},
				MarkAsOptional: tfbridge.True(),
			},
			"azure_auth": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"client_id": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"ARM_CLIENT_ID", "AZURE_CLIENT_ID"},
						},
						MarkAsOptional: tfbridge.True(),
					},
					"client_secret": {
						Secret: tfbridge.True(),
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"ARM_CLIENT_SECRET", "AZURE_CLIENT_SECRET"},
						},
						MarkAsOptional: tfbridge.True(),
					},
					"tenant_id": {
						Default: &tfbridge.DefaultInfo{
							EnvVars: []string{"ARM_TENANT_ID", "AZURE_TENANT_ID"},
						},
						MarkAsOptional: tfbridge.True(),
					},
				},
				MarkAsOptional: tfbridge.True(),
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"mssql_azuread_service_principal": {
				Tok: makeResource(mainMod, "mssql_azuread_service_principal"),
			},
			"mssql_azuread_user": {
				Tok: makeResource(mainMod, "mssql_azuread_user"),
			},
			"mssql_database": {
				Tok: makeResource(mainMod, "mssql_database"),
			},
			"mssql_database_permission": {
				Tok: makeResource(mainMod, "mssql_database_permission"),
			},
			"mssql_database_role": {
				Tok: makeResource(mainMod, "mssql_database_role"),
			},
			"mssql_database_role_member": {
				Tok: makeResource(mainMod, "mssql_database_role_member"),
			},
			"mssql_schema": {
				Tok: makeResource(mainMod, "mssql_schema"),
			},
			"mssql_schema_permission": {
				Tok: makeResource(mainMod, "mssql_schema_permission"),
			},
			"mssql_script": {
				Tok: makeResource(mainMod, "mssql_script"),
			},
			"mssql_server_permission": {
				Tok: makeResource(mainMod, "mssql_server_permission"),
			},
			"mssql_server_role": {
				Tok: makeResource(mainMod, "mssql_server_role"),
			},
			"mssql_server_role_member": {
				Tok: makeResource(mainMod, "mssql_server_role_member"),
			},
			"mssql_sql_login": {
				Tok: makeResource(mainMod, "mssql_sql_login"),
			},
			"mssql_sql_user": {
				Tok: makeResource(mainMod, "mssql_sql_user"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"mssql_azuread_service_principal": {
				Tok: makeDataSource(mainMod, "mssql_azuread_service_principal"),
			},
			"mssql_azuread_user": {
				Tok: makeDataSource(mainMod, "mssql_azuread_user"),
			},
			"mssql_database": {
				Tok: makeDataSource(mainMod, "mssql_database"),
			},
			"mssql_database_permissions": {
				Tok: makeDataSource(mainMod, "mssql_database_permissions"),
			},
			"mssql_database_role": {
				Tok: makeDataSource(mainMod, "mssql_database_role"),
			},
			"mssql_database_roles": {
				Tok: makeDataSource(mainMod, "mssql_database_roles"),
			},
			"mssql_databases": {
				Tok: makeDataSource(mainMod, "mssql_databases"),
			},
			"mssql_query": {
				Tok: makeDataSource(mainMod, "mssql_query"),
			},
			"mssql_schema": {
				Tok: makeDataSource(mainMod, "mssql_schema"),
			},
			"mssql_schema_permissions": {
				Tok: makeDataSource(mainMod, "mssql_schema_permissions"),
			},
			"mssql_schemas": {
				Tok: makeDataSource(mainMod, "mssql_schemas"),
			},
			"mssql_server_permissions": {
				Tok: makeDataSource(mainMod, "mssql_server_permissions"),
			},
			"mssql_server_role": {
				Tok: makeDataSource(mainMod, "mssql_server_role"),
			},
			"mssql_server_roles": {
				Tok: makeDataSource(mainMod, "mssql_server_roles"),
			},
			"mssql_sql_login": {
				Tok: makeDataSource(mainMod, "mssql_sql_login"),
			},
			"mssql_sql_logins": {
				Tok: makeDataSource(mainMod, "mssql_sql_logins"),
			},
			"mssql_sql_user": {
				Tok: makeDataSource(mainMod, "mssql_sql_user"),
			},
			"mssql_sql_users": {
				Tok: makeDataSource(mainMod, "mssql_sql_users"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/mssql",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_mssql",

			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				"github.com/pulumiverse/pulumi-mssql/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"mssql",
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumiverse",

			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumiverse",
		},

		Version: version.Version,
	}

	prov.SetAutonaming(255, "-")

	return prov
}
