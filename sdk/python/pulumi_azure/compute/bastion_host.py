# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class BastionHost(pulumi.CustomResource):
    dns_name: pulumi.Output[str]
    """
    The FQDN for the Bastion Host.
    """
    ip_configuration: pulumi.Output[dict]
    """
    A `ip_configuration` block as defined below.

      * `name` (`str`) - The name of the IP configuration.
      * `publicIpAddressId` (`str`) - Reference to a Public IP Address to associate with this Bastion Host.
      * `subnet_id` (`str`) - Reference to a subnet in which this Bastion Host has been created.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Bastion Host.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, ip_configuration=None, location=None, name=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Bastion Host.

        > **Note:** Bastion Hosts are a preview feature in Azure, and therefore are only supported in a select number of regions. [Read more](https://docs.microsoft.com/en-us/azure/bastion/bastion-faq).



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] ip_configuration: A `ip_configuration` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bastion Host.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **ip_configuration** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the IP configuration.
          * `publicIpAddressId` (`pulumi.Input[str]`) - Reference to a Public IP Address to associate with this Bastion Host.
          * `subnet_id` (`pulumi.Input[str]`) - Reference to a subnet in which this Bastion Host has been created.
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

            __props__['ip_configuration'] = ip_configuration
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['dns_name'] = None
        super(BastionHost, __self__).__init__(
            'azure:compute/bastionHost:BastionHost',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, dns_name=None, ip_configuration=None, location=None, name=None, resource_group_name=None, tags=None):
        """
        Get an existing BastionHost resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_name: The FQDN for the Bastion Host.
        :param pulumi.Input[dict] ip_configuration: A `ip_configuration` block as defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Bastion Host. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Bastion Host.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        The **ip_configuration** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the IP configuration.
          * `publicIpAddressId` (`pulumi.Input[str]`) - Reference to a Public IP Address to associate with this Bastion Host.
          * `subnet_id` (`pulumi.Input[str]`) - Reference to a subnet in which this Bastion Host has been created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dns_name"] = dns_name
        __props__["ip_configuration"] = ip_configuration
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return BastionHost(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

