// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Obtains information about all SQL logins found in SQL Server instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const example = mssql.getSqlLogins({});
 * export const databases = example.then(example => example.logins);
 * ```
 */
export function getSqlLogins(opts?: pulumi.InvokeOptions): Promise<GetSqlLoginsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getSqlLogins:getSqlLogins", {
    }, opts);
}

/**
 * A collection of values returned by getSqlLogins.
 */
export interface GetSqlLoginsResult {
    /**
     * ID of the resource used only internally by the provider.
     */
    readonly id: string;
    /**
     * Set of SQL login objects
     */
    readonly logins: outputs.GetSqlLoginsLogin[];
}
/**
 * Obtains information about all SQL logins found in SQL Server instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const example = mssql.getSqlLogins({});
 * export const databases = example.then(example => example.logins);
 * ```
 */
export function getSqlLoginsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetSqlLoginsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("mssql:index/getSqlLogins:getSqlLogins", {
    }, opts);
}
