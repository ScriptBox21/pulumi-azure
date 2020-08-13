# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class InterationServiceEnvironment(pulumi.CustomResource):
    access_endpoint_type: pulumi.Output[str]
    """
    The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
    """
    connector_endpoint_ip_addresses: pulumi.Output[list]
    """
    The list of access endpoint ip addresses of connector.
    """
    connector_outbound_ip_addresses: pulumi.Output[list]
    """
    The list of outgoing ip addresses of connector.
    """
    location: pulumi.Output[str]
    """
    The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
    """
    sku_name: pulumi.Output[str]
    """
    The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags which should be assigned to the Integration Service Environment.
    """
    virtual_network_subnet_ids: pulumi.Output[list]
    """
    A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
    """
    workflow_endpoint_ip_addresses: pulumi.Output[list]
    """
    The list of access endpoint ip addresses of workflow.
    """
    workflow_outbound_ip_addresses: pulumi.Output[list]
    """
    The list of outgoing ip addresses of workflow.
    """
    def __init__(__self__, resource_name, opts=None, access_endpoint_type=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None, virtual_network_subnet_ids=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages private and isolated Logic App instances within an Azure virtual network.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="westeurope")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            address_spaces=["10.0.0.0/22"])
        isesubnet1 = azure.network.Subnet("isesubnet1",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.0/26"],
            delegations=[{
                "name": "integrationServiceEnvironments",
                "serviceDelegation": {
                    "name": "Microsoft.Logic/integrationServiceEnvironments",
                },
            }])
        isesubnet2 = azure.network.Subnet("isesubnet2",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.64/26"])
        isesubnet3 = azure.network.Subnet("isesubnet3",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.128/26"])
        isesubnet4 = azure.network.Subnet("isesubnet4",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.1.192/26"])
        example_interation_service_environment = azure.logicapps.InterationServiceEnvironment("exampleInterationServiceEnvironment",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku_name="Developer_0",
            access_endpoint_type="Internal",
            virtual_network_subnet_ids=[
                isesubnet1.id,
                isesubnet2.id,
                isesubnet3.id,
                isesubnet4.id,
            ],
            tags={
                "environment": "development",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_endpoint_type: The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[str] location: The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[str] name: The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[str] sku_name: The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to the Integration Service Environment.
        :param pulumi.Input[list] virtual_network_subnet_ids: A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
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

            if access_endpoint_type is None:
                raise TypeError("Missing required property 'access_endpoint_type'")
            __props__['access_endpoint_type'] = access_endpoint_type
            __props__['location'] = location
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku_name'] = sku_name
            __props__['tags'] = tags
            if virtual_network_subnet_ids is None:
                raise TypeError("Missing required property 'virtual_network_subnet_ids'")
            __props__['virtual_network_subnet_ids'] = virtual_network_subnet_ids
            __props__['connector_endpoint_ip_addresses'] = None
            __props__['connector_outbound_ip_addresses'] = None
            __props__['workflow_endpoint_ip_addresses'] = None
            __props__['workflow_outbound_ip_addresses'] = None
        super(InterationServiceEnvironment, __self__).__init__(
            'azure:logicapps/interationServiceEnvironment:InterationServiceEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_endpoint_type=None, connector_endpoint_ip_addresses=None, connector_outbound_ip_addresses=None, location=None, name=None, resource_group_name=None, sku_name=None, tags=None, virtual_network_subnet_ids=None, workflow_endpoint_ip_addresses=None, workflow_outbound_ip_addresses=None):
        """
        Get an existing InterationServiceEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_endpoint_type: The type of access endpoint to use for the Integration Service Environment. Possible Values are `Internal` and `External`. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[list] connector_endpoint_ip_addresses: The list of access endpoint ip addresses of connector.
        :param pulumi.Input[list] connector_outbound_ip_addresses: The list of outgoing ip addresses of connector.
        :param pulumi.Input[str] location: The Azure Region where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[str] name: The name of the Integration Service Environment. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Integration Service Environment should exist. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[str] sku_name: The sku name and capacity of the Integration Service Environment. Possible Values for `sku` element are `Developer` and `Premium` and possible values for the `capacity` element are from `0` to `10`.  Defaults to `sku` of `Developer` with a `Capacity` of `0` (e.g. `Developer_0`). Changing this forces a new Integration Service Environment to be created when `sku` element is not the same with existing one.
        :param pulumi.Input[dict] tags: A mapping of tags which should be assigned to the Integration Service Environment.
        :param pulumi.Input[list] virtual_network_subnet_ids: A list of virtual network subnet ids to be used by Integration Service Environment. Exactly four distinct ids to subnets must be provided. Changing this forces a new Integration Service Environment to be created.
        :param pulumi.Input[list] workflow_endpoint_ip_addresses: The list of access endpoint ip addresses of workflow.
        :param pulumi.Input[list] workflow_outbound_ip_addresses: The list of outgoing ip addresses of workflow.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_endpoint_type"] = access_endpoint_type
        __props__["connector_endpoint_ip_addresses"] = connector_endpoint_ip_addresses
        __props__["connector_outbound_ip_addresses"] = connector_outbound_ip_addresses
        __props__["location"] = location
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sku_name"] = sku_name
        __props__["tags"] = tags
        __props__["virtual_network_subnet_ids"] = virtual_network_subnet_ids
        __props__["workflow_endpoint_ip_addresses"] = workflow_endpoint_ip_addresses
        __props__["workflow_outbound_ip_addresses"] = workflow_outbound_ip_addresses
        return InterationServiceEnvironment(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
