# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['BgpConnection']


class BgpConnection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_asn: Optional[pulumi.Input[int]] = None,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 virtual_hub_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Bgp Connection for a Virtual Hub.

        ## Import

        Virtual Hub Bgp Connections can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:network/bgpConnection:BgpConnection example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/virtualHub1/bgpConnections/connection1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[int] peer_asn: The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] peer_ip: The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_hub_id: The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
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

            __props__['name'] = name
            if peer_asn is None:
                raise TypeError("Missing required property 'peer_asn'")
            __props__['peer_asn'] = peer_asn
            if peer_ip is None:
                raise TypeError("Missing required property 'peer_ip'")
            __props__['peer_ip'] = peer_ip
            if virtual_hub_id is None:
                raise TypeError("Missing required property 'virtual_hub_id'")
            __props__['virtual_hub_id'] = virtual_hub_id
        super(BgpConnection, __self__).__init__(
            'azure:network/bgpConnection:BgpConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            peer_asn: Optional[pulumi.Input[int]] = None,
            peer_ip: Optional[pulumi.Input[str]] = None,
            virtual_hub_id: Optional[pulumi.Input[str]] = None) -> 'BgpConnection':
        """
        Get an existing BgpConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[int] peer_asn: The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] peer_ip: The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        :param pulumi.Input[str] virtual_hub_id: The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["peer_asn"] = peer_asn
        __props__["peer_ip"] = peer_ip
        __props__["virtual_hub_id"] = virtual_hub_id
        return BgpConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> pulumi.Output[int]:
        """
        The peer autonomous system number for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "peer_asn")

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> pulumi.Output[str]:
        """
        The peer ip address for the Virtual Hub Bgp Connection. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "peer_ip")

    @property
    @pulumi.getter(name="virtualHubId")
    def virtual_hub_id(self) -> pulumi.Output[str]:
        """
        The ID of the Virtual Hub within which this Bgp connection should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_hub_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

