// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Managed database-level user mapped to Azure AD identity (service principal or managed identity).
//
// > **Note** When using this resource, Azure SQL server managed identity does not need any [AzureAD role assignments](https://docs.microsoft.com/en-us/azure/azure-sql/database/authentication-aad-service-principal?view=azuresql).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
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
//			exampleServicePrincipal, err := azuread.LookupServicePrincipal(ctx, &azuread.LookupServicePrincipalArgs{
//				DisplayName: pulumi.StringRef("test-application"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleAzureadServicePrincipal, err := mssql.NewAzureadServicePrincipal(ctx, "exampleAzureadServicePrincipal", &mssql.AzureadServicePrincipalArgs{
//				DatabaseId: pulumi.String(exampleDatabase.Id),
//				ClientId:   pulumi.String(exampleServicePrincipal.ApplicationId),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("userId", exampleAzureadServicePrincipal.ID())
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// import using <db_id>/<user_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', principal_id) FROM sys.database_principals WHERE [name] = '<username>'`
//
// ```sh
// $ pulumi import mssql:index/azureadServicePrincipal:AzureadServicePrincipal example '7/5'
// ```
type AzureadServicePrincipal struct {
	pulumi.CustomResourceState

	// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// User name. Cannot be longer than 128 chars.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAzureadServicePrincipal registers a new resource with the given unique name, arguments, and options.
func NewAzureadServicePrincipal(ctx *pulumi.Context,
	name string, args *AzureadServicePrincipalArgs, opts ...pulumi.ResourceOption) (*AzureadServicePrincipal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.DatabaseId == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AzureadServicePrincipal
	err := ctx.RegisterResource("mssql:index/azureadServicePrincipal:AzureadServicePrincipal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAzureadServicePrincipal gets an existing AzureadServicePrincipal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAzureadServicePrincipal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AzureadServicePrincipalState, opts ...pulumi.ResourceOption) (*AzureadServicePrincipal, error) {
	var resource AzureadServicePrincipal
	err := ctx.ReadResource("mssql:index/azureadServicePrincipal:AzureadServicePrincipal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AzureadServicePrincipal resources.
type azureadServicePrincipalState struct {
	// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
	ClientId *string `pulumi:"clientId"`
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId *string `pulumi:"databaseId"`
	// User name. Cannot be longer than 128 chars.
	Name *string `pulumi:"name"`
}

type AzureadServicePrincipalState struct {
	// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
	ClientId pulumi.StringPtrInput
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId pulumi.StringPtrInput
	// User name. Cannot be longer than 128 chars.
	Name pulumi.StringPtrInput
}

func (AzureadServicePrincipalState) ElementType() reflect.Type {
	return reflect.TypeOf((*azureadServicePrincipalState)(nil)).Elem()
}

type azureadServicePrincipalArgs struct {
	// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
	ClientId string `pulumi:"clientId"`
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId string `pulumi:"databaseId"`
	// User name. Cannot be longer than 128 chars.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AzureadServicePrincipal resource.
type AzureadServicePrincipalArgs struct {
	// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
	ClientId pulumi.StringInput
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
	DatabaseId pulumi.StringInput
	// User name. Cannot be longer than 128 chars.
	Name pulumi.StringPtrInput
}

func (AzureadServicePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*azureadServicePrincipalArgs)(nil)).Elem()
}

type AzureadServicePrincipalInput interface {
	pulumi.Input

	ToAzureadServicePrincipalOutput() AzureadServicePrincipalOutput
	ToAzureadServicePrincipalOutputWithContext(ctx context.Context) AzureadServicePrincipalOutput
}

func (*AzureadServicePrincipal) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureadServicePrincipal)(nil)).Elem()
}

func (i *AzureadServicePrincipal) ToAzureadServicePrincipalOutput() AzureadServicePrincipalOutput {
	return i.ToAzureadServicePrincipalOutputWithContext(context.Background())
}

func (i *AzureadServicePrincipal) ToAzureadServicePrincipalOutputWithContext(ctx context.Context) AzureadServicePrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureadServicePrincipalOutput)
}

// AzureadServicePrincipalArrayInput is an input type that accepts AzureadServicePrincipalArray and AzureadServicePrincipalArrayOutput values.
// You can construct a concrete instance of `AzureadServicePrincipalArrayInput` via:
//
//	AzureadServicePrincipalArray{ AzureadServicePrincipalArgs{...} }
type AzureadServicePrincipalArrayInput interface {
	pulumi.Input

	ToAzureadServicePrincipalArrayOutput() AzureadServicePrincipalArrayOutput
	ToAzureadServicePrincipalArrayOutputWithContext(context.Context) AzureadServicePrincipalArrayOutput
}

