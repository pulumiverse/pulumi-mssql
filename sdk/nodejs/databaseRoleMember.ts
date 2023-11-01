// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages database role membership.
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
 * const owner = exampleDatabase.then(exampleDatabase => mssql.getSqlUser({
 *     name: "example_user",
 *     databaseId: exampleDatabase.id,
 * }));
 * const member = exampleDatabase.then(exampleDatabase => mssql.getSqlUser({
 *     name: "member_user",
 *     databaseId: exampleDatabase.id,
 * }));
 * const exampleDatabaseRole = new mssql.DatabaseRole("exampleDatabaseRole", {
 *     databaseId: exampleDatabase.then(exampleDatabase => exampleDatabase.id),
 *     ownerId: owner.then(owner => owner.id),
 * });
 * const exampleDatabaseRoleMember = new mssql.DatabaseRoleMember("exampleDatabaseRoleMember", {
 *     roleId: exampleDatabaseRole.id,
 *     memberId: member.then(member => member.id),
 * });
 * ```
 *
 * ## Import
 *
 * import using <db_id>/<role_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', DATABASE_PRINCIPAL_ID('<role_name>'), '/', DATABASE_PRINCIPAL_ID('<member_name'))`
 *
 * ```sh
 *  $ pulumi import mssql:index/databaseRoleMember:DatabaseRoleMember example '7/5/9'
 * ```
 */
export class DatabaseRoleMember extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseRoleMember resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseRoleMemberState, opts?: pulumi.CustomResourceOptions): DatabaseRoleMember {
        return new DatabaseRoleMember(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mssql:index/databaseRoleMember:DatabaseRoleMember';

    /**
     * Returns true if the given object is an instance of DatabaseRoleMember.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseRoleMember {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseRoleMember.__pulumiType;
    }

    /**
     * Can be either user or role ID in format `<database_id>/<member_id>`. Can be retrieved using `mssql.SqlUser` or `mssqlDatabaseMember`.
     */
    public readonly memberId!: pulumi.Output<string>;
    /**
     * `<database_id>/<role_id>`
     */
    public readonly roleId!: pulumi.Output<string>;

    /**
     * Create a DatabaseRoleMember resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseRoleMemberArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseRoleMemberArgs | DatabaseRoleMemberState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseRoleMemberState | undefined;
            resourceInputs["memberId"] = state ? state.memberId : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
        } else {
            const args = argsOrState as DatabaseRoleMemberArgs | undefined;
            if ((!args || args.memberId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memberId'");
            }
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            resourceInputs["memberId"] = args ? args.memberId : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatabaseRoleMember.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseRoleMember resources.
 */
export interface DatabaseRoleMemberState {
    /**
     * Can be either user or role ID in format `<database_id>/<member_id>`. Can be retrieved using `mssql.SqlUser` or `mssqlDatabaseMember`.
     */
    memberId?: pulumi.Input<string>;
    /**
     * `<database_id>/<role_id>`
     */
    roleId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseRoleMember resource.
 */
export interface DatabaseRoleMemberArgs {
    /**
     * Can be either user or role ID in format `<database_id>/<member_id>`. Can be retrieved using `mssql.SqlUser` or `mssqlDatabaseMember`.
     */
    memberId: pulumi.Input<string>;
    /**
     * `<database_id>/<role_id>`
     */
    roleId: pulumi.Input<string>;
}
