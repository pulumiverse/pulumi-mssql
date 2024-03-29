# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

azureAuth: Optional[str]
"""
When provided, Azure AD authentication will be used when connecting.
"""

hostname: Optional[str]
"""
FQDN or IP address of the SQL endpoint. Can be also set using `MSSQL_HOSTNAME` environment variable.
"""

port: int
"""
TCP port of SQL endpoint. Defaults to `1433`. Can be also set using `MSSQL_PORT` environment variable.
"""

sqlAuth: Optional[str]
"""
When provided, SQL authentication will be used when connecting.
"""

