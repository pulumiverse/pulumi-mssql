// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Manages database-level user, based on SQL login.
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
//			exampleSqlLogin, err := mssql.NewSqlLogin(ctx, "exampleSqlLogin", &mssql.SqlLoginArgs{
//				Password:                pulumi.String("Str0ngPa$$word12"),
//				MustChangePassword:      pulumi.Bool(true),
//				DefaultDatabaseId:       *pulumi.String(exampleDatabase.Id),
//				DefaultLanguage:         pulumi.String("english"),
//				CheckPasswordExpiration: pulumi.Bool(true),
//				CheckPasswordPolicy:     pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSqlUser, err := mssql.NewSqlUser(ctx, "exampleSqlUser", &mssql.SqlUserArgs{
//				DatabaseId: *pulumi.String(exampleDatabase.Id),
//				LoginId:    exampleSqlLogin.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("userId", exampleSqlUser.ID())
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// import using <db_id>/<user_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', DATABASE_PRINCIPAL_ID('<username>'))`
//
// ```sh
//
//	$ pulumi import mssql:index/sqlUser:SqlUser example '7/5'
//
// ```
type SqlUser struct {
	pulumi.CustomResourceState

	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// SID of SQL login. Can be retrieved using `SqlLogin` or `SELECT SUSER_SID('<login_name>')`.
	LoginId pulumi.StringOutput `pulumi:"loginId"`
	// User name. Cannot be longer than 128 chars.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSqlUser registers a new resource with the given unique name, arguments, and options.
func NewSqlUser(ctx *pulumi.Context,
	name string, args *SqlUserArgs, opts ...pulumi.ResourceOption) (*SqlUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoginId == nil {
		return nil, errors.New("invalid value for required argument 'LoginId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SqlUser
	err := ctx.RegisterResource("mssql:index/sqlUser:SqlUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlUser gets an existing SqlUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlUserState, opts ...pulumi.ResourceOption) (*SqlUser, error) {
	var resource SqlUser
	err := ctx.ReadResource("mssql:index/sqlUser:SqlUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlUser resources.
type sqlUserState struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// SID of SQL login. Can be retrieved using `SqlLogin` or `SELECT SUSER_SID('<login_name>')`.
	LoginId *string `pulumi:"loginId"`
	// User name. Cannot be longer than 128 chars.
	Name *string `pulumi:"name"`
}

type SqlUserState struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringPtrInput
	// SID of SQL login. Can be retrieved using `SqlLogin` or `SELECT SUSER_SID('<login_name>')`.
	LoginId pulumi.StringPtrInput
	// User name. Cannot be longer than 128 chars.
	Name pulumi.StringPtrInput
}

func (SqlUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlUserState)(nil)).Elem()
}

type sqlUserArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId *string `pulumi:"databaseId"`
	// SID of SQL login. Can be retrieved using `SqlLogin` or `SELECT SUSER_SID('<login_name>')`.
	LoginId string `pulumi:"loginId"`
	// User name. Cannot be longer than 128 chars.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SqlUser resource.
type SqlUserArgs struct {
	// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
	DatabaseId pulumi.StringPtrInput
	// SID of SQL login. Can be retrieved using `SqlLogin` or `SELECT SUSER_SID('<login_name>')`.
	LoginId pulumi.StringInput
	// User name. Cannot be longer than 128 chars.
	Name pulumi.StringPtrInput
}

func (SqlUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlUserArgs)(nil)).Elem()
}

type SqlUserInput interface {
	pulumi.Input

	ToSqlUserOutput() SqlUserOutput
	ToSqlUserOutputWithContext(ctx context.Context) SqlUserOutput
}

func (*SqlUser) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUser)(nil)).Elem()
}

func (i *SqlUser) ToSqlUserOutput() SqlUserOutput {
	return i.ToSqlUserOutputWithContext(context.Background())
}

func (i *SqlUser) ToSqlUserOutputWithContext(ctx context.Context) SqlUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserOutput)
}

func (i *SqlUser) ToOutput(ctx context.Context) pulumix.Output[*SqlUser] {
	return pulumix.Output[*SqlUser]{
		OutputState: i.ToSqlUserOutputWithContext(ctx).OutputState,
	}
}

// SqlUserArrayInput is an input type that accepts SqlUserArray and SqlUserArrayOutput values.
// You can construct a concrete instance of `SqlUserArrayInput` via:
//
//	SqlUserArray{ SqlUserArgs{...} }
type SqlUserArrayInput interface {
	pulumi.Input

	ToSqlUserArrayOutput() SqlUserArrayOutput
	ToSqlUserArrayOutputWithContext(context.Context) SqlUserArrayOutput
}

