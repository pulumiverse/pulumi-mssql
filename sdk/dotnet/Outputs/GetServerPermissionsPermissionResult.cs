// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Mssql.Outputs
{

    [OutputType]
    public sealed class GetServerPermissionsPermissionResult
    {
        /// <summary>
        /// Name of server-level SQL permission. For full list of supported permissions see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-server-permissions-transact-sql?view=azuresqldb-current#remarks)
        /// </summary>
        public readonly string Permission;
        /// <summary>
        /// When set to `true`, `principal_id` will be allowed to grant the `permission` to other principals.
        /// </summary>
        public readonly bool WithGrantOption;

        [OutputConstructor]
        private GetServerPermissionsPermissionResult(
            string permission,

            bool withGrantOption)
        {
            Permission = permission;
            WithGrantOption = withGrantOption;
        }
    }
}