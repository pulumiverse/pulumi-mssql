import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
import * as azuread from "@pulumi/azuread";
import * as mssql from "@pulumiverse/mssql";
import * as random from "@pulumi/random";

const res = (async () => {
    const provider = new azure.Provider("provider", { skipProviderRegistration: true });
    const providerAzuread = new azuread.Provider("provider-azuread", {});

    const config = new pulumi.Config();
    const resourcegroupLocation = config.require("resourcegroup-location");
    const resourcegroupName = config.require("resourcegroup-name");
    const localIp = config.require("local-ip");
    const sqladmin = config.require("sqladmin");
    const dbName = config.require("db-name");
    const roleName = "db_owner";

    const randomPrefix = new random.RandomId("random-prefix", {
        byteLength: 8,
        keepers: {
            keep: resourcegroupName,
        },
    });

    const serverName = pulumi.interpolate`sql${randomPrefix.hex}`;
    const current = await azure.core.getClientConfig({
        provider: provider
    });
    const aadSqladmin = await azuread.getUser({
        userPrincipalName: sqladmin,
    }, {
        provider: providerAzuread
    });
    const aadGroupSqladmins = new azuread.Group("aad-group-sqladmins", {
        displayName: pulumi.interpolate`AZ.RESOURCES.${serverName}.Admins`,
        securityEnabled: true,
        members: [
            current.objectId,
            aadSqladmin.objectId,
        ],
    }, {
        provider: providerAzuread,
    });
    const resourceGroup = new azure.core.ResourceGroup("resource-group", {
        name: resourcegroupName,
        location: resourcegroupLocation,
    }, {
        provider: provider,
    });
    const serverPassword = new random.RandomPassword("server-password", {
        length: 32,
        special: true,
    });
    const server = new azure.mssql.Server("server", {
        name: serverName,
        resourceGroupName: resourceGroup.name,
        location: resourceGroup.location,
        version: "12.0",
        administratorLogin: "sadmin",
        administratorLoginPassword: serverPassword.result,
        minimumTlsVersion: "1.2",
        azureadAdministrator: {
            azureadAuthenticationOnly: true,
            loginUsername: aadGroupSqladmins.displayName,
            objectId: aadGroupSqladmins.objectId,
            tenantId: current.tenantId,
        },
    }, {
        provider: provider,
        parent: resourceGroup,
    });

    const database = new azure.mssql.Database("database", {
        name: dbName,
        serverId: server.id,
        licenseType: "LicenseIncluded",
        maxSizeGb: 2,
        skuName: "Basic",
    }, {
        provider: provider,
        parent: server,
    });

    const databaseFirewallRule = new azure.mssql.FirewallRule("database-firewall-rule", {
        name: "client",
        serverId: server.id,
        startIpAddress: localIp,
        endIpAddress: localIp,
    }, {
        provider: provider,
        parent: server,
    });

    const providerMssql = new mssql.Provider("provider-mssql", {
        hostname: server.fullyQualifiedDomainName,
        azureAuth: {},
    }, {
        dependsOn: [
            database,
            databaseFirewallRule,
        ],
    });

    const databaseId = await mssql.getDatabase({
        name: dbName,
    }, {
        provider: providerMssql,
        parent: database,
    });
    const databaseRoleOwner = await mssql.getDatabaseRole({
        databaseId: databaseId.id,
        name: roleName,
    }, {
        provider: providerMssql,
        parent: database,
    });

    const aadGroupDbOwner = new azuread.Group("aad-group-db-owner", {
        displayName: pulumi.interpolate`AZ.RESOURCES.${server.name}.Database.${database.name}.${roleName}`,
        securityEnabled: true,
    }, {
        provider: providerAzuread,
        parent: database,
    });
    const dbUser = new mssql.AzureadUser("db-user", {
        databaseId: databaseId.id || "0",
        name: aadGroupDbOwner.displayName,
        userObjectId: aadGroupDbOwner.objectId,
    }, {
        provider: providerMssql,
        parent: database,
    });
    const dbRoleMember = new mssql.DatabaseRoleMember("db-role-member", {
        roleId: databaseRoleOwner.id || "0",
        memberId: dbUser.id,
    }, {
        provider: providerMssql,
        parent: database,
    });
})() // async-wrapper

export default res
