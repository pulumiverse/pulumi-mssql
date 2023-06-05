// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Obtains information about single Azure AD Service Principal database user.
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
 * const exampleAzureadServicePrincipal = exampleDatabase.then(exampleDatabase => mssql.getAzureadServicePrincipal({
 *     name: "example",
 *     databaseId: exampleDatabase.id,
 * }));
 * export const appClientId = exampleAzureadServicePrincipal.then(exampleAzureadServicePrincipal => exampleAzureadServicePrincipal.clientId);
 * ```
 */
export function getAzureadServicePrincipal(args: GetAzureadServicePrincipalArgs, opts?: pulumi.InvokeOptions): Promise<GetAzureadServicePrincipalResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getAzureadServicePrincipal:getAzureadServicePrincipal", {
        "clientId": args.clientId,
        "databaseId": args.databaseId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getAzureadServicePrincipal.
 */
export interface GetAzureadServicePrincipalArgs {
    /**
     * Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
     */
    clientId?: string;
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: string;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    name?: string;
}

/**
 * A collection of values returned by getAzureadServicePrincipal.
 */
export interface GetAzureadServicePrincipalResult {
    /**
     * Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
     */
    readonly clientId: string;
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    readonly databaseId: string;
    /**
     * `<database_id>/<user_id>`. User ID can be retrieved using `sys.database_principals` view.
     */
    readonly id: string;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    readonly name: string;
}
/**
 * Obtains information about single Azure AD Service Principal database user.
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
 * const exampleAzureadServicePrincipal = exampleDatabase.then(exampleDatabase => mssql.getAzureadServicePrincipal({
 *     name: "example",
 *     databaseId: exampleDatabase.id,
 * }));
 * export const appClientId = exampleAzureadServicePrincipal.then(exampleAzureadServicePrincipal => exampleAzureadServicePrincipal.clientId);
 * ```
 */
export function getAzureadServicePrincipalOutput(args: GetAzureadServicePrincipalOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAzureadServicePrincipalResult> {
    return pulumi.output(args).apply((a: any) => getAzureadServicePrincipal(a, opts))
}

/**
 * A collection of arguments for invoking getAzureadServicePrincipal.
 */
export interface GetAzureadServicePrincipalOutputArgs {
    /**
     * Azure AD clientId of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
     */
    clientId?: pulumi.Input<string>;
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: pulumi.Input<string>;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    name?: pulumi.Input<string>;
}
