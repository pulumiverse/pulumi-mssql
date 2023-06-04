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
    public static class GetDatabaseRoles
    {
        /// <summary>
        /// Obtains information about all roles defined in a database.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Mssql = Pulumi.Mssql;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var master = Mssql.GetDatabase.Invoke(new()
        ///     {
        ///         Name = "master",
        ///     });
        /// 
        ///     var example = Mssql.GetDatabaseRoles.Invoke(new()
        ///     {
        ///         DatabaseId = master.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["roles"] = example.Apply(getDatabaseRolesResult =&gt; getDatabaseRolesResult.Roles),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseRolesResult> InvokeAsync(GetDatabaseRolesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseRolesResult>("mssql:index/getDatabaseRoles:getDatabaseRoles", args ?? new GetDatabaseRolesArgs(), options.WithDefaults());

        /// <summary>
        /// Obtains information about all roles defined in a database.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Mssql = Pulumi.Mssql;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var master = Mssql.GetDatabase.Invoke(new()
        ///     {
        ///         Name = "master",
        ///     });
        /// 
        ///     var example = Mssql.GetDatabaseRoles.Invoke(new()
        ///     {
        ///         DatabaseId = master.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["roles"] = example.Apply(getDatabaseRolesResult =&gt; getDatabaseRolesResult.Roles),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatabaseRolesResult> Invoke(GetDatabaseRolesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseRolesResult>("mssql:index/getDatabaseRoles:getDatabaseRoles", args ?? new GetDatabaseRolesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseRolesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        [Input("databaseId")]
        public string? DatabaseId { get; set; }

        public GetDatabaseRolesArgs()
        {
        }
        public static new GetDatabaseRolesArgs Empty => new GetDatabaseRolesArgs();
    }

    public sealed class GetDatabaseRolesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        public GetDatabaseRolesInvokeArgs()
        {
        }
        public static new GetDatabaseRolesInvokeArgs Empty => new GetDatabaseRolesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseRolesResult
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        public readonly string? DatabaseId;
        /// <summary>
        /// ID of the resource, equals to database ID
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set of database role objects
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseRolesRoleResult> Roles;

        [OutputConstructor]
        private GetDatabaseRolesResult(
            string? databaseId,

            string id,

            ImmutableArray<Outputs.GetDatabaseRolesRoleResult> roles)
        {
            DatabaseId = databaseId;
            Id = id;
            Roles = roles;
        }
    }
}