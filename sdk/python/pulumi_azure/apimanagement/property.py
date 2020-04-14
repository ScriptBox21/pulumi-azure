# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Property(pulumi.CustomResource):
    api_management_name: pulumi.Output[str]
    """
    The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
    """
    display_name: pulumi.Output[str]
    """
    The display name of this API Management Property.
    """
    name: pulumi.Output[str]
    """
    The name of the API Management Property. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
    """
    secret: pulumi.Output[bool]
    """
    Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
    """
    tags: pulumi.Output[list]
    """
    A list of tags to be applied to the API Management Property.
    """
    value: pulumi.Output[str]
    """
    The value of this API Management Property.
    """
    def __init__(__self__, resource_name, opts=None, api_management_name=None, display_name=None, name=None, resource_group_name=None, secret=None, tags=None, value=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an API Management Property.




        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The display name of this API Management Property.
        :param pulumi.Input[str] name: The name of the API Management Property. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] secret: Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
        :param pulumi.Input[list] tags: A list of tags to be applied to the API Management Property.
        :param pulumi.Input[str] value: The value of this API Management Property.
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

            if api_management_name is None:
                raise TypeError("Missing required property 'api_management_name'")
            __props__['api_management_name'] = api_management_name
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['secret'] = secret
            __props__['tags'] = tags
            if value is None:
                raise TypeError("Missing required property 'value'")
            __props__['value'] = value
        super(Property, __self__).__init__(
            'azure:apimanagement/property:Property',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, api_management_name=None, display_name=None, name=None, resource_group_name=None, secret=None, tags=None, value=None):
        """
        Get an existing Property resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_management_name: The name of the API Management Service in which the API Management Property should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The display name of this API Management Property.
        :param pulumi.Input[str] name: The name of the API Management Property. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the API Management Property should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] secret: Specifies whether the API Management Property is secret. Valid values are `true` or `false`. The default value is `false`.
        :param pulumi.Input[list] tags: A list of tags to be applied to the API Management Property.
        :param pulumi.Input[str] value: The value of this API Management Property.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_management_name"] = api_management_name
        __props__["display_name"] = display_name
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["secret"] = secret
        __props__["tags"] = tags
        __props__["value"] = value
        return Property(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

