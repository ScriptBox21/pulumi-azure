# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DefinitionAuthorization',
]

@pulumi.output_type
class DefinitionAuthorization(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "principalId":
            suggest = "principal_id"
        elif key == "roleDefinitionId":
            suggest = "role_definition_id"
        elif key == "delegatedRoleDefinitionIds":
            suggest = "delegated_role_definition_ids"
        elif key == "principalDisplayName":
            suggest = "principal_display_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in DefinitionAuthorization. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        DefinitionAuthorization.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        DefinitionAuthorization.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 principal_id: str,
                 role_definition_id: str,
                 delegated_role_definition_ids: Optional[Sequence[str]] = None,
                 principal_display_name: Optional[str] = None):
        """
        :param str principal_id: Principal ID of the security group/service principal/user that would be assigned permissions to the projected subscription.
        :param str role_definition_id: The role definition identifier. This role will define the permissions that are granted to the principal. This cannot be an `Owner` role.
        :param Sequence[str] delegated_role_definition_ids: The set of role definition ids which define all the permissions that the principal id can assign.
        :param str principal_display_name: The display name of the security group/service principal/user that would be assigned permissions to the projected subscription.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "role_definition_id", role_definition_id)
        if delegated_role_definition_ids is not None:
            pulumi.set(__self__, "delegated_role_definition_ids", delegated_role_definition_ids)
        if principal_display_name is not None:
            pulumi.set(__self__, "principal_display_name", principal_display_name)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        Principal ID of the security group/service principal/user that would be assigned permissions to the projected subscription.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> str:
        """
        The role definition identifier. This role will define the permissions that are granted to the principal. This cannot be an `Owner` role.
        """
        return pulumi.get(self, "role_definition_id")

    @property
    @pulumi.getter(name="delegatedRoleDefinitionIds")
    def delegated_role_definition_ids(self) -> Optional[Sequence[str]]:
        """
        The set of role definition ids which define all the permissions that the principal id can assign.
        """
        return pulumi.get(self, "delegated_role_definition_ids")

    @property
    @pulumi.getter(name="principalDisplayName")
    def principal_display_name(self) -> Optional[str]:
        """
        The display name of the security group/service principal/user that would be assigned permissions to the projected subscription.
        """
        return pulumi.get(self, "principal_display_name")


