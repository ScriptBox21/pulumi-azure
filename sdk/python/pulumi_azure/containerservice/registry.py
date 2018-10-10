# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Registry(pulumi.CustomResource):
    """
    Manages an Azure Container Registry.
    
    ~> **Note:** All arguments including the access key will be stored in the raw state as plain-text.
    [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    """
    def __init__(__self__, __name__, __opts__=None, admin_enabled=None, location=None, name=None, resource_group_name=None, sku=None, storage_account=None, storage_account_id=None, tags=None):
        """Create a Registry resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if admin_enabled and not isinstance(admin_enabled, bool):
            raise TypeError('Expected property admin_enabled to be a bool')
        __self__.admin_enabled = admin_enabled
        """
        Specifies whether the admin user is enabled. Defaults to `false`.
        """
        __props__['adminEnabled'] = admin_enabled

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        Specifies the name of the Container Registry. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to create the Container Registry. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if sku and not isinstance(sku, basestring):
            raise TypeError('Expected property sku to be a basestring')
        __self__.sku = sku
        """
        The SKU name of the the container registry. Possible values are `Classic` (which was previously `Basic`), `Basic`, `Standard` and `Premium`.
        """
        __props__['sku'] = sku

        if storage_account and not isinstance(storage_account, dict):
            raise TypeError('Expected property storage_account to be a dict')
        __self__.storage_account = storage_account
        __props__['storageAccount'] = storage_account

        if storage_account_id and not isinstance(storage_account_id, basestring):
            raise TypeError('Expected property storage_account_id to be a basestring')
        __self__.storage_account_id = storage_account_id
        """
        The ID of a Storage Account which must be located in the same Azure Region as the Container Registry.
        """
        __props__['storageAccountId'] = storage_account_id

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        __self__.admin_password = pulumi.runtime.UNKNOWN
        """
        The Password associated with the Container Registry Admin account - if the admin account is enabled.
        """
        __self__.admin_username = pulumi.runtime.UNKNOWN
        """
        The Username associated with the Container Registry Admin account - if the admin account is enabled.
        """
        __self__.login_server = pulumi.runtime.UNKNOWN
        """
        The URL that can be used to log into the container registry.
        """

        super(Registry, __self__).__init__(
            'azure:containerservice/registry:Registry',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'adminEnabled' in outs:
            self.admin_enabled = outs['adminEnabled']
        if 'adminPassword' in outs:
            self.admin_password = outs['adminPassword']
        if 'adminUsername' in outs:
            self.admin_username = outs['adminUsername']
        if 'location' in outs:
            self.location = outs['location']
        if 'loginServer' in outs:
            self.login_server = outs['loginServer']
        if 'name' in outs:
            self.name = outs['name']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'sku' in outs:
            self.sku = outs['sku']
        if 'storageAccount' in outs:
            self.storage_account = outs['storageAccount']
        if 'storageAccountId' in outs:
            self.storage_account_id = outs['storageAccountId']
        if 'tags' in outs:
            self.tags = outs['tags']