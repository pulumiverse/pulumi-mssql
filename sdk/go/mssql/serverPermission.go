// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Grants server-level permission.
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
//			example, err := mssql.LookupSqlLogin(ctx, &mssql.LookupSqlLoginArgs{
//				Name: "example_login",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = mssql.NewServerPermission(ctx, "connectToExample", &mssql.ServerPermissionArgs{
//				PrincipalId:     *pulumi.String(example.PrincipalId),
//				Permission:      pulumi.String("CONNECT SQL"),
//				WithGrantOption: pulumi.Bool(true),
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
// import using <principal_id>/<permission>
//
// ```sh
//
//	$ pulumi import mssql:index/serverPermission:ServerPermission example '7/CONNECT SQL'
//
// ```
type ServerPermission struct {
	pulumi.CustomResourceState

	// Name of server-level SQL permission. For full list of supported permissions see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-server-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringOutput `pulumi:"permission"`
	// ID of the principal who will be granted `permission`. Can be retrieved using `ServerRole` or `SqlLogin`.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`
	WithGrantOption pulumi.BoolOutput `pulumi:"withGrantOption"`
}

// NewServerPermission registers a new resource with the given unique name, arguments, and options.
func NewServerPermission(ctx *pulumi.Context,
	name string, args *ServerPermissionArgs, opts ...pulumi.ResourceOption) (*ServerPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	if args.PrincipalId == nil {
		return nil, errors.New("invalid value for required argument 'PrincipalId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServerPermission
	err := ctx.RegisterResource("mssql:index/serverPermission:ServerPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerPermission gets an existing ServerPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerPermissionState, opts ...pulumi.ResourceOption) (*ServerPermission, error) {
	var resource ServerPermission
	err := ctx.ReadResource("mssql:index/serverPermission:ServerPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerPermission resources.
type serverPermissionState struct {
	// Name of server-level SQL permission. For full list of supported permissions see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-server-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission *string `pulumi:"permission"`
	// ID of the principal who will be granted `permission`. Can be retrieved using `ServerRole` or `SqlLogin`.
	PrincipalId *string `pulumi:"principalId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

type ServerPermissionState struct {
	// Name of server-level SQL permission. For full list of supported permissions see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-server-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringPtrInput
	// ID of the principal who will be granted `permission`. Can be retrieved using `ServerRole` or `SqlLogin`.
	PrincipalId pulumi.StringPtrInput
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`
	WithGrantOption pulumi.BoolPtrInput
}

func (ServerPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverPermissionState)(nil)).Elem()
}

type serverPermissionArgs struct {
	// Name of server-level SQL permission. For full list of supported permissions see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-server-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission string `pulumi:"permission"`
	// ID of the principal who will be granted `permission`. Can be retrieved using `ServerRole` or `SqlLogin`.
	PrincipalId string `pulumi:"principalId"`
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`
	WithGrantOption *bool `pulumi:"withGrantOption"`
}

// The set of arguments for constructing a ServerPermission resource.
type ServerPermissionArgs struct {
	// Name of server-level SQL permission. For full list of supported permissions see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-server-permissions-transact-sql?view=azuresqldb-current#remarks)
	Permission pulumi.StringInput
	// ID of the principal who will be granted `permission`. Can be retrieved using `ServerRole` or `SqlLogin`.
	PrincipalId pulumi.StringInput
	// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`
	WithGrantOption pulumi.BoolPtrInput
}

func (ServerPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverPermissionArgs)(nil)).Elem()
}

type ServerPermissionInput interface {
	pulumi.Input

	ToServerPermissionOutput() ServerPermissionOutput
	ToServerPermissionOutputWithContext(ctx context.Context) ServerPermissionOutput
}

func (*ServerPermission) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPermission)(nil)).Elem()
}

func (i *ServerPermission) ToServerPermissionOutput() ServerPermissionOutput {
	return i.ToServerPermissionOutputWithContext(context.Background())
}

func (i *ServerPermission) ToServerPermissionOutputWithContext(ctx context.Context) ServerPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPermissionOutput)
}

// ServerPermissionArrayInput is an input type that accepts ServerPermissionArray and ServerPermissionArrayOutput values.
// You can construct a concrete instance of `ServerPermissionArrayInput` via:
//
//	ServerPermissionArray{ ServerPermissionArgs{...} }
type ServerPermissionArrayInput interface {
	pulumi.Input

	ToServerPermissionArrayOutput() ServerPermissionArrayOutput
	ToServerPermissionArrayOutputWithContext(context.Context) ServerPermissionArrayOutput
}

