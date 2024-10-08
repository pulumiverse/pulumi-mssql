// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Retrieves arbitrary SQL query result.
 *
 * > **Note** This data source is meant to be an escape hatch for all cases not supported by the provider's data sources. Whenever possible, use dedicated data sources, which offer better plan, validation and error reporting.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const test = mssql.getDatabase({
 *     name: "test",
 * });
 * const column = test.then(test => mssql.getQuery({
 *     databaseId: test.id,
 *     query: "SELECT [column_id], [name] FROM sys.columns WHERE [object_id] = OBJECT_ID('test_table')",
 * }));
 * export const columnNames = column.then(column => column.results.map(__item => __item.name));
 * ```
 */
export function getQuery(args: GetQueryArgs, opts?: pulumi.InvokeOptions): Promise<GetQueryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getQuery:getQuery", {
        "databaseId": args.databaseId,
        "query": args.query,
    }, opts);
}

/**
 * A collection of arguments for invoking getQuery.
 */
export interface GetQueryArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: string;
    /**
     * SQL query returning single result set, with any number of rows, where all columns are strings
     */
    query: string;
}

/**
 * A collection of values returned by getQuery.
 */
export interface GetQueryResult {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    readonly databaseId: string;
    readonly id: string;
    /**
     * SQL query returning single result set, with any number of rows, where all columns are strings
     */
    readonly query: string;
    /**
     * Results of the SQL query, represented as list of maps, where the map key corresponds to column name and the value is the value of column in given row.
     */
    readonly results: {[key: string]: string}[];
}
/**
 * Retrieves arbitrary SQL query result.
 *
 * > **Note** This data source is meant to be an escape hatch for all cases not supported by the provider's data sources. Whenever possible, use dedicated data sources, which offer better plan, validation and error reporting.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 *
 * const test = mssql.getDatabase({
 *     name: "test",
 * });
 * const column = test.then(test => mssql.getQuery({
 *     databaseId: test.id,
 *     query: "SELECT [column_id], [name] FROM sys.columns WHERE [object_id] = OBJECT_ID('test_table')",
 * }));
 * export const columnNames = column.then(column => column.results.map(__item => __item.name));
 * ```
 */
export function getQueryOutput(args: GetQueryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQueryResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("mssql:index/getQuery:getQuery", {
        "databaseId": args.databaseId,
        "query": args.query,
    }, opts);
}

/**
 * A collection of arguments for invoking getQuery.
 */
export interface GetQueryOutputArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: pulumi.Input<string>;
    /**
     * SQL query returning single result set, with any number of rows, where all columns are strings
     */
    query: pulumi.Input<string>;
}
