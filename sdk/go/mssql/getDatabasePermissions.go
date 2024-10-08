// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Returns all permissions granted in a DB to given principal
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDatabase, err := mssql.LookupDatabase(ctx, &mssql.LookupDatabaseArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleSqlUser, err := mssql.LookupSqlUser(ctx, &mssql.LookupSqlUserArgs{
//				Name:       "example_user",
//				DatabaseId: pulumi.StringRef(exampleDatabase.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleDatabasePermissions, err := mssql.GetDatabasePermissions(ctx, &mssql.GetDatabasePermissionsArgs{
//				PrincipalId: exampleSqlUser.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("permissions", exampleDatabasePermissions.Permissions)
//			return nil
//		})
//	}
//
// ```
func GetDatabasePermissions(ctx *pulumi.Context, args *GetDatabasePermissionsArgs, opts ...pulumi.InvokeOption) (*GetDatabasePermissionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasePermissionsResult
	err := ctx.Invoke("mssql:index/getDatabasePermissions:getDatabasePermissions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabasePermissions.
type GetDatabasePermissionsArgs struct {
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId string `pulumi:"principalId"`
}

// A collection of values returned by getDatabasePermissions.
type GetDatabasePermissionsResult struct {
	// `<database_id>/<principal_id>`.
	Id string `pulumi:"id"`
	// Set of permissions granted to the principal
	Permissions []GetDatabasePermissionsPermission `pulumi:"permissions"`
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId string `pulumi:"principalId"`
}

func GetDatabasePermissionsOutput(ctx *pulumi.Context, args GetDatabasePermissionsOutputArgs, opts ...pulumi.InvokeOption) GetDatabasePermissionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabasePermissionsResultOutput, error) {
			args := v.(GetDatabasePermissionsArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetDatabasePermissionsResult
			secret, err := ctx.InvokePackageRaw("mssql:index/getDatabasePermissions:getDatabasePermissions", args, &rv, "", opts...)
			if err != nil {
				return GetDatabasePermissionsResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetDatabasePermissionsResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetDatabasePermissionsResultOutput), nil
			}
			return output, nil
		}).(GetDatabasePermissionsResultOutput)
}

// A collection of arguments for invoking getDatabasePermissions.
type GetDatabasePermissionsOutputArgs struct {
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (GetDatabasePermissionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasePermissionsArgs)(nil)).Elem()
}

// A collection of values returned by getDatabasePermissions.
type GetDatabasePermissionsResultOutput struct{ *pulumi.OutputState }

func (GetDatabasePermissionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasePermissionsResult)(nil)).Elem()
}

func (o GetDatabasePermissionsResultOutput) ToGetDatabasePermissionsResultOutput() GetDatabasePermissionsResultOutput {
	return o
}

func (o GetDatabasePermissionsResultOutput) ToGetDatabasePermissionsResultOutputWithContext(ctx context.Context) GetDatabasePermissionsResultOutput {
	return o
}

// `<database_id>/<principal_id>`.
func (o GetDatabasePermissionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasePermissionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of permissions granted to the principal
func (o GetDatabasePermissionsResultOutput) Permissions() GetDatabasePermissionsPermissionArrayOutput {
	return o.ApplyT(func(v GetDatabasePermissionsResult) []GetDatabasePermissionsPermission { return v.Permissions }).(GetDatabasePermissionsPermissionArrayOutput)
}

// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
func (o GetDatabasePermissionsResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasePermissionsResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasePermissionsResultOutput{})
}
