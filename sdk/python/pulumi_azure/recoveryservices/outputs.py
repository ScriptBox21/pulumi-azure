# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'VaultIdentity',
]

@pulumi.output_type
class VaultIdentity(dict):
    def __init__(__self__, *,
                 type: str,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        :param str type: The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The Type of Identity which should be used for this Recovery Services Vault. At this time the only possible value is `SystemAssigned`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

