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
    /// <summary>
    /// Manages single DB schema.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Mssql = Pulumi.Mssql;
    /// using Mssql = Pulumiverse.Mssql;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleDatabase = Mssql.GetDatabase.Invoke(new()
    ///     {
    ///         Name = "example",
    ///     });
    /// 
    ///     var owner = Mssql.GetSqlUser.Invoke(new()
    ///     {
    ///         Name = "example_user",
    ///     });
    /// 
    ///     var exampleSchema = new Mssql.Schema("exampleSchema", new()
    ///     {
    ///         Name = "example",
    ///         DatabaseId = exampleDatabase.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         OwnerId = owner.Apply(getSqlUserResult =&gt; getSqlUserResult.Id),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import using &lt;db_id&gt;/&lt;schema_id&gt; - can be retrieved using `SELECT CONCAT(DB_ID(), '/', SCHEMA_ID('&lt;schema_name&gt;'))`
    /// 
    /// ```sh
    ///  $ pulumi import mssql:index/schema:Schema example '7/5'
    /// ```
    /// </summary>
    [MssqlResourceType("mssql:index/schema:Schema")]
    public partial class Schema : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// Schema name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;


        /// <summary>
        /// Create a Schema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Schema(string name, SchemaArgs args, CustomResourceOptions? options = null)
            : base("mssql:index/schema:Schema", name, args ?? new SchemaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Schema(string name, Input<string> id, SchemaState? state = null, CustomResourceOptions? options = null)
            : base("mssql:index/schema:Schema", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-mssql",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Schema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Schema Get(string name, Input<string> id, SchemaState? state = null, CustomResourceOptions? options = null)
        {
            return new Schema(name, id, state, options);
        }
    }

    public sealed class SchemaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// Schema name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        public SchemaArgs()
        {
        }
        public static new SchemaArgs Empty => new SchemaArgs();
    }

    public sealed class SchemaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`. Defaults to ID of `master`.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// Schema name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        public SchemaState()
        {
        }
        public static new SchemaState Empty => new SchemaState();
    }
}