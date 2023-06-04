// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Grants database-level permission.
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
//			exampleDatabase, err := mssql.LookupDatabase(ctx, &mssql.LookupDatabaseArgs{
//				Name: "example",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleSqlUser, err := mssql.LookupSqlUser(ctx, &mssql.LookupSqlUserArgs{
//				Name:       "example_user",
//				DatabaseId: pulumi.StringRef(exampleDatabase.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleSchema, err := mssql.LookupSchema(ctx, &mssql.LookupSchemaArgs{
//				Name:       pulumi.StringRef("example_schema"),
//				DatabaseId: pulumi.StringRef(exampleDatabase.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = mssql.NewSchemaPermission(ctx, "deleteToExample", &mssql.SchemaPermissionArgs{
//				SchemaId:    *pulumi.String(exampleSchema.Id),
//				PrincipalId: *pulumi.String(exampleSqlUser.Id),
//				Permission:  pulumi.String("DELETE"),
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
// import using <db_id>/<schema_id>/<principal_id>/<permission> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', SCHEMA_ID('<schema_name>'), '/', DATABASE_PRINCIPAL_ID('<principal_name>'), '/DELETE')`
//
// ```sh
//
//	$ pulumi import mssql:index/schemaPermission:SchemaPermission example '7/5/8/DELETE'
//
// ```
type SchemaPermission struct {
	pulumi.CustomResourceState

	// Name of schema SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-schema-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringOutput `pulumi:"permission"`
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// `<database_id>/<schema_id>`. Can be retrieved using `Schema`.
	SchemaId pulumi.StringOutput `pulumi:"schemaId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption pulumi.BoolOutput `pulumi:"withGrantOption"`
}

// NewSchemaPermission registers a new resource with the given unique name, arguments, and options.
func NewSchemaPermission(ctx *pulumi.Context,
	name string, args *SchemaPermissionArgs, opts ...pulumi.ResourceOption) (*SchemaPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	if args.SchemaId == nil {
		return nil, errors.New("invalid value for required argument 'SchemaId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SchemaPermission
	err := ctx.RegisterResource("mssql:index/schemaPermission:SchemaPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchemaPermission gets an existing SchemaPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchemaPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SchemaPermissionState, opts ...pulumi.ResourceOption) (*SchemaPermission, error) {
	var resource SchemaPermission
	err := ctx.ReadResource("mssql:index/schemaPermission:SchemaPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SchemaPermission resources.
type schemaPermissionState struct {
	// Name of schema SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-schema-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission *string `pulumi:"permission"`
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId *string `pulumi:"principalId"`
	// `<database_id>/<schema_id>`. Can be retrieved using `Schema`.
	SchemaId *string `pulumi:"schemaId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

type SchemaPermissionState struct {
	// Name of schema SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-schema-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringPtrInput
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId pulumi.StringPtrInput
	// `<database_id>/<schema_id>`. Can be retrieved using `Schema`.
	SchemaId pulumi.StringPtrInput
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption pulumi.BoolPtrInput
}

func (SchemaPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaPermissionState)(nil)).Elem()
}

