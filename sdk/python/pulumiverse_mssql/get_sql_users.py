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
    'GetSqlUsersResult',
    'AwaitableGetSqlUsersResult',
    'get_sql_users',
    'get_sql_users_output',
]

@pulumi.output_type
class GetSqlUsersResult:
    """
    A collection of values returned by getSqlUsers.
    """
    def __init__(__self__, database_id=None, id=None, users=None):
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

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
    def users(self) -> Sequence['outputs.GetSqlUsersUserResult']:
        """
        Set of SQL user objects
        """
        return pulumi.get(self, "users")


class AwaitableGetSqlUsersResult(GetSqlUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSqlUsersResult(
            database_id=self.database_id,
            id=self.id,
            users=self.users)


def get_sql_users(database_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSqlUsersResult:
    """
    Obtains information about all SQL users found in a database

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    master = mssql.get_database(name="master")
    example = mssql.get_sql_users(database_id=master.id)
    pulumi.export("users", example.users)
    ```


    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mssql:index/getSqlUsers:getSqlUsers', __args__, opts=opts, typ=GetSqlUsersResult).value

    return AwaitableGetSqlUsersResult(
        database_id=pulumi.get(__ret__, 'database_id'),
        id=pulumi.get(__ret__, 'id'),
        users=pulumi.get(__ret__, 'users'))


@_utilities.lift_output_func(get_sql_users)
def get_sql_users_output(database_id: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSqlUsersResult]:
    """
    Obtains information about all SQL users found in a database

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    master = mssql.get_database(name="master")
    example = mssql.get_sql_users(database_id=master.id)
    pulumi.export("users", example.users)
    ```


    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
    """
    ...
