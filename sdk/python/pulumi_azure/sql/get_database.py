# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDatabaseResult:
    """
    A collection of values returned by getDatabase.
    """
    def __init__(__self__, collation=None, default_secondary_location=None, edition=None, elastic_pool_name=None, failover_group_id=None, id=None, location=None, name=None, read_scale=None, resource_group_name=None, server_name=None, tags=None):
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        __self__.collation = collation
        """
        The name of the collation. 
        """
        if default_secondary_location and not isinstance(default_secondary_location, str):
            raise TypeError("Expected argument 'default_secondary_location' to be a str")
        __self__.default_secondary_location = default_secondary_location
        """
        The default secondary location of the SQL Database.
        """
        if edition and not isinstance(edition, str):
            raise TypeError("Expected argument 'edition' to be a str")
        __self__.edition = edition
        """
        The edition of the database.
        """
        if elastic_pool_name and not isinstance(elastic_pool_name, str):
            raise TypeError("Expected argument 'elastic_pool_name' to be a str")
        __self__.elastic_pool_name = elastic_pool_name
        """
        The name of the elastic database pool the database belongs to.
        """
        if failover_group_id and not isinstance(failover_group_id, str):
            raise TypeError("Expected argument 'failover_group_id' to be a str")
        __self__.failover_group_id = failover_group_id
        """
        The ID of the failover group the database belongs to.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The location of the Resource Group in which the SQL Server exists.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the database.
        """
        if read_scale and not isinstance(read_scale, bool):
            raise TypeError("Expected argument 'read_scale' to be a bool")
        __self__.read_scale = read_scale
        """
        Indicate if read-only connections will be redirected to a high-available replica.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which the database resides. This will always be the same resource group as the Database Server.
        """
        if server_name and not isinstance(server_name, str):
            raise TypeError("Expected argument 'server_name' to be a str")
        __self__.server_name = server_name
        """
        The name of the SQL Server on which to create the database.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resource.
        """
class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            collation=self.collation,
            default_secondary_location=self.default_secondary_location,
            edition=self.edition,
            elastic_pool_name=self.elastic_pool_name,
            failover_group_id=self.failover_group_id,
            id=self.id,
            location=self.location,
            name=self.name,
            read_scale=self.read_scale,
            resource_group_name=self.resource_group_name,
            server_name=self.server_name,
            tags=self.tags)

def get_database(name=None,resource_group_name=None,server_name=None,tags=None,opts=None):
    """
    Use this data source to access information about an existing SQL Azure Database.




    :param str name: The name of the SQL Database.
    :param str resource_group_name: Specifies the name of the Resource Group where the Azure SQL Database exists.
    :param str server_name: The name of the SQL Server.
    :param dict tags: A mapping of tags assigned to the resource.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:sql/getDatabase:getDatabase', __args__, opts=opts).value

    return AwaitableGetDatabaseResult(
        collation=__ret__.get('collation'),
        default_secondary_location=__ret__.get('defaultSecondaryLocation'),
        edition=__ret__.get('edition'),
        elastic_pool_name=__ret__.get('elasticPoolName'),
        failover_group_id=__ret__.get('failoverGroupId'),
        id=__ret__.get('id'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        read_scale=__ret__.get('readScale'),
        resource_group_name=__ret__.get('resourceGroupName'),
        server_name=__ret__.get('serverName'),
        tags=__ret__.get('tags'))
