// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages single DB schema.
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
 * const exampleSchema = new mssql.Schema("exampleSchema", {
 *     databaseId: exampleDatabase.then(exampleDatabase => exampleDatabase.id),
 *     ownerId: owner.then(owner => owner.id),
 * });
 * ```
 *
 * ## Import
 *
 * import using <db_id>/<schema_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', SCHEMA_ID('<schema_name>'))`
 *
 * ```sh
 * $ pulumi import mssql:index/schema:Schema example '7/5'
 * ```
 */
export class Schema extends pulumi.CustomResource {
    /**
     * Get an existing Schema resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SchemaState, opts?: pulumi.CustomResourceOptions): Schema {
        return new Schema(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mssql:index/schema:Schema';

    /**
     * Returns true if the given object is an instance of Schema.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schema {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schema.__pulumiType;
    }

    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    public readonly databaseId!: pulumi.Output<string>;
    /**
     * Schema name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
     */
    public readonly ownerId!: pulumi.Output<string>;

    /**
     * Create a Schema resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SchemaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SchemaArgs | SchemaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SchemaState | undefined;
            resourceInputs["databaseId"] = state ? state.databaseId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
        } else {
            const args = argsOrState as SchemaArgs | undefined;
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Schema.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Schema resources.
 */
export interface SchemaState {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * Schema name.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
     */
    ownerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Schema resource.
 */
export interface SchemaArgs {
    /**
     * ID of database. Can be retrieved using `mssql.Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * Schema name.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of database role or user owning this schema. Can be retrieved using `mssql.DatabaseRole`, `mssql.SqlUser`, `mssql.AzureadUser` or `mssql.AzureadServicePrincipal`
     */
    ownerId?: pulumi.Input<string>;
}
