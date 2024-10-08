// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Obtains information about single server role.
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
//			_, err := mssql.LookupServerRole(ctx, &mssql.LookupServerRoleArgs{
//				Name: pulumi.StringRef("example"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = mssql.LookupServerRole(ctx, &mssql.LookupServerRoleArgs{
//				Id: pulumi.StringRef("8"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupServerRole(ctx *pulumi.Context, args *LookupServerRoleArgs, opts ...pulumi.InvokeOption) (*LookupServerRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServerRoleResult
	err := ctx.Invoke("mssql:index/getServerRole:getServerRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerRole.
type LookupServerRoleArgs struct {
	// Role principal ID. Either `name` or `id` must be provided.
	Id *string `pulumi:"id"`
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars. Either `name` or `id` must be provided.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getServerRole.
type LookupServerRoleResult struct {
	// Role principal ID. Either `name` or `id` must be provided.
	Id string `pulumi:"id"`
	// Set of role members
	Members []GetServerRoleMemberType `pulumi:"members"`
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars. Either `name` or `id` must be provided.
	Name string `pulumi:"name"`
	// ID of another server role or login owning this role. Can be retrieved using `ServerRole` or `SqlLogin`.
	OwnerId string `pulumi:"ownerId"`
}

func LookupServerRoleOutput(ctx *pulumi.Context, args LookupServerRoleOutputArgs, opts ...pulumi.InvokeOption) LookupServerRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerRoleResultOutput, error) {
			args := v.(LookupServerRoleArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupServerRoleResult
			secret, err := ctx.InvokePackageRaw("mssql:index/getServerRole:getServerRole", args, &rv, "", opts...)
			if err != nil {
				return LookupServerRoleResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupServerRoleResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupServerRoleResultOutput), nil
			}
			return output, nil
		}).(LookupServerRoleResultOutput)
}

// A collection of arguments for invoking getServerRole.
type LookupServerRoleOutputArgs struct {
	// Role principal ID. Either `name` or `id` must be provided.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars. Either `name` or `id` must be provided.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupServerRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerRoleArgs)(nil)).Elem()
}

// A collection of values returned by getServerRole.
type LookupServerRoleResultOutput struct{ *pulumi.OutputState }

func (LookupServerRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerRoleResult)(nil)).Elem()
}

func (o LookupServerRoleResultOutput) ToLookupServerRoleResultOutput() LookupServerRoleResultOutput {
	return o
}

func (o LookupServerRoleResultOutput) ToLookupServerRoleResultOutputWithContext(ctx context.Context) LookupServerRoleResultOutput {
	return o
}

// Role principal ID. Either `name` or `id` must be provided.
func (o LookupServerRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of role members
func (o LookupServerRoleResultOutput) Members() GetServerRoleMemberTypeArrayOutput {
	return o.ApplyT(func(v LookupServerRoleResult) []GetServerRoleMemberType { return v.Members }).(GetServerRoleMemberTypeArrayOutput)
}

// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars. Either `name` or `id` must be provided.
func (o LookupServerRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of another server role or login owning this role. Can be retrieved using `ServerRole` or `SqlLogin`.
func (o LookupServerRoleResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerRoleResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerRoleResultOutput{})
}
