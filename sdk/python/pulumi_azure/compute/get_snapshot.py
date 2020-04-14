# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSnapshotResult:
    """
    A collection of values returned by getSnapshot.
    """
    def __init__(__self__, creation_option=None, disk_size_gb=None, encryption_settings=None, id=None, name=None, os_type=None, resource_group_name=None, source_resource_id=None, source_uri=None, storage_account_id=None, time_created=None):
        if creation_option and not isinstance(creation_option, str):
            raise TypeError("Expected argument 'creation_option' to be a str")
        __self__.creation_option = creation_option
        if disk_size_gb and not isinstance(disk_size_gb, float):
            raise TypeError("Expected argument 'disk_size_gb' to be a float")
        __self__.disk_size_gb = disk_size_gb
        """
        The size of the Snapshotted Disk in GB.
        """
        if encryption_settings and not isinstance(encryption_settings, list):
            raise TypeError("Expected argument 'encryption_settings' to be a list")
        __self__.encryption_settings = encryption_settings
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if os_type and not isinstance(os_type, str):
            raise TypeError("Expected argument 'os_type' to be a str")
        __self__.os_type = os_type
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if source_resource_id and not isinstance(source_resource_id, str):
            raise TypeError("Expected argument 'source_resource_id' to be a str")
        __self__.source_resource_id = source_resource_id
        """
        The reference to an existing snapshot.
        """
        if source_uri and not isinstance(source_uri, str):
            raise TypeError("Expected argument 'source_uri' to be a str")
        __self__.source_uri = source_uri
        """
        The URI to a Managed or Unmanaged Disk.
        """
        if storage_account_id and not isinstance(storage_account_id, str):
            raise TypeError("Expected argument 'storage_account_id' to be a str")
        __self__.storage_account_id = storage_account_id
        """
        The ID of an storage account.
        """
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        __self__.time_created = time_created
class AwaitableGetSnapshotResult(GetSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotResult(
            creation_option=self.creation_option,
            disk_size_gb=self.disk_size_gb,
            encryption_settings=self.encryption_settings,
            id=self.id,
            name=self.name,
            os_type=self.os_type,
            resource_group_name=self.resource_group_name,
            source_resource_id=self.source_resource_id,
            source_uri=self.source_uri,
            storage_account_id=self.storage_account_id,
            time_created=self.time_created)

def get_snapshot(name=None,resource_group_name=None,opts=None):
    """
    Use this data source to access information about an existing Snapshot.




    :param str name: Specifies the name of the Snapshot.
    :param str resource_group_name: Specifies the name of the resource group the Snapshot is located in.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:compute/getSnapshot:getSnapshot', __args__, opts=opts).value

    return AwaitableGetSnapshotResult(
        creation_option=__ret__.get('creationOption'),
        disk_size_gb=__ret__.get('diskSizeGb'),
        encryption_settings=__ret__.get('encryptionSettings'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        os_type=__ret__.get('osType'),
        resource_group_name=__ret__.get('resourceGroupName'),
        source_resource_id=__ret__.get('sourceResourceId'),
        source_uri=__ret__.get('sourceUri'),
        storage_account_id=__ret__.get('storageAccountId'),
        time_created=__ret__.get('timeCreated'))
