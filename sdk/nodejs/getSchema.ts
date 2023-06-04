// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves information about DB schema.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const example = mssql.getDatabase({
 *     name: "example",
 * });
 * const byName = example.then(example => mssql.getSchema({
 *     databaseId: example.id,
 *     name: "dbo",
 * }));
 * ```
 */
export function getSchema(args?: GetSchemaArgs, opts?: pulumi.InvokeOptions): Promise<GetSchemaResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getSchema:getSchema", {
        "databaseId": args.databaseId,
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSchema.
 */
export interface GetSchemaArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId?: string;
    /**
     * `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`. Either `id` or `name` must be provided.
     */
    id?: string;
    /**
     * Schema name. Either `id` or `name` must be provided.
     */
    name?: string;
}

/**
 * A collection of values returned by getSchema.
 */
export interface GetSchemaResult {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    readonly databaseId: string;
    /**
     * `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`. Either `id` or `name` must be provided.
     */
    readonly id: string;
    /**
     * Schema name. Either `id` or `name` must be provided.
     */
    readonly name: string;
    /**
     * ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
     */
    readonly ownerId: string;
}
/**
 * Retrieves information about DB schema.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const example = mssql.getDatabase({
 *     name: "example",
 * });
 * const byName = example.then(example => mssql.getSchema({
 *     databaseId: example.id,
 *     name: "dbo",
 * }));
 * ```
 */
export function getSchemaOutput(args?: GetSchemaOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSchemaResult> {
    return pulumi.output(args).apply((a: any) => getSchema(a, opts))
}

/**
 * A collection of arguments for invoking getSchema.
 */
export interface GetSchemaOutputArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * `<database_id>/<schema_id>`. Schema ID can be retrieved using `SELECT SCHEMA_ID('<schema_name>')`. Either `id` or `name` must be provided.
     */
    id?: pulumi.Input<string>;
    /**
     * Schema name. Either `id` or `name` must be provided.
     */
    name?: pulumi.Input<string>;
}