# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetDatabaseRolesResult',
    'AwaitableGetDatabaseRolesResult',
    'get_database_roles',
    'get_database_roles_output',
]

@pulumi.output_type
class GetDatabaseRolesResult:
    """
    A collection of values returned by getDatabaseRoles.
    """
    def __init__(__self__, database_id=None, id=None, roles=None):
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> Optional[str]:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the resource, equals to database ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetDatabaseRolesRoleResult']:
        """
        Set of database role objects
        """
        return pulumi.get(self, "roles")


class AwaitableGetDatabaseRolesResult(GetDatabaseRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseRolesResult(
            database_id=self.database_id,
            id=self.id,
            roles=self.roles)


def get_database_roles(database_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseRolesResult:
    """
    Obtains information about all roles defined in a database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    master = mssql.get_database(name="master")
    example = mssql.get_database_roles(database_id=master.id)
    pulumi.export("roles", example.roles)
    ```


    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mssql:index/getDatabaseRoles:getDatabaseRoles', __args__, opts=opts, typ=GetDatabaseRolesResult).value

    return AwaitableGetDatabaseRolesResult(
        database_id=__ret__.database_id,
        id=__ret__.id,
        roles=__ret__.roles)


@_utilities.lift_output_func(get_database_roles)
def get_database_roles_output(database_id: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseRolesResult]:
    """
    Obtains information about all roles defined in a database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    master = mssql.get_database(name="master")
    example = mssql.get_database_roles(database_id=master.id)
    pulumi.export("roles", example.roles)
    ```


    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
    """
    ...
