// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Obtains information about single Azure AD Service Principal database user.
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
//			exampleAzureadServicePrincipal, err := mssql.LookupAzureadServicePrincipal(ctx, &mssql.LookupAzureadServicePrincipalArgs{
//				Name:       pulumi.StringRef("example"),
//				DatabaseId: exampleDatabase.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("appClientId", exampleAzureadServicePrincipal.ClientId)
//			return nil
//		})
//	}
//
// ```
func LookupAzureadServicePrincipal(ctx *pulumi.Context, args *LookupAzureadServicePrincipalArgs, opts ...pulumi.InvokeOption) (*LookupAzureadServicePrincipalResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAzureadServicePrincipalResult
	err := ctx.Invoke("mssql:index/getAzureadServicePrincipal:getAzureadServicePrincipal", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAzureadServicePrincipal.
type LookupAzureadServicePrincipalArgs struct {
	// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
	ClientId *string `pulumi:"clientId"`
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId string `pulumi:"databaseId"`
	// User name. Cannot be longer than 128 chars.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAzureadServicePrincipal.
type LookupAzureadServicePrincipalResult struct {
	// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
	ClientId string `pulumi:"clientId"`
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId string `pulumi:"databaseId"`
	// `<database_id>/<user_id>`. User ID can be retrieved using `sys.database_principals` view.
	Id string `pulumi:"id"`
	// User name. Cannot be longer than 128 chars.
	Name string `pulumi:"name"`
}

func LookupAzureadServicePrincipalOutput(ctx *pulumi.Context, args LookupAzureadServicePrincipalOutputArgs, opts ...pulumi.InvokeOption) LookupAzureadServicePrincipalResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAzureadServicePrincipalResult, error) {
			args := v.(LookupAzureadServicePrincipalArgs)
			r, err := LookupAzureadServicePrincipal(ctx, &args, opts...)
			var s LookupAzureadServicePrincipalResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAzureadServicePrincipalResultOutput)
}

// A collection of arguments for invoking getAzureadServicePrincipal.
type LookupAzureadServicePrincipalOutputArgs struct {
	// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId pulumi.StringInput `pulumi:"databaseId"`
	// User name. Cannot be longer than 128 chars.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAzureadServicePrincipalOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureadServicePrincipalArgs)(nil)).Elem()
}

// A collection of values returned by getAzureadServicePrincipal.
type LookupAzureadServicePrincipalResultOutput struct{ *pulumi.OutputState }

func (LookupAzureadServicePrincipalResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAzureadServicePrincipalResult)(nil)).Elem()
}

func (o LookupAzureadServicePrincipalResultOutput) ToLookupAzureadServicePrincipalResultOutput() LookupAzureadServicePrincipalResultOutput {
	return o
}

func (o LookupAzureadServicePrincipalResultOutput) ToLookupAzureadServicePrincipalResultOutputWithContext(ctx context.Context) LookupAzureadServicePrincipalResultOutput {
	return o
}

// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
func (o LookupAzureadServicePrincipalResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureadServicePrincipalResult) string { return v.ClientId }).(pulumi.StringOutput)
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
func (o LookupAzureadServicePrincipalResultOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureadServicePrincipalResult) string { return v.DatabaseId }).(pulumi.StringOutput)
}

// `<database_id>/<user_id>`. User ID can be retrieved using `sys.database_principals` view.
func (o LookupAzureadServicePrincipalResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureadServicePrincipalResult) string { return v.Id }).(pulumi.StringOutput)
}

// User name. Cannot be longer than 128 chars.
func (o LookupAzureadServicePrincipalResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAzureadServicePrincipalResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAzureadServicePrincipalResultOutput{})
}
