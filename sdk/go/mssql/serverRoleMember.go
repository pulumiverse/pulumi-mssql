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

// Manages server role membership.
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
//			member, err := mssql.LookupSqlLogin(ctx, &mssql.LookupSqlLoginArgs{
//				Name: "member_login",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleServerRole, err := mssql.NewServerRole(ctx, "exampleServerRole", nil)
//			if err != nil {
//				return err
//			}
//			_, err = mssql.NewServerRoleMember(ctx, "exampleServerRoleMember", &mssql.ServerRoleMemberArgs{
//				RoleId:   exampleServerRole.ID(),
//				MemberId: pulumi.String(member.Id),
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
// import using <role_id>/<member_id> - can be retrieved using `sys.server_principals` view
//
// ```sh
// $ pulumi import mssql:index/serverRoleMember:ServerRoleMember example '7/5'
// ```
type ServerRoleMember struct {
	pulumi.CustomResourceState

	// ID of the member. Can be retrieved using `ServerRole` or `SqlLogin`
	MemberId pulumi.StringOutput `pulumi:"memberId"`
	// ID of the server role. Can be retrieved using `ServerRole`
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewServerRoleMember registers a new resource with the given unique name, arguments, and options.
func NewServerRoleMember(ctx *pulumi.Context,
	name string, args *ServerRoleMemberArgs, opts ...pulumi.ResourceOption) (*ServerRoleMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MemberId == nil {
		return nil, errors.New("invalid value for required argument 'MemberId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerRoleMember
	err := ctx.RegisterResource("mssql:index/serverRoleMember:ServerRoleMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerRoleMember gets an existing ServerRoleMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerRoleMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerRoleMemberState, opts ...pulumi.ResourceOption) (*ServerRoleMember, error) {
	var resource ServerRoleMember
	err := ctx.ReadResource("mssql:index/serverRoleMember:ServerRoleMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerRoleMember resources.
type serverRoleMemberState struct {
	// ID of the member. Can be retrieved using `ServerRole` or `SqlLogin`
	MemberId *string `pulumi:"memberId"`
	// ID of the server role. Can be retrieved using `ServerRole`
	RoleId *string `pulumi:"roleId"`
}

type ServerRoleMemberState struct {
	// ID of the member. Can be retrieved using `ServerRole` or `SqlLogin`
	MemberId pulumi.StringPtrInput
	// ID of the server role. Can be retrieved using `ServerRole`
	RoleId pulumi.StringPtrInput
}

func (ServerRoleMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverRoleMemberState)(nil)).Elem()
}

type serverRoleMemberArgs struct {
	// ID of the member. Can be retrieved using `ServerRole` or `SqlLogin`
	MemberId string `pulumi:"memberId"`
	// ID of the server role. Can be retrieved using `ServerRole`
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a ServerRoleMember resource.
type ServerRoleMemberArgs struct {
	// ID of the member. Can be retrieved using `ServerRole` or `SqlLogin`
	MemberId pulumi.StringInput
	// ID of the server role. Can be retrieved using `ServerRole`
	RoleId pulumi.StringInput
}

func (ServerRoleMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverRoleMemberArgs)(nil)).Elem()
}

type ServerRoleMemberInput interface {
	pulumi.Input

	ToServerRoleMemberOutput() ServerRoleMemberOutput
	ToServerRoleMemberOutputWithContext(ctx context.Context) ServerRoleMemberOutput
}

func (*ServerRoleMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerRoleMember)(nil)).Elem()
}

func (i *ServerRoleMember) ToServerRoleMemberOutput() ServerRoleMemberOutput {
	return i.ToServerRoleMemberOutputWithContext(context.Background())
}

func (i *ServerRoleMember) ToServerRoleMemberOutputWithContext(ctx context.Context) ServerRoleMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRoleMemberOutput)
}

// ServerRoleMemberArrayInput is an input type that accepts ServerRoleMemberArray and ServerRoleMemberArrayOutput values.
// You can construct a concrete instance of `ServerRoleMemberArrayInput` via:
//
//	ServerRoleMemberArray{ ServerRoleMemberArgs{...} }
type ServerRoleMemberArrayInput interface {
	pulumi.Input

	ToServerRoleMemberArrayOutput() ServerRoleMemberArrayOutput
	ToServerRoleMemberArrayOutputWithContext(context.Context) ServerRoleMemberArrayOutput
}

