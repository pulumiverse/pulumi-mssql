// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Obtains information about single database role.
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
//			master, err := mssql.LookupDatabase(ctx, &mssql.LookupDatabaseArgs{
//				Name: "master",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			example, err := mssql.LookupDatabaseRole(ctx, &mssql.LookupDatabaseRoleArgs{
//				Name:       "public",
//				DatabaseId: pulumi.StringRef(master.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("id", example.Id)
//			return nil
//		})
//	}
//
// ```
func LookupDatabaseRole(ctx *pulumi.Context, args *LookupDatabaseRoleArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseRoleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatabaseRoleResult
	err := ctx.Invoke("mssql:index/getDatabaseRole:getDatabaseRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseRole.
type LookupDatabaseRoleArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// Name of the database principal.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDatabaseRole.
type LookupDatabaseRoleResult struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// `<database_id>/<role_id>`. Role ID can be retrieved using `SELECT DATABASE_PRINCIPAL_ID('<role_name>')`
	Id string `pulumi:"id"`
	// Set of role members
	Members []GetDatabaseRoleMemberType `pulumi:"members"`
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name string `pulumi:"name"`
	// ID of another database role or user owning this role. Can be retrieved using `DatabaseRole` or `SqlUser`.
	OwnerId string `pulumi:"ownerId"`
}

func LookupDatabaseRoleOutput(ctx *pulumi.Context, args LookupDatabaseRoleOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseRoleResult, error) {
			args := v.(LookupDatabaseRoleArgs)
			r, err := LookupDatabaseRole(ctx, &args, opts...)
			var s LookupDatabaseRoleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatabaseRoleResultOutput)
}

// A collection of arguments for invoking getDatabaseRole.
type LookupDatabaseRoleOutputArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringPtrInput `pulumi:"databaseId"`
	// Name of the database principal.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDatabaseRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseRoleArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseRole.
type LookupDatabaseRoleResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseRoleResult)(nil)).Elem()
}

func (o LookupDatabaseRoleResultOutput) ToLookupDatabaseRoleResultOutput() LookupDatabaseRoleResultOutput {
	return o
}

func (o LookupDatabaseRoleResultOutput) ToLookupDatabaseRoleResultOutputWithContext(ctx context.Context) LookupDatabaseRoleResultOutput {
	return o
}

func (o LookupDatabaseRoleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDatabaseRoleResult] {
	return pulumix.Output[LookupDatabaseRoleResult]{
		OutputState: o.OutputState,
	}
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
func (o LookupDatabaseRoleResultOutput) DatabaseId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatabaseRoleResult) *string { return v.DatabaseId }).(pulumi.StringPtrOutput)
}

// `<database_id>/<role_id>`. Role ID can be retrieved using `SELECT DATABASE_PRINCIPAL_ID('<role_name>')`
func (o LookupDatabaseRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of role members
func (o LookupDatabaseRoleResultOutput) Members() GetDatabaseRoleMemberTypeArrayOutput {
	return o.ApplyT(func(v LookupDatabaseRoleResult) []GetDatabaseRoleMemberType { return v.Members }).(GetDatabaseRoleMemberTypeArrayOutput)
}

// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
func (o LookupDatabaseRoleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseRoleResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of another database role or user owning this role. Can be retrieved using `DatabaseRole` or `SqlUser`.
func (o LookupDatabaseRoleResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseRoleResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseRoleResultOutput{})
}
