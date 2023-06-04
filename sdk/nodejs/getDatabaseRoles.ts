// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Obtains information about all roles defined in a database.
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
 * const example = master.then(master => mssql.getDatabaseRoles({
 *     databaseId: master.id,
 * }));
 * export const roles = example.then(example => example.roles);
 * ```
 */
export function getDatabaseRoles(args?: GetDatabaseRolesArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseRolesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getDatabaseRoles:getDatabaseRoles", {
        "databaseId": args.databaseId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseRoles.
 */
export interface GetDatabaseRolesArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    databaseId?: string;
}

/**
 * A collection of values returned by getDatabaseRoles.
 */
export interface GetDatabaseRolesResult {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    readonly databaseId?: string;
    /**
     * ID of the resource, equals to database ID
     */
    readonly id: string;
    /**
     * Set of database role objects
     */
    readonly roles: outputs.GetDatabaseRolesRole[];
}
/**
 * Obtains information about all roles defined in a database.
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
 * const example = master.then(master => mssql.getDatabaseRoles({
 *     databaseId: master.id,
 * }));
 * export const roles = example.then(example => example.roles);
 * ```
 */
export function getDatabaseRolesOutput(args?: GetDatabaseRolesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseRolesResult> {
    return pulumi.output(args).apply((a: any) => getDatabaseRoles(a, opts))
}

/**
 * A collection of arguments for invoking getDatabaseRoles.
 */
export interface GetDatabaseRolesOutputArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    databaseId?: pulumi.Input<string>;
}