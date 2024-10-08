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
    /// Managed database-level user mapped to Azure AD identity (user or group).
    /// 
    /// &gt; **Note** When using this resource, Azure SQL server managed identity does not need any [AzureAD role assignments](https://docs.microsoft.com/en-us/azure/azure-sql/database/authentication-aad-service-principal?view=azuresql).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
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
    ///     var exampleUser = AzureAD.GetUser.Invoke(new()
    ///     {
    ///         UserPrincipalName = "user@example.com",
    ///     });
    /// 
    ///     var exampleAzureadUser = new Mssql.AzureadUser("exampleAzureadUser", new()
    ///     {
    ///         DatabaseId = exampleDatabase.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         UserObjectId = exampleUser.Apply(getUserResult =&gt; getUserResult.ObjectId),
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["userId"] = exampleAzureadUser.Id,
    ///     };
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// import using &lt;db_id&gt;/&lt;user_id&gt; - can be retrieved using `SELECT CONCAT(DB_ID(), '/', principal_id) FROM sys.database_principals WHERE [name] = '&lt;username&gt;'`
    /// 
    /// ```sh
    /// $ pulumi import mssql:index/azureadUser:AzureadUser example '7/5'
    /// ```
    /// </summary>
    [MssqlResourceType("mssql:index/azureadUser:AzureadUser")]
    public partial class AzureadUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// User name. Cannot be longer than 128 chars.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure AD object_id of the user. This can be either regular user or a group.
        /// </summary>
        [Output("userObjectId")]
        public Output<string> UserObjectId { get; private set; } = null!;


        /// <summary>
        /// Create a AzureadUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AzureadUser(string name, AzureadUserArgs args, CustomResourceOptions? options = null)
            : base("mssql:index/azureadUser:AzureadUser", name, args ?? new AzureadUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AzureadUser(string name, Input<string> id, AzureadUserState? state = null, CustomResourceOptions? options = null)
            : base("mssql:index/azureadUser:AzureadUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AzureadUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AzureadUser Get(string name, Input<string> id, AzureadUserState? state = null, CustomResourceOptions? options = null)
        {
            return new AzureadUser(name, id, state, options);
        }
    }

    public sealed class AzureadUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        /// <summary>
        /// User name. Cannot be longer than 128 chars.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Azure AD object_id of the user. This can be either regular user or a group.
        /// </summary>
        [Input("userObjectId", required: true)]
        public Input<string> UserObjectId { get; set; } = null!;

        public AzureadUserArgs()
        {
        }
        public static new AzureadUserArgs Empty => new AzureadUserArgs();
    }

    public sealed class AzureadUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('&lt;db_name&gt;')`.
        /// </summary>
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// User name. Cannot be longer than 128 chars.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Azure AD object_id of the user. This can be either regular user or a group.
        /// </summary>
        [Input("userObjectId")]
        public Input<string>? UserObjectId { get; set; }

        public AzureadUserState()
        {
        }
        public static new AzureadUserState Empty => new AzureadUserState();
    }
}