type SqlUserArray []SqlUserInput

func (SqlUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlUser)(nil)).Elem()
}

func (i SqlUserArray) ToSqlUserArrayOutput() SqlUserArrayOutput {
	return i.ToSqlUserArrayOutputWithContext(context.Background())
}

func (i SqlUserArray) ToSqlUserArrayOutputWithContext(ctx context.Context) SqlUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserArrayOutput)
}

func (i SqlUserArray) ToOutput(ctx context.Context) pulumix.Output[[]*SqlUser] {
	return pulumix.Output[[]*SqlUser]{
		OutputState: i.ToSqlUserArrayOutputWithContext(ctx).OutputState,
	}
}

// SqlUserMapInput is an input type that accepts SqlUserMap and SqlUserMapOutput values.
// You can construct a concrete instance of `SqlUserMapInput` via:
//
//	SqlUserMap{ "key": SqlUserArgs{...} }
type SqlUserMapInput interface {
	pulumi.Input

	ToSqlUserMapOutput() SqlUserMapOutput
	ToSqlUserMapOutputWithContext(context.Context) SqlUserMapOutput
}

type SqlUserMap map[string]SqlUserInput

func (SqlUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlUser)(nil)).Elem()
}

func (i SqlUserMap) ToSqlUserMapOutput() SqlUserMapOutput {
	return i.ToSqlUserMapOutputWithContext(context.Background())
}

func (i SqlUserMap) ToSqlUserMapOutputWithContext(ctx context.Context) SqlUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserMapOutput)
}

func (i SqlUserMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SqlUser] {
	return pulumix.Output[map[string]*SqlUser]{
		OutputState: i.ToSqlUserMapOutputWithContext(ctx).OutputState,
	}
}

type SqlUserOutput struct{ *pulumi.OutputState }

func (SqlUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUser)(nil)).Elem()
}

func (o SqlUserOutput) ToSqlUserOutput() SqlUserOutput {
	return o
}

func (o SqlUserOutput) ToSqlUserOutputWithContext(ctx context.Context) SqlUserOutput {
	return o
}

func (o SqlUserOutput) ToOutput(ctx context.Context) pulumix.Output[*SqlUser] {
	return pulumix.Output[*SqlUser]{
		OutputState: o.OutputState,
	}
}

// ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
func (o SqlUserOutput) DatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringOutput { return v.DatabaseId }).(pulumi.StringOutput)
}

// SID of SQL login. Can be retrieved using `SqlLogin` or `SELECT SUSER_SID('<login_name>')`.
func (o SqlUserOutput) LoginId() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringOutput { return v.LoginId }).(pulumi.StringOutput)
}

// User name. Cannot be longer than 128 chars.
func (o SqlUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type SqlUserArrayOutput struct{ *pulumi.OutputState }

func (SqlUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SqlUser)(nil)).Elem()
}

func (o SqlUserArrayOutput) ToSqlUserArrayOutput() SqlUserArrayOutput {
	return o
}

func (o SqlUserArrayOutput) ToSqlUserArrayOutputWithContext(ctx context.Context) SqlUserArrayOutput {
	return o
}

func (o SqlUserArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SqlUser] {
	return pulumix.Output[[]*SqlUser]{
		OutputState: o.OutputState,
	}
}

func (o SqlUserArrayOutput) Index(i pulumi.IntInput) SqlUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SqlUser {
		return vs[0].([]*SqlUser)[vs[1].(int)]
	}).(SqlUserOutput)
}

type SqlUserMapOutput struct{ *pulumi.OutputState }

func (SqlUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SqlUser)(nil)).Elem()
}

func (o SqlUserMapOutput) ToSqlUserMapOutput() SqlUserMapOutput {
	return o
}

func (o SqlUserMapOutput) ToSqlUserMapOutputWithContext(ctx context.Context) SqlUserMapOutput {
	return o
}

func (o SqlUserMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SqlUser] {
	return pulumix.Output[map[string]*SqlUser]{
		OutputState: o.OutputState,
	}
}

func (o SqlUserMapOutput) MapIndex(k pulumi.StringInput) SqlUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SqlUser {
		return vs[0].(map[string]*SqlUser)[vs[1].(string)]
	}).(SqlUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlUserInput)(nil)).Elem(), &SqlUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlUserArrayInput)(nil)).Elem(), SqlUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlUserMapInput)(nil)).Elem(), SqlUserMap{})
	pulumi.RegisterOutputType(SqlUserOutput{})
	pulumi.RegisterOutputType(SqlUserArrayOutput{})
	pulumi.RegisterOutputType(SqlUserMapOutput{})
}
