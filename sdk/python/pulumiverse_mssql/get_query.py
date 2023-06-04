# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetQueryResult',
    'AwaitableGetQueryResult',
    'get_query',
    'get_query_output',
]

@pulumi.output_type
class GetQueryResult:
    """
    A collection of values returned by getQuery.
    """
    def __init__(__self__, database_id=None, id=None, query=None, results=None):
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if query and not isinstance(query, str):
            raise TypeError("Expected argument 'query' to be a str")
        pulumi.set(__self__, "query", query)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> str:
        """
        ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def query(self) -> str:
        """
        SQL query returning single result set, with any number of rows, where all columns are strings
        """
        return pulumi.get(self, "query")

    @property
    @pulumi.getter
    def results(self) -> Sequence[Mapping[str, str]]:
        """
        Results of the SQL query, represented as list of maps, where the map key corresponds to column name and the value is the value of column in given row.
        """
        return pulumi.get(self, "results")


class AwaitableGetQueryResult(GetQueryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQueryResult(
            database_id=self.database_id,
            id=self.id,
            query=self.query,
            results=self.results)


def get_query(database_id: Optional[str] = None,
              query: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQueryResult:
    """
    Retrieves arbitrary SQL query result.

    > **Note** This data source is meant to be an escape hatch for all cases not supported by the provider's data sources. Whenever possible, use dedicated data sources, which offer better plan, validation and error reporting.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    test = mssql.get_database(name="test")
    column = mssql.get_query(database_id=test.id,
        query="SELECT [column_id], [name] FROM sys.columns WHERE [object_id] = OBJECT_ID('test_table')")
    pulumi.export("columnNames", [__item["name"] for __item in column.results])
    ```


    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
    :param str query: SQL query returning single result set, with any number of rows, where all columns are strings
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    __args__['query'] = query
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mssql:index/getQuery:getQuery', __args__, opts=opts, typ=GetQueryResult).value

    return AwaitableGetQueryResult(
        database_id=__ret__.database_id,
        id=__ret__.id,
        query=__ret__.query,
        results=__ret__.results)


@_utilities.lift_output_func(get_query)
def get_query_output(database_id: Optional[pulumi.Input[str]] = None,
                     query: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetQueryResult]:
    """
    Retrieves arbitrary SQL query result.

    > **Note** This data source is meant to be an escape hatch for all cases not supported by the provider's data sources. Whenever possible, use dedicated data sources, which offer better plan, validation and error reporting.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    test = mssql.get_database(name="test")
    column = mssql.get_query(database_id=test.id,
        query="SELECT [column_id], [name] FROM sys.columns WHERE [object_id] = OBJECT_ID('test_table')")
    pulumi.export("columnNames", [__item["name"] for __item in column.results])
    ```


    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
    :param str query: SQL query returning single result set, with any number of rows, where all columns are strings
    """
    ...
