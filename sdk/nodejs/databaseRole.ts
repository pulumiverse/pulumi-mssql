// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages database-level role.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as mssql from "@pulumi/mssql";
 * import * as mssql from "@pulumiverse/mssql";
 *
 * const exampleDatabase = mssql.getDatabase({
 *     name: "example",
 * });
 * const owner = mssql.getSqlUser({
 *     name: "example_user",
 * });
 * const exampleDatabaseRole = new mssql.DatabaseRole("exampleDatabaseRole", {
 *     name: "example",
 *     databaseId: exampleDatabase.then(exampleDatabase => exampleDatabase.id),
 *     ownerId: owner.then(owner => owner.id),
 * });
 * ```
 *
 * ## Import
 *
 * import using <db_id>/<role_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', DATABASE_PRINCIPAL_ID('<role_name>'))`
 *
 * ```sh
 *  $ pulumi import mssql:index/databaseRole:DatabaseRole example '7/5'
 * ```
 */
export class DatabaseRole extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseRoleState, opts?: pulumi.CustomResourceOptions): DatabaseRole {
        return new DatabaseRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mssql:index/databaseRole:DatabaseRole';

    /**
     * Returns true if the given object is an instance of DatabaseRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseRole.__pulumiType;
    }

    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    public readonly databaseId!: pulumi.Output<string>;
    /**
     * Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
     * Defaults to ID of current user, used to authorize the Terraform provider.
     */
    public readonly ownerId!: pulumi.Output<string>;

    /**
     * Create a DatabaseRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseRoleArgs | DatabaseRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseRoleState | undefined;
            resourceInputs["databaseId"] = state ? state.databaseId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
        } else {
            const args = argsOrState as DatabaseRoleArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseRole resources.
 */
export interface DatabaseRoleState {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
     * Defaults to ID of current user, used to authorize the Terraform provider.
     */
    ownerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseRole resource.
 */
export interface DatabaseRoleArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
     */
    name: pulumi.Input<string>;
    /**
     * ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
     * Defaults to ID of current user, used to authorize the Terraform provider.
     */
    ownerId?: pulumi.Input<string>;
}
