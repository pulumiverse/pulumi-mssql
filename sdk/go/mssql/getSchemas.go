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

// Obtains information about all schemas found in SQL database.
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
//			example, err := mssql.LookupDatabase(ctx, &mssql.LookupDatabaseArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			all, err := mssql.GetSchemas(ctx, &mssql.GetSchemasArgs{
//				DatabaseId: pulumi.StringRef(example.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			var splat0 []*string
//			for _, val0 := range all.Schemas {
//				splat0 = append(splat0, val0.Name)
//			}
//			ctx.Export("allSchemaNames", splat0)
//			return nil
//		})
//	}
//
// ```
func GetSchemas(ctx *pulumi.Context, args *GetSchemasArgs, opts ...pulumi.InvokeOption) (*GetSchemasResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSchemasResult
	err := ctx.Invoke("mssql:index/getSchemas:getSchemas", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSchemas.
type GetSchemasArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId *string `pulumi:"databaseId"`
}

// A collection of values returned by getSchemas.
type GetSchemasResult struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId string `pulumi:"databaseId"`
	// ID of the data source, equals to database ID
	Id string `pulumi:"id"`
	// Set of schemas found in the DB.
	Schemas []GetSchemasSchema `pulumi:"schemas"`
}

func GetSchemasOutput(ctx *pulumi.Context, args GetSchemasOutputArgs, opts ...pulumi.InvokeOption) GetSchemasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSchemasResult, error) {
			args := v.(GetSchemasArgs)
			r, err := GetSchemas(ctx, &args, opts...)
			var s GetSchemasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSchemasResultOutput)
}

// A collection of arguments for invoking getSchemas.
type GetSchemasOutputArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId pulumi.StringPtrInput `pulumi:"databaseId"`
}

func (GetSchemasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSchemasArgs)(nil)).Elem()
}

// A collection of values returned by getSchemas.
type GetSchemasResultOutput struct{ *pulumi.OutputState }

func (GetSchemasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSchemasResult)(nil)).Elem()
}

func (o GetSchemasResultOutput) ToGetSchemasResultOutput() GetSchemasResultOutput {
	return o
}

func (o GetSchemasResultOutput) ToGetSchemasResultOutputWithContext(ctx context.Context) GetSchemasResultOutput {
	return o
}

func (o GetSchemasResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSchemasResult] {
	return pulumix.Output[GetSchemasResult]{
		OutputState: o.OutputState,
	}
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
func (o GetSchemasResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSchemasResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

// ID of the data source, equals to database ID
func (o GetSchemasResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSchemasResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of schemas found in the DB.
func (o GetSchemasResultOutput) Schemas() GetSchemasSchemaArrayOutput {
	return o.ApplyT(func(v GetSchemasResult) []GetSchemasSchema { return v.Schemas }).(GetSchemasSchemaArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSchemasResultOutput{})
}
