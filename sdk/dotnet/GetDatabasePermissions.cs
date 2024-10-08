// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Mssql
{
    public static class GetDatabasePermissions
    {
        /// <summary>
        /// Returns all permissions granted in a DB to given principal
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Mssql = Pulumi.Mssql;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleDatabase = Mssql.GetDatabase.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        ///     var exampleSqlUser = Mssql.GetSqlUser.Invoke(new()
        ///     {
        ///         Name = "example_user",
        ///         DatabaseId = exampleDatabase.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
        ///     });
        /// 
        ///     var exampleDatabasePermissions = Mssql.GetDatabasePermissions.Invoke(new()
        ///     {
        ///         PrincipalId = exampleSqlUser.Apply(getSqlUserResult =&gt; getSqlUserResult.Id),
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["permissions"] = exampleDatabasePermissions.Apply(getDatabasePermissionsResult =&gt; getDatabasePermissionsResult.Permissions),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatabasePermissionsResult> InvokeAsync(GetDatabasePermissionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasePermissionsResult>("mssql:index/getDatabasePermissions:getDatabasePermissions", args ?? new GetDatabasePermissionsArgs(), options.WithDefaults());

        /// <summary>
        /// Returns all permissions granted in a DB to given principal
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Mssql = Pulumi.Mssql;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var exampleDatabase = Mssql.GetDatabase.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        ///     var exampleSqlUser = Mssql.GetSqlUser.Invoke(new()
        ///     {
        ///         Name = "example_user",
        ///         DatabaseId = exampleDatabase.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
        ///     });
        /// 
        ///     var exampleDatabasePermissions = Mssql.GetDatabasePermissions.Invoke(new()
        ///     {
        ///         PrincipalId = exampleSqlUser.Apply(getSqlUserResult =&gt; getSqlUserResult.Id),
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["permissions"] = exampleDatabasePermissions.Apply(getDatabasePermissionsResult =&gt; getDatabasePermissionsResult.Permissions),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabasePermissionsResult> Invoke(GetDatabasePermissionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasePermissionsResult>("mssql:index/getDatabasePermissions:getDatabasePermissions", args ?? new GetDatabasePermissionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabasePermissionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// `&lt;database_id&gt;/&lt;principal_id&gt;`. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`.
        /// </summary>
        [Input("principalId", required: true)]
        public string PrincipalId { get; set; } = null!;

        public GetDatabasePermissionsArgs()
        {
        }
        public static new GetDatabasePermissionsArgs Empty => new GetDatabasePermissionsArgs();
    }

    public sealed class GetDatabasePermissionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// `&lt;database_id&gt;/&lt;principal_id&gt;`. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`.
        /// </summary>
        [Input("principalId", required: true)]
        public Input<string> PrincipalId { get; set; } = null!;

        public GetDatabasePermissionsInvokeArgs()
        {
        }
        public static new GetDatabasePermissionsInvokeArgs Empty => new GetDatabasePermissionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabasePermissionsResult
    {
        /// <summary>
        /// `&lt;database_id&gt;/&lt;principal_id&gt;`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set of permissions granted to the principal
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabasePermissionsPermissionResult> Permissions;
        /// <summary>
        /// `&lt;database_id&gt;/&lt;principal_id&gt;`. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`.
        /// </summary>
        public readonly string PrincipalId;

        [OutputConstructor]
        private GetDatabasePermissionsResult(
            string id,

            ImmutableArray<Outputs.GetDatabasePermissionsPermissionResult> permissions,

            string principalId)
        {
            Id = id;
            Permissions = permissions;
            PrincipalId = principalId;
        }
    }
}
