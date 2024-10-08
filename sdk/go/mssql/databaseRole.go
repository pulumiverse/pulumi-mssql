// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Manages database-level role.
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
//			_, err = mssql.NewDatabaseRole(ctx, "exampleDatabaseRole", &mssql.DatabaseRoleArgs{
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
// import using <db_id>/<role_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', DATABASE_PRINCIPAL_ID('<role_name>'))`
//
// ```sh
// $ pulumi import mssql:index/databaseRole:DatabaseRole example '7/5'
// ```
type DatabaseRole struct {
	pulumi.CustomResourceState

	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name    pulumi.StringOutput `pulumi:"name"`
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
}

// NewDatabaseRole registers a new resource with the given unique name, arguments, and options.
func NewDatabaseRole(ctx *pulumi.Context,
	name string, args *DatabaseRoleArgs, opts ...pulumi.ResourceOption) (*DatabaseRole, error) {
	if args == nil {
		args = &DatabaseRoleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseRole
	err := ctx.RegisterResource("mssql:index/databaseRole:DatabaseRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseRole gets an existing DatabaseRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseRoleState, opts ...pulumi.ResourceOption) (*DatabaseRole, error) {
	var resource DatabaseRole
	err := ctx.ReadResource("mssql:index/databaseRole:DatabaseRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseRole resources.
type databaseRoleState struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name    *string `pulumi:"name"`
	OwnerId *string `pulumi:"ownerId"`
}

type DatabaseRoleState struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringPtrInput
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name    pulumi.StringPtrInput
	OwnerId pulumi.StringPtrInput
}

func (DatabaseRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseRoleState)(nil)).Elem()
}

type databaseRoleArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name    *string `pulumi:"name"`
	OwnerId *string `pulumi:"ownerId"`
}

// The set of arguments for constructing a DatabaseRole resource.
type DatabaseRoleArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringPtrInput
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name    pulumi.StringPtrInput
	OwnerId pulumi.StringPtrInput
}

func (DatabaseRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseRoleArgs)(nil)).Elem()
}

type DatabaseRoleInput interface {
	pulumi.Input

	ToDatabaseRoleOutput() DatabaseRoleOutput
	ToDatabaseRoleOutputWithContext(ctx context.Context) DatabaseRoleOutput
}

func (*DatabaseRole) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseRole)(nil)).Elem()
}

func (i *DatabaseRole) ToDatabaseRoleOutput() DatabaseRoleOutput {
	return i.ToDatabaseRoleOutputWithContext(context.Background())
}

func (i *DatabaseRole) ToDatabaseRoleOutputWithContext(ctx context.Context) DatabaseRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRoleOutput)
}

// DatabaseRoleArrayInput is an input type that accepts DatabaseRoleArray and DatabaseRoleArrayOutput values.
// You can construct a concrete instance of `DatabaseRoleArrayInput` via:
//
//	DatabaseRoleArray{ DatabaseRoleArgs{...} }
type DatabaseRoleArrayInput interface {
	pulumi.Input

	ToDatabaseRoleArrayOutput() DatabaseRoleArrayOutput
	ToDatabaseRoleArrayOutputWithContext(context.Context) DatabaseRoleArrayOutput
}

type DatabaseRoleArray []DatabaseRoleInput

func (DatabaseRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseRole)(nil)).Elem()
}

func (i DatabaseRoleArray) ToDatabaseRoleArrayOutput() DatabaseRoleArrayOutput {
	return i.ToDatabaseRoleArrayOutputWithContext(context.Background())
}

func (i DatabaseRoleArray) ToDatabaseRoleArrayOutputWithContext(ctx context.Context) DatabaseRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRoleArrayOutput)
}

// DatabaseRoleMapInput is an input type that accepts DatabaseRoleMap and DatabaseRoleMapOutput values.
// You can construct a concrete instance of `DatabaseRoleMapInput` via:
//
//	DatabaseRoleMap{ "key": DatabaseRoleArgs{...} }
type DatabaseRoleMapInput interface {
	pulumi.Input

	ToDatabaseRoleMapOutput() DatabaseRoleMapOutput
	ToDatabaseRoleMapOutputWithContext(context.Context) DatabaseRoleMapOutput
}

type DatabaseRoleMap map[string]DatabaseRoleInput

func (DatabaseRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseRole)(nil)).Elem()
}

func (i DatabaseRoleMap) ToDatabaseRoleMapOutput() DatabaseRoleMapOutput {
	return i.ToDatabaseRoleMapOutputWithContext(context.Background())
}

func (i DatabaseRoleMap) ToDatabaseRoleMapOutputWithContext(ctx context.Context) DatabaseRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRoleMapOutput)
}

type DatabaseRoleOutput struct{ *pulumi.OutputState }

func (DatabaseRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseRole)(nil)).Elem()
}

func (o DatabaseRoleOutput) ToDatabaseRoleOutput() DatabaseRoleOutput {
	return o
}

func (o DatabaseRoleOutput) ToDatabaseRoleOutputWithContext(ctx context.Context) DatabaseRoleOutput {
	return o
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
func (o DatabaseRoleOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseRole) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
func (o DatabaseRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DatabaseRoleOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseRole) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

type DatabaseRoleArrayOutput struct{ *pulumi.OutputState }

func (DatabaseRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseRole)(nil)).Elem()
}

func (o DatabaseRoleArrayOutput) ToDatabaseRoleArrayOutput() DatabaseRoleArrayOutput {
	return o
}

func (o DatabaseRoleArrayOutput) ToDatabaseRoleArrayOutputWithContext(ctx context.Context) DatabaseRoleArrayOutput {
	return o
}

func (o DatabaseRoleArrayOutput) Index(i pulumi.IntInput) DatabaseRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseRole {
		return vs[0].([]*DatabaseRole)[vs[1].(int)]
	}).(DatabaseRoleOutput)
}

type DatabaseRoleMapOutput struct{ *pulumi.OutputState }

func (DatabaseRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseRole)(nil)).Elem()
}

func (o DatabaseRoleMapOutput) ToDatabaseRoleMapOutput() DatabaseRoleMapOutput {
	return o
}

func (o DatabaseRoleMapOutput) ToDatabaseRoleMapOutputWithContext(ctx context.Context) DatabaseRoleMapOutput {
	return o
}

func (o DatabaseRoleMapOutput) MapIndex(k pulumi.StringInput) DatabaseRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseRole {
		return vs[0].(map[string]*DatabaseRole)[vs[1].(string)]
	}).(DatabaseRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseRoleInput)(nil)).Elem(), &DatabaseRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseRoleArrayInput)(nil)).Elem(), DatabaseRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseRoleMapInput)(nil)).Elem(), DatabaseRoleMap{})
	pulumi.RegisterOutputType(DatabaseRoleOutput{})
	pulumi.RegisterOutputType(DatabaseRoleArrayOutput{})
	pulumi.RegisterOutputType(DatabaseRoleMapOutput{})
}
