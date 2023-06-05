// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
//	"github.com/pulumi/pulumi-mssql/sdk/go/mssql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
	opts = pkgInvokeDefaultOpts(opts)
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
