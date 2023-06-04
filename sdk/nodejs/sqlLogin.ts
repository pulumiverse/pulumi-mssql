// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages single login.
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
 * const exampleSqlLogin = new mssql.SqlLogin("exampleSqlLogin", {
 *     name: "example",
 *     password: "Str0ngPa$$word12",
 *     mustChangePassword: true,
 *     defaultDatabaseId: exampleDatabase.then(exampleDatabase => exampleDatabase.id),
 *     defaultLanguage: "english",
 *     checkPasswordExpiration: true,
 *     checkPasswordPolicy: true,
 * });
 * export const loginId = exampleSqlLogin.id;
 * ```
 *
 * ## Import
 *
 * import using login ID - can be retrieved using `SELECT SUSER_SID('<login_name>')`
 *
 * ```sh
 *  $ pulumi import mssql:index/sqlLogin:SqlLogin example 0x27578D8516843E4094EFA2CEED085C82
 * ```
 */
export class SqlLogin extends pulumi.CustomResource {
    /**
     * Get an existing SqlLogin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SqlLoginState, opts?: pulumi.CustomResourceOptions): SqlLogin {
        return new SqlLogin(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'mssql:index/sqlLogin:SqlLogin';

    /**
     * Returns true if the given object is an instance of SqlLogin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SqlLogin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SqlLogin.__pulumiType;
    }

    /**
     * When `true`, password expiration policy is enforced for this login. Defaults to `false`. -> **Note** In case of Azure
     * SQL, which does not support this feature, the flag will be ignored.
     */
    public readonly checkPasswordExpiration!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, the Windows password policies of the computer on which SQL Server is running are enforced on this login.
     * Defaults to `true`. -> **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    public readonly checkPasswordPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * ID of login's default DB. The ID can be retrieved using `mssql_database` data resource. Defaults to ID of `master`. ->
     * **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    public readonly defaultDatabaseId!: pulumi.Output<string | undefined>;
    /**
     * Default language assigned to login. Defaults to current default language of the server. If the default language of the
     * server is later changed, the default language of the login remains unchanged. -> **Note** In case of Azure SQL, which
     * does not support this feature, the flag will be ignored.
     */
    public readonly defaultLanguage!: pulumi.Output<string | undefined>;
    /**
     * When true, password change will be forced on first logon. Defaults to `false`. -> **Note** After password is changed,
     * this flag is being reset to `false`, which will show as changes in Terraform plan. Use `ignore_changes` block to prevent
     * this behavior. -> **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    public readonly mustChangePassword!: pulumi.Output<boolean | undefined>;
    /**
     * Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for the login. Must follow strong password policies defined for SQL server. Passwords are case-sensitive, length must be 8-128 chars, can include all characters except `'` or `name`.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * ID used to reference SQL Login in other resources, e.g. `server_role`. Can be retrieved from `sys.sql_logins`.
     */
    public /*out*/ readonly principalId!: pulumi.Output<string>;

    /**
     * Create a SqlLogin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SqlLoginArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SqlLoginArgs | SqlLoginState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SqlLoginState | undefined;
            resourceInputs["checkPasswordExpiration"] = state ? state.checkPasswordExpiration : undefined;
            resourceInputs["checkPasswordPolicy"] = state ? state.checkPasswordPolicy : undefined;
            resourceInputs["defaultDatabaseId"] = state ? state.defaultDatabaseId : undefined;
            resourceInputs["defaultLanguage"] = state ? state.defaultLanguage : undefined;
            resourceInputs["mustChangePassword"] = state ? state.mustChangePassword : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["principalId"] = state ? state.principalId : undefined;
        } else {
            const args = argsOrState as SqlLoginArgs | undefined;
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["checkPasswordExpiration"] = args ? args.checkPasswordExpiration : undefined;
            resourceInputs["checkPasswordPolicy"] = args ? args.checkPasswordPolicy : undefined;
            resourceInputs["defaultDatabaseId"] = args ? args.defaultDatabaseId : undefined;
            resourceInputs["defaultLanguage"] = args ? args.defaultLanguage : undefined;
            resourceInputs["mustChangePassword"] = args ? args.mustChangePassword : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["principalId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SqlLogin.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SqlLogin resources.
 */
export interface SqlLoginState {
    /**
     * When `true`, password expiration policy is enforced for this login. Defaults to `false`. -> **Note** In case of Azure
     * SQL, which does not support this feature, the flag will be ignored.
     */
    checkPasswordExpiration?: pulumi.Input<boolean>;
    /**
     * When `true`, the Windows password policies of the computer on which SQL Server is running are enforced on this login.
     * Defaults to `true`. -> **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    checkPasswordPolicy?: pulumi.Input<boolean>;
    /**
     * ID of login's default DB. The ID can be retrieved using `mssql_database` data resource. Defaults to ID of `master`. ->
     * **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    defaultDatabaseId?: pulumi.Input<string>;
    /**
     * Default language assigned to login. Defaults to current default language of the server. If the default language of the
     * server is later changed, the default language of the login remains unchanged. -> **Note** In case of Azure SQL, which
     * does not support this feature, the flag will be ignored.
     */
    defaultLanguage?: pulumi.Input<string>;
    /**
     * When true, password change will be forced on first logon. Defaults to `false`. -> **Note** After password is changed,
     * this flag is being reset to `false`, which will show as changes in Terraform plan. Use `ignore_changes` block to prevent
     * this behavior. -> **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    mustChangePassword?: pulumi.Input<boolean>;
    /**
     * Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
     */
    name?: pulumi.Input<string>;
    /**
     * Password for the login. Must follow strong password policies defined for SQL server. Passwords are case-sensitive, length must be 8-128 chars, can include all characters except `'` or `name`.
     */
    password?: pulumi.Input<string>;
    /**
     * ID used to reference SQL Login in other resources, e.g. `server_role`. Can be retrieved from `sys.sql_logins`.
     */
    principalId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SqlLogin resource.
 */
export interface SqlLoginArgs {
    /**
     * When `true`, password expiration policy is enforced for this login. Defaults to `false`. -> **Note** In case of Azure
     * SQL, which does not support this feature, the flag will be ignored.
     */
    checkPasswordExpiration?: pulumi.Input<boolean>;
    /**
     * When `true`, the Windows password policies of the computer on which SQL Server is running are enforced on this login.
     * Defaults to `true`. -> **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    checkPasswordPolicy?: pulumi.Input<boolean>;
    /**
     * ID of login's default DB. The ID can be retrieved using `mssql_database` data resource. Defaults to ID of `master`. ->
     * **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    defaultDatabaseId?: pulumi.Input<string>;
    /**
     * Default language assigned to login. Defaults to current default language of the server. If the default language of the
     * server is later changed, the default language of the login remains unchanged. -> **Note** In case of Azure SQL, which
     * does not support this feature, the flag will be ignored.
     */
    defaultLanguage?: pulumi.Input<string>;
    /**
     * When true, password change will be forced on first logon. Defaults to `false`. -> **Note** After password is changed,
     * this flag is being reset to `false`, which will show as changes in Terraform plan. Use `ignore_changes` block to prevent
     * this behavior. -> **Note** In case of Azure SQL, which does not support this feature, the flag will be ignored.
     */
    mustChangePassword?: pulumi.Input<boolean>;
    /**
     * Login name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot contain `\`
     */
    name: pulumi.Input<string>;
    /**
     * Password for the login. Must follow strong password policies defined for SQL server. Passwords are case-sensitive, length must be 8-128 chars, can include all characters except `'` or `name`.
     */
    password: pulumi.Input<string>;
}