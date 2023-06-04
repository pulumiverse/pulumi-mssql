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
    public static class GetSchema
    {
        /// <summary>
        /// Retrieves information about DB schema.
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
        ///     var example = Mssql.GetDatabase.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        ///     var byName = Mssql.GetSchema.Invoke(new()
        ///     {
        ///         DatabaseId = example.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
        ///         Name = "dbo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSchemaResult> InvokeAsync(GetSchemaArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSchemaResult>("mssql:index/getSchema:getSchema", args ?? new GetSchemaArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves information about DB schema.
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
        ///     var example = Mssql.GetDatabase.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        ///     var byName = Mssql.GetSchema.Invoke(new()
        ///     {
        ///         DatabaseId = example.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
        ///         Name = "dbo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSchemaResult> Invoke(GetSchemaInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSchemaResult>("mssql:index/getSchema:getSchema", args ?? new GetSchemaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSchemaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`.
        /// </summary>
        [Input("databaseId")]
        public string? DatabaseId { get; set; }

        /// <summary>
        /// `&lt;database_id&gt;/&lt;schema_id&gt;`. Schema ID can be retrieved using `SELECT SCHEMA_ID('&lt;schema_name&gt;')`. Either `id` or `name` must be provided.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Schema name. Either `id` or `name` must be provided.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetSchemaArgs()
        {
        }
        public static new GetSchemaArgs Empty => new GetSchemaArgs();
    }

    public sealed class GetSchemaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// `&lt;database_id&gt;/&lt;schema_id&gt;`. Schema ID can be retrieved using `SELECT SCHEMA_ID('&lt;schema_name&gt;')`. Either `id` or `name` must be provided.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Schema name. Either `id` or `name` must be provided.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetSchemaInvokeArgs()
        {
        }
        public static new GetSchemaInvokeArgs Empty => new GetSchemaInvokeArgs();
    }


    [OutputType]
    public sealed class GetSchemaResult
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`.
        /// </summary>
        public readonly string DatabaseId;
        /// <summary>
        /// `&lt;database_id&gt;/&lt;schema_id&gt;`. Schema ID can be retrieved using `SELECT SCHEMA_ID('&lt;schema_name&gt;')`. Either `id` or `name` must be provided.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Schema name. Either `id` or `name` must be provided.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
        /// </summary>
        public readonly string OwnerId;

        [OutputConstructor]
        private GetSchemaResult(
            string databaseId,

            string id,

            string name,

            string ownerId)
        {
            DatabaseId = databaseId;
            Id = id;
            Name = name;
            OwnerId = ownerId;
        }
    }
}
