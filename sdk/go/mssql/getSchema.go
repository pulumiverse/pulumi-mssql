// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves information about DB schema.
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
//			example, err := mssql.LookupDatabase(ctx, &mssql.LookupDatabaseArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = mssql.LookupSchema(ctx, &mssql.LookupSchemaArgs{
//				DatabaseId: pulumi.StringRef(example.Id),
//				Name:       pulumi.StringRef("dbo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSchema(ctx *pulumi.Context, args *LookupSchemaArgs, opts ...pulumi.InvokeOption) (*LookupSchemaResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSchemaResult
	err := ctx.Invoke("mssql:index/getSchema:getSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSchema.
type LookupSchemaArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId *string `pulumi:"databaseId"`
	// `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`. Either `id` or `name` must be provided.
	Id *string `pulumi:"id"`
	// Schema name. Either `id` or `name` must be provided.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSchema.
type LookupSchemaResult struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId string `pulumi:"databaseId"`
	// `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`. Either `id` or `name` must be provided.
	Id string `pulumi:"id"`
	// Schema name. Either `id` or `name` must be provided.
	Name string `pulumi:"name"`
	// ID of database role or user owning this schema. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`
	OwnerId string `pulumi:"ownerId"`
}

func LookupSchemaOutput(ctx *pulumi.Context, args LookupSchemaOutputArgs, opts ...pulumi.InvokeOption) LookupSchemaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSchemaResult, error) {
			args := v.(LookupSchemaArgs)
			r, err := LookupSchema(ctx, &args, opts...)
			var s LookupSchemaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSchemaResultOutput)
}

// A collection of arguments for invoking getSchema.
type LookupSchemaOutputArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId pulumi.StringPtrInput `pulumi:"databaseId"`
	// `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`. Either `id` or `name` must be provided.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Schema name. Either `id` or `name` must be provided.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaArgs)(nil)).Elem()
}

// A collection of values returned by getSchema.
type LookupSchemaResultOutput struct{ *pulumi.OutputState }

func (LookupSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchemaResult)(nil)).Elem()
}

func (o LookupSchemaResultOutput) ToLookupSchemaResultOutput() LookupSchemaResultOutput {
	return o
}

func (o LookupSchemaResultOutput) ToLookupSchemaResultOutputWithContext(ctx context.Context) LookupSchemaResultOutput {
	return o
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
func (o LookupSchemaResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

// `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`. Either `id` or `name` must be provided.
func (o LookupSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

// Schema name. Either `id` or `name` must be provided.
func (o LookupSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of database role or user owning this schema. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`
func (o LookupSchemaResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSchemaResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchemaResultOutput{})
}
