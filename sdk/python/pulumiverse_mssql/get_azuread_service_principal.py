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
    'GetAzureadServicePrincipalResult',
    'AwaitableGetAzureadServicePrincipalResult',
    'get_azuread_service_principal',
    'get_azuread_service_principal_output',
]

@pulumi.output_type
class GetAzureadServicePrincipalResult:
    """
    A collection of values returned by getAzureadServicePrincipal.
    """
    def __init__(__self__, client_id=None, database_id=None, id=None, name=None):
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if database_id and not isinstance(database_id, str):
            raise TypeError("Expected argument 'database_id' to be a str")
        pulumi.set(__self__, "database_id", database_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        Azure AD client_id of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
        """
        return pulumi.get(self, "client_id")

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
        """
        `<database_id>/<user_id>`. User ID can be retrieved using `sys.database_principals` view.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        User name. Cannot be longer than 128 chars.
        """
        return pulumi.get(self, "name")


class AwaitableGetAzureadServicePrincipalResult(GetAzureadServicePrincipalResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAzureadServicePrincipalResult(
            client_id=self.client_id,
            database_id=self.database_id,
            id=self.id,
            name=self.name)


def get_azuread_service_principal(client_id: Optional[str] = None,
                                  database_id: Optional[str] = None,
                                  name: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAzureadServicePrincipalResult:
    """
    Obtains information about single Azure AD Service Principal database user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    example_database = mssql.get_database(name="example")
    example_azuread_service_principal = mssql.get_azuread_service_principal(name="example",
        database_id=example_database.id)
    pulumi.export("appClientId", example_azuread_service_principal.client_id)
    ```


    :param str client_id: Azure AD client_id of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
    :param str name: User name. Cannot be longer than 128 chars.
    """
    __args__ = dict()
    __args__['clientId'] = client_id
    __args__['databaseId'] = database_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('mssql:index/getAzureadServicePrincipal:getAzureadServicePrincipal', __args__, opts=opts, typ=GetAzureadServicePrincipalResult).value

    return AwaitableGetAzureadServicePrincipalResult(
        client_id=__ret__.client_id,
        database_id=__ret__.database_id,
        id=__ret__.id,
        name=__ret__.name)


@_utilities.lift_output_func(get_azuread_service_principal)
def get_azuread_service_principal_output(client_id: Optional[pulumi.Input[Optional[str]]] = None,
                                         database_id: Optional[pulumi.Input[str]] = None,
                                         name: Optional[pulumi.Input[Optional[str]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAzureadServicePrincipalResult]:
    """
    Obtains information about single Azure AD Service Principal database user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_mssql as mssql

    example_database = mssql.get_database(name="example")
    example_azuread_service_principal = mssql.get_azuread_service_principal(name="example",
        database_id=example_database.id)
    pulumi.export("appClientId", example_azuread_service_principal.client_id)
    ```


    :param str client_id: Azure AD client_id of the Service Principal. This can be either regular Service Principal or Managed Service Identity.
    :param str database_id: ID of database. Can be retrieved using `Database` or `SELECT DB_ID('<db_name>')`.
    :param str name: User name. Cannot be longer than 128 chars.
    """
    ...
