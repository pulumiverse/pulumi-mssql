// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mssql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-mssql/sdk/go/mssql/internal"
)

// Obtains information about single SQL login.
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
//			sa, err := mssql.LookupSqlLogin(ctx, &mssql.LookupSqlLoginArgs{
//				Name: "sa",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("id", sa.Id)
//			ctx.Export("dbId", sa.DefaultDatabaseId)
//			return nil
//		})
//	}
//
// ```
func LookupSqlLogin(ctx *pulumi.Context, args *LookupSqlLoginArgs, opts ...pulumi.InvokeOption) (*LookupSqlLoginResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSqlLoginResult
	err := ctx.Invoke("mssql:index/getSqlLogin:getSqlLogin", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSqlLogin.
type LookupSqlLoginArgs struct {
	// Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
	Name string `pulumi:"name"`
}

// A collection of values returned by getSqlLogin.
type LookupSqlLoginResult struct {
	// When `true`, password expiration policy is enforced for this login.
	CheckPasswordExpiration bool `pulumi:"checkPasswordExpiration"`
	// When `true`, the Windows password policies of the computer on which SQL Server is running are enforced on this login.
	CheckPasswordPolicy bool `pulumi:"checkPasswordPolicy"`
	// ID of login's default DB. The ID can be retrieved using `Database` data resource.
	DefaultDatabaseId string `pulumi:"defaultDatabaseId"`
	// Default language assigned to login.
	DefaultLanguage string `pulumi:"defaultLanguage"`
	// Login SID. Can be retrieved using `SELECT SUSER_SID('<login_name>')`.
	Id string `pulumi:"id"`
	// When true, password change will be forced on first logon.
	MustChangePassword bool `pulumi:"mustChangePassword"`
	// Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
	Name string `pulumi:"name"`
	// ID used to reference SQL Login in other resources, e.g. `serverRole`. Can be retrieved from `sys.sql_logins`.
	PrincipalId string `pulumi:"principalId"`
}

func LookupSqlLoginOutput(ctx *pulumi.Context, args LookupSqlLoginOutputArgs, opts ...pulumi.InvokeOption) LookupSqlLoginResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSqlLoginResult, error) {
			args := v.(LookupSqlLoginArgs)
			r, err := LookupSqlLogin(ctx, &args, opts...)
			var s LookupSqlLoginResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSqlLoginResultOutput)
}

// A collection of arguments for invoking getSqlLogin.
type LookupSqlLoginOutputArgs struct {
	// Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupSqlLoginOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlLoginArgs)(nil)).Elem()
}

// A collection of values returned by getSqlLogin.
type LookupSqlLoginResultOutput struct{ *pulumi.OutputState }

func (LookupSqlLoginResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSqlLoginResult)(nil)).Elem()
}

func (o LookupSqlLoginResultOutput) ToLookupSqlLoginResultOutput() LookupSqlLoginResultOutput {
	return o
}

func (o LookupSqlLoginResultOutput) ToLookupSqlLoginResultOutputWithContext(ctx context.Context) LookupSqlLoginResultOutput {
	return o
}

func (o LookupSqlLoginResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSqlLoginResult] {
	return pulumix.Output[LookupSqlLoginResult]{
		OutputState: o.OutputState,
	}
}

// When `true`, password expiration policy is enforced for this login.
func (o LookupSqlLoginResultOutput) CheckPasswordExpiration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSqlLoginResult) bool { return v.CheckPasswordExpiration }).(pulumi.BoolOutput)
}

// When `true`, the Windows password policies of the computer on which SQL Server is running are enforced on this login.
func (o LookupSqlLoginResultOutput) CheckPasswordPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSqlLoginResult) bool { return v.CheckPasswordPolicy }).(pulumi.BoolOutput)
}

// ID of login's default DB. The ID can be retrieved using `Database` data resource.
func (o LookupSqlLoginResultOutput) DefaultDatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlLoginResult) string { return v.DefaultDatabaseId }).(pulumi.StringOutput)
}

// Default language assigned to login.
func (o LookupSqlLoginResultOutput) DefaultLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlLoginResult) string { return v.DefaultLanguage }).(pulumi.StringOutput)
}

// Login SID. Can be retrieved using `SELECT SUSER_SID('<login_name>')`.
func (o LookupSqlLoginResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlLoginResult) string { return v.Id }).(pulumi.StringOutput)
}

// When true, password change will be forced on first logon.
func (o LookupSqlLoginResultOutput) MustChangePassword() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSqlLoginResult) bool { return v.MustChangePassword }).(pulumi.BoolOutput)
}

// Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
func (o LookupSqlLoginResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlLoginResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID used to reference SQL Login in other resources, e.g. `serverRole`. Can be retrieved from `sys.sql_logins`.
func (o LookupSqlLoginResultOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSqlLoginResult) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSqlLoginResultOutput{})
}
