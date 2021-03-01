# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Volume']


class Volume(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 data_protection_replication: Optional[pulumi.Input[pulumi.InputType['VolumeDataProtectionReplicationArgs']]] = None,
                 export_policy_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeExportPolicyRuleArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pool_name: Optional[pulumi.Input[str]] = None,
                 protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_level: Optional[pulumi.Input[str]] = None,
                 storage_quota_in_gb: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 volume_path: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a NetApp Volume.

        ## NetApp Volume Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_virtual_network = azure.network.VirtualNetwork("exampleVirtualNetwork",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            address_spaces=["10.0.0.0/16"])
        example_subnet = azure.network.Subnet("exampleSubnet",
            resource_group_name=example_resource_group.name,
            virtual_network_name=example_virtual_network.name,
            address_prefixes=["10.0.2.0/24"],
            delegations=[azure.network.SubnetDelegationArgs(
                name="netapp",
                service_delegation=azure.network.SubnetDelegationServiceDelegationArgs(
                    name="Microsoft.Netapp/volumes",
                    actions=[
                        "Microsoft.Network/networkinterfaces/*",
                        "Microsoft.Network/virtualNetworks/subnets/join/action",
                    ],
                ),
            )])
        example_account = azure.netapp.Account("exampleAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name)
        example_pool = azure.netapp.Pool("examplePool",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            account_name=example_account.name,
            service_level="Premium",
            size_in_tb=4)
        example_volume = azure.netapp.Volume("exampleVolume",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            account_name=example_account.name,
            pool_name=example_pool.name,
            volume_path="my-unique-file-path",
            service_level="Premium",
            subnet_id=example_subnet.id,
            protocols=["NFSv4.1"],
            storage_quota_in_gb=100,
            data_protection_replication=azure.netapp.VolumeDataProtectionReplicationArgs(
                endpoint_type="dst",
                remote_volume_location=azurerm_resource_group["example_primary"]["location"],
                remote_volume_resource_id=azurerm_netapp_volume["example_primary"]["id"],
                replication_frequency="_10minutely",
            ))
        ```

        ## Import

        NetApp Volumes can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:netapp/volume:Volume example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeExportPolicyRuleArgs']]]] export_policy_rules: One or more `export_policy_rule` block defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the NetApp Volume. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pool_name: The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] protocols: The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[int] storage_quota_in_gb: The maximum Storage Quota allowed for a file system in Gigabytes.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] volume_path: A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
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

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['data_protection_replication'] = data_protection_replication
            __props__['export_policy_rules'] = export_policy_rules
            __props__['location'] = location
            __props__['name'] = name
            if pool_name is None and not opts.urn:
                raise TypeError("Missing required property 'pool_name'")
            __props__['pool_name'] = pool_name
            __props__['protocols'] = protocols
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_level is None and not opts.urn:
                raise TypeError("Missing required property 'service_level'")
            __props__['service_level'] = service_level
            if storage_quota_in_gb is None and not opts.urn:
                raise TypeError("Missing required property 'storage_quota_in_gb'")
            __props__['storage_quota_in_gb'] = storage_quota_in_gb
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            if volume_path is None and not opts.urn:
                raise TypeError("Missing required property 'volume_path'")
            __props__['volume_path'] = volume_path
            __props__['mount_ip_addresses'] = None
        super(Volume, __self__).__init__(
            'azure:netapp/volume:Volume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            data_protection_replication: Optional[pulumi.Input[pulumi.InputType['VolumeDataProtectionReplicationArgs']]] = None,
            export_policy_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeExportPolicyRuleArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            mount_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pool_name: Optional[pulumi.Input[str]] = None,
            protocols: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            service_level: Optional[pulumi.Input[str]] = None,
            storage_quota_in_gb: Optional[pulumi.Input[int]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            volume_path: Optional[pulumi.Input[str]] = None) -> 'Volume':
        """
        Get an existing Volume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VolumeExportPolicyRuleArgs']]]] export_policy_rules: One or more `export_policy_rule` block defined below.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] mount_ip_addresses: A list of IPv4 Addresses which should be used to mount the volume.
        :param pulumi.Input[str] name: The name of the NetApp Volume. Changing this forces a new resource to be created.
        :param pulumi.Input[str] pool_name: The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] protocols: The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
        :param pulumi.Input[str] resource_group_name: The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] service_level: The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        :param pulumi.Input[int] storage_quota_in_gb: The maximum Storage Quota allowed for a file system in Gigabytes.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] volume_path: A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_name"] = account_name
        __props__["data_protection_replication"] = data_protection_replication
        __props__["export_policy_rules"] = export_policy_rules
        __props__["location"] = location
        __props__["mount_ip_addresses"] = mount_ip_addresses
        __props__["name"] = name
        __props__["pool_name"] = pool_name
        __props__["protocols"] = protocols
        __props__["resource_group_name"] = resource_group_name
        __props__["service_level"] = service_level
        __props__["storage_quota_in_gb"] = storage_quota_in_gb
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        __props__["volume_path"] = volume_path
        return Volume(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        The name of the NetApp account in which the NetApp Pool should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="dataProtectionReplication")
    def data_protection_replication(self) -> pulumi.Output[Optional['outputs.VolumeDataProtectionReplication']]:
        return pulumi.get(self, "data_protection_replication")

    @property
    @pulumi.getter(name="exportPolicyRules")
    def export_policy_rules(self) -> pulumi.Output[Optional[Sequence['outputs.VolumeExportPolicyRule']]]:
        """
        One or more `export_policy_rule` block defined below.
        """
        return pulumi.get(self, "export_policy_rules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="mountIpAddresses")
    def mount_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of IPv4 Addresses which should be used to mount the volume.
        """
        return pulumi.get(self, "mount_ip_addresses")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the NetApp Volume. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="poolName")
    def pool_name(self) -> pulumi.Output[str]:
        """
        The name of the NetApp pool in which the NetApp Volume should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "pool_name")

    @property
    @pulumi.getter
    def protocols(self) -> pulumi.Output[Sequence[str]]:
        """
        The target volume protocol expressed as a list. Supported single value include `CIFS`, `NFSv3`, or `NFSv4.1`. If argument is not defined it will default to `NFSv3`. Changing this forces a new resource to be created and data will be lost.
        """
        return pulumi.get(self, "protocols")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group where the NetApp Volume should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="serviceLevel")
    def service_level(self) -> pulumi.Output[str]:
        """
        The target performance of the file system. Valid values include `Premium`, `Standard`, or `Ultra`.
        """
        return pulumi.get(self, "service_level")

    @property
    @pulumi.getter(name="storageQuotaInGb")
    def storage_quota_in_gb(self) -> pulumi.Output[int]:
        """
        The maximum Storage Quota allowed for a file system in Gigabytes.
        """
        return pulumi.get(self, "storage_quota_in_gb")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the Subnet the NetApp Volume resides in, which must have the `Microsoft.NetApp/volumes` delegation. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumePath")
    def volume_path(self) -> pulumi.Output[str]:
        """
        A unique file path for the volume. Used when creating mount targets. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "volume_path")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

