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

__all__ = ['KubernetesCluster']


class KubernetesCluster(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 addon_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterAddonProfileArgs']]] = None,
                 api_server_authorized_ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 auto_scaler_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterAutoScalerProfileArgs']]] = None,
                 default_node_pool: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterDefaultNodePoolArgs']]] = None,
                 disk_encryption_set_id: Optional[pulumi.Input[str]] = None,
                 dns_prefix: Optional[pulumi.Input[str]] = None,
                 enable_pod_security_policy: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterIdentityArgs']]] = None,
                 kubernetes_version: Optional[pulumi.Input[str]] = None,
                 linux_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterLinuxProfileArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterNetworkProfileArgs']]] = None,
                 node_resource_group: Optional[pulumi.Input[str]] = None,
                 private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
                 private_link_enabled: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 role_based_access_control: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterRoleBasedAccessControlArgs']]] = None,
                 service_principal: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterServicePrincipalArgs']]] = None,
                 sku_tier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 windows_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterWindowsProfileArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Managed Kubernetes Cluster (also known as AKS / Azure Kubernetes Service)

        ## Example Usage

        This example provisions a basic Managed Kubernetes Cluster.

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_kubernetes_cluster = azure.containerservice.KubernetesCluster("exampleKubernetesCluster",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            dns_prefix="exampleaks1",
            default_node_pool=azure.containerservice.KubernetesClusterDefaultNodePoolArgs(
                name="default",
                node_count=1,
                vm_size="Standard_D2_v2",
            ),
            identity=azure.containerservice.KubernetesClusterIdentityArgs(
                type="SystemAssigned",
            ),
            tags={
                "Environment": "Production",
            })
        pulumi.export("clientCertificate", example_kubernetes_cluster.kube_configs[0].client_certificate)
        pulumi.export("kubeConfig", example_kubernetes_cluster.kube_config_raw)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterAddonProfileArgs']] addon_profile: A `addon_profile` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_server_authorized_ip_ranges: The IP ranges to whitelist for incoming traffic to the masters.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterAutoScalerProfileArgs']] auto_scaler_profile: A `auto_scaler_profile` block as defined below.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterDefaultNodePoolArgs']] default_node_pool: A `default_node_pool` block as defined below.
        :param pulumi.Input[str] disk_encryption_set_id: The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
        :param pulumi.Input[str] dns_prefix: DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterIdentityArgs']] identity: A `identity` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
        :param pulumi.Input[pulumi.InputType['KubernetesClusterLinuxProfileArgs']] linux_profile: A `linux_profile` block as defined below.
        :param pulumi.Input[str] location: The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterNetworkProfileArgs']] network_profile: A `network_profile` block as defined below.
        :param pulumi.Input[str] node_resource_group: The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] private_cluster_enabled: Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterRoleBasedAccessControlArgs']] role_based_access_control: A `role_based_access_control` block. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterServicePrincipalArgs']] service_principal: A `service_principal` block as documented below.
        :param pulumi.Input[str] sku_tier: The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterWindowsProfileArgs']] windows_profile: A `windows_profile` block as defined below.
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

            __props__['addon_profile'] = addon_profile
            __props__['api_server_authorized_ip_ranges'] = api_server_authorized_ip_ranges
            __props__['auto_scaler_profile'] = auto_scaler_profile
            if default_node_pool is None:
                raise TypeError("Missing required property 'default_node_pool'")
            __props__['default_node_pool'] = default_node_pool
            __props__['disk_encryption_set_id'] = disk_encryption_set_id
            if dns_prefix is None:
                raise TypeError("Missing required property 'dns_prefix'")
            __props__['dns_prefix'] = dns_prefix
            __props__['enable_pod_security_policy'] = enable_pod_security_policy
            __props__['identity'] = identity
            __props__['kubernetes_version'] = kubernetes_version
            __props__['linux_profile'] = linux_profile
            __props__['location'] = location
            __props__['name'] = name
            __props__['network_profile'] = network_profile
            __props__['node_resource_group'] = node_resource_group
            __props__['private_cluster_enabled'] = private_cluster_enabled
            if private_link_enabled is not None:
                warnings.warn("Deprecated in favor of `private_cluster_enabled`", DeprecationWarning)
                pulumi.log.warn("private_link_enabled is deprecated: Deprecated in favor of `private_cluster_enabled`")
            __props__['private_link_enabled'] = private_link_enabled
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['role_based_access_control'] = role_based_access_control
            __props__['service_principal'] = service_principal
            __props__['sku_tier'] = sku_tier
            __props__['tags'] = tags
            __props__['windows_profile'] = windows_profile
            __props__['fqdn'] = None
            __props__['kube_admin_config_raw'] = None
            __props__['kube_admin_configs'] = None
            __props__['kube_config_raw'] = None
            __props__['kube_configs'] = None
            __props__['kubelet_identities'] = None
            __props__['private_fqdn'] = None
        super(KubernetesCluster, __self__).__init__(
            'azure:containerservice/kubernetesCluster:KubernetesCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            addon_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterAddonProfileArgs']]] = None,
            api_server_authorized_ip_ranges: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            auto_scaler_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterAutoScalerProfileArgs']]] = None,
            default_node_pool: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterDefaultNodePoolArgs']]] = None,
            disk_encryption_set_id: Optional[pulumi.Input[str]] = None,
            dns_prefix: Optional[pulumi.Input[str]] = None,
            enable_pod_security_policy: Optional[pulumi.Input[bool]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            identity: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterIdentityArgs']]] = None,
            kube_admin_config_raw: Optional[pulumi.Input[str]] = None,
            kube_admin_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesClusterKubeAdminConfigArgs']]]]] = None,
            kube_config_raw: Optional[pulumi.Input[str]] = None,
            kube_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesClusterKubeConfigArgs']]]]] = None,
            kubelet_identities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesClusterKubeletIdentityArgs']]]]] = None,
            kubernetes_version: Optional[pulumi.Input[str]] = None,
            linux_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterLinuxProfileArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterNetworkProfileArgs']]] = None,
            node_resource_group: Optional[pulumi.Input[str]] = None,
            private_cluster_enabled: Optional[pulumi.Input[bool]] = None,
            private_fqdn: Optional[pulumi.Input[str]] = None,
            private_link_enabled: Optional[pulumi.Input[bool]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None,
            role_based_access_control: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterRoleBasedAccessControlArgs']]] = None,
            service_principal: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterServicePrincipalArgs']]] = None,
            sku_tier: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            windows_profile: Optional[pulumi.Input[pulumi.InputType['KubernetesClusterWindowsProfileArgs']]] = None) -> 'KubernetesCluster':
        """
        Get an existing KubernetesCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterAddonProfileArgs']] addon_profile: A `addon_profile` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] api_server_authorized_ip_ranges: The IP ranges to whitelist for incoming traffic to the masters.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterAutoScalerProfileArgs']] auto_scaler_profile: A `auto_scaler_profile` block as defined below.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterDefaultNodePoolArgs']] default_node_pool: A `default_node_pool` block as defined below.
        :param pulumi.Input[str] disk_encryption_set_id: The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
        :param pulumi.Input[str] dns_prefix: DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[str] fqdn: The FQDN of the Azure Kubernetes Managed Cluster.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterIdentityArgs']] identity: A `identity` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[str] kube_admin_config_raw: Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesClusterKubeAdminConfigArgs']]]] kube_admin_configs: A `kube_admin_config` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        :param pulumi.Input[str] kube_config_raw: Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesClusterKubeConfigArgs']]]] kube_configs: A `kube_config` block as defined below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KubernetesClusterKubeletIdentityArgs']]]] kubelet_identities: A `kubelet_identity` block as defined below.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
        :param pulumi.Input[pulumi.InputType['KubernetesClusterLinuxProfileArgs']] linux_profile: A `linux_profile` block as defined below.
        :param pulumi.Input[str] location: The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterNetworkProfileArgs']] network_profile: A `network_profile` block as defined below.
        :param pulumi.Input[str] node_resource_group: The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] private_cluster_enabled: Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] private_fqdn: The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
        :param pulumi.Input[str] resource_group_name: Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterRoleBasedAccessControlArgs']] role_based_access_control: A `role_based_access_control` block. Changing this forces a new resource to be created.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterServicePrincipalArgs']] service_principal: A `service_principal` block as documented below.
        :param pulumi.Input[str] sku_tier: The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['KubernetesClusterWindowsProfileArgs']] windows_profile: A `windows_profile` block as defined below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["addon_profile"] = addon_profile
        __props__["api_server_authorized_ip_ranges"] = api_server_authorized_ip_ranges
        __props__["auto_scaler_profile"] = auto_scaler_profile
        __props__["default_node_pool"] = default_node_pool
        __props__["disk_encryption_set_id"] = disk_encryption_set_id
        __props__["dns_prefix"] = dns_prefix
        __props__["enable_pod_security_policy"] = enable_pod_security_policy
        __props__["fqdn"] = fqdn
        __props__["identity"] = identity
        __props__["kube_admin_config_raw"] = kube_admin_config_raw
        __props__["kube_admin_configs"] = kube_admin_configs
        __props__["kube_config_raw"] = kube_config_raw
        __props__["kube_configs"] = kube_configs
        __props__["kubelet_identities"] = kubelet_identities
        __props__["kubernetes_version"] = kubernetes_version
        __props__["linux_profile"] = linux_profile
        __props__["location"] = location
        __props__["name"] = name
        __props__["network_profile"] = network_profile
        __props__["node_resource_group"] = node_resource_group
        __props__["private_cluster_enabled"] = private_cluster_enabled
        __props__["private_fqdn"] = private_fqdn
        __props__["private_link_enabled"] = private_link_enabled
        __props__["resource_group_name"] = resource_group_name
        __props__["role_based_access_control"] = role_based_access_control
        __props__["service_principal"] = service_principal
        __props__["sku_tier"] = sku_tier
        __props__["tags"] = tags
        __props__["windows_profile"] = windows_profile
        return KubernetesCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addonProfile")
    def addon_profile(self) -> pulumi.Output['outputs.KubernetesClusterAddonProfile']:
        """
        A `addon_profile` block as defined below.
        """
        return pulumi.get(self, "addon_profile")

    @property
    @pulumi.getter(name="apiServerAuthorizedIpRanges")
    def api_server_authorized_ip_ranges(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IP ranges to whitelist for incoming traffic to the masters.
        """
        return pulumi.get(self, "api_server_authorized_ip_ranges")

    @property
    @pulumi.getter(name="autoScalerProfile")
    def auto_scaler_profile(self) -> pulumi.Output['outputs.KubernetesClusterAutoScalerProfile']:
        """
        A `auto_scaler_profile` block as defined below.
        """
        return pulumi.get(self, "auto_scaler_profile")

    @property
    @pulumi.getter(name="defaultNodePool")
    def default_node_pool(self) -> pulumi.Output['outputs.KubernetesClusterDefaultNodePool']:
        """
        A `default_node_pool` block as defined below.
        """
        return pulumi.get(self, "default_node_pool")

    @property
    @pulumi.getter(name="diskEncryptionSetId")
    def disk_encryption_set_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the Disk Encryption Set which should be used for the Nodes and Volumes. More information [can be found in the documentation](https://docs.microsoft.com/en-us/azure/aks/azure-disk-customer-managed-keys).
        """
        return pulumi.get(self, "disk_encryption_set_id")

    @property
    @pulumi.getter(name="dnsPrefix")
    def dns_prefix(self) -> pulumi.Output[str]:
        """
        DNS prefix specified when creating the managed cluster. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "dns_prefix")

    @property
    @pulumi.getter(name="enablePodSecurityPolicy")
    def enable_pod_security_policy(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_pod_security_policy")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        The FQDN of the Azure Kubernetes Managed Cluster.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.KubernetesClusterIdentity']]:
        """
        A `identity` block as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="kubeAdminConfigRaw")
    def kube_admin_config_raw(self) -> pulumi.Output[str]:
        """
        Raw Kubernetes config for the admin account to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        """
        return pulumi.get(self, "kube_admin_config_raw")

    @property
    @pulumi.getter(name="kubeAdminConfigs")
    def kube_admin_configs(self) -> pulumi.Output[Sequence['outputs.KubernetesClusterKubeAdminConfig']]:
        """
        A `kube_admin_config` block as defined below. This is only available when Role Based Access Control with Azure Active Directory is enabled.
        """
        return pulumi.get(self, "kube_admin_configs")

    @property
    @pulumi.getter(name="kubeConfigRaw")
    def kube_config_raw(self) -> pulumi.Output[str]:
        """
        Raw Kubernetes config to be used by [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) and other compatible tools
        """
        return pulumi.get(self, "kube_config_raw")

    @property
    @pulumi.getter(name="kubeConfigs")
    def kube_configs(self) -> pulumi.Output[Sequence['outputs.KubernetesClusterKubeConfig']]:
        """
        A `kube_config` block as defined below.
        """
        return pulumi.get(self, "kube_configs")

    @property
    @pulumi.getter(name="kubeletIdentities")
    def kubelet_identities(self) -> pulumi.Output[Sequence['outputs.KubernetesClusterKubeletIdentity']]:
        """
        A `kubelet_identity` block as defined below.
        """
        return pulumi.get(self, "kubelet_identities")

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> pulumi.Output[str]:
        """
        Version of Kubernetes specified when creating the AKS managed cluster. If not specified, the latest recommended version will be used at provisioning time (but won't auto-upgrade).
        """
        return pulumi.get(self, "kubernetes_version")

    @property
    @pulumi.getter(name="linuxProfile")
    def linux_profile(self) -> pulumi.Output[Optional['outputs.KubernetesClusterLinuxProfile']]:
        """
        A `linux_profile` block as defined below.
        """
        return pulumi.get(self, "linux_profile")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location where the Managed Kubernetes Cluster should be created. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Managed Kubernetes Cluster to create. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkProfile")
    def network_profile(self) -> pulumi.Output['outputs.KubernetesClusterNetworkProfile']:
        """
        A `network_profile` block as defined below.
        """
        return pulumi.get(self, "network_profile")

    @property
    @pulumi.getter(name="nodeResourceGroup")
    def node_resource_group(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Kubernetes Nodes should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "node_resource_group")

    @property
    @pulumi.getter(name="privateClusterEnabled")
    def private_cluster_enabled(self) -> pulumi.Output[bool]:
        """
        Should this Kubernetes Cluster have its API server only exposed on internal IP addresses? This provides a Private IP Address for the Kubernetes API on the Virtual Network where the Kubernetes Cluster is located. Defaults to `false`. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "private_cluster_enabled")

    @property
    @pulumi.getter(name="privateFqdn")
    def private_fqdn(self) -> pulumi.Output[str]:
        """
        The FQDN for the Kubernetes Cluster when private link has been enabled, which is only resolvable inside the Virtual Network used by the Kubernetes Cluster.
        """
        return pulumi.get(self, "private_fqdn")

    @property
    @pulumi.getter(name="privateLinkEnabled")
    def private_link_enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "private_link_enabled")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        Specifies the Resource Group where the Managed Kubernetes Cluster should exist. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "resource_group_name")

    @property
    @pulumi.getter(name="roleBasedAccessControl")
    def role_based_access_control(self) -> pulumi.Output['outputs.KubernetesClusterRoleBasedAccessControl']:
        """
        A `role_based_access_control` block. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "role_based_access_control")

    @property
    @pulumi.getter(name="servicePrincipal")
    def service_principal(self) -> pulumi.Output[Optional['outputs.KubernetesClusterServicePrincipal']]:
        """
        A `service_principal` block as documented below.
        """
        return pulumi.get(self, "service_principal")

    @property
    @pulumi.getter(name="skuTier")
    def sku_tier(self) -> pulumi.Output[Optional[str]]:
        """
        The SKU Tier that should be used for this Kubernetes Cluster. Possible values are `Free` and `Paid` (which includes the Uptime SLA). Defaults to `Free`.
        """
        return pulumi.get(self, "sku_tier")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="windowsProfile")
    def windows_profile(self) -> pulumi.Output['outputs.KubernetesClusterWindowsProfile']:
        """
        A `windows_profile` block as defined below.
        """
        return pulumi.get(self, "windows_profile")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

