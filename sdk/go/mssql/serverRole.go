// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages server-level role.
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
//			owner, err := mssql.NewServerRole(ctx, "owner", &mssql.ServerRoleArgs{
//				Name: pulumi.String("owner_role"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = mssql.NewServerRole(ctx, "example", &mssql.ServerRoleArgs{
//				Name:    pulumi.String("example"),
//				OwnerId: owner.ID(),
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
// import using <role_id> - can be retrieved using `SELECT [principal_id] FROM sys.server_principals WHERE [name]='<role_name>'`
//
// ```sh
//
//	$ pulumi import mssql:index/serverRole:ServerRole example 7
//
// ```
type ServerRole struct {
	pulumi.CustomResourceState

	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of another server role or login owning this role. Can be retrieved using `ServerRole` or `SqlLogin`.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
}

// NewServerRole registers a new resource with the given unique name, arguments, and options.
func NewServerRole(ctx *pulumi.Context,
	name string, args *ServerRoleArgs, opts ...pulumi.ResourceOption) (*ServerRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServerRole
	err := ctx.RegisterResource("mssql:index/serverRole:ServerRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerRole gets an existing ServerRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerRoleState, opts ...pulumi.ResourceOption) (*ServerRole, error) {
	var resource ServerRole
	err := ctx.ReadResource("mssql:index/serverRole:ServerRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerRole resources.
type serverRoleState struct {
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name *string `pulumi:"name"`
	// ID of another server role or login owning this role. Can be retrieved using `ServerRole` or `SqlLogin`.
	OwnerId *string `pulumi:"ownerId"`
}

type ServerRoleState struct {
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name pulumi.StringPtrInput
	// ID of another server role or login owning this role. Can be retrieved using `ServerRole` or `SqlLogin`.
	OwnerId pulumi.StringPtrInput
}

func (ServerRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverRoleState)(nil)).Elem()
}

type serverRoleArgs struct {
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name string `pulumi:"name"`
	// ID of another server role or login owning this role. Can be retrieved using `ServerRole` or `SqlLogin`.
	OwnerId *string `pulumi:"ownerId"`
}

// The set of arguments for constructing a ServerRole resource.
type ServerRoleArgs struct {
	// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
	Name pulumi.StringInput
	// ID of another server role or login owning this role. Can be retrieved using `ServerRole` or `SqlLogin`.
	OwnerId pulumi.StringPtrInput
}

func (ServerRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverRoleArgs)(nil)).Elem()
}

type ServerRoleInput interface {
	pulumi.Input

	ToServerRoleOutput() ServerRoleOutput
	ToServerRoleOutputWithContext(ctx context.Context) ServerRoleOutput
}

func (*ServerRole) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerRole)(nil)).Elem()
}

func (i *ServerRole) ToServerRoleOutput() ServerRoleOutput {
	return i.ToServerRoleOutputWithContext(context.Background())
}

func (i *ServerRole) ToServerRoleOutputWithContext(ctx context.Context) ServerRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRoleOutput)
}

// ServerRoleArrayInput is an input type that accepts ServerRoleArray and ServerRoleArrayOutput values.
// You can construct a concrete instance of `ServerRoleArrayInput` via:
//
//	ServerRoleArray{ ServerRoleArgs{...} }
type ServerRoleArrayInput interface {
	pulumi.Input

	ToServerRoleArrayOutput() ServerRoleArrayOutput
	ToServerRoleArrayOutputWithContext(context.Context) ServerRoleArrayOutput
}

type ServerRoleArray []ServerRoleInput

func (ServerRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerRole)(nil)).Elem()
}

func (i ServerRoleArray) ToServerRoleArrayOutput() ServerRoleArrayOutput {
	return i.ToServerRoleArrayOutputWithContext(context.Background())
}

func (i ServerRoleArray) ToServerRoleArrayOutputWithContext(ctx context.Context) ServerRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRoleArrayOutput)
}

// ServerRoleMapInput is an input type that accepts ServerRoleMap and ServerRoleMapOutput values.
// You can construct a concrete instance of `ServerRoleMapInput` via:
//
//	ServerRoleMap{ "key": ServerRoleArgs{...} }
type ServerRoleMapInput interface {
	pulumi.Input

	ToServerRoleMapOutput() ServerRoleMapOutput
	ToServerRoleMapOutputWithContext(context.Context) ServerRoleMapOutput
}

type ServerRoleMap map[string]ServerRoleInput

func (ServerRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerRole)(nil)).Elem()
}

func (i ServerRoleMap) ToServerRoleMapOutput() ServerRoleMapOutput {
	return i.ToServerRoleMapOutputWithContext(context.Background())
}

func (i ServerRoleMap) ToServerRoleMapOutputWithContext(ctx context.Context) ServerRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRoleMapOutput)
}

type ServerRoleOutput struct{ *pulumi.OutputState }

func (ServerRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerRole)(nil)).Elem()
}

func (o ServerRoleOutput) ToServerRoleOutput() ServerRoleOutput {
	return o
}

func (o ServerRoleOutput) ToServerRoleOutputWithContext(ctx context.Context) ServerRoleOutput {
	return o
}

// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
func (o ServerRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of another server role or login owning this role. Can be retrieved using `ServerRole` or `SqlLogin`.
func (o ServerRoleOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRole) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

type ServerRoleArrayOutput struct{ *pulumi.OutputState }

func (ServerRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerRole)(nil)).Elem()
}

func (o ServerRoleArrayOutput) ToServerRoleArrayOutput() ServerRoleArrayOutput {
	return o
}

func (o ServerRoleArrayOutput) ToServerRoleArrayOutputWithContext(ctx context.Context) ServerRoleArrayOutput {
	return o
}

func (o ServerRoleArrayOutput) Index(i pulumi.IntInput) ServerRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerRole {
		return vs[0].([]*ServerRole)[vs[1].(int)]
	}).(ServerRoleOutput)
}

type ServerRoleMapOutput struct{ *pulumi.OutputState }

func (ServerRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerRole)(nil)).Elem()
}

func (o ServerRoleMapOutput) ToServerRoleMapOutput() ServerRoleMapOutput {
	return o
}

func (o ServerRoleMapOutput) ToServerRoleMapOutputWithContext(ctx context.Context) ServerRoleMapOutput {
	return o
}

func (o ServerRoleMapOutput) MapIndex(k pulumi.StringInput) ServerRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerRole {
		return vs[0].(map[string]*ServerRole)[vs[1].(string)]
	}).(ServerRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRoleInput)(nil)).Elem(), &ServerRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRoleArrayInput)(nil)).Elem(), ServerRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRoleMapInput)(nil)).Elem(), ServerRoleMap{})
	pulumi.RegisterOutputType(ServerRoleOutput{})
	pulumi.RegisterOutputType(ServerRoleArrayOutput{})
	pulumi.RegisterOutputType(ServerRoleMapOutput{})
}