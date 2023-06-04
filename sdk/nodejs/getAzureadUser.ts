// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Obtains information about single Azure AD database user.
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
 * const exampleAzureadUser = exampleDatabase.then(exampleDatabase => mssql.getAzureadUser({
 *     name: "example",
 *     databaseId: exampleDatabase.id,
 * }));
 * export const userObjectId = exampleAzureadUser.then(exampleAzureadUser => exampleAzureadUser.userObjectId);
 * ```
 */
export function getAzureadUser(args: GetAzureadUserArgs, opts?: pulumi.InvokeOptions): Promise<GetAzureadUserResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("mssql:index/getAzureadUser:getAzureadUser", {
        "databaseId": args.databaseId,
        "name": args.name,
        "userObjectId": args.userObjectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAzureadUser.
 */
export interface GetAzureadUserArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: string;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    name?: string;
    /**
     * Azure AD objectId of the user. This can be either regular user or a group.
     */
    userObjectId?: string;
}

/**
 * A collection of values returned by getAzureadUser.
 */
export interface GetAzureadUserResult {
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
    /**
     * Azure AD objectId of the user. This can be either regular user or a group.
     */
    readonly userObjectId: string;
}
/**
 * Obtains information about single Azure AD database user.
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
 * const exampleAzureadUser = exampleDatabase.then(exampleDatabase => mssql.getAzureadUser({
 *     name: "example",
 *     databaseId: exampleDatabase.id,
 * }));
 * export const userObjectId = exampleAzureadUser.then(exampleAzureadUser => exampleAzureadUser.userObjectId);
 * ```
 */
export function getAzureadUserOutput(args: GetAzureadUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAzureadUserResult> {
    return pulumi.output(args).apply((a: any) => getAzureadUser(a, opts))
}

/**
 * A collection of arguments for invoking getAzureadUser.
 */
export interface GetAzureadUserOutputArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId: pulumi.Input<string>;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    name?: pulumi.Input<string>;
    /**
     * Azure AD objectId of the user. This can be either regular user or a group.
     */
    userObjectId?: pulumi.Input<string>;
}