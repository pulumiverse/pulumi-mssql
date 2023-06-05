// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Obtains information about all roles defined in a database.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-mssql/sdk/go/mssql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			master, err := mssql.LookupDatabase(ctx, &mssql.LookupDatabaseArgs{
//				Name: "master",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			example, err := mssql.GetDatabaseRoles(ctx, &mssql.GetDatabaseRolesArgs{
//				DatabaseId: pulumi.StringRef(master.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("roles", example.Roles)
//			return nil
//		})
//	}
//
// ```
func GetDatabaseRoles(ctx *pulumi.Context, args *GetDatabaseRolesArgs, opts ...pulumi.InvokeOption) (*GetDatabaseRolesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatabaseRolesResult
	err := ctx.Invoke("mssql:index/getDatabaseRoles:getDatabaseRoles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseRoles.
type GetDatabaseRolesArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
}

// A collection of values returned by getDatabaseRoles.
type GetDatabaseRolesResult struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// ID of the resource, equals to database ID
	Id string `pulumi:"id"`
	// Set of database role objects
	Roles []GetDatabaseRolesRole `pulumi:"roles"`
}

func GetDatabaseRolesOutput(ctx *pulumi.Context, args GetDatabaseRolesOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseRolesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseRolesResult, error) {
			args := v.(GetDatabaseRolesArgs)
			r, err := GetDatabaseRoles(ctx, &args, opts...)
			var s GetDatabaseRolesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabaseRolesResultOutput)
}

// A collection of arguments for invoking getDatabaseRoles.
type GetDatabaseRolesOutputArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringPtrInput `pulumi:"databaseId"`
}

func (GetDatabaseRolesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseRolesArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseRoles.
type GetDatabaseRolesResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseRolesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseRolesResult)(nil)).Elem()
}

func (o GetDatabaseRolesResultOutput) ToGetDatabaseRolesResultOutput() GetDatabaseRolesResultOutput {
	return o
}

func (o GetDatabaseRolesResultOutput) ToGetDatabaseRolesResultOutputWithContext(ctx context.Context) GetDatabaseRolesResultOutput {
	return o
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
func (o GetDatabaseRolesResultOutput) DatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseRolesResult) *string { return v.DatabaseId }).(pulumi.StringPtrOutput)
}

// ID of the resource, equals to database ID
func (o GetDatabaseRolesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseRolesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of database role objects
func (o GetDatabaseRolesResultOutput) Roles() GetDatabaseRolesRoleArrayOutput {
	return o.ApplyT(func(v GetDatabaseRolesResult) []GetDatabaseRolesRole { return v.Roles }).(GetDatabaseRolesRoleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseRolesResultOutput{})
}
