// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Obtains information about all databases found in SQL Server instance.
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
//			example, err := mssql.GetDatabases(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("databases", example.Databases)
//			return nil
//		})
//	}
//
// ```
func GetDatabases(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesResult
	err := ctx.Invoke("mssql:index/getDatabases:getDatabases", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getDatabases.
type GetDatabasesResult struct {
	// Set of database objects
	Databases []GetDatabasesDatabase `pulumi:"databases"`
	// ID of the resource used only internally by the provider.
	Id string `pulumi:"id"`
}

func GetDatabasesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetDatabasesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetDatabasesResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv GetDatabasesResult
		secret, err := ctx.InvokePackageRaw("mssql:index/getDatabases:getDatabases", nil, &rv, "", opts...)
		if err != nil {
			return GetDatabasesResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(GetDatabasesResultOutput)
		if secret {
			return pulumi.ToSecret(output).(GetDatabasesResultOutput), nil
		}
		return output, nil
	}).(GetDatabasesResultOutput)
}

// A collection of values returned by getDatabases.
type GetDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesResult)(nil)).Elem()
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutput() GetDatabasesResultOutput {
	return o
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutputWithContext(ctx context.Context) GetDatabasesResultOutput {
	return o
}

// Set of database objects
func (o GetDatabasesResultOutput) Databases() GetDatabasesDatabaseArrayOutput {
	return o.ApplyT(func(v GetDatabasesResult) []GetDatabasesDatabase { return v.Databases }).(GetDatabasesDatabaseArrayOutput)
}

// ID of the resource used only internally by the provider.
func (o GetDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesResultOutput{})
}
