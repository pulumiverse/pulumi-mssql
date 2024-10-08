// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

var _ = internal.GetEnvOrDefault

type AzureAuth struct {
	// Service Principal client (application) ID. When omitted, default, chained set of credentials will be used.
	ClientId *string `pulumi:"clientId"`
	// Service Principal secret. When omitted, default, chained set of credentials will be used.
	ClientSecret *string `pulumi:"clientSecret"`
	// Azure AD tenant ID. Required only if Azure SQL Server's tenant is different than Service Principal's.
	TenantId *string `pulumi:"tenantId"`
}

// AzureAuthInput is an input type that accepts AzureAuthArgs and AzureAuthOutput values.
// You can construct a concrete instance of `AzureAuthInput` via:
//
//	AzureAuthArgs{...}
type AzureAuthInput interface {
	pulumi.Input

	ToAzureAuthOutput() AzureAuthOutput
	ToAzureAuthOutputWithContext(context.Context) AzureAuthOutput
}

type AzureAuthArgs struct {
	// Service Principal client (application) ID. When omitted, default, chained set of credentials will be used.
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	// Service Principal secret. When omitted, default, chained set of credentials will be used.
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
	// Azure AD tenant ID. Required only if Azure SQL Server's tenant is different than Service Principal's.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AzureAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureAuth)(nil)).Elem()
}

func (i AzureAuthArgs) ToAzureAuthOutput() AzureAuthOutput {
	return i.ToAzureAuthOutputWithContext(context.Background())
}

func (i AzureAuthArgs) ToAzureAuthOutputWithContext(ctx context.Context) AzureAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureAuthOutput)
}

type AzureAuthOutput struct{ *pulumi.OutputState }

func (AzureAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureAuth)(nil)).Elem()
}

func (o AzureAuthOutput) ToAzureAuthOutput() AzureAuthOutput {
	return o
}

func (o AzureAuthOutput) ToAzureAuthOutputWithContext(ctx context.Context) AzureAuthOutput {
	return o
}

// Service Principal client (application) ID. When omitted, default, chained set of credentials will be used.
func (o AzureAuthOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureAuth) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

// Service Principal secret. When omitted, default, chained set of credentials will be used.
func (o AzureAuthOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureAuth) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// Azure AD tenant ID. Required only if Azure SQL Server's tenant is different than Service Principal's.
func (o AzureAuthOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureAuth) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type SqlAuth struct {
	// Password for SQL authentication.
	Password string `pulumi:"password"`
	// User name for SQL authentication.
	Username string `pulumi:"username"`
}

// SqlAuthInput is an input type that accepts SqlAuthArgs and SqlAuthOutput values.
// You can construct a concrete instance of `SqlAuthInput` via:
//
//	SqlAuthArgs{...}
type SqlAuthInput interface {
	pulumi.Input

	ToSqlAuthOutput() SqlAuthOutput
	ToSqlAuthOutputWithContext(context.Context) SqlAuthOutput
}

type SqlAuthArgs struct {
	// Password for SQL authentication.
	Password pulumi.StringInput `pulumi:"password"`
	// User name for SQL authentication.
	Username pulumi.StringInput `pulumi:"username"`
}

func (SqlAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlAuth)(nil)).Elem()
}

func (i SqlAuthArgs) ToSqlAuthOutput() SqlAuthOutput {
	return i.ToSqlAuthOutputWithContext(context.Background())
}

func (i SqlAuthArgs) ToSqlAuthOutputWithContext(ctx context.Context) SqlAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlAuthOutput)
}

type SqlAuthOutput struct{ *pulumi.OutputState }

func (SqlAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlAuth)(nil)).Elem()
}

func (o SqlAuthOutput) ToSqlAuthOutput() SqlAuthOutput {
	return o
}

func (o SqlAuthOutput) ToSqlAuthOutputWithContext(ctx context.Context) SqlAuthOutput {
	return o
}

// Password for SQL authentication.
func (o SqlAuthOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v SqlAuth) string { return v.Password }).(pulumi.StringOutput)
}

// User name for SQL authentication.
func (o SqlAuthOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v SqlAuth) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureAuthInput)(nil)).Elem(), AzureAuthArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlAuthInput)(nil)).Elem(), SqlAuthArgs{})
	pulumi.RegisterOutputType(AzureAuthOutput{})
	pulumi.RegisterOutputType(SqlAuthOutput{})
}
