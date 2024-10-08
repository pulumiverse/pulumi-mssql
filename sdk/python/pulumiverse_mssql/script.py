# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ScriptArgs', 'Script']

@pulumi.input_type
class ScriptArgs:
    def __init__(__self__, *,
                 database_id: pulumi.Input[str],
                 read_script: pulumi.Input[str],
                 state: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 update_script: pulumi.Input[str],
                 create_script: Optional[pulumi.Input[str]] = None,
                 delete_script: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Script resource.
        :param pulumi.Input[str] database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
        :param pulumi.Input[str] read_script: SQL script returning current state of the DB. It must return single-row result set where column names match the keys of `state` map and all values are strings that will be compared against `state` to determine if the resource state matches DB state.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] state: Desired state of the DB. It is arbitrary map of string values that will be compared against the values returned by the `read_script`.
        :param pulumi.Input[str] update_script: SQL script executed when the desired state specified in `state` attribute does not match the state returned by `read_script`
        :param pulumi.Input[str] delete_script: SQL script executed when the resource is being destroyed. When not provided, no action will be taken during resource destruction.
        """
        pulumi.set(__self__, "database_id", database_id)
        pulumi.set(__self__, "read_script", read_script)
        pulumi.set(__self__, "state", state)
        pulumi.set(__self__, "update_script", update_script)
        if create_script is not None:
            pulumi.set(__self__, "create_script", create_script)
        if delete_script is not None:
            pulumi.set(__self__, "delete_script", delete_script)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Input[str]:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter(name="readScript")
    def read_script(self) -> pulumi.Input[str]:
        """
        SQL script returning current state of the DB. It must return single-row result set where column names match the keys of `state` map and all values are strings that will be compared against `state` to determine if the resource state matches DB state.
        """
        return pulumi.get(self, "read_script")

    @read_script.setter
    def read_script(self, value: pulumi.Input[str]):
        pulumi.set(self, "read_script", value)

    @property
    @pulumi.getter
    def state(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        Desired state of the DB. It is arbitrary map of string values that will be compared against the values returned by the `read_script`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="updateScript")
    def update_script(self) -> pulumi.Input[str]:
        """
        SQL script executed when the desired state specified in `state` attribute does not match the state returned by `read_script`
        """
        return pulumi.get(self, "update_script")

    @update_script.setter
    def update_script(self, value: pulumi.Input[str]):
        pulumi.set(self, "update_script", value)

    @property
    @pulumi.getter(name="createScript")
    def create_script(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "create_script")

    @create_script.setter
    def create_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_script", value)

    @property
    @pulumi.getter(name="deleteScript")
    def delete_script(self) -> Optional[pulumi.Input[str]]:
        """
        SQL script executed when the resource is being destroyed. When not provided, no action will be taken during resource destruction.
        """
        return pulumi.get(self, "delete_script")

    @delete_script.setter
    def delete_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_script", value)


@pulumi.input_type
class _ScriptState:
    def __init__(__self__, *,
                 create_script: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 delete_script: Optional[pulumi.Input[str]] = None,
                 read_script: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 update_script: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Script resources.
        :param pulumi.Input[str] database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
        :param pulumi.Input[str] delete_script: SQL script executed when the resource is being destroyed. When not provided, no action will be taken during resource destruction.
        :param pulumi.Input[str] read_script: SQL script returning current state of the DB. It must return single-row result set where column names match the keys of `state` map and all values are strings that will be compared against `state` to determine if the resource state matches DB state.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] state: Desired state of the DB. It is arbitrary map of string values that will be compared against the values returned by the `read_script`.
        :param pulumi.Input[str] update_script: SQL script executed when the desired state specified in `state` attribute does not match the state returned by `read_script`
        """
        if create_script is not None:
            pulumi.set(__self__, "create_script", create_script)
        if database_id is not None:
            pulumi.set(__self__, "database_id", database_id)
        if delete_script is not None:
            pulumi.set(__self__, "delete_script", delete_script)
        if read_script is not None:
            pulumi.set(__self__, "read_script", read_script)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if update_script is not None:
            pulumi.set(__self__, "update_script", update_script)

    @property
    @pulumi.getter(name="createScript")
    def create_script(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "create_script")

    @create_script.setter
    def create_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_script", value)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter(name="deleteScript")
    def delete_script(self) -> Optional[pulumi.Input[str]]:
        """
        SQL script executed when the resource is being destroyed. When not provided, no action will be taken during resource destruction.
        """
        return pulumi.get(self, "delete_script")

    @delete_script.setter
    def delete_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_script", value)

    @property
    @pulumi.getter(name="readScript")
    def read_script(self) -> Optional[pulumi.Input[str]]:
        """
        SQL script returning current state of the DB. It must return single-row result set where column names match the keys of `state` map and all values are strings that will be compared against `state` to determine if the resource state matches DB state.
        """
        return pulumi.get(self, "read_script")

    @read_script.setter
    def read_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "read_script", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Desired state of the DB. It is arbitrary map of string values that will be compared against the values returned by the `read_script`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="updateScript")
    def update_script(self) -> Optional[pulumi.Input[str]]:
        """
        SQL script executed when the desired state specified in `state` attribute does not match the state returned by `read_script`
        """
        return pulumi.get(self, "update_script")

    @update_script.setter
    def update_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_script", value)