type schemaPermissionArgs struct {
	// Name of schema SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-schema-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission string `pulumi:"permission"`
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId string `pulumi:"principalId"`
	// `<database_id>/<schema_id>`. Can be retrieved using `Schema`.
	SchemaId string `pulumi:"schemaId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

// The set of arguments for constructing a SchemaPermission resource.
type SchemaPermissionArgs struct {
	// Name of schema SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-schema-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringInput
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId pulumi.StringInput
	// `<database_id>/<schema_id>`. Can be retrieved using `Schema`.
	SchemaId pulumi.StringInput
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption pulumi.BoolPtrInput
}

func (SchemaPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*schemaPermissionArgs)(nil)).Elem()
}

type SchemaPermissionInput interface {
	pulumi.Input

	ToSchemaPermissionOutput() SchemaPermissionOutput
	ToSchemaPermissionOutputWithContext(ctx context.Context) SchemaPermissionOutput
}

func (*SchemaPermission) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaPermission)(nil)).Elem()
}

func (i *SchemaPermission) ToSchemaPermissionOutput() SchemaPermissionOutput {
	return i.ToSchemaPermissionOutputWithContext(context.Background())
}

func (i *SchemaPermission) ToSchemaPermissionOutputWithContext(ctx context.Context) SchemaPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPermissionOutput)
}

// SchemaPermissionArrayInput is an input type that accepts SchemaPermissionArray and SchemaPermissionArrayOutput values.
// You can construct a concrete instance of `SchemaPermissionArrayInput` via:
//
//	SchemaPermissionArray{ SchemaPermissionArgs{...} }
type SchemaPermissionArrayInput interface {
	pulumi.Input

	ToSchemaPermissionArrayOutput() SchemaPermissionArrayOutput
	ToSchemaPermissionArrayOutputWithContext(context.Context) SchemaPermissionArrayOutput
}

type SchemaPermissionArray []SchemaPermissionInput

func (SchemaPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SchemaPermission)(nil)).Elem()
}

func (i SchemaPermissionArray) ToSchemaPermissionArrayOutput() SchemaPermissionArrayOutput {
	return i.ToSchemaPermissionArrayOutputWithContext(context.Background())
}

func (i SchemaPermissionArray) ToSchemaPermissionArrayOutputWithContext(ctx context.Context) SchemaPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPermissionArrayOutput)
}

// SchemaPermissionMapInput is an input type that accepts SchemaPermissionMap and SchemaPermissionMapOutput values.
// You can construct a concrete instance of `SchemaPermissionMapInput` via:
//
//	SchemaPermissionMap{ "key": SchemaPermissionArgs{...} }
type SchemaPermissionMapInput interface {
	pulumi.Input

	ToSchemaPermissionMapOutput() SchemaPermissionMapOutput
	ToSchemaPermissionMapOutputWithContext(context.Context) SchemaPermissionMapOutput
}

type SchemaPermissionMap map[string]SchemaPermissionInput

func (SchemaPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SchemaPermission)(nil)).Elem()
}

func (i SchemaPermissionMap) ToSchemaPermissionMapOutput() SchemaPermissionMapOutput {
	return i.ToSchemaPermissionMapOutputWithContext(context.Background())
}

func (i SchemaPermissionMap) ToSchemaPermissionMapOutputWithContext(ctx context.Context) SchemaPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchemaPermissionMapOutput)
}

type SchemaPermissionOutput struct{ *pulumi.OutputState }

func (SchemaPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaPermission)(nil)).Elem()
}

func (o SchemaPermissionOutput) ToSchemaPermissionOutput() SchemaPermissionOutput {
	return o
}

func (o SchemaPermissionOutput) ToSchemaPermissionOutputWithContext(ctx context.Context) SchemaPermissionOutput {
	return o
}

// Name of schema SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-schema-permissions-transact-sql?view=azuresqldb-current#remarks)
func (o SchemaPermissionOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaPermission) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
func (o SchemaPermissionOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaPermission) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// `<database_id>/<schema_id>`. Can be retrieved using `Schema`.
func (o SchemaPermissionOutput) SchemaId() pulumi.StringOutput {
	return o.ApplyT(func(v *SchemaPermission) pulumi.StringOutput { return v.SchemaId }).(pulumi.StringOutput)
}

// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
func (o SchemaPermissionOutput) WithGrantOption() pulumi.BoolOutput {
	return o.ApplyT(func(v *SchemaPermission) pulumi.BoolOutput { return v.WithGrantOption }).(pulumi.BoolOutput)
}

type SchemaPermissionArrayOutput struct{ *pulumi.OutputState }

func (SchemaPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SchemaPermission)(nil)).Elem()
}

func (o SchemaPermissionArrayOutput) ToSchemaPermissionArrayOutput() SchemaPermissionArrayOutput {
	return o
}

func (o SchemaPermissionArrayOutput) ToSchemaPermissionArrayOutputWithContext(ctx context.Context) SchemaPermissionArrayOutput {
	return o
}

func (o SchemaPermissionArrayOutput) Index(i pulumi.IntInput) SchemaPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SchemaPermission {
		return vs[0].([]*SchemaPermission)[vs[1].(int)]
	}).(SchemaPermissionOutput)
}

type SchemaPermissionMapOutput struct{ *pulumi.OutputState }

func (SchemaPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SchemaPermission)(nil)).Elem()
}

func (o SchemaPermissionMapOutput) ToSchemaPermissionMapOutput() SchemaPermissionMapOutput {
	return o
}

func (o SchemaPermissionMapOutput) ToSchemaPermissionMapOutputWithContext(ctx context.Context) SchemaPermissionMapOutput {
	return o
}

func (o SchemaPermissionMapOutput) MapIndex(k pulumi.StringInput) SchemaPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SchemaPermission {
		return vs[0].(map[string]*SchemaPermission)[vs[1].(string)]
	}).(SchemaPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaPermissionInput)(nil)).Elem(), &SchemaPermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaPermissionArrayInput)(nil)).Elem(), SchemaPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaPermissionMapInput)(nil)).Elem(), SchemaPermissionMap{})
	pulumi.RegisterOutputType(SchemaPermissionOutput{})
	pulumi.RegisterOutputType(SchemaPermissionArrayOutput{})
	pulumi.RegisterOutputType(SchemaPermissionMapOutput{})
}