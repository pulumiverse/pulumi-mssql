// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Managed database-level user mapped to Azure AD identity (user or group).
 *
 * > **Note** When using this resource, Azure SQL server managed identity does not need any [AzureAD role assignments](https://docs.microsoft.com/en-us/azure/azure-sql/database/authentication-aad-service-principal?view=azuresql).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 * import * as mssql from "@pulumi/mssql";
 * import * as mssql from "@pulumiverse/mssql";
 *
 * const exampleDatabase = mssql.getDatabase({
 *     name: "example",
 * });
 * const exampleUser = azuread.getUser({
 *     userPrincipalName: "user@example.com",
 * });
 * const exampleAzureadUser = new mssql.AzureadUser("exampleAzureadUser", {
 *     databaseId: exampleDatabase.then(exampleDatabase => exampleDatabase.id),
 *     userObjectId: exampleUser.then(exampleUser => exampleUser.objectId),
 * });
 * export const userId = exampleAzureadUser.id;
 * ```
 *
 * ## Import
 *
 * import using <db_id>/<user_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', principal_id) FROM sys.database_principals WHERE [name] = '<username>'`
 *
 * ```sh
 * $ pulumi import mssql:index/azureadUser:AzureadUser example '7/5'
 * ```
 */
export class AzureadUser extends pulumi.CustomResource {
    /**
     * Get an existing AzureadUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AzureadUserState, opts?: pulumi.CustomResourceOptions): AzureadUser {
        return new AzureadUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mssql:index/azureadUser:AzureadUser';

    /**
     * Returns true if the given object is an instance of AzureadUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AzureadUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AzureadUser.__pulumiType;
    }

    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    public readonly databaseId!: pulumi.Output<string>;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Azure AD objectId of the user. This can be either regular user or a group.
     */
    public readonly userObjectId!: pulumi.Output<string>;

    /**
     * Create a AzureadUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AzureadUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AzureadUserArgs | AzureadUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AzureadUserState | undefined;
            resourceInputs["databaseId"] = state ? state.databaseId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["userObjectId"] = state ? state.userObjectId : undefined;
        } else {
            const args = argsOrState as AzureadUserArgs | undefined;
            if ((!args || args.databaseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseId'");
            }
            if ((!args || args.userObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userObjectId'");
            }
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["userObjectId"] = args ? args.userObjectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AzureadUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AzureadUser resources.
 */
export interface AzureadUserState {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * User name. Cannot be longer than 128 chars.
     */
    name?: pulumi.Input<string>;
    /**
     * Azure AD objectId of the user. This can be either regular user or a group.
     */
    userObjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AzureadUser resource.
 */
export interface AzureadUserArgs {
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
    userObjectId: pulumi.Input<string>;
}
