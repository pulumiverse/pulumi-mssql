# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DatabaseRoleArgs', 'DatabaseRole']

@pulumi.input_type
class DatabaseRoleArgs:
    def __init__(__self__, *,
                 database_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabaseRole resource.
        :param pulumi.Input[str] database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        :param pulumi.Input[str] name: Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        :param pulumi.Input[str] owner_id: ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
               Defaults to ID of current user, used to authorize the Terraform provider.
        """
        if database_id is not None:
            pulumi.set(__self__, "database_id", database_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
        Defaults to ID of current user, used to authorize the Terraform provider.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)


@pulumi.input_type
class _DatabaseRoleState:
    def __init__(__self__, *,
                 database_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatabaseRole resources.
        :param pulumi.Input[str] database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        :param pulumi.Input[str] name: Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        :param pulumi.Input[str] owner_id: ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
               Defaults to ID of current user, used to authorize the Terraform provider.
        """
        if database_id is not None:
            pulumi.set(__self__, "database_id", database_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
        Defaults to ID of current user, used to authorize the Terraform provider.
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)


class DatabaseRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages database-level role.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_mssql as mssql
        import pulumiverse_mssql as mssql

        example_database = mssql.get_database(name="example")
        owner = mssql.get_sql_user(name="example_user")
        example_database_role = mssql.DatabaseRole("exampleDatabaseRole",
            database_id=example_database.id,
            owner_id=owner.id)
        ```

        ## Import

        import using <db_id>/<role_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', DATABASE_PRINCIPAL_ID('<role_name>'))`

        ```sh
         $ pulumi import mssql:index/databaseRole:DatabaseRole example '7/5'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        :param pulumi.Input[str] name: Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        :param pulumi.Input[str] owner_id: ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
               Defaults to ID of current user, used to authorize the Terraform provider.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DatabaseRoleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages database-level role.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_mssql as mssql
        import pulumiverse_mssql as mssql

        example_database = mssql.get_database(name="example")
        owner = mssql.get_sql_user(name="example_user")
        example_database_role = mssql.DatabaseRole("exampleDatabaseRole",
            database_id=example_database.id,
            owner_id=owner.id)
        ```

        ## Import

        import using <db_id>/<role_id> - can be retrieved using `SELECT CONCAT(DB_ID(), '/', DATABASE_PRINCIPAL_ID('<role_name>'))`

        ```sh
         $ pulumi import mssql:index/databaseRole:DatabaseRole example '7/5'
        ```

        :param str resource_name: The name of the resource.
        :param DatabaseRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseRoleArgs.__new__(DatabaseRoleArgs)

            __props__.__dict__["database_id"] = database_id
            __props__.__dict__["name"] = name
            __props__.__dict__["owner_id"] = owner_id
        super(DatabaseRole, __self__).__init__(
            'mssql:index/databaseRole:DatabaseRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None) -> 'DatabaseRole':
        """
        Get an existing DatabaseRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        :param pulumi.Input[str] name: Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        :param pulumi.Input[str] owner_id: ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
               Defaults to ID of current user, used to authorize the Terraform provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseRoleState.__new__(_DatabaseRoleState)

        __props__.__dict__["database_id"] = database_id
        __props__.__dict__["name"] = name
        __props__.__dict__["owner_id"] = owner_id
        return DatabaseRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Output[str]:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Role name. Must follow [Regular Identifiers rules](https://docs.microsoft.com/en-us/sql/relational-databases/databases/database-identifiers#rules-for-regular-identifiers) and cannot be longer than 128 chars.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        ID of another database role or user owning this role. Can be retrieved using `mssql_database_role` or `mssql_sql_user`.
        Defaults to ID of current user, used to authorize the Terraform provider.
        """
        return pulumi.get(self, "owner_id")

