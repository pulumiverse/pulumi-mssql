// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Returns all permissions granted in a schema to given principal
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const exampleDatabase = mssql.getDatabase({
 *     name: "example",
 * });
 * const exampleSqlUser = exampleDatabase.then(exampleDatabase => mssql.getSqlUser({
 *     name: "example_user",
 *     databaseId: exampleDatabase.id,
 * }));
 * const exampleSchema = exampleDatabase.then(exampleDatabase => mssql.getSchema({
 *     name: "example_schema",
 *     databaseId: exampleDatabase.id,
 * }));
 * const exampleSchemaPermissions = Promise.all([exampleSchema, exampleSqlUser]).then(([exampleSchema, exampleSqlUser]) => mssql.getSchemaPermissions({
 *     schemaId: exampleSchema.id,
 *     principalId: exampleSqlUser.id,
 * }));
 * export const permissions = exampleSchemaPermissions.then(exampleSchemaPermissions => exampleSchemaPermissions.permissions);
 * ```
 */
export function getSchemaPermissions(args: GetSchemaPermissionsArgs, opts?: pulumi.InvokeOptions): Promise<GetSchemaPermissionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getSchemaPermissions:getSchemaPermissions", {
        "principalId": args.principalId,
        "schemaId": args.schemaId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSchemaPermissions.
 */
export interface GetSchemaPermissionsArgs {
    /**
     * `<database_id>/<principal_id>`. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`.
     */
    principalId: string;
    /**
     * `<database_id>/<schema_id>`. Can be retrieved using `mssql.Schema`.
     */
    schemaId: string;
}

/**
 * A collection of values returned by getSchemaPermissions.
 */
export interface GetSchemaPermissionsResult {
    /**
     * `<database_id>/<schema_id>/<principal_id>`.
     */
    readonly id: string;
    /**
     * Set of permissions granted to the principal
     */
    readonly permissions: outputs.GetSchemaPermissionsPermission[];
    /**
     * `<database_id>/<principal_id>`. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`.
     */
    readonly principalId: string;
    /**
     * `<database_id>/<schema_id>`. Can be retrieved using `mssql.Schema`.
     */
    readonly schemaId: string;
}
/**
 * Returns all permissions granted in a schema to given principal
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const exampleDatabase = mssql.getDatabase({
 *     name: "example",
 * });
 * const exampleSqlUser = exampleDatabase.then(exampleDatabase => mssql.getSqlUser({
 *     name: "example_user",
 *     databaseId: exampleDatabase.id,
 * }));
 * const exampleSchema = exampleDatabase.then(exampleDatabase => mssql.getSchema({
 *     name: "example_schema",
 *     databaseId: exampleDatabase.id,
 * }));
 * const exampleSchemaPermissions = Promise.all([exampleSchema, exampleSqlUser]).then(([exampleSchema, exampleSqlUser]) => mssql.getSchemaPermissions({
 *     schemaId: exampleSchema.id,
 *     principalId: exampleSqlUser.id,
 * }));
 * export const permissions = exampleSchemaPermissions.then(exampleSchemaPermissions => exampleSchemaPermissions.permissions);
 * ```
 */
export function getSchemaPermissionsOutput(args: GetSchemaPermissionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSchemaPermissionsResult> {
    return pulumi.output(args).apply((a: any) => getSchemaPermissions(a, opts))
}

/**
 * A collection of arguments for invoking getSchemaPermissions.
 */
export interface GetSchemaPermissionsOutputArgs {
    /**
     * `<database_id>/<principal_id>`. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`.
     */
    principalId: pulumi.Input<string>;
    /**
     * `<database_id>/<schema_id>`. Can be retrieved using `mssql.Schema`.
     */
    schemaId: pulumi.Input<string>;
}
