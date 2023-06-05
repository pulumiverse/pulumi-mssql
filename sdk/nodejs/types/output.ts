// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetDatabasePermissionsPermission {
    /**
     * Name of database-level SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-database-permissions-transact-sql?view=azuresqldb-current#remarks)
     */
    permission: string;
    /**
     * When set to `true`, `principalId` will be allowed to grant the `permission` to other principals.
     */
    withGrantOption: boolean;
}

export interface GetDatabaseRoleMember {
    /**
     * `<database_id>/<member_id>`. Member ID can be retrieved using `SELECT DATABASE*PRINCIPAL*ID('\n\n')
     */
    id: string;
    /**
     * Name of the database principal.
     */
    name: string;
    /**
     * One of: `SQL_USER`, `DATABASE_ROLE`, `AZUREAD_USER`
     */
    type: string;
}

export interface GetDatabaseRolesRole {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: string;
    /**
     * `<database_id>/<role_id>`. Role ID can be retrieved using `SELECT DATABASE_PRINCIPAL_ID('<role_name>')`
     */
    id: string;
    /**
     * Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
     */
    name: string;
    /**
     * ID of another database role or user owning this role. Can be retrieved using `mssql.DatabaseRole` or `mssql.SqlUser`.
     */
    ownerId: string;
}

export interface GetDatabasesDatabase {
    /**
     * Default collation name. Can be either a Windows collation name or a SQL collation name.
     */
    collation: string;
    /**
     * Database ID. Can be retrieved using `SELECT DB_ID('<db_name>')`.
     */
    id: string;
    /**
     * Database name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers).
     */
    name: string;
}

export interface GetSchemaPermissionsPermission {
    /**
     * Name of schema SQL permission. For full list of supported permissions, see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-schema-permissions-transact-sql?view=azuresqldb-current#remarks)
     */
    permission: string;
    /**
     * When set to `true`, `principalId` will be allowed to grant the `permission` to other principals.
     */
    withGrantOption: boolean;
}

export interface GetSchemasSchema {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: string;
    /**
     * `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`.
     */
    id: string;
    /**
     * Schema name.
     */
    name: string;
    /**
     * ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
     */
    ownerId: string;
}

export interface GetServerPermissionsPermission {
    /**
     * Name of server-level SQL permission. For full list of supported permissions see [docs](https://learn.microsoft.com/en-us/sql/t-sql/statements/grant-server-permissions-transact-sql?view=azuresqldb-current#remarks)
     */
    permission: string;
    /**
     * When set to `true`, `principalId` will be allowed to grant the `permission` to other principals.
     */
    withGrantOption: boolean;
}

export interface GetServerRoleMember {
    /**
     * ID of the member principal
     */
    id: string;
    /**
     * Name of the server principal
     */
    name: string;
    /**
     * One of: `SQL_LOGIN`, `SERVER_ROLE`
     */
    type: string;
}

export interface GetServerRolesRole {
    /**
     * Role principal ID.
     */
    id: string;
    /**
     * Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
     */
    name: string;
    /**
     * ID of another server role or login owning this role. Can be retrieved using `mssql.ServerRole` or `mssql.SqlLogin`.
     */
    ownerId: string;
}

export interface GetSqlLoginsLogin {
    /**
     * When `true`, password expiration policy is enforced for this login.
     */
    checkPasswordExpiration: boolean;
    /**
     * When `true`, the Windows password policies of the computer on which SQL Server is running are enforced on this login.
     */
    checkPasswordPolicy: boolean;
    /**
     * ID of login's default DB. The ID can be retrieved using `mssql.Database` data resource.
     */
    defaultDatabaseId: string;
    /**
     * Default language assigned to login.
     */
    defaultLanguage: string;
    /**
     * Login SID. Can be retrieved using `SELECT SUSER_SID('<login_name>')`.
     */
    id: string;
    /**
     * When true, password change will be forced on first logon.
     */
    mustChangePassword: boolean;
    /**
     * Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
     */
    name: string;
    /**
     * ID used to reference SQL Login in other resources, e.g. `serverRole`. Can be retrieved from `sys.sql_logins`.
     */
    principalId: string;
}

export interface GetSqlUsersUser {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: string;
    /**
     * `<database_id>/<user_id>`. User ID can be retrieved using `SELECT DATABASE_PRINCIPAL_ID('<user_name>')`.
     */
    id: string;
    /**
     * SID of SQL login. Can be retrieved using `mssql.SqlLogin` or `SELECT SUSER_SID('<login_name>')`.
     */
    loginId: string;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    name: string;
}

export namespace config {
    export interface AzureAuth {
        clientId?: string;
        clientSecret?: string;
        tenantId?: string;
    }

    export interface SqlAuth {
        password: string;
        username: string;
    }

}
