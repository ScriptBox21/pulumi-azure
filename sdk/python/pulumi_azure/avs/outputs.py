# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'PrivateCloudCircuit',
    'PrivateCloudManagementCluster',
    'GetPrivateCloudCircuitResult',
    'GetPrivateCloudManagementClusterResult',
]

@pulumi.output_type
class PrivateCloudCircuit(dict):
    def __init__(__self__, *,
                 express_route_id: Optional[str] = None,
                 express_route_private_peering_id: Optional[str] = None,
                 primary_subnet_cidr: Optional[str] = None,
                 secondary_subnet_cidr: Optional[str] = None):
        """
        :param str express_route_id: The ID of the ExpressRoute Circuit.
        :param str express_route_private_peering_id: The ID of the ExpressRoute Circuit private peering.
        :param str primary_subnet_cidr: The CIDR of the primary subnet.
        :param str secondary_subnet_cidr: The CIDR of the secondary subnet.
        """
        if express_route_id is not None:
            pulumi.set(__self__, "express_route_id", express_route_id)
        if express_route_private_peering_id is not None:
            pulumi.set(__self__, "express_route_private_peering_id", express_route_private_peering_id)
        if primary_subnet_cidr is not None:
            pulumi.set(__self__, "primary_subnet_cidr", primary_subnet_cidr)
        if secondary_subnet_cidr is not None:
            pulumi.set(__self__, "secondary_subnet_cidr", secondary_subnet_cidr)

    @property
    @pulumi.getter(name="expressRouteId")
    def express_route_id(self) -> Optional[str]:
        """
        The ID of the ExpressRoute Circuit.
        """
        return pulumi.get(self, "express_route_id")

    @property
    @pulumi.getter(name="expressRoutePrivatePeeringId")
    def express_route_private_peering_id(self) -> Optional[str]:
        """
        The ID of the ExpressRoute Circuit private peering.
        """
        return pulumi.get(self, "express_route_private_peering_id")

    @property
    @pulumi.getter(name="primarySubnetCidr")
    def primary_subnet_cidr(self) -> Optional[str]:
        """
        The CIDR of the primary subnet.
        """
        return pulumi.get(self, "primary_subnet_cidr")

    @property
    @pulumi.getter(name="secondarySubnetCidr")
    def secondary_subnet_cidr(self) -> Optional[str]:
        """
        The CIDR of the secondary subnet.
        """
        return pulumi.get(self, "secondary_subnet_cidr")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PrivateCloudManagementCluster(dict):
    def __init__(__self__, *,
                 size: int,
                 hosts: Optional[Sequence[str]] = None,
                 id: Optional[int] = None):
        """
        :param int size: The size of the management cluster. This field can not updated with `internet_connection_enabled` together.
        :param Sequence[str] hosts: A list of hosts in the management cluster.
        :param int id: The ID of the  management cluster.
        """
        pulumi.set(__self__, "size", size)
        if hosts is not None:
            pulumi.set(__self__, "hosts", hosts)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size of the management cluster. This field can not updated with `internet_connection_enabled` together.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def hosts(self) -> Optional[Sequence[str]]:
        """
        A list of hosts in the management cluster.
        """
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter
    def id(self) -> Optional[int]:
        """
        The ID of the  management cluster.
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetPrivateCloudCircuitResult(dict):
    def __init__(__self__, *,
                 express_route_id: str,
                 express_route_private_peering_id: str,
                 primary_subnet_cidr: str,
                 secondary_subnet_cidr: str):
        """
        :param str express_route_id: The ID of the ExpressRoute Circuit.
        :param str express_route_private_peering_id: The ID of the ExpressRoute Circuit private peering.
        :param str primary_subnet_cidr: The CIDR of the primary subnet.
        :param str secondary_subnet_cidr: The CIDR of the secondary subnet.
        """
        pulumi.set(__self__, "express_route_id", express_route_id)
        pulumi.set(__self__, "express_route_private_peering_id", express_route_private_peering_id)
        pulumi.set(__self__, "primary_subnet_cidr", primary_subnet_cidr)
        pulumi.set(__self__, "secondary_subnet_cidr", secondary_subnet_cidr)

    @property
    @pulumi.getter(name="expressRouteId")
    def express_route_id(self) -> str:
        """
        The ID of the ExpressRoute Circuit.
        """
        return pulumi.get(self, "express_route_id")

    @property
    @pulumi.getter(name="expressRoutePrivatePeeringId")
    def express_route_private_peering_id(self) -> str:
        """
        The ID of the ExpressRoute Circuit private peering.
        """
        return pulumi.get(self, "express_route_private_peering_id")

    @property
    @pulumi.getter(name="primarySubnetCidr")
    def primary_subnet_cidr(self) -> str:
        """
        The CIDR of the primary subnet.
        """
        return pulumi.get(self, "primary_subnet_cidr")

    @property
    @pulumi.getter(name="secondarySubnetCidr")
    def secondary_subnet_cidr(self) -> str:
        """
        The CIDR of the secondary subnet.
        """
        return pulumi.get(self, "secondary_subnet_cidr")


@pulumi.output_type
class GetPrivateCloudManagementClusterResult(dict):
    def __init__(__self__, *,
                 hosts: Sequence[str],
                 id: int,
                 size: int):
        """
        :param Sequence[str] hosts: The list of the hosts in the management cluster.
        :param int id: The ID of the management cluster.
        :param int size: The size of the management cluster.
        """
        pulumi.set(__self__, "hosts", hosts)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter
    def hosts(self) -> Sequence[str]:
        """
        The list of the hosts in the management cluster.
        """
        return pulumi.get(self, "hosts")

    @property
    @pulumi.getter
    def id(self) -> int:
        """
        The ID of the management cluster.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size of the management cluster.
        """
        return pulumi.get(self, "size")

