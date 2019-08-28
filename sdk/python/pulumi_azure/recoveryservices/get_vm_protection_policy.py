# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetVMProtectionPolicyResult:
    """
    A collection of values returned by getVMProtectionPolicy.
    """
    def __init__(__self__, name=None, recovery_vault_name=None, resource_group_name=None, tags=None, id=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if recovery_vault_name and not isinstance(recovery_vault_name, str):
            raise TypeError("Expected argument 'recovery_vault_name' to be a str")
        __self__.recovery_vault_name = recovery_vault_name
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetVMProtectionPolicyResult(GetVMProtectionPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVMProtectionPolicyResult(
            name=self.name,
            recovery_vault_name=self.recovery_vault_name,
            resource_group_name=self.resource_group_name,
            tags=self.tags,
            id=self.id)

def get_vm_protection_policy(name=None,recovery_vault_name=None,resource_group_name=None,opts=None):
    """
    Use this data source to access information about an existing Recovery Services VM Protection Policy.
    
    :param str name: Specifies the name of the Recovery Services VM Protection Policy.
    :param str recovery_vault_name: Specifies the name of the Recovery Services Vault.
    :param str resource_group_name: The name of the resource group in which the Recovery Services VM Protection Policy resides.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/recovery_services_protection_policy_vm.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['recoveryVaultName'] = recovery_vault_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:recoveryservices/getVMProtectionPolicy:getVMProtectionPolicy', __args__, opts=opts).value

    return AwaitableGetVMProtectionPolicyResult(
        name=__ret__.get('name'),
        recovery_vault_name=__ret__.get('recoveryVaultName'),
        resource_group_name=__ret__.get('resourceGroupName'),
        tags=__ret__.get('tags'),
        id=__ret__.get('id'))
