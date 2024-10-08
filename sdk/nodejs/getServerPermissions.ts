// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Returns all permissions grated to given principal
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const exampleSqlLogin = mssql.getSqlLogin({
 *     name: "example_login",
 * });
 * const exampleServerPermissions = exampleSqlLogin.then(exampleSqlLogin => mssql.getServerPermissions({
 *     principalId: exampleSqlLogin.principalId,
 * }));
 * export const permissions = exampleServerPermissions.then(exampleServerPermissions => exampleServerPermissions.permissions);
 * ```
 */
export function getServerPermissions(args: GetServerPermissionsArgs, opts?: pulumi.InvokeOptions): Promise<GetServerPermissionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getServerPermissions:getServerPermissions", {
        "principalId": args.principalId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerPermissions.
 */
export interface GetServerPermissionsArgs {
    /**
     * ID of the principal who will be granted `permission`. Can be retrieved using `mssql.ServerRole` or `mssql.SqlLogin`.
     */
    principalId: string;
}

/**
 * A collection of values returned by getServerPermissions.
 */
export interface GetServerPermissionsResult {
    /**
     * Equals to `principalId`.
     */
    readonly id: string;
    /**
     * Set of permissions granted to the principal
     */
    readonly permissions: outputs.GetServerPermissionsPermission[];
    /**
     * ID of the principal who will be granted `permission`. Can be retrieved using `mssql.ServerRole` or `mssql.SqlLogin`.
     */
    readonly principalId: string;
}
/**
 * Returns all permissions grated to given principal
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const exampleSqlLogin = mssql.getSqlLogin({
 *     name: "example_login",
 * });
 * const exampleServerPermissions = exampleSqlLogin.then(exampleSqlLogin => mssql.getServerPermissions({
 *     principalId: exampleSqlLogin.principalId,
 * }));
 * export const permissions = exampleServerPermissions.then(exampleServerPermissions => exampleServerPermissions.permissions);
 * ```
 */
export function getServerPermissionsOutput(args: GetServerPermissionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServerPermissionsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("mssql:index/getServerPermissions:getServerPermissions", {
        "principalId": args.principalId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerPermissions.
 */
export interface GetServerPermissionsOutputArgs {
    /**
     * ID of the principal who will be granted `permission`. Can be retrieved using `mssql.ServerRole` or `mssql.SqlLogin`.
     */
    principalId: pulumi.Input<string>;
}
