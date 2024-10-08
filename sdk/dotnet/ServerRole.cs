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
    /// Manages server-level role.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Mssql = Pulumiverse.Mssql;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var owner = new Mssql.ServerRole("owner");
    /// 
    ///     var example = new Mssql.ServerRole("example", new()
    ///     {
    ///         OwnerId = owner.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import using &lt;role_id&gt; - can be retrieved using `SELECT [principal_id] FROM sys.server_principals WHERE [name]='&lt;role_name&gt;'`
    /// 
    /// ```sh
    /// $ pulumi import mssql:index/serverRole:ServerRole example 7
    /// ```
    /// </summary>
    [MssqlResourceType("mssql:index/serverRole:ServerRole")]
    public partial class ServerRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of another server role or login owning this role. Can be retrieved using `mssql.ServerRole` or `mssql.SqlLogin`.
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;


        /// <summary>
        /// Create a ServerRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerRole(string name, ServerRoleArgs? args = null, CustomResourceOptions? options = null)
            : base("mssql:index/serverRole:ServerRole", name, args ?? new ServerRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerRole(string name, Input<string> id, ServerRoleState? state = null, CustomResourceOptions? options = null)
            : base("mssql:index/serverRole:ServerRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerRole Get(string name, Input<string> id, ServerRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerRole(name, id, state, options);
        }
    }

    public sealed class ServerRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of another server role or login owning this role. Can be retrieved using `mssql.ServerRole` or `mssql.SqlLogin`.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        public ServerRoleArgs()
        {
        }
        public static new ServerRoleArgs Empty => new ServerRoleArgs();
    }

    public sealed class ServerRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of another server role or login owning this role. Can be retrieved using `mssql.ServerRole` or `mssql.SqlLogin`.
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        public ServerRoleState()
        {
        }
        public static new ServerRoleState Empty => new ServerRoleState();
    }
}