class Script(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_script: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 delete_script: Optional[pulumi.Input[str]] = None,
                 read_script: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 update_script: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows execution of arbitrary SQL scripts to check state and apply desired state.

        > **Note** This resource is meant to be an escape hatch for all cases not supported by the provider's resources. Whenever possible, use dedicated resources, which offer better plan, validation and error reporting.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_mssql as mssql
        import pulumiverse_mssql as mssql

        test = mssql.get_database(name="test")
        cdc = mssql.Script("cdc",
            database_id=test.id,
            read_script=f"SELECT COUNT(*) AS [is_enabled] FROM sys.change_tracking_databases WHERE database_id={test.id}",
            delete_script=f"ALTER DATABASE [{test.name}] SET CHANGE_TRACKING = OFF",
            update_script=f\"\"\"IF (SELECT COUNT(*) FROM sys.change_tracking_databases WHERE database_id={test.id}) = 0
          ALTER DATABASE [{test.name}] SET CHANGE_TRACKING = ON
        \"\"\",
            state={
                "is_enabled": "1",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
        :param pulumi.Input[str] delete_script: SQL script executed when the resource is being destroyed. When not provided, no action will be taken during resource destruction.
        :param pulumi.Input[str] read_script: SQL script returning current state of the DB. It must return single-row result set where column names match the keys of `state` map and all values are strings that will be compared against `state` to determine if the resource state matches DB state.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] state: Desired state of the DB. It is arbitrary map of string values that will be compared against the values returned by the `read_script`.
        :param pulumi.Input[str] update_script: SQL script executed when the desired state specified in `state` attribute does not match the state returned by `read_script`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScriptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows execution of arbitrary SQL scripts to check state and apply desired state.

        > **Note** This resource is meant to be an escape hatch for all cases not supported by the provider's resources. Whenever possible, use dedicated resources, which offer better plan, validation and error reporting.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_mssql as mssql
        import pulumiverse_mssql as mssql

        test = mssql.get_database(name="test")
        cdc = mssql.Script("cdc",
            database_id=test.id,
            read_script=f"SELECT COUNT(*) AS [is_enabled] FROM sys.change_tracking_databases WHERE database_id={test.id}",
            delete_script=f"ALTER DATABASE [{test.name}] SET CHANGE_TRACKING = OFF",
            update_script=f\"\"\"IF (SELECT COUNT(*) FROM sys.change_tracking_databases WHERE database_id={test.id}) = 0
          ALTER DATABASE [{test.name}] SET CHANGE_TRACKING = ON
        \"\"\",
            state={
                "is_enabled": "1",
            })
        ```

        :param str resource_name: The name of the resource.
        :param ScriptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScriptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_script: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 delete_script: Optional[pulumi.Input[str]] = None,
                 read_script: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 update_script: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScriptArgs.__new__(ScriptArgs)

            __props__.__dict__["create_script"] = create_script
            if database_id is None and not opts.urn:
                raise TypeError("Missing required property 'database_id'")
            __props__.__dict__["database_id"] = database_id
            __props__.__dict__["delete_script"] = delete_script
            if read_script is None and not opts.urn:
                raise TypeError("Missing required property 'read_script'")
            __props__.__dict__["read_script"] = read_script
            if state is None and not opts.urn:
                raise TypeError("Missing required property 'state'")
            __props__.__dict__["state"] = state
            if update_script is None and not opts.urn:
                raise TypeError("Missing required property 'update_script'")
            __props__.__dict__["update_script"] = update_script
        super(Script, __self__).__init__(
            'mssql:index/script:Script',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_script: Optional[pulumi.Input[str]] = None,
            database_id: Optional[pulumi.Input[str]] = None,
            delete_script: Optional[pulumi.Input[str]] = None,
            read_script: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            update_script: Optional[pulumi.Input[str]] = None) -> 'Script':
        """
        Get an existing Script resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
        :param pulumi.Input[str] delete_script: SQL script executed when the resource is being destroyed. When not provided, no action will be taken during resource destruction.
        :param pulumi.Input[str] read_script: SQL script returning current state of the DB. It must return single-row result set where column names match the keys of `state` map and all values are strings that will be compared against `state` to determine if the resource state matches DB state.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] state: Desired state of the DB. It is arbitrary map of string values that will be compared against the values returned by the `read_script`.
        :param pulumi.Input[str] update_script: SQL script executed when the desired state specified in `state` attribute does not match the state returned by `read_script`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ScriptState.__new__(_ScriptState)

        __props__.__dict__["create_script"] = create_script
        __props__.__dict__["database_id"] = database_id
        __props__.__dict__["delete_script"] = delete_script
        __props__.__dict__["read_script"] = read_script
        __props__.__dict__["state"] = state
        __props__.__dict__["update_script"] = update_script
        return Script(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createScript")
    def create_script(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "create_script")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Output[str]:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter(name="deleteScript")
    def delete_script(self) -> pulumi.Output[Optional[str]]:
        """
        SQL script executed when the resource is being destroyed. When not provided, no action will be taken during resource destruction.
        """
        return pulumi.get(self, "delete_script")

    @property
    @pulumi.getter(name="readScript")
    def read_script(self) -> pulumi.Output[str]:
        """
        SQL script returning current state of the DB. It must return single-row result set where column names match the keys of `state` map and all values are strings that will be compared against `state` to determine if the resource state matches DB state.
        """
        return pulumi.get(self, "read_script")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Desired state of the DB. It is arbitrary map of string values that will be compared against the values returned by the `read_script`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateScript")
    def update_script(self) -> pulumi.Output[str]:
        """
        SQL script executed when the desired state specified in `state` attribute does not match the state returned by `read_script`
        """
        return pulumi.get(self, "update_script")

