# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Snapshot(pulumi.CustomResource):
    account_name: pulumi.Output[str]
    """
    The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the NetApp Snapshot. Changing this forces a new resource to be created.
    """
    pool_name: pulumi.Output[str]
    """
    The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    volume_name: pulumi.Output[str]
    """
    The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, location=None, name=None, pool_name=None, resource_group_name=None, tags=None, volume_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a NetApp Snapshot.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Snapshot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pool_name: The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] volume_name: The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['location'] = location
            __props__['name'] = name
            if pool_name is None:
                raise TypeError("Missing required property 'pool_name'")
            __props__['pool_name'] = pool_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if volume_name is None:
                raise TypeError("Missing required property 'volume_name'")
            __props__['volume_name'] = volume_name
        super(Snapshot, __self__).__init__(
            'azure:netapp/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_name=None, location=None, name=None, pool_name=None, resource_group_name=None, tags=None, volume_name=None):
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Snapshot. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pool_name: The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] volume_name: The name of the NetApp volume in which the NetApp Snapshot should be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["location"] = location
        __props__["name"] = name
        __props__["pool_name"] = pool_name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        __props__["volume_name"] = volume_name
        return Snapshot(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

