// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Obtains information about single SQL database user.
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
 * const example = master.then(master => mssql.getSqlUser({
 *     name: "dbo",
 *     databaseId: master.id,
 * }));
 * export const id = example.then(example => example.id);
 * ```
 */
export function getSqlUser(args: GetSqlUserArgs, opts?: pulumi.InvokeOptions): Promise<GetSqlUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getSqlUser:getSqlUser", {
        "databaseId": args.databaseId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSqlUser.
 */
export interface GetSqlUserArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId?: string;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    name: string;
}

/**
 * A collection of values returned by getSqlUser.
 */
export interface GetSqlUserResult {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    readonly databaseId: string;
    /**
     * `<database_id>/<user_id>`. User ID can be retrieved using `SELECT DATABASE_PRINCIPAL_ID('<user_name>')`.
     */
    readonly id: string;
    /**
     * SID of SQL login. Can be retrieved using `mssql.SqlLogin` or `SELECT SUSER_SID('<login_name>')`.
     */
    readonly loginId: string;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    readonly name: string;
}
/**
 * Obtains information about single SQL database user.
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
 * const example = master.then(master => mssql.getSqlUser({
 *     name: "dbo",
 *     databaseId: master.id,
 * }));
 * export const id = example.then(example => example.id);
 * ```
 */
export function getSqlUserOutput(args: GetSqlUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSqlUserResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("mssql:index/getSqlUser:getSqlUser", {
        "databaseId": args.databaseId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSqlUser.
 */
export interface GetSqlUserOutputArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    name: pulumi.Input<string>;
}
