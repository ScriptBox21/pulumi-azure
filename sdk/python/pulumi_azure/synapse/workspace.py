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

__all__ = ['Workspace']


class Workspace(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aad_admin: Optional[pulumi.Input[pulumi.InputType['WorkspaceAadAdminArgs']]] = None,
                 azure_devops_repo: Optional[pulumi.Input[pulumi.InputType['WorkspaceAzureDevopsRepoArgs']]] = None,
                 github_repo: Optional[pulumi.Input[pulumi.InputType['WorkspaceGithubRepoArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_resource_group_name: Optional[pulumi.Input[str]] = None,
                 managed_virtual_network_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sql_administrator_login: Optional[pulumi.Input[str]] = None,
                 sql_administrator_login_password: Optional[pulumi.Input[str]] = None,
                 sql_identity_control_enabled: Optional[pulumi.Input[bool]] = None,
                 storage_data_lake_gen2_filesystem_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Synapse Workspace.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS",
            account_kind="StorageV2",
            is_hns_enabled=True)
        example_data_lake_gen2_filesystem = azure.storage.DataLakeGen2Filesystem("exampleDataLakeGen2Filesystem", storage_account_id=example_account.id)
        example_workspace = azure.synapse.Workspace("exampleWorkspace",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            storage_data_lake_gen2_filesystem_id=example_data_lake_gen2_filesystem.id,
            sql_administrator_login="sqladminuser",
            sql_administrator_login_password="H@Sh1CoR3!",
            aad_admin=azure.synapse.WorkspaceAadAdminArgs(
                login="AzureAD Admin",
                object_id="00000000-0000-0000-0000-000000000000",
                tenant_id="00000000-0000-0000-0000-000000000000",
            ),
            tags={
                "Env": "production",
            })
        ```

        ## Import

        Synapse Workspace can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:synapse/workspace:Workspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['WorkspaceAadAdminArgs']] aad_admin: An `aad_admin` block as defined below.
        :param pulumi.Input[pulumi.InputType['WorkspaceAzureDevopsRepoArgs']] azure_devops_repo: An `azure_devops_repo` block as defined below.
        :param pulumi.Input[pulumi.InputType['WorkspaceGithubRepoArgs']] github_repo: A `github_repo` block as defined below.
        :param pulumi.Input[str] location: Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_resource_group_name: Workspace managed resource group.
        :param pulumi.Input[bool] managed_virtual_network_enabled: Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_administrator_login: Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_administrator_login_password: The Password associated with the `sql_administrator_login` for the SQL administrator.
        :param pulumi.Input[bool] sql_identity_control_enabled: Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
        :param pulumi.Input[str] storage_data_lake_gen2_filesystem_id: Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Synapse Workspace.
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

            __props__['aad_admin'] = aad_admin
            __props__['azure_devops_repo'] = azure_devops_repo
            __props__['github_repo'] = github_repo
            __props__['location'] = location
            __props__['managed_resource_group_name'] = managed_resource_group_name
            __props__['managed_virtual_network_enabled'] = managed_virtual_network_enabled
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sql_administrator_login is None and not opts.urn:
                raise TypeError("Missing required property 'sql_administrator_login'")
            __props__['sql_administrator_login'] = sql_administrator_login
            if sql_administrator_login_password is None and not opts.urn:
                raise TypeError("Missing required property 'sql_administrator_login_password'")
            __props__['sql_administrator_login_password'] = sql_administrator_login_password
            __props__['sql_identity_control_enabled'] = sql_identity_control_enabled
            if storage_data_lake_gen2_filesystem_id is None and not opts.urn:
                raise TypeError("Missing required property 'storage_data_lake_gen2_filesystem_id'")
            __props__['storage_data_lake_gen2_filesystem_id'] = storage_data_lake_gen2_filesystem_id
            __props__['tags'] = tags
            __props__['connectivity_endpoints'] = None
            __props__['identities'] = None
        super(Workspace, __self__).__init__(
            'azure:synapse/workspace:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            aad_admin: Optional[pulumi.Input[pulumi.InputType['WorkspaceAadAdminArgs']]] = None,
            azure_devops_repo: Optional[pulumi.Input[pulumi.InputType['WorkspaceAzureDevopsRepoArgs']]] = None,
            connectivity_endpoints: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            github_repo: Optional[pulumi.Input[pulumi.InputType['WorkspaceGithubRepoArgs']]] = None,
            identities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkspaceIdentityArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            managed_resource_group_name: Optional[pulumi.Input[str]] = None,
            managed_virtual_network_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            sql_administrator_login: Optional[pulumi.Input[str]] = None,
            sql_administrator_login_password: Optional[pulumi.Input[str]] = None,
            sql_identity_control_enabled: Optional[pulumi.Input[bool]] = None,
            storage_data_lake_gen2_filesystem_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Workspace':
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['WorkspaceAadAdminArgs']] aad_admin: An `aad_admin` block as defined below.
        :param pulumi.Input[pulumi.InputType['WorkspaceAzureDevopsRepoArgs']] azure_devops_repo: An `azure_devops_repo` block as defined below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] connectivity_endpoints: A list of Connectivity endpoints for this Synapse Workspace.
        :param pulumi.Input[pulumi.InputType['WorkspaceGithubRepoArgs']] github_repo: A `github_repo` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkspaceIdentityArgs']]]] identities: An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
        :param pulumi.Input[str] location: Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] managed_resource_group_name: Workspace managed resource group.
        :param pulumi.Input[bool] managed_virtual_network_enabled: Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_administrator_login: Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
        :param pulumi.Input[str] sql_administrator_login_password: The Password associated with the `sql_administrator_login` for the SQL administrator.
        :param pulumi.Input[bool] sql_identity_control_enabled: Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
        :param pulumi.Input[str] storage_data_lake_gen2_filesystem_id: Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags which should be assigned to the Synapse Workspace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["aad_admin"] = aad_admin
        __props__["azure_devops_repo"] = azure_devops_repo
        __props__["connectivity_endpoints"] = connectivity_endpoints
        __props__["github_repo"] = github_repo
        __props__["identities"] = identities
        __props__["location"] = location
        __props__["managed_resource_group_name"] = managed_resource_group_name
        __props__["managed_virtual_network_enabled"] = managed_virtual_network_enabled
        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["sql_administrator_login"] = sql_administrator_login
        __props__["sql_administrator_login_password"] = sql_administrator_login_password
        __props__["sql_identity_control_enabled"] = sql_identity_control_enabled
        __props__["storage_data_lake_gen2_filesystem_id"] = storage_data_lake_gen2_filesystem_id
        __props__["tags"] = tags
        return Workspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aadAdmin")
    def aad_admin(self) -> pulumi.Output['outputs.WorkspaceAadAdmin']:
        """
        An `aad_admin` block as defined below.
        """
        return pulumi.get(self, "aad_admin")

    @property
    @pulumi.getter(name="azureDevopsRepo")
    def azure_devops_repo(self) -> pulumi.Output[Optional['outputs.WorkspaceAzureDevopsRepo']]:
        """
        An `azure_devops_repo` block as defined below.
        """
        return pulumi.get(self, "azure_devops_repo")

    @property
    @pulumi.getter(name="connectivityEndpoints")
    def connectivity_endpoints(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A list of Connectivity endpoints for this Synapse Workspace.
        """
        return pulumi.get(self, "connectivity_endpoints")

    @property
    @pulumi.getter(name="githubRepo")
    def github_repo(self) -> pulumi.Output[Optional['outputs.WorkspaceGithubRepo']]:
        """
        A `github_repo` block as defined below.
        """
        return pulumi.get(self, "github_repo")

    @property
    @pulumi.getter
    def identities(self) -> pulumi.Output[Sequence['outputs.WorkspaceIdentity']]:
        """
        An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
        """
        return pulumi.get(self, "identities")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedResourceGroupName")
    def managed_resource_group_name(self) -> pulumi.Output[str]:
        """
        Workspace managed resource group.
        """
        return pulumi.get(self, "managed_resource_group_name")

    @property
    @pulumi.getter(name="managedVirtualNetworkEnabled")
    def managed_virtual_network_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "managed_virtual_network_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="sqlAdministratorLogin")
    def sql_administrator_login(self) -> pulumi.Output[str]:
        """
        Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "sql_administrator_login")

    @property
    @pulumi.getter(name="sqlAdministratorLoginPassword")
    def sql_administrator_login_password(self) -> pulumi.Output[str]:
        """
        The Password associated with the `sql_administrator_login` for the SQL administrator.
        """
        return pulumi.get(self, "sql_administrator_login_password")

    @property
    @pulumi.getter(name="sqlIdentityControlEnabled")
    def sql_identity_control_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
        """
        return pulumi.get(self, "sql_identity_control_enabled")

    @property
    @pulumi.getter(name="storageDataLakeGen2FilesystemId")
    def storage_data_lake_gen2_filesystem_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "storage_data_lake_gen2_filesystem_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags which should be assigned to the Synapse Workspace.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

