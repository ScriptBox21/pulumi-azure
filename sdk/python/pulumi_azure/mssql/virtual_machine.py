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

__all__ = ['VirtualMachine']


class VirtualMachine(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_patching: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']]] = None,
                 key_vault_credential: Optional[pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']]] = None,
                 r_services_enabled: Optional[pulumi.Input[bool]] = None,
                 sql_connectivity_port: Optional[pulumi.Input[float]] = None,
                 sql_connectivity_type: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_password: Optional[pulumi.Input[str]] = None,
                 sql_connectivity_update_username: Optional[pulumi.Input[str]] = None,
                 sql_license_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_machine_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Microsoft SQL Virtual Machine

        ## Example Usage

        This example provisions a brief Managed MsSql Virtual Machine.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_virtual_machine = azure.compute.get_virtual_machine(name="example-vm",
            resource_group_name="example-resources")
        example_mssql_virtual_machine_virtual_machine = azure.mssql.VirtualMachine("exampleMssql/virtualMachineVirtualMachine",
            virtual_machine_id=example_virtual_machine.id,
            sql_license_type="PAYG",
            r_services_enabled=True,
            sql_connectivity_port=1433,
            sql_connectivity_type="PRIVATE",
            sql_connectivity_update_password="Password1234!",
            sql_connectivity_update_username="sqllogin",
            auto_patching=azure.mssql.VirtualMachineAutoPatchingArgs(
                day_of_week="Sunday",
                maintenance_window_duration_in_minutes=60,
                maintenance_window_starting_hour=2,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']] auto_patching: An `auto_patching` block as defined below.
        :param pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']] key_vault_credential: (Optional) An `key_vault_credential` block as defined below.
        :param pulumi.Input[bool] r_services_enabled: Should R Services be enabled?
        :param pulumi.Input[float] sql_connectivity_port: The SQL Server port. Defaults to `1433`.
        :param pulumi.Input[str] sql_connectivity_type: The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        :param pulumi.Input[str] sql_connectivity_update_password: The SQL Server sysadmin login password.
        :param pulumi.Input[str] sql_connectivity_update_username: The SQL Server sysadmin login to create.
        :param pulumi.Input[str] sql_license_type: The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created.
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

            __props__['auto_patching'] = auto_patching
            __props__['key_vault_credential'] = key_vault_credential
            __props__['r_services_enabled'] = r_services_enabled
            __props__['sql_connectivity_port'] = sql_connectivity_port
            __props__['sql_connectivity_type'] = sql_connectivity_type
            __props__['sql_connectivity_update_password'] = sql_connectivity_update_password
            __props__['sql_connectivity_update_username'] = sql_connectivity_update_username
            if sql_license_type is None:
                raise TypeError("Missing required property 'sql_license_type'")
            __props__['sql_license_type'] = sql_license_type
            __props__['tags'] = tags
            if virtual_machine_id is None:
                raise TypeError("Missing required property 'virtual_machine_id'")
            __props__['virtual_machine_id'] = virtual_machine_id
        super(VirtualMachine, __self__).__init__(
            'azure:mssql/virtualMachine:VirtualMachine',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_patching: Optional[pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']]] = None,
            key_vault_credential: Optional[pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']]] = None,
            r_services_enabled: Optional[pulumi.Input[bool]] = None,
            sql_connectivity_port: Optional[pulumi.Input[float]] = None,
            sql_connectivity_type: Optional[pulumi.Input[str]] = None,
            sql_connectivity_update_password: Optional[pulumi.Input[str]] = None,
            sql_connectivity_update_username: Optional[pulumi.Input[str]] = None,
            sql_license_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            virtual_machine_id: Optional[pulumi.Input[str]] = None) -> 'VirtualMachine':
        """
        Get an existing VirtualMachine resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['VirtualMachineAutoPatchingArgs']] auto_patching: An `auto_patching` block as defined below.
        :param pulumi.Input[pulumi.InputType['VirtualMachineKeyVaultCredentialArgs']] key_vault_credential: (Optional) An `key_vault_credential` block as defined below.
        :param pulumi.Input[bool] r_services_enabled: Should R Services be enabled?
        :param pulumi.Input[float] sql_connectivity_port: The SQL Server port. Defaults to `1433`.
        :param pulumi.Input[str] sql_connectivity_type: The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        :param pulumi.Input[str] sql_connectivity_update_password: The SQL Server sysadmin login password.
        :param pulumi.Input[str] sql_connectivity_update_username: The SQL Server sysadmin login to create.
        :param pulumi.Input[str] sql_license_type: The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] virtual_machine_id: The ID of the Virtual Machine. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_patching"] = auto_patching
        __props__["key_vault_credential"] = key_vault_credential
        __props__["r_services_enabled"] = r_services_enabled
        __props__["sql_connectivity_port"] = sql_connectivity_port
        __props__["sql_connectivity_type"] = sql_connectivity_type
        __props__["sql_connectivity_update_password"] = sql_connectivity_update_password
        __props__["sql_connectivity_update_username"] = sql_connectivity_update_username
        __props__["sql_license_type"] = sql_license_type
        __props__["tags"] = tags
        __props__["virtual_machine_id"] = virtual_machine_id
        return VirtualMachine(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoPatching")
    def auto_patching(self) -> Optional['outputs.VirtualMachineAutoPatching']:
        """
        An `auto_patching` block as defined below.
        """
        return pulumi.get(self, "auto_patching")

    @property
    @pulumi.getter(name="keyVaultCredential")
    def key_vault_credential(self) -> Optional['outputs.VirtualMachineKeyVaultCredential']:
        """
        (Optional) An `key_vault_credential` block as defined below.
        """
        return pulumi.get(self, "key_vault_credential")

    @property
    @pulumi.getter(name="rServicesEnabled")
    def r_services_enabled(self) -> Optional[bool]:
        """
        Should R Services be enabled?
        """
        return pulumi.get(self, "r_services_enabled")

    @property
    @pulumi.getter(name="sqlConnectivityPort")
    def sql_connectivity_port(self) -> Optional[float]:
        """
        The SQL Server port. Defaults to `1433`.
        """
        return pulumi.get(self, "sql_connectivity_port")

    @property
    @pulumi.getter(name="sqlConnectivityType")
    def sql_connectivity_type(self) -> Optional[str]:
        """
        The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
        """
        return pulumi.get(self, "sql_connectivity_type")

    @property
    @pulumi.getter(name="sqlConnectivityUpdatePassword")
    def sql_connectivity_update_password(self) -> Optional[str]:
        """
        The SQL Server sysadmin login password.
        """
        return pulumi.get(self, "sql_connectivity_update_password")

    @property
    @pulumi.getter(name="sqlConnectivityUpdateUsername")
    def sql_connectivity_update_username(self) -> Optional[str]:
        """
        The SQL Server sysadmin login to create.
        """
        return pulumi.get(self, "sql_connectivity_update_username")

    @property
    @pulumi.getter(name="sqlLicenseType")
    def sql_license_type(self) -> str:
        """
        The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sql_license_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="virtualMachineId")
    def virtual_machine_id(self) -> str:
        """
        The ID of the Virtual Machine. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "virtual_machine_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

