# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetManagedDiskResult:
    """
    A collection of values returned by getManagedDisk.
    """
    def __init__(__self__, create_option=None, disk_size_gb=None, name=None, os_type=None, resource_group_name=None, source_resource_id=None, source_uri=None, storage_account_type=None, tags=None, zones=None, id=None):
        if create_option and not isinstance(create_option, str):
            raise TypeError("Expected argument 'create_option' to be a str")
        __self__.create_option = create_option
        if disk_size_gb and not isinstance(disk_size_gb, float):
            raise TypeError("Expected argument 'disk_size_gb' to be a float")
        __self__.disk_size_gb = disk_size_gb
        """
        The size of the managed disk in gigabytes.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        __self__.os_type = os_type
        """
        The operating system for managed disk. Valid values are `Linux` or `Windows`
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if source_resource_id and not isinstance(source_resource_id, str):
            raise TypeError("Expected argument 'source_resource_id' to be a str")
        __self__.source_resource_id = source_resource_id
        """
        ID of an existing managed disk that the current resource was created from.
        """
        if source_uri and not isinstance(source_uri, str):
            raise TypeError("Expected argument 'source_uri' to be a str")
        __self__.source_uri = source_uri
        """
        The source URI for the managed disk
        """
        if storage_account_type and not isinstance(storage_account_type, str):
            raise TypeError("Expected argument 'storage_account_type' to be a str")
        __self__.storage_account_type = storage_account_type
        """
        The storage account type for the managed disk.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        __self__.zones = zones
        """
        A collection containing the availability zone the managed disk is allocated in.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetManagedDiskResult(GetManagedDiskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedDiskResult(
            create_option=self.create_option,
            disk_size_gb=self.disk_size_gb,
            name=self.name,
            os_type=self.os_type,
            resource_group_name=self.resource_group_name,
            source_resource_id=self.source_resource_id,
            source_uri=self.source_uri,
            storage_account_type=self.storage_account_type,
            tags=self.tags,
            zones=self.zones,
            id=self.id)

def get_managed_disk(name=None,resource_group_name=None,tags=None,zones=None,opts=None):
    """
    Use this data source to access information about an existing Managed Disk.
    
    :param str name: Specifies the name of the Managed Disk.
    :param str resource_group_name: Specifies the name of the resource group.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/managed_disk.html.markdown.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['tags'] = tags
    __args__['zones'] = zones
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:compute/getManagedDisk:getManagedDisk', __args__, opts=opts).value

    return AwaitableGetManagedDiskResult(
        create_option=__ret__.get('createOption'),
        disk_size_gb=__ret__.get('diskSizeGb'),
        name=__ret__.get('name'),
        os_type=__ret__.get('osType'),
        resource_group_name=__ret__.get('resourceGroupName'),
        source_resource_id=__ret__.get('sourceResourceId'),
        source_uri=__ret__.get('sourceUri'),
        storage_account_type=__ret__.get('storageAccountType'),
        tags=__ret__.get('tags'),
        zones=__ret__.get('zones'),
        id=__ret__.get('id'))