type ServerPermissionArray []ServerPermissionInput

func (ServerPermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerPermission)(nil)).Elem()
}

func (i ServerPermissionArray) ToServerPermissionArrayOutput() ServerPermissionArrayOutput {
	return i.ToServerPermissionArrayOutputWithContext(context.Background())
}

func (i ServerPermissionArray) ToServerPermissionArrayOutputWithContext(ctx context.Context) ServerPermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPermissionArrayOutput)
}

// ServerPermissionMapInput is an input type that accepts ServerPermissionMap and ServerPermissionMapOutput values.
// You can construct a concrete instance of `ServerPermissionMapInput` via:
//
//	ServerPermissionMap{ "key": ServerPermissionArgs{...} }
type ServerPermissionMapInput interface {
	pulumi.Input

	ToServerPermissionMapOutput() ServerPermissionMapOutput
	ToServerPermissionMapOutputWithContext(context.Context) ServerPermissionMapOutput
}

type ServerPermissionMap map[string]ServerPermissionInput

func (ServerPermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerPermission)(nil)).Elem()
}

func (i ServerPermissionMap) ToServerPermissionMapOutput() ServerPermissionMapOutput {
	return i.ToServerPermissionMapOutputWithContext(context.Background())
}

func (i ServerPermissionMap) ToServerPermissionMapOutputWithContext(ctx context.Context) ServerPermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPermissionMapOutput)
}

type ServerPermissionOutput struct{ *pulumi.OutputState }

func (ServerPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPermission)(nil)).Elem()
}

func (o ServerPermissionOutput) ToServerPermissionOutput() ServerPermissionOutput {
	return o
}

func (o ServerPermissionOutput) ToServerPermissionOutputWithContext(ctx context.Context) ServerPermissionOutput {
	return o
}

// Name of server-level SQL permission. For full list of supported permissions see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-server-permissions-transact-sql?view=azuresqldb-current#remarks)
func (o ServerPermissionOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPermission) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// ID of the principal who will be granted `permission`. Can be retrieved using `ServerRole` or `SqlLogin`.
func (o ServerPermissionOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerPermission) pulumi.StringOutput { return v.PrincipalId }).(pulumi.StringOutput)
}

// When set to `true`, `principalId` will be allowed to grant the `permission` to other principals. Defaults to `false`
func (o ServerPermissionOutput) WithGrantOption() pulumi.BoolOutput {
	return o.ApplyT(func(v *ServerPermission) pulumi.BoolOutput { return v.WithGrantOption }).(pulumi.BoolOutput)
}

type ServerPermissionArrayOutput struct{ *pulumi.OutputState }

func (ServerPermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerPermission)(nil)).Elem()
}

func (o ServerPermissionArrayOutput) ToServerPermissionArrayOutput() ServerPermissionArrayOutput {
	return o
}

func (o ServerPermissionArrayOutput) ToServerPermissionArrayOutputWithContext(ctx context.Context) ServerPermissionArrayOutput {
	return o
}

func (o ServerPermissionArrayOutput) Index(i pulumi.IntInput) ServerPermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerPermission {
		return vs[0].([]*ServerPermission)[vs[1].(int)]
	}).(ServerPermissionOutput)
}

type ServerPermissionMapOutput struct{ *pulumi.OutputState }

func (ServerPermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerPermission)(nil)).Elem()
}

func (o ServerPermissionMapOutput) ToServerPermissionMapOutput() ServerPermissionMapOutput {
	return o
}

func (o ServerPermissionMapOutput) ToServerPermissionMapOutputWithContext(ctx context.Context) ServerPermissionMapOutput {
	return o
}

func (o ServerPermissionMapOutput) MapIndex(k pulumi.StringInput) ServerPermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerPermission {
		return vs[0].(map[string]*ServerPermission)[vs[1].(string)]
	}).(ServerPermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPermissionInput)(nil)).Elem(), &ServerPermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPermissionArrayInput)(nil)).Elem(), ServerPermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerPermissionMapInput)(nil)).Elem(), ServerPermissionMap{})
	pulumi.RegisterOutputType(ServerPermissionOutput{})
	pulumi.RegisterOutputType(ServerPermissionArrayOutput{})
	pulumi.RegisterOutputType(ServerPermissionMapOutput{})
}
