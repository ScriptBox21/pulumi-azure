# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ExpressRouteCircuitPeering']


class ExpressRouteCircuitPeering(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 express_route_circuit_name: Optional[pulumi.Input[str]] = None,
                 microsoft_peering_config: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringMicrosoftPeeringConfigArgs']]] = None,
                 peer_asn: Optional[pulumi.Input[float]] = None,
                 peering_type: Optional[pulumi.Input[str]] = None,
                 primary_peer_address_prefix: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 route_filter_id: Optional[pulumi.Input[str]] = None,
                 secondary_peer_address_prefix: Optional[pulumi.Input[str]] = None,
                 shared_key: Optional[pulumi.Input[str]] = None,
                 vlan_id: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an ExpressRoute Circuit Peering.

        ## Example Usage
        ### Creating A Microsoft Peering)

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_express_route_circuit = azure.network.ExpressRouteCircuit("exampleExpressRouteCircuit",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            service_provider_name="Equinix",
            peering_location="Silicon Valley",
            bandwidth_in_mbps=50,
            sku=azure.network.ExpressRouteCircuitSkuArgs(
                tier="Standard",
                family="MeteredData",
            ),
            allow_classic_operations=False,
            tags={
                "environment": "Production",
            })
        example_express_route_circuit_peering = azure.network.ExpressRouteCircuitPeering("exampleExpressRouteCircuitPeering",
            peering_type="MicrosoftPeering",
            express_route_circuit_name=example_express_route_circuit.name,
            resource_group_name=example_resource_group.name,
            peer_asn=100,
            primary_peer_address_prefix="123.0.0.0/30",
            secondary_peer_address_prefix="123.0.0.4/30",
            vlan_id=300,
            microsoft_peering_config=azure.network.ExpressRouteCircuitPeeringMicrosoftPeeringConfigArgs(
                advertised_public_prefixes=["123.1.0.0/24"],
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] express_route_circuit_name: The name of the ExpressRoute Circuit in which to create the Peering.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringMicrosoftPeeringConfigArgs']] microsoft_peering_config: A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.
        :param pulumi.Input[float] peer_asn: The Either a 16-bit or a 32-bit ASN. Can either be public or private..
        :param pulumi.Input[str] peering_type: The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_peer_address_prefix: A `/30` subnet for the primary link.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the Express Route Circuit Peering. Changing this forces a new resource to be created.
        :param pulumi.Input[str] route_filter_id: The ID of the Route Filter. Only available when `peering_type` is set to `MicrosoftPeering`.
        :param pulumi.Input[str] secondary_peer_address_prefix: A `/30` subnet for the secondary link.
        :param pulumi.Input[str] shared_key: The shared key. Can be a maximum of 25 characters.
        :param pulumi.Input[float] vlan_id: A valid VLAN ID to establish this peering on.
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

            if express_route_circuit_name is None:
                raise TypeError("Missing required property 'express_route_circuit_name'")
            __props__['express_route_circuit_name'] = express_route_circuit_name
            __props__['microsoft_peering_config'] = microsoft_peering_config
            __props__['peer_asn'] = peer_asn
            if peering_type is None:
                raise TypeError("Missing required property 'peering_type'")
            __props__['peering_type'] = peering_type
            if primary_peer_address_prefix is None:
                raise TypeError("Missing required property 'primary_peer_address_prefix'")
            __props__['primary_peer_address_prefix'] = primary_peer_address_prefix
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['route_filter_id'] = route_filter_id
            if secondary_peer_address_prefix is None:
                raise TypeError("Missing required property 'secondary_peer_address_prefix'")
            __props__['secondary_peer_address_prefix'] = secondary_peer_address_prefix
            __props__['shared_key'] = shared_key
            if vlan_id is None:
                raise TypeError("Missing required property 'vlan_id'")
            __props__['vlan_id'] = vlan_id
            __props__['azure_asn'] = None
            __props__['primary_azure_port'] = None
            __props__['secondary_azure_port'] = None
        super(ExpressRouteCircuitPeering, __self__).__init__(
            'azure:network/expressRouteCircuitPeering:ExpressRouteCircuitPeering',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            azure_asn: Optional[pulumi.Input[float]] = None,
            express_route_circuit_name: Optional[pulumi.Input[str]] = None,
            microsoft_peering_config: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringMicrosoftPeeringConfigArgs']]] = None,
            peer_asn: Optional[pulumi.Input[float]] = None,
            peering_type: Optional[pulumi.Input[str]] = None,
            primary_azure_port: Optional[pulumi.Input[str]] = None,
            primary_peer_address_prefix: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            route_filter_id: Optional[pulumi.Input[str]] = None,
            secondary_azure_port: Optional[pulumi.Input[str]] = None,
            secondary_peer_address_prefix: Optional[pulumi.Input[str]] = None,
            shared_key: Optional[pulumi.Input[str]] = None,
            vlan_id: Optional[pulumi.Input[float]] = None) -> 'ExpressRouteCircuitPeering':
        """
        Get an existing ExpressRouteCircuitPeering resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] azure_asn: The ASN used by Azure.
        :param pulumi.Input[str] express_route_circuit_name: The name of the ExpressRoute Circuit in which to create the Peering.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringMicrosoftPeeringConfigArgs']] microsoft_peering_config: A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.
        :param pulumi.Input[float] peer_asn: The Either a 16-bit or a 32-bit ASN. Can either be public or private..
        :param pulumi.Input[str] peering_type: The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] primary_azure_port: The Primary Port used by Azure for this Peering.
        :param pulumi.Input[str] primary_peer_address_prefix: A `/30` subnet for the primary link.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to
               create the Express Route Circuit Peering. Changing this forces a new resource to be created.
        :param pulumi.Input[str] route_filter_id: The ID of the Route Filter. Only available when `peering_type` is set to `MicrosoftPeering`.
        :param pulumi.Input[str] secondary_azure_port: The Secondary Port used by Azure for this Peering.
        :param pulumi.Input[str] secondary_peer_address_prefix: A `/30` subnet for the secondary link.
        :param pulumi.Input[str] shared_key: The shared key. Can be a maximum of 25 characters.
        :param pulumi.Input[float] vlan_id: A valid VLAN ID to establish this peering on.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["azure_asn"] = azure_asn
        __props__["express_route_circuit_name"] = express_route_circuit_name
        __props__["microsoft_peering_config"] = microsoft_peering_config
        __props__["peer_asn"] = peer_asn
        __props__["peering_type"] = peering_type
        __props__["primary_azure_port"] = primary_azure_port
        __props__["primary_peer_address_prefix"] = primary_peer_address_prefix
        __props__["resource_group_name"] = resource_group_name
        __props__["route_filter_id"] = route_filter_id
        __props__["secondary_azure_port"] = secondary_azure_port
        __props__["secondary_peer_address_prefix"] = secondary_peer_address_prefix
        __props__["shared_key"] = shared_key
        __props__["vlan_id"] = vlan_id
        return ExpressRouteCircuitPeering(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureAsn")
    def azure_asn(self) -> float:
        """
        The ASN used by Azure.
        """
        return pulumi.get(self, "azure_asn")

    @property
    @pulumi.getter(name="expressRouteCircuitName")
    def express_route_circuit_name(self) -> str:
        """
        The name of the ExpressRoute Circuit in which to create the Peering.
        """
        return pulumi.get(self, "express_route_circuit_name")

    @property
    @pulumi.getter(name="microsoftPeeringConfig")
    def microsoft_peering_config(self) -> Optional['outputs.ExpressRouteCircuitPeeringMicrosoftPeeringConfig']:
        """
        A `microsoft_peering_config` block as defined below. Required when `peering_type` is set to `MicrosoftPeering`.
        """
        return pulumi.get(self, "microsoft_peering_config")

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> float:
        """
        The Either a 16-bit or a 32-bit ASN. Can either be public or private..
        """
        return pulumi.get(self, "peer_asn")

    @property
    @pulumi.getter(name="peeringType")
    def peering_type(self) -> str:
        """
        The type of the ExpressRoute Circuit Peering. Acceptable values include `AzurePrivatePeering`, `AzurePublicPeering` and `MicrosoftPeering`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "peering_type")

    @property
    @pulumi.getter(name="primaryAzurePort")
    def primary_azure_port(self) -> str:
        """
        The Primary Port used by Azure for this Peering.
        """
        return pulumi.get(self, "primary_azure_port")

    @property
    @pulumi.getter(name="primaryPeerAddressPrefix")
    def primary_peer_address_prefix(self) -> str:
        """
        A `/30` subnet for the primary link.
        """
        return pulumi.get(self, "primary_peer_address_prefix")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> str:
        """
        The name of the resource group in which to
        create the Express Route Circuit Peering. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="routeFilterId")
    def route_filter_id(self) -> Optional[str]:
        """
        The ID of the Route Filter. Only available when `peering_type` is set to `MicrosoftPeering`.
        """
        return pulumi.get(self, "route_filter_id")

    @property
    @pulumi.getter(name="secondaryAzurePort")
    def secondary_azure_port(self) -> str:
        """
        The Secondary Port used by Azure for this Peering.
        """
        return pulumi.get(self, "secondary_azure_port")

    @property
    @pulumi.getter(name="secondaryPeerAddressPrefix")
    def secondary_peer_address_prefix(self) -> str:
        """
        A `/30` subnet for the secondary link.
        """
        return pulumi.get(self, "secondary_peer_address_prefix")

    @property
    @pulumi.getter(name="sharedKey")
    def shared_key(self) -> Optional[str]:
        """
        The shared key. Can be a maximum of 25 characters.
        """
        return pulumi.get(self, "shared_key")

    @property
    @pulumi.getter(name="vlanId")
    def vlan_id(self) -> float:
        """
        A valid VLAN ID to establish this peering on.
        """
        return pulumi.get(self, "vlan_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

