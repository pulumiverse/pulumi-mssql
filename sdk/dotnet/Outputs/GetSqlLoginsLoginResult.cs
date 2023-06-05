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
    public sealed class GetSqlLoginsLoginResult
    {
        /// <summary>
        /// When `true`, password expiration policy is enforced for this login.
        /// </summary>
        public readonly bool CheckPasswordExpiration;
        /// <summary>
        /// When `true`, the Windows password policies of the computer on which SQL Server is running are enforced on this login.
        /// </summary>
        public readonly bool CheckPasswordPolicy;
        /// <summary>
        /// ID of login's default DB. The ID can be retrieved using `mssql.Database` data resource.
        /// </summary>
        public readonly string DefaultDatabaseId;
        /// <summary>
        /// Default language assigned to login.
        /// </summary>
        public readonly string DefaultLanguage;
        /// <summary>
        /// Login SID. Can be retrieved using `SELECT SUSER_SID('&lt;login_name&gt;')`.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// When true, password change will be forced on first logon.
        /// </summary>
        public readonly bool MustChangePassword;
        /// <summary>
        /// Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ID used to reference SQL Login in other resources, e.g. `server_role`. Can be retrieved from `sys.sql_logins`.
        /// </summary>
        public readonly string PrincipalId;

        [OutputConstructor]
        private GetSqlLoginsLoginResult(
            bool checkPasswordExpiration,

            bool checkPasswordPolicy,

            string defaultDatabaseId,

            string defaultLanguage,

            string id,

            bool mustChangePassword,

            string name,

            string principalId)
        {
            CheckPasswordExpiration = checkPasswordExpiration;
            CheckPasswordPolicy = checkPasswordPolicy;
            DefaultDatabaseId = defaultDatabaseId;
            DefaultLanguage = defaultLanguage;
            Id = id;
            MustChangePassword = mustChangePassword;
            Name = name;
            PrincipalId = principalId;
        }
    }
}
