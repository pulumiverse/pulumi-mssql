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
    public sealed class GetDatabaseRoleMemberResult
    {
        /// <summary>
        /// `&lt;database_id&gt;/&lt;member_id&gt;`. Member ID can be retrieved using `SELECT DATABASE*PRINCIPAL*ID('\n\n')
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the database principal.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// One of: `SQL_USER`, `DATABASE_ROLE`, `AZUREAD_USER`
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDatabaseRoleMemberResult(
            string id,

            string name,

            string type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}