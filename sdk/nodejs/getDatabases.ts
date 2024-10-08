// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Obtains information about all databases found in SQL Server instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const example = mssql.getDatabases({});
 * export const databases = example.then(example => example.databases);
 * ```
 */
export function getDatabases(opts?: pulumi.InvokeOptions): Promise<GetDatabasesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getDatabases:getDatabases", {
    }, opts);
}

/**
 * A collection of values returned by getDatabases.
 */
export interface GetDatabasesResult {
    /**
     * Set of database objects
     */
    readonly databases: outputs.GetDatabasesDatabase[];
    /**
     * ID of the resource used only internally by the provider.
     */
    readonly id: string;
}
/**
 * Obtains information about all databases found in SQL Server instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const example = mssql.getDatabases({});
 * export const databases = example.then(example => example.databases);
 * ```
 */
export function getDatabasesOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabasesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("mssql:index/getDatabases:getDatabases", {
    }, opts);
}
