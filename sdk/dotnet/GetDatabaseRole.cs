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
    public static class GetDatabaseRole
    {
        /// <summary>
        /// Obtains information about single database role.
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
        ///     var example = Mssql.GetDatabaseRole.Invoke(new()
        ///     {
        ///         Name = "public",
        ///         DatabaseId = master.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["id"] = example.Apply(getDatabaseRoleResult =&gt; getDatabaseRoleResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatabaseRoleResult> InvokeAsync(GetDatabaseRoleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseRoleResult>("mssql:index/getDatabaseRole:getDatabaseRole", args ?? new GetDatabaseRoleArgs(), options.WithDefaults());

        /// <summary>
        /// Obtains information about single database role.
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
        ///     var example = Mssql.GetDatabaseRole.Invoke(new()
        ///     {
        ///         Name = "public",
        ///         DatabaseId = master.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["id"] = example.Apply(getDatabaseRoleResult =&gt; getDatabaseRoleResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDatabaseRoleResult> Invoke(GetDatabaseRoleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseRoleResult>("mssql:index/getDatabaseRole:getDatabaseRole", args ?? new GetDatabaseRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseRoleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        [Input("databaseId")]
        public string? DatabaseId { get; set; }

        /// <summary>
        /// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDatabaseRoleArgs()
        {
        }
        public static new GetDatabaseRoleArgs Empty => new GetDatabaseRoleArgs();
    }

    public sealed class GetDatabaseRoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDatabaseRoleInvokeArgs()
        {
        }
        public static new GetDatabaseRoleInvokeArgs Empty => new GetDatabaseRoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseRoleResult
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        public readonly string? DatabaseId;
        /// <summary>
        /// `&lt;database_id&gt;/&lt;role_id&gt;`. Role ID can be retrieved using `SELECT DATABASE_PRINCIPAL_ID('&lt;role_name&gt;')`
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set of role members
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseRoleMemberResult> Members;
        /// <summary>
        /// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of another database role or user owning this role. Can be retrieved using `mssql.DatabaseRole` or `mssql.SqlUser`.
        /// </summary>
        public readonly string OwnerId;

        [OutputConstructor]
        private GetDatabaseRoleResult(
            string? databaseId,

            string id,

            ImmutableArray<Outputs.GetDatabaseRoleMemberResult> members,

            string name,

            string ownerId)
        {
            DatabaseId = databaseId;
            Id = id;
            Members = members;
            Name = name;
            OwnerId = ownerId;
        }
    }
}
