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

// Grants database-level permission.
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
//			exampleSqlUser, err := mssql.LookupSqlUser(ctx, &mssql.LookupSqlUserArgs{
//				Name:       "example_user",
//				DatabaseId: pulumi.StringRef(exampleDatabase.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = mssql.NewDatabasePermission(ctx, "deleteToExample", &mssql.DatabasePermissionArgs{
//				PrincipalId: pulumi.String(exampleSqlUser.Id),
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
// import using <db_id>/<principal_id>/<permission> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', DATABASE_PRINCIPAL_ID('<principal_name>'), '/DELETE')`
//
// ```sh
// $ pulumi import mssql:index/databasePermission:DatabasePermission example '7/5/DELETE'
// ```
type DatabasePermission struct {
	pulumi.CustomResourceState

	// Name of database-level SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-database-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringOutput `pulumi:"permission"`
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption pulumi.BoolOutput `pulumi:"withGrantOption"`
}

// NewDatabasePermission registers a new resource with the given unique name, arguments, and options.
func NewDatabasePermission(ctx *pulumi.Context,
	name string, args *DatabasePermissionArgs, opts ...pulumi.ResourceOption) (*DatabasePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabasePermission
	err := ctx.RegisterResource("mssql:index/databasePermission:DatabasePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabasePermission gets an existing DatabasePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabasePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabasePermissionState, opts ...pulumi.ResourceOption) (*DatabasePermission, error) {
	var resource DatabasePermission
	err := ctx.ReadResource("mssql:index/databasePermission:DatabasePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabasePermission resources.
type databasePermissionState struct {
	// Name of database-level SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-database-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission *string `pulumi:"permission"`
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId *string `pulumi:"principalId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

type DatabasePermissionState struct {
	// Name of database-level SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-database-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringPtrInput
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId pulumi.StringPtrInput
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption pulumi.BoolPtrInput
}

func (DatabasePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePermissionState)(nil)).Elem()
}

type databasePermissionArgs struct {
	// Name of database-level SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-database-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission string `pulumi:"permission"`
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId string `pulumi:"principalId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

// The set of arguments for constructing a DatabasePermission resource.
type DatabasePermissionArgs struct {
	// Name of database-level SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-database-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringInput
	// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
	PrincipalId pulumi.StringInput
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
	WithGrantOption pulumi.BoolPtrInput
}

func (DatabasePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databasePermissionArgs)(nil)).Elem()
}

type DatabasePermissionInput interface {
	pulumi.Input

	ToDatabasePermissionOutput() DatabasePermissionOutput
	ToDatabasePermissionOutputWithContext(ctx context.Context) DatabasePermissionOutput
}

func (*DatabasePermission) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePermission)(nil)).Elem()
}

func (i *DatabasePermission) ToDatabasePermissionOutput() DatabasePermissionOutput {
	return i.ToDatabasePermissionOutputWithContext(context.Background())
}

func (i *DatabasePermission) ToDatabasePermissionOutputWithContext(ctx context.Context) DatabasePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePermissionOutput)
}

// DatabasePermissionArrayInput is an input type that accepts DatabasePermissionArray and DatabasePermissionArrayOutput values.
// You can construct a concrete instance of `DatabasePermissionArrayInput` via:
//
//	DatabasePermissionArray{ DatabasePermissionArgs{...} }
type DatabasePermissionArrayInput interface {
	pulumi.Input

	ToDatabasePermissionArrayOutput() DatabasePermissionArrayOutput
	ToDatabasePermissionArrayOutputWithContext(context.Context) DatabasePermissionArrayOutput
}

type DatabasePermissionArray []DatabasePermissionInput

func (DatabasePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabasePermission)(nil)).Elem()
}

func (i DatabasePermissionArray) ToDatabasePermissionArrayOutput() DatabasePermissionArrayOutput {
	return i.ToDatabasePermissionArrayOutputWithContext(context.Background())
}

func (i DatabasePermissionArray) ToDatabasePermissionArrayOutputWithContext(ctx context.Context) DatabasePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePermissionArrayOutput)
}

// DatabasePermissionMapInput is an input type that accepts DatabasePermissionMap and DatabasePermissionMapOutput values.
// You can construct a concrete instance of `DatabasePermissionMapInput` via:
//
//	DatabasePermissionMap{ "key": DatabasePermissionArgs{...} }
type DatabasePermissionMapInput interface {
	pulumi.Input

	ToDatabasePermissionMapOutput() DatabasePermissionMapOutput
	ToDatabasePermissionMapOutputWithContext(context.Context) DatabasePermissionMapOutput
}

type DatabasePermissionMap map[string]DatabasePermissionInput

func (DatabasePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabasePermission)(nil)).Elem()
}

func (i DatabasePermissionMap) ToDatabasePermissionMapOutput() DatabasePermissionMapOutput {
	return i.ToDatabasePermissionMapOutputWithContext(context.Background())
}

func (i DatabasePermissionMap) ToDatabasePermissionMapOutputWithContext(ctx context.Context) DatabasePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePermissionMapOutput)
}

type DatabasePermissionOutput struct{ *pulumi.OutputState }

func (DatabasePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePermission)(nil)).Elem()
}

func (o DatabasePermissionOutput) ToDatabasePermissionOutput() DatabasePermissionOutput {
	return o
}

func (o DatabasePermissionOutput) ToDatabasePermissionOutputWithContext(ctx context.Context) DatabasePermissionOutput {
	return o
}

// Name of database-level SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-database-permissions-transact-sql?view=azuresqldb-current#remarks)
func (o DatabasePermissionOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePermission) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// `<database_id>/<principal_id>`. Can be retrieved using `DatabaseRole`, `SqlUser`, `AzureadUser` or `AzureadServicePrincipal`.
func (o DatabasePermissionOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabasePermission) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`.
func (o DatabasePermissionOutput) WithGrantOption() pulumi.BoolOutput {
	return o.ApplyT(func(v *DatabasePermission) pulumi.BoolOutput { return v.WithGrantOption }).(pulumi.BoolOutput)
}

type DatabasePermissionArrayOutput struct{ *pulumi.OutputState }

func (DatabasePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabasePermission)(nil)).Elem()
}

func (o DatabasePermissionArrayOutput) ToDatabasePermissionArrayOutput() DatabasePermissionArrayOutput {
	return o
}

func (o DatabasePermissionArrayOutput) ToDatabasePermissionArrayOutputWithContext(ctx context.Context) DatabasePermissionArrayOutput {
	return o
}

func (o DatabasePermissionArrayOutput) Index(i pulumi.IntInput) DatabasePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabasePermission {
		return vs[0].([]*DatabasePermission)[vs[1].(int)]
	}).(DatabasePermissionOutput)
}

type DatabasePermissionMapOutput struct{ *pulumi.OutputState }

func (DatabasePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabasePermission)(nil)).Elem()
}

func (o DatabasePermissionMapOutput) ToDatabasePermissionMapOutput() DatabasePermissionMapOutput {
	return o
}

func (o DatabasePermissionMapOutput) ToDatabasePermissionMapOutputWithContext(ctx context.Context) DatabasePermissionMapOutput {
	return o
}

func (o DatabasePermissionMapOutput) MapIndex(k pulumi.StringInput) DatabasePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabasePermission {
		return vs[0].(map[string]*DatabasePermission)[vs[1].(string)]
	}).(DatabasePermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePermissionInput)(nil)).Elem(), &DatabasePermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePermissionArrayInput)(nil)).Elem(), DatabasePermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasePermissionMapInput)(nil)).Elem(), DatabasePermissionMap{})
	pulumi.RegisterOutputType(DatabasePermissionOutput{})
	pulumi.RegisterOutputType(DatabasePermissionArrayOutput{})
	pulumi.RegisterOutputType(DatabasePermissionMapOutput{})
}
