# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ServerKey']


class ServerKey(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_vault_key_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Customer Managed Key for a PostgreSQL Server.

        ## Import

        A PostgreSQL Server Key can be imported using the `resource id` of the PostgreSQL Server Key, e.g.

        ```sh
         $ pulumi import azure:postgresql/serverKey:ServerKey example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_vault_key_id: The URL to a Key Vault Key.
        :param pulumi.Input[str] server_id: The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
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

            if key_vault_key_id is None:
                raise TypeError("Missing required property 'key_vault_key_id'")
            __props__['key_vault_key_id'] = key_vault_key_id
            if server_id is None:
                raise TypeError("Missing required property 'server_id'")
            __props__['server_id'] = server_id
        super(ServerKey, __self__).__init__(
            'azure:postgresql/serverKey:ServerKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_vault_key_id: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None) -> 'ServerKey':
        """
        Get an existing ServerKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_vault_key_id: The URL to a Key Vault Key.
        :param pulumi.Input[str] server_id: The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["key_vault_key_id"] = key_vault_key_id
        __props__["server_id"] = server_id
        return ServerKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyVaultKeyId")
    def key_vault_key_id(self) -> pulumi.Output[str]:
        """
        The URL to a Key Vault Key.
        """
        return pulumi.get(self, "key_vault_key_id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "server_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

