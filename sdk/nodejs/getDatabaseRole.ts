// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Obtains information about single database role.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const master = mssql.getDatabase({
 *     name: "master",
 * });
 * const example = master.then(master => mssql.getDatabaseRole({
 *     name: "public",
 *     databaseId: master.id,
 * }));
 * export const id = example.then(example => example.id);
 * ```
 */
export function getDatabaseRole(args: GetDatabaseRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getDatabaseRole:getDatabaseRole", {
        "databaseId": args.databaseId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseRole.
 */
export interface GetDatabaseRoleArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    databaseId?: string;
    /**
     * Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
     */
    name: string;
}

/**
 * A collection of values returned by getDatabaseRole.
 */
export interface GetDatabaseRoleResult {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    readonly databaseId?: string;
    /**
     * `<database_id>/<role_id>`. Role ID can be retrieved using `SELECT DATABASE_PRINCIPAL_ID('<role_name>')`
     */
    readonly id: string;
    /**
     * Set of role members
     */
    readonly members: outputs.GetDatabaseRoleMember[];
    /**
     * Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
     */
    readonly name: string;
    /**
     * ID of another database role or user owning this role. Can be retrieved using `mssql.DatabaseRole` or `mssql.SqlUser`.
     */
    readonly ownerId: string;
}
/**
 * Obtains information about single database role.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const master = mssql.getDatabase({
 *     name: "master",
 * });
 * const example = master.then(master => mssql.getDatabaseRole({
 *     name: "public",
 *     databaseId: master.id,
 * }));
 * export const id = example.then(example => example.id);
 * ```
 */
export function getDatabaseRoleOutput(args: GetDatabaseRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("mssql:index/getDatabaseRole:getDatabaseRole", {
        "databaseId": args.databaseId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseRole.
 */
export interface GetDatabaseRoleOutputArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
     */
    name: pulumi.Input<string>;
}