type AzureadServicePrincipalArray []AzureadServicePrincipalInput

func (AzureadServicePrincipalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureadServicePrincipal)(nil)).Elem()
}

func (i AzureadServicePrincipalArray) ToAzureadServicePrincipalArrayOutput() AzureadServicePrincipalArrayOutput {
	return i.ToAzureadServicePrincipalArrayOutputWithContext(context.Background())
}

func (i AzureadServicePrincipalArray) ToAzureadServicePrincipalArrayOutputWithContext(ctx context.Context) AzureadServicePrincipalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureadServicePrincipalArrayOutput)
}

// AzureadServicePrincipalMapInput is an input type that accepts AzureadServicePrincipalMap and AzureadServicePrincipalMapOutput values.
// You can construct a concrete instance of `AzureadServicePrincipalMapInput` via:
//
//	AzureadServicePrincipalMap{ "key": AzureadServicePrincipalArgs{...} }
type AzureadServicePrincipalMapInput interface {
	pulumi.Input

	ToAzureadServicePrincipalMapOutput() AzureadServicePrincipalMapOutput
	ToAzureadServicePrincipalMapOutputWithContext(context.Context) AzureadServicePrincipalMapOutput
}

type AzureadServicePrincipalMap map[string]AzureadServicePrincipalInput

func (AzureadServicePrincipalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureadServicePrincipal)(nil)).Elem()
}

func (i AzureadServicePrincipalMap) ToAzureadServicePrincipalMapOutput() AzureadServicePrincipalMapOutput {
	return i.ToAzureadServicePrincipalMapOutputWithContext(context.Background())
}

func (i AzureadServicePrincipalMap) ToAzureadServicePrincipalMapOutputWithContext(ctx context.Context) AzureadServicePrincipalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureadServicePrincipalMapOutput)
}

type AzureadServicePrincipalOutput struct{ *pulumi.OutputState }

func (AzureadServicePrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureadServicePrincipal)(nil)).Elem()
}

func (o AzureadServicePrincipalOutput) ToAzureadServicePrincipalOutput() AzureadServicePrincipalOutput {
	return o
}

func (o AzureadServicePrincipalOutput) ToAzureadServicePrincipalOutputWithContext(ctx context.Context) AzureadServicePrincipalOutput {
	return o
}

// Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
func (o AzureadServicePrincipalOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureadServicePrincipal) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
func (o AzureadServicePrincipalOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureadServicePrincipal) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// User name. Cannot be longer than 128 chars.
func (o AzureadServicePrincipalOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AzureadServicePrincipal) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type AzureadServicePrincipalArrayOutput struct{ *pulumi.OutputState }

func (AzureadServicePrincipalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AzureadServicePrincipal)(nil)).Elem()
}

func (o AzureadServicePrincipalArrayOutput) ToAzureadServicePrincipalArrayOutput() AzureadServicePrincipalArrayOutput {
	return o
}

func (o AzureadServicePrincipalArrayOutput) ToAzureadServicePrincipalArrayOutputWithContext(ctx context.Context) AzureadServicePrincipalArrayOutput {
	return o
}

func (o AzureadServicePrincipalArrayOutput) Index(i pulumi.IntInput) AzureadServicePrincipalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AzureadServicePrincipal {
		return vs[0].([]*AzureadServicePrincipal)[vs[1].(int)]
	}).(AzureadServicePrincipalOutput)
}

type AzureadServicePrincipalMapOutput struct{ *pulumi.OutputState }

func (AzureadServicePrincipalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AzureadServicePrincipal)(nil)).Elem()
}

func (o AzureadServicePrincipalMapOutput) ToAzureadServicePrincipalMapOutput() AzureadServicePrincipalMapOutput {
	return o
}

func (o AzureadServicePrincipalMapOutput) ToAzureadServicePrincipalMapOutputWithContext(ctx context.Context) AzureadServicePrincipalMapOutput {
	return o
}

func (o AzureadServicePrincipalMapOutput) MapIndex(k pulumi.StringInput) AzureadServicePrincipalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AzureadServicePrincipal {
		return vs[0].(map[string]*AzureadServicePrincipal)[vs[1].(string)]
	}).(AzureadServicePrincipalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureadServicePrincipalInput)(nil)).Elem(), &AzureadServicePrincipal{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureadServicePrincipalArrayInput)(nil)).Elem(), AzureadServicePrincipalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureadServicePrincipalMapInput)(nil)).Elem(), AzureadServicePrincipalMap{})
	pulumi.RegisterOutputType(AzureadServicePrincipalOutput{})
	pulumi.RegisterOutputType(AzureadServicePrincipalArrayOutput{})
	pulumi.RegisterOutputType(AzureadServicePrincipalMapOutput{})
}
