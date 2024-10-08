// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Manages single DB schema.
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
//			owner, err := mssql.LookupSqlUser(ctx, &mssql.LookupSqlUserArgs{
//				Name: "example_user",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = mssql.NewSchema(ctx, "exampleSchema", &mssql.SchemaArgs{
//				DatabaseId: pulumi.String(exampleDatabase.Id),
//				OwnerId:    pulumi.String(owner.Id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// import using <db_id>/<schema_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', SCHEMA_ID('<schema_name>'))`
//
// ```sh
// $ pulumi import mssql:index/schema:Schema example '7/5'
// ```
type Schema struct {
	pulumi.CustomResourceState

	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// Schema name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of database role or user owning this schema. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
}

// NewSchema registers a new resource with the given unique name, arguments, and options.
func NewSchema(ctx *pulumi.Context,
	name string, args *SchemaArgs, opts ...pulumi.ResourceOption) (*Schema, error) {
	if args == nil {
		args = &SchemaArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Schema
	err := ctx.RegisterResource("mssql:index/schema:Schema", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchema gets an existing Schema resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchema(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaState, opts ...pulumi.ResourceOption) (*Schema, error) {
	var resource Schema
	err := ctx.ReadResource("mssql:index/schema:Schema", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schema resources.
type schemaState struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// Schema name.
	Name *string `pulumi:"name"`
	// ID of database role or user owning this schema. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`
	OwnerId *string `pulumi:"ownerId"`
}

type SchemaState struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringPtrInput
	// Schema name.
	Name pulumi.StringPtrInput
	// ID of database role or user owning this schema. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`
	OwnerId pulumi.StringPtrInput
}

func (SchemaState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaState)(nil)).Elem()
}

type schemaArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// Schema name.
	Name *string `pulumi:"name"`
	// ID of database role or user owning this schema. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`
	OwnerId *string `pulumi:"ownerId"`
}

// The set of arguments for constructing a Schema resource.
type SchemaArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringPtrInput
	// Schema name.
	Name pulumi.StringPtrInput
	// ID of database role or user owning this schema. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`
	OwnerId pulumi.StringPtrInput
}

func (SchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaArgs)(nil)).Elem()
}

type SchemaInput interface {
	pulumi.Input

	ToSchemaOutput() SchemaOutput
	ToSchemaOutputWithContext(ctx context.Context) SchemaOutput
}

func (*Schema) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (i *Schema) ToSchemaOutput() SchemaOutput {
	return i.ToSchemaOutputWithContext(context.Background())
}

func (i *Schema) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaOutput)
}

// SchemaArrayInput is an input type that accepts SchemaArray and SchemaArrayOutput values.
// You can construct a concrete instance of `SchemaArrayInput` via:
//
//	SchemaArray{ SchemaArgs{...} }
type SchemaArrayInput interface {
	pulumi.Input

	ToSchemaArrayOutput() SchemaArrayOutput
	ToSchemaArrayOutputWithContext(context.Context) SchemaArrayOutput
}

type SchemaArray []SchemaInput

func (SchemaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schema)(nil)).Elem()
}

func (i SchemaArray) ToSchemaArrayOutput() SchemaArrayOutput {
	return i.ToSchemaArrayOutputWithContext(context.Background())
}

func (i SchemaArray) ToSchemaArrayOutputWithContext(ctx context.Context) SchemaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaArrayOutput)
}

// SchemaMapInput is an input type that accepts SchemaMap and SchemaMapOutput values.
// You can construct a concrete instance of `SchemaMapInput` via:
//
//	SchemaMap{ "key": SchemaArgs{...} }
type SchemaMapInput interface {
	pulumi.Input

	ToSchemaMapOutput() SchemaMapOutput
	ToSchemaMapOutputWithContext(context.Context) SchemaMapOutput
}

type SchemaMap map[string]SchemaInput

func (SchemaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schema)(nil)).Elem()
}

func (i SchemaMap) ToSchemaMapOutput() SchemaMapOutput {
	return i.ToSchemaMapOutputWithContext(context.Background())
}

func (i SchemaMap) ToSchemaMapOutputWithContext(ctx context.Context) SchemaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaMapOutput)
}

type SchemaOutput struct{ *pulumi.OutputState }

func (SchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schema)(nil)).Elem()
}

func (o SchemaOutput) ToSchemaOutput() SchemaOutput {
	return o
}

func (o SchemaOutput) ToSchemaOutputWithContext(ctx context.Context) SchemaOutput {
	return o
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
func (o SchemaOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// Schema name.
func (o SchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of database role or user owning this schema. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`
func (o SchemaOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Schema) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

type SchemaArrayOutput struct{ *pulumi.OutputState }

func (SchemaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Schema)(nil)).Elem()
}

func (o SchemaArrayOutput) ToSchemaArrayOutput() SchemaArrayOutput {
	return o
}

func (o SchemaArrayOutput) ToSchemaArrayOutputWithContext(ctx context.Context) SchemaArrayOutput {
	return o
}

func (o SchemaArrayOutput) Index(i pulumi.IntInput) SchemaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Schema {
		return vs[0].([]*Schema)[vs[1].(int)]
	}).(SchemaOutput)
}

type SchemaMapOutput struct{ *pulumi.OutputState }

func (SchemaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Schema)(nil)).Elem()
}

func (o SchemaMapOutput) ToSchemaMapOutput() SchemaMapOutput {
	return o
}

func (o SchemaMapOutput) ToSchemaMapOutputWithContext(ctx context.Context) SchemaMapOutput {
	return o
}

func (o SchemaMapOutput) MapIndex(k pulumi.StringInput) SchemaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Schema {
		return vs[0].(map[string]*Schema)[vs[1].(string)]
	}).(SchemaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaInput)(nil)).Elem(), &Schema{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaArrayInput)(nil)).Elem(), SchemaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaMapInput)(nil)).Elem(), SchemaMap{})
	pulumi.RegisterOutputType(SchemaOutput{})
	pulumi.RegisterOutputType(SchemaArrayOutput{})
	pulumi.RegisterOutputType(SchemaMapOutput{})
}
