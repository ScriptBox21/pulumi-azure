# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Group(pulumi.CustomResource):
    containers: pulumi.Output[list]
    """
    The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
    """
    diagnostics: pulumi.Output[dict]
    """
    A `diagnostics` block as documented below.
    """
    dns_name_label: pulumi.Output[str]
    """
    The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
    """
    fqdn: pulumi.Output[str]
    """
    The FQDN of the container group derived from `dns_name_label`.
    """
    identity: pulumi.Output[dict]
    """
    An `identity` block as defined below.
    """
    image_registry_credentials: pulumi.Output[list]
    """
    A `image_registry_credential` block as documented below. Changing this forces a new resource to be created.
    """
    ip_address: pulumi.Output[str]
    """
    The IP address allocated to the container group.
    """
    ip_address_type: pulumi.Output[str]
    """
    Specifies the ip address type of the container. `Public` is the only acceptable value at this time. Changing this forces a new resource to be created.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the Container Group. Changing this forces a new resource to be created.
    """
    os_type: pulumi.Output[str]
    """
    The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
    """
    restart_policy: pulumi.Output[str]
    """
    Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, containers=None, diagnostics=None, dns_name_label=None, identity=None, image_registry_credentials=None, ip_address_type=None, location=None, name=None, os_type=None, resource_group_name=None, restart_policy=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manage as an Azure Container Group instance.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] containers: The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] diagnostics: A `diagnostics` block as documented below.
        :param pulumi.Input[str] dns_name_label: The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] identity: An `identity` block as defined below.
        :param pulumi.Input[list] image_registry_credentials: A `image_registry_credential` block as documented below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ip_address_type: Specifies the ip address type of the container. `Public` is the only acceptable value at this time. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Container Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] restart_policy: Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/container_group.html.markdown.
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

            if containers is None:
                raise TypeError("Missing required property 'containers'")
            __props__['containers'] = containers
            __props__['diagnostics'] = diagnostics
            __props__['dns_name_label'] = dns_name_label
            __props__['identity'] = identity
            __props__['image_registry_credentials'] = image_registry_credentials
            __props__['ip_address_type'] = ip_address_type
            __props__['location'] = location
            __props__['name'] = name
            if os_type is None:
                raise TypeError("Missing required property 'os_type'")
            __props__['os_type'] = os_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['restart_policy'] = restart_policy
            __props__['tags'] = tags
            __props__['fqdn'] = None
            __props__['ip_address'] = None
        super(Group, __self__).__init__(
            'azure:containerservice/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, containers=None, diagnostics=None, dns_name_label=None, fqdn=None, identity=None, image_registry_credentials=None, ip_address=None, ip_address_type=None, location=None, name=None, os_type=None, resource_group_name=None, restart_policy=None, tags=None):
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] containers: The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] diagnostics: A `diagnostics` block as documented below.
        :param pulumi.Input[str] dns_name_label: The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
        :param pulumi.Input[str] fqdn: The FQDN of the container group derived from `dns_name_label`.
        :param pulumi.Input[dict] identity: An `identity` block as defined below.
        :param pulumi.Input[list] image_registry_credentials: A `image_registry_credential` block as documented below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] ip_address: The IP address allocated to the container group.
        :param pulumi.Input[str] ip_address_type: Specifies the ip address type of the container. `Public` is the only acceptable value at this time. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the Container Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] os_type: The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] restart_policy: Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource. Changing this forces a new resource to be created.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/container_group.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["containers"] = containers
        __props__["diagnostics"] = diagnostics
        __props__["dns_name_label"] = dns_name_label
        __props__["fqdn"] = fqdn
        __props__["identity"] = identity
        __props__["image_registry_credentials"] = image_registry_credentials
        __props__["ip_address"] = ip_address
        __props__["ip_address_type"] = ip_address_type
        __props__["location"] = location
        __props__["name"] = name
        __props__["os_type"] = os_type
        __props__["resource_group_name"] = resource_group_name
        __props__["restart_policy"] = restart_policy
        __props__["tags"] = tags
        return Group(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

