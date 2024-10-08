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
    'GetSchemasResult',
    'AwaitableGetSchemasResult',
    'get_schemas',
    'get_schemas_output',
]

@pulumi.output_type
class GetSchemasResult:
    """
    A collection of values returned by getSchemas.
    """
    def __init__(__self__, database_id=None, id=None, schemas=None):
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if schemas and not isinstance(schemas, list):
            raise TypeError("Expected argument 'schemas' to be a list")
        pulumi.set(__self__, "schemas", schemas)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> str:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the data source, equals to database ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def schemas(self) -> Sequence['outputs.GetSchemasSchemaResult']:
        """
        Set of schemas found in the DB.
        """
        return pulumi.get(self, "schemas")


class AwaitableGetSchemasResult(GetSchemasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSchemasResult(
            database_id=self.database_id,
            id=self.id,
            schemas=self.schemas)


def get_schemas(database_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSchemasResult:
    """
    Obtains information about all schemas found in SQL database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    example = mssql.get_database(name="example")
    all = mssql.get_schemas(database_id=example.id)
    pulumi.export("allSchemaNames", [__item.name for __item in all.schemas])
    ```


    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mssql:index/getSchemas:getSchemas', __args__, opts=opts, typ=GetSchemasResult).value

    return AwaitableGetSchemasResult(
        database_id=pulumi.get(__ret__, 'database_id'),
        id=pulumi.get(__ret__, 'id'),
        schemas=pulumi.get(__ret__, 'schemas'))


@_utilities.lift_output_func(get_schemas)
def get_schemas_output(database_id: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSchemasResult]:
    """
    Obtains information about all schemas found in SQL database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    example = mssql.get_database(name="example")
    all = mssql.get_schemas(database_id=example.id)
    pulumi.export("allSchemaNames", [__item.name for __item in all.schemas])
    ```


    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`. Defaults to ID of `master`.
    """
    ...
