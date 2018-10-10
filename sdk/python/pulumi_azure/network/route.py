# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Route(pulumi.CustomResource):
    """
    Manages a Route within a Route Table.
    """
    def __init__(__self__, __name__, __opts__=None, address_prefix=None, name=None, next_hop_in_ip_address=None, next_hop_type=None, resource_group_name=None, route_table_name=None):
        """Create a Route resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not address_prefix:
            raise TypeError('Missing required property address_prefix')
        elif not isinstance(address_prefix, basestring):
            raise TypeError('Expected property address_prefix to be a basestring')
        __self__.address_prefix = address_prefix
        """
        The destination CIDR to which the route applies, such as `10.1.0.0/16`
        """
        __props__['addressPrefix'] = address_prefix

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the route. Changing this forces a new resource to be created.
        """
        __props__['name'] = name

        if next_hop_in_ip_address and not isinstance(next_hop_in_ip_address, basestring):
            raise TypeError('Expected property next_hop_in_ip_address to be a basestring')
        __self__.next_hop_in_ip_address = next_hop_in_ip_address
        """
        Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is `VirtualAppliance`.
        """
        __props__['nextHopInIpAddress'] = next_hop_in_ip_address

        if not next_hop_type:
            raise TypeError('Missing required property next_hop_type')
        elif not isinstance(next_hop_type, basestring):
            raise TypeError('Expected property next_hop_type to be a basestring')
        __self__.next_hop_type = next_hop_type
        """
        The type of Azure hop the packet should be sent to. Possible values are `VirtualNetworkGateway`, `VnetLocal`, `Internet`, `VirtualAppliance` and `None`
        """
        __props__['nextHopType'] = next_hop_type

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to create the route. Changing this forces a new resource to be created.
        """
        __props__['resourceGroupName'] = resource_group_name

        if not route_table_name:
            raise TypeError('Missing required property route_table_name')
        elif not isinstance(route_table_name, basestring):
            raise TypeError('Expected property route_table_name to be a basestring')
        __self__.route_table_name = route_table_name
        """
        The name of the route table within which create the route. Changing this forces a new resource to be created.
        """
        __props__['routeTableName'] = route_table_name

        super(Route, __self__).__init__(
            'azure:network/route:Route',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'addressPrefix' in outs:
            self.address_prefix = outs['addressPrefix']
        if 'name' in outs:
            self.name = outs['name']
        if 'nextHopInIpAddress' in outs:
            self.next_hop_in_ip_address = outs['nextHopInIpAddress']
        if 'nextHopType' in outs:
            self.next_hop_type = outs['nextHopType']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'routeTableName' in outs:
            self.route_table_name = outs['routeTableName']