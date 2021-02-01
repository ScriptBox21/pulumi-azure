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

__all__ = ['Account']


class Account(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 analytical_storage_enabled: Optional[pulumi.Input[bool]] = None,
                 capabilities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountCapabilityArgs']]]]] = None,
                 consistency_policy: Optional[pulumi.Input[pulumi.InputType['AccountConsistencyPolicyArgs']]] = None,
                 enable_automatic_failover: Optional[pulumi.Input[bool]] = None,
                 enable_free_tier: Optional[pulumi.Input[bool]] = None,
                 enable_multiple_write_locations: Optional[pulumi.Input[bool]] = None,
                 geo_locations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountGeoLocationArgs']]]]] = None,
                 ip_range_filter: Optional[pulumi.Input[str]] = None,
                 is_virtual_network_filter_enabled: Optional[pulumi.Input[bool]] = None,
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 offer_type: Optional[pulumi.Input[str]] = None,
                 public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_network_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountVirtualNetworkRuleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a CosmosDB (formally DocumentDB) Account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure
        import pulumi_random as random

        rg = azure.core.ResourceGroup("rg", location="westus")
        ri = random.RandomInteger("ri",
            min=10000,
            max=99999)
        db = azure.cosmosdb.Account("db",
            location=rg.location,
            resource_group_name=rg.name,
            offer_type="Standard",
            kind="GlobalDocumentDB",
            enable_automatic_failover=True,
            capabilities=[
                azure.cosmosdb.AccountCapabilityArgs(
                    name="EnableAggregationPipeline",
                ),
                azure.cosmosdb.AccountCapabilityArgs(
                    name="mongoEnableDocLevelTTL",
                ),
                azure.cosmosdb.AccountCapabilityArgs(
                    name="MongoDBv3.4",
                ),
            ],
            consistency_policy=azure.cosmosdb.AccountConsistencyPolicyArgs(
                consistency_level="BoundedStaleness",
                max_interval_in_seconds=10,
                max_staleness_prefix=200,
            ),
            geo_locations=[
                azure.cosmosdb.AccountGeoLocationArgs(
                    location=var["failover_location"],
                    failover_priority=1,
                ),
                azure.cosmosdb.AccountGeoLocationArgs(
                    location=rg.location,
                    failover_priority=0,
                ),
            ])
        ```

        ## Import

        CosmosDB Accounts can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:cosmosdb/account:Account account1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] analytical_storage_enabled: Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountCapabilityArgs']]]] capabilities: The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, `EnableServerless`, and `mongoEnableDocLevelTTL`.
        :param pulumi.Input[pulumi.InputType['AccountConsistencyPolicyArgs']] consistency_policy: Specifies a `consistency_policy` resource, used to define the consistency policy for this CosmosDB account.
        :param pulumi.Input[bool] enable_automatic_failover: Enable automatic fail over for this Cosmos DB account.
        :param pulumi.Input[bool] enable_free_tier: Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enable_multiple_write_locations: Enable multi-master support for this Cosmos DB account.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountGeoLocationArgs']]]] geo_locations: Specifies a `geo_location` resource, used to define where data should be replicated with the `failover_priority` 0 specifying the primary location.
        :param pulumi.Input[str] ip_range_filter: CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        :param pulumi.Input[bool] is_virtual_network_filter_enabled: Enables virtual network filtering for this Cosmos DB account.
        :param pulumi.Input[str] key_vault_key_id: A Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
        :param pulumi.Input[str] kind: Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] offer_type: Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this CosmosDB account.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountVirtualNetworkRuleArgs']]]] virtual_network_rules: Specifies a `virtual_network_rules` resource, used to define which subnets are allowed to access this CosmosDB account.
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

            __props__['analytical_storage_enabled'] = analytical_storage_enabled
            __props__['capabilities'] = capabilities
            if consistency_policy is None and not opts.urn:
                raise TypeError("Missing required property 'consistency_policy'")
            __props__['consistency_policy'] = consistency_policy
            __props__['enable_automatic_failover'] = enable_automatic_failover
            __props__['enable_free_tier'] = enable_free_tier
            __props__['enable_multiple_write_locations'] = enable_multiple_write_locations
            if geo_locations is None and not opts.urn:
                raise TypeError("Missing required property 'geo_locations'")
            __props__['geo_locations'] = geo_locations
            __props__['ip_range_filter'] = ip_range_filter
            __props__['is_virtual_network_filter_enabled'] = is_virtual_network_filter_enabled
            __props__['key_vault_key_id'] = key_vault_key_id
            __props__['kind'] = kind
            __props__['location'] = location
            __props__['name'] = name
            if offer_type is None and not opts.urn:
                raise TypeError("Missing required property 'offer_type'")
            __props__['offer_type'] = offer_type
            __props__['public_network_access_enabled'] = public_network_access_enabled
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_network_rules'] = virtual_network_rules
            __props__['connection_strings'] = None
            __props__['endpoint'] = None
            __props__['primary_key'] = None
            __props__['primary_master_key'] = None
            __props__['primary_readonly_key'] = None
            __props__['primary_readonly_master_key'] = None
            __props__['read_endpoints'] = None
            __props__['secondary_key'] = None
            __props__['secondary_master_key'] = None
            __props__['secondary_readonly_key'] = None
            __props__['secondary_readonly_master_key'] = None
            __props__['write_endpoints'] = None
        super(Account, __self__).__init__(
            'azure:cosmosdb/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            analytical_storage_enabled: Optional[pulumi.Input[bool]] = None,
            capabilities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountCapabilityArgs']]]]] = None,
            connection_strings: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            consistency_policy: Optional[pulumi.Input[pulumi.InputType['AccountConsistencyPolicyArgs']]] = None,
            enable_automatic_failover: Optional[pulumi.Input[bool]] = None,
            enable_free_tier: Optional[pulumi.Input[bool]] = None,
            enable_multiple_write_locations: Optional[pulumi.Input[bool]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            geo_locations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountGeoLocationArgs']]]]] = None,
            ip_range_filter: Optional[pulumi.Input[str]] = None,
            is_virtual_network_filter_enabled: Optional[pulumi.Input[bool]] = None,
            key_vault_key_id: Optional[pulumi.Input[str]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            offer_type: Optional[pulumi.Input[str]] = None,
            primary_key: Optional[pulumi.Input[str]] = None,
            primary_master_key: Optional[pulumi.Input[str]] = None,
            primary_readonly_key: Optional[pulumi.Input[str]] = None,
            primary_readonly_master_key: Optional[pulumi.Input[str]] = None,
            public_network_access_enabled: Optional[pulumi.Input[bool]] = None,
            read_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            secondary_key: Optional[pulumi.Input[str]] = None,
            secondary_master_key: Optional[pulumi.Input[str]] = None,
            secondary_readonly_key: Optional[pulumi.Input[str]] = None,
            secondary_readonly_master_key: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            virtual_network_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountVirtualNetworkRuleArgs']]]]] = None,
            write_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] analytical_storage_enabled: Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountCapabilityArgs']]]] capabilities: The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, `EnableServerless`, and `mongoEnableDocLevelTTL`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] connection_strings: A list of connection strings available for this CosmosDB account.
        :param pulumi.Input[pulumi.InputType['AccountConsistencyPolicyArgs']] consistency_policy: Specifies a `consistency_policy` resource, used to define the consistency policy for this CosmosDB account.
        :param pulumi.Input[bool] enable_automatic_failover: Enable automatic fail over for this Cosmos DB account.
        :param pulumi.Input[bool] enable_free_tier: Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] enable_multiple_write_locations: Enable multi-master support for this Cosmos DB account.
        :param pulumi.Input[str] endpoint: The endpoint used to connect to the CosmosDB account.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountGeoLocationArgs']]]] geo_locations: Specifies a `geo_location` resource, used to define where data should be replicated with the `failover_priority` 0 specifying the primary location.
        :param pulumi.Input[str] ip_range_filter: CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        :param pulumi.Input[bool] is_virtual_network_filter_enabled: Enables virtual network filtering for this Cosmos DB account.
        :param pulumi.Input[str] key_vault_key_id: A Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
        :param pulumi.Input[str] kind: Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
        :param pulumi.Input[str] offer_type: Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
        :param pulumi.Input[str] primary_key: The Primary master key for the CosmosDB Account.
        :param pulumi.Input[str] primary_readonly_key: The Primary read-only master Key for the CosmosDB Account.
        :param pulumi.Input[bool] public_network_access_enabled: Whether or not public network access is allowed for this CosmosDB account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] read_endpoints: A list of read endpoints available for this CosmosDB account.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] secondary_key: The Secondary master key for the CosmosDB Account.
        :param pulumi.Input[str] secondary_readonly_key: The Secondary read-only master key for the CosmosDB Account.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountVirtualNetworkRuleArgs']]]] virtual_network_rules: Specifies a `virtual_network_rules` resource, used to define which subnets are allowed to access this CosmosDB account.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] write_endpoints: A list of write endpoints available for this CosmosDB account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["analytical_storage_enabled"] = analytical_storage_enabled
        __props__["capabilities"] = capabilities
        __props__["connection_strings"] = connection_strings
        __props__["consistency_policy"] = consistency_policy
        __props__["enable_automatic_failover"] = enable_automatic_failover
        __props__["enable_free_tier"] = enable_free_tier
        __props__["enable_multiple_write_locations"] = enable_multiple_write_locations
        __props__["endpoint"] = endpoint
        __props__["geo_locations"] = geo_locations
        __props__["ip_range_filter"] = ip_range_filter
        __props__["is_virtual_network_filter_enabled"] = is_virtual_network_filter_enabled
        __props__["key_vault_key_id"] = key_vault_key_id
        __props__["kind"] = kind
        __props__["location"] = location
        __props__["name"] = name
        __props__["offer_type"] = offer_type
        __props__["primary_key"] = primary_key
        __props__["primary_master_key"] = primary_master_key
        __props__["primary_readonly_key"] = primary_readonly_key
        __props__["primary_readonly_master_key"] = primary_readonly_master_key
        __props__["public_network_access_enabled"] = public_network_access_enabled
        __props__["read_endpoints"] = read_endpoints
        __props__["resource_group_name"] = resource_group_name
        __props__["secondary_key"] = secondary_key
        __props__["secondary_master_key"] = secondary_master_key
        __props__["secondary_readonly_key"] = secondary_readonly_key
        __props__["secondary_readonly_master_key"] = secondary_readonly_master_key
        __props__["tags"] = tags
        __props__["virtual_network_rules"] = virtual_network_rules
        __props__["write_endpoints"] = write_endpoints
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="analyticalStorageEnabled")
    def analytical_storage_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable Analytical Storage option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "analytical_storage_enabled")

    @property
    @pulumi.getter
    def capabilities(self) -> pulumi.Output[Optional[Sequence['outputs.AccountCapability']]]:
        """
        The capabilities which should be enabled for this Cosmos DB account. Possible values are `EnableAggregationPipeline`, `EnableCassandra`, `EnableGremlin`, `EnableTable`, `MongoDBv3.4`, `EnableServerless`, and `mongoEnableDocLevelTTL`.
        """
        return pulumi.get(self, "capabilities")

    @property
    @pulumi.getter(name="connectionStrings")
    def connection_strings(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of connection strings available for this CosmosDB account.
        """
        return pulumi.get(self, "connection_strings")

    @property
    @pulumi.getter(name="consistencyPolicy")
    def consistency_policy(self) -> pulumi.Output['outputs.AccountConsistencyPolicy']:
        """
        Specifies a `consistency_policy` resource, used to define the consistency policy for this CosmosDB account.
        """
        return pulumi.get(self, "consistency_policy")

    @property
    @pulumi.getter(name="enableAutomaticFailover")
    def enable_automatic_failover(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable automatic fail over for this Cosmos DB account.
        """
        return pulumi.get(self, "enable_automatic_failover")

    @property
    @pulumi.getter(name="enableFreeTier")
    def enable_free_tier(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable Free Tier pricing option for this Cosmos DB account. Defaults to `false`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "enable_free_tier")

    @property
    @pulumi.getter(name="enableMultipleWriteLocations")
    def enable_multiple_write_locations(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable multi-master support for this Cosmos DB account.
        """
        return pulumi.get(self, "enable_multiple_write_locations")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint used to connect to the CosmosDB account.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="geoLocations")
    def geo_locations(self) -> pulumi.Output[Sequence['outputs.AccountGeoLocation']]:
        """
        Specifies a `geo_location` resource, used to define where data should be replicated with the `failover_priority` 0 specifying the primary location.
        """
        return pulumi.get(self, "geo_locations")

    @property
    @pulumi.getter(name="ipRangeFilter")
    def ip_range_filter(self) -> pulumi.Output[Optional[str]]:
        """
        CosmosDB Firewall Support: This value specifies the set of IP addresses or IP address ranges in CIDR form to be included as the allowed list of client IP's for a given database account. IP addresses/ranges must be comma separated and must not contain any spaces.
        """
        return pulumi.get(self, "ip_range_filter")

    @property
    @pulumi.getter(name="isVirtualNetworkFilterEnabled")
    def is_virtual_network_filter_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables virtual network filtering for this Cosmos DB account.
        """
        return pulumi.get(self, "is_virtual_network_filter_enabled")

    @property
    @pulumi.getter(name="keyVaultKeyId")
    def key_vault_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        A Key Vault Key ID for CMK encryption. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "key_vault_key_id")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Kind of CosmosDB to create - possible values are `GlobalDocumentDB` and `MongoDB`. Defaults to `GlobalDocumentDB`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the CosmosDB Account. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="offerType")
    def offer_type(self) -> pulumi.Output[str]:
        """
        Specifies the Offer Type to use for this CosmosDB Account - currently this can only be set to `Standard`.
        """
        return pulumi.get(self, "offer_type")

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> pulumi.Output[str]:
        """
        The Primary master key for the CosmosDB Account.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="primaryMasterKey")
    def primary_master_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "primary_master_key")

    @property
    @pulumi.getter(name="primaryReadonlyKey")
    def primary_readonly_key(self) -> pulumi.Output[str]:
        """
        The Primary read-only master Key for the CosmosDB Account.
        """
        return pulumi.get(self, "primary_readonly_key")

    @property
    @pulumi.getter(name="primaryReadonlyMasterKey")
    def primary_readonly_master_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "primary_readonly_master_key")

    @property
    @pulumi.getter(name="publicNetworkAccessEnabled")
    def public_network_access_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not public network access is allowed for this CosmosDB account.
        """
        return pulumi.get(self, "public_network_access_enabled")

    @property
    @pulumi.getter(name="readEndpoints")
    def read_endpoints(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of read endpoints available for this CosmosDB account.
        """
        return pulumi.get(self, "read_endpoints")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the resource group in which the CosmosDB Account is created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> pulumi.Output[str]:
        """
        The Secondary master key for the CosmosDB Account.
        """
        return pulumi.get(self, "secondary_key")

    @property
    @pulumi.getter(name="secondaryMasterKey")
    def secondary_master_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secondary_master_key")

    @property
    @pulumi.getter(name="secondaryReadonlyKey")
    def secondary_readonly_key(self) -> pulumi.Output[str]:
        """
        The Secondary read-only master key for the CosmosDB Account.
        """
        return pulumi.get(self, "secondary_readonly_key")

    @property
    @pulumi.getter(name="secondaryReadonlyMasterKey")
    def secondary_readonly_master_key(self) -> pulumi.Output[str]:
        return pulumi.get(self, "secondary_readonly_master_key")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualNetworkRules")
    def virtual_network_rules(self) -> pulumi.Output[Optional[Sequence['outputs.AccountVirtualNetworkRule']]]:
        """
        Specifies a `virtual_network_rules` resource, used to define which subnets are allowed to access this CosmosDB account.
        """
        return pulumi.get(self, "virtual_network_rules")

    @property
    @pulumi.getter(name="writeEndpoints")
    def write_endpoints(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of write endpoints available for this CosmosDB account.
        """
        return pulumi.get(self, "write_endpoints")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

