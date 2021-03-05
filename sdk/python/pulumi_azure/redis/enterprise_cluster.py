# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['EnterpriseCluster']


class EnterpriseCluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Redis Enterprise Cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_enterprise_cluster = azure.redis.EnterpriseCluster("exampleEnterpriseCluster",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            sku_name="EnterpriseFlash_F300-3")
        ```

        ## Import

        Redis Enterprise Clusters can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:redis/enterpriseCluster:EnterpriseCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Cache/redisEnterprise/cluster1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[str] name: The name which should be used for this Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[str] sku_name: The `sku_name` is comprised of two segments separated by a hyphen (e.g. `Enterprise_E10-2`). The first segment of the `sku_name` defines the `name` of the sku, possible values are `Enterprise_E10`, `Enterprise_E20"`, `Enterprise_E50`, `Enterprise_E100`, `EnterpriseFlash_F300`, `EnterpriseFlash_F700` or `EnterpriseFlash_F1500`. The second segment defines the `capacity` of the `sku_name`, possible values for `Enteprise` skus are (`2`, `4`, `6`, ...). Possible values for `EnterpriseFlash` skus are (`3`, `9`, `15`, ...). Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: A list of a one or more Availability Zones, where the Redis Cache should be allocated. Possible values are: `1`, `2` and `3`. Changing this forces a new Redis Enterprise Cluster to be created.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku_name is None and not opts.urn:
                raise TypeError("Missing required property 'sku_name'")
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['hostname'] = None
            __props__['version'] = None
        super(EnterpriseCluster, __self__).__init__(
            'azure:redis/enterpriseCluster:EnterpriseCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sku_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            version: Optional[pulumi.Input[str]] = None,
            zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'EnterpriseCluster':
        """
        Get an existing EnterpriseCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hostname: DNS name of the cluster endpoint.
        :param pulumi.Input[str] location: The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[str] name: The name which should be used for this Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[str] sku_name: The `sku_name` is comprised of two segments separated by a hyphen (e.g. `Enterprise_E10-2`). The first segment of the `sku_name` defines the `name` of the sku, possible values are `Enterprise_E10`, `Enterprise_E20"`, `Enterprise_E50`, `Enterprise_E100`, `EnterpriseFlash_F300`, `EnterpriseFlash_F700` or `EnterpriseFlash_F1500`. The second segment defines the `capacity` of the `sku_name`, possible values for `Enteprise` skus are (`2`, `4`, `6`, ...). Possible values for `EnterpriseFlash` skus are (`3`, `9`, `15`, ...). Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
        :param pulumi.Input[str] version: Version of redis the cluster supports, e.g. '6'.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] zones: A list of a one or more Availability Zones, where the Redis Cache should be allocated. Possible values are: `1`, `2` and `3`. Changing this forces a new Redis Enterprise Cluster to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["hostname"] = hostname
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        __props__["version"] = version
        __props__["zones"] = zones
        return EnterpriseCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        DNS name of the cluster endpoint.
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Azure Region where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Redis Enterprise Cluster should exist. Changing this forces a new Redis Enterprise Cluster to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="skuName")
    def sku_name(self) -> pulumi.Output[str]:
        """
        The `sku_name` is comprised of two segments separated by a hyphen (e.g. `Enterprise_E10-2`). The first segment of the `sku_name` defines the `name` of the sku, possible values are `Enterprise_E10`, `Enterprise_E20"`, `Enterprise_E50`, `Enterprise_E100`, `EnterpriseFlash_F300`, `EnterpriseFlash_F700` or `EnterpriseFlash_F1500`. The second segment defines the `capacity` of the `sku_name`, possible values for `Enteprise` skus are (`2`, `4`, `6`, ...). Possible values for `EnterpriseFlash` skus are (`3`, `9`, `15`, ...). Changing this forces a new Redis Enterprise Cluster to be created.
        """
        return pulumi.get(self, "sku_name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Redis Enterprise Cluster. Changing this forces a new Redis Enterprise Cluster to be created.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Version of redis the cluster supports, e.g. '6'.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of a one or more Availability Zones, where the Redis Cache should be allocated. Possible values are: `1`, `2` and `3`. Changing this forces a new Redis Enterprise Cluster to be created.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