type ServerRoleMemberArray []ServerRoleMemberInput

func (ServerRoleMemberArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerRoleMember)(nil)).Elem()
}

func (i ServerRoleMemberArray) ToServerRoleMemberArrayOutput() ServerRoleMemberArrayOutput {
	return i.ToServerRoleMemberArrayOutputWithContext(context.Background())
}

func (i ServerRoleMemberArray) ToServerRoleMemberArrayOutputWithContext(ctx context.Context) ServerRoleMemberArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRoleMemberArrayOutput)
}

// ServerRoleMemberMapInput is an input type that accepts ServerRoleMemberMap and ServerRoleMemberMapOutput values.
// You can construct a concrete instance of `ServerRoleMemberMapInput` via:
//
//	ServerRoleMemberMap{ "key": ServerRoleMemberArgs{...} }
type ServerRoleMemberMapInput interface {
	pulumi.Input

	ToServerRoleMemberMapOutput() ServerRoleMemberMapOutput
	ToServerRoleMemberMapOutputWithContext(context.Context) ServerRoleMemberMapOutput
}

type ServerRoleMemberMap map[string]ServerRoleMemberInput

func (ServerRoleMemberMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerRoleMember)(nil)).Elem()
}

func (i ServerRoleMemberMap) ToServerRoleMemberMapOutput() ServerRoleMemberMapOutput {
	return i.ToServerRoleMemberMapOutputWithContext(context.Background())
}

func (i ServerRoleMemberMap) ToServerRoleMemberMapOutputWithContext(ctx context.Context) ServerRoleMemberMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRoleMemberMapOutput)
}

type ServerRoleMemberOutput struct{ *pulumi.OutputState }

func (ServerRoleMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerRoleMember)(nil)).Elem()
}

func (o ServerRoleMemberOutput) ToServerRoleMemberOutput() ServerRoleMemberOutput {
	return o
}

func (o ServerRoleMemberOutput) ToServerRoleMemberOutputWithContext(ctx context.Context) ServerRoleMemberOutput {
	return o
}

// ID of the member. Can be retrieved using `ServerRole` or `SqlLogin`
func (o ServerRoleMemberOutput) MemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRoleMember) pulumi.StringOutput { return v.MemberId }).(pulumi.StringOutput)
}

// ID of the server role. Can be retrieved using `ServerRole`
func (o ServerRoleMemberOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRoleMember) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

type ServerRoleMemberArrayOutput struct{ *pulumi.OutputState }

func (ServerRoleMemberArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerRoleMember)(nil)).Elem()
}

func (o ServerRoleMemberArrayOutput) ToServerRoleMemberArrayOutput() ServerRoleMemberArrayOutput {
	return o
}

func (o ServerRoleMemberArrayOutput) ToServerRoleMemberArrayOutputWithContext(ctx context.Context) ServerRoleMemberArrayOutput {
	return o
}

func (o ServerRoleMemberArrayOutput) Index(i pulumi.IntInput) ServerRoleMemberOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerRoleMember {
		return vs[0].([]*ServerRoleMember)[vs[1].(int)]
	}).(ServerRoleMemberOutput)
}

type ServerRoleMemberMapOutput struct{ *pulumi.OutputState }

func (ServerRoleMemberMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerRoleMember)(nil)).Elem()
}

func (o ServerRoleMemberMapOutput) ToServerRoleMemberMapOutput() ServerRoleMemberMapOutput {
	return o
}

func (o ServerRoleMemberMapOutput) ToServerRoleMemberMapOutputWithContext(ctx context.Context) ServerRoleMemberMapOutput {
	return o
}

func (o ServerRoleMemberMapOutput) MapIndex(k pulumi.StringInput) ServerRoleMemberOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerRoleMember {
		return vs[0].(map[string]*ServerRoleMember)[vs[1].(string)]
	}).(ServerRoleMemberOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRoleMemberInput)(nil)).Elem(), &ServerRoleMember{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRoleMemberArrayInput)(nil)).Elem(), ServerRoleMemberArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRoleMemberMapInput)(nil)).Elem(), ServerRoleMemberMap{})
	pulumi.RegisterOutputType(ServerRoleMemberOutput{})
	pulumi.RegisterOutputType(ServerRoleMemberArrayOutput{})
	pulumi.RegisterOutputType(ServerRoleMemberMapOutput{})
}
