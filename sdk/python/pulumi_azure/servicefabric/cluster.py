# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Cluster(pulumi.CustomResource):
    add_on_features: pulumi.Output[list]
    """
    A List of one or more features which should be enabled, such as `DnsService`.
    """
    azure_active_directory: pulumi.Output[dict]
    """
    An `azure_active_directory` block as defined below.

      * `clientApplicationId` (`str`) - The Azure Active Directory Client ID which should be used for the Client Application.
      * `clusterApplicationId` (`str`) - The Azure Active Directory Cluster Application ID.
      * `tenant_id` (`str`) - The Azure Active Directory Tenant ID.
    """
    certificate: pulumi.Output[dict]
    """
    A `certificate` block as defined below. Conflicts with `certificate_common_names`.

      * `thumbprint` (`str`) - The Thumbprint of the Certificate.
      * `thumbprintSecondary` (`str`) - The Secondary Thumbprint of the Certificate.
      * `x509StoreName` (`str`) - The X509 Store where the Certificate Exists, such as `My`.
    """
    certificate_common_names: pulumi.Output[dict]
    """
    A `certificate_common_names` block as defined below. Conflicts with `certificate`.

      * `commonNames` (`list`) - A `common_names` block as defined below.
        * `certificateCommonName` (`str`) - The common or subject name of the certificate.
        * `certificateIssuerThumbprint` (`str`) - The Issuer Thumbprint of the Certificate.

      * `x509StoreName` (`str`) - The X509 Store where the Certificate Exists, such as `My`.
    """
    client_certificate_thumbprints: pulumi.Output[list]
    """
    One or two `client_certificate_thumbprint` blocks as defined below.

      * `isAdmin` (`bool`) - Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
      * `thumbprint` (`str`) - The Thumbprint associated with the Client Certificate.
    """
    cluster_code_version: pulumi.Output[str]
    """
    Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
    """
    cluster_endpoint: pulumi.Output[str]
    """
    The Cluster Endpoint for this Service Fabric Cluster.
    """
    diagnostics_config: pulumi.Output[dict]
    """
    A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.

      * `blobEndpoint` (`str`) - The Blob Endpoint of the Storage Account.
      * `protectedAccountKeyName` (`str`) - The protected diagnostics storage key name, such as `StorageAccountKey1`.
      * `queueEndpoint` (`str`) - The Queue Endpoint of the Storage Account.
      * `storage_account_name` (`str`) - The name of the Storage Account where the Diagnostics should be sent to.
      * `tableEndpoint` (`str`) - The Table Endpoint of the Storage Account.
    """
    fabric_settings: pulumi.Output[list]
    """
    One or more `fabric_settings` blocks as defined below.

      * `name` (`str`) - The name of the Fabric Setting, such as `Security` or `Federation`.
      * `parameters` (`dict`) - A map containing settings for the specified Fabric Setting.
    """
    location: pulumi.Output[str]
    """
    Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
    """
    management_endpoint: pulumi.Output[str]
    """
    Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
    """
    node_types: pulumi.Output[list]
    """
    One or more `node_type` blocks as defined below.

      * `applicationPorts` (`dict`) - A `application_ports` block as defined below.
        * `endPort` (`float`) - The end of the Application Port Range on this Node Type.
        * `startPort` (`float`) - The start of the Application Port Range on this Node Type.

      * `capacities` (`dict`) - The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
      * `clientEndpointPort` (`float`) - The Port used for the Client Endpoint for this Node Type. Changing this forces a new resource to be created.
      * `durabilityLevel` (`str`) - The Durability Level for this Node Type. Possible values include `Bronze`, `Gold` and `Silver`. Defaults to `Bronze`. Changing this forces a new resource to be created.
      * `ephemeralPorts` (`dict`) - A `ephemeral_ports` block as defined below.
        * `endPort` (`float`) - The end of the Ephemeral Port Range on this Node Type.
        * `startPort` (`float`) - The start of the Ephemeral Port Range on this Node Type.

      * `httpEndpointPort` (`float`) - The Port used for the HTTP Endpoint for this Node Type. Changing this forces a new resource to be created.
      * `instanceCount` (`float`) - The number of nodes for this Node Type.
      * `isPrimary` (`bool`) - Is this the Primary Node Type? Changing this forces a new resource to be created.
      * `name` (`str`) - The name of the Node Type. Changing this forces a new resource to be created.
      * `placementProperties` (`dict`) - The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
      * `reverseProxyEndpointPort` (`float`) - The Port used for the Reverse Proxy Endpoint  for this Node Type. Changing this will upgrade the cluster.
    """
    reliability_level: pulumi.Output[str]
    """
    Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
    """
    reverse_proxy_certificate: pulumi.Output[dict]
    """
    A `reverse_proxy_certificate` block as defined below.

      * `thumbprint` (`str`) - The Thumbprint of the Certificate.
      * `thumbprintSecondary` (`str`) - The Secondary Thumbprint of the Certificate.
      * `x509StoreName` (`str`) - The X509 Store where the Certificate Exists, such as `My`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    upgrade_mode: pulumi.Output[str]
    """
    Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
    """
    vm_image: pulumi.Output[str]
    """
    Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, add_on_features=None, azure_active_directory=None, certificate=None, certificate_common_names=None, client_certificate_thumbprints=None, cluster_code_version=None, diagnostics_config=None, fabric_settings=None, location=None, management_endpoint=None, name=None, node_types=None, reliability_level=None, resource_group_name=None, reverse_proxy_certificate=None, tags=None, upgrade_mode=None, vm_image=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Service Fabric Cluster.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/service_fabric_cluster.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] add_on_features: A List of one or more features which should be enabled, such as `DnsService`.
        :param pulumi.Input[dict] azure_active_directory: An `azure_active_directory` block as defined below.
        :param pulumi.Input[dict] certificate: A `certificate` block as defined below. Conflicts with `certificate_common_names`.
        :param pulumi.Input[dict] certificate_common_names: A `certificate_common_names` block as defined below. Conflicts with `certificate`.
        :param pulumi.Input[list] client_certificate_thumbprints: One or two `client_certificate_thumbprint` blocks as defined below.
        :param pulumi.Input[str] cluster_code_version: Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
        :param pulumi.Input[dict] diagnostics_config: A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[list] fabric_settings: One or more `fabric_settings` blocks as defined below.
        :param pulumi.Input[str] location: Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_endpoint: Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[list] node_types: One or more `node_type` blocks as defined below.
        :param pulumi.Input[str] reliability_level: Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] reverse_proxy_certificate: A `reverse_proxy_certificate` block as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] upgrade_mode: Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
        :param pulumi.Input[str] vm_image: Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.

        The **azure_active_directory** object supports the following:

          * `clientApplicationId` (`pulumi.Input[str]`) - The Azure Active Directory Client ID which should be used for the Client Application.
          * `clusterApplicationId` (`pulumi.Input[str]`) - The Azure Active Directory Cluster Application ID.
          * `tenant_id` (`pulumi.Input[str]`) - The Azure Active Directory Tenant ID.

        The **certificate** object supports the following:

          * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint of the Certificate.
          * `thumbprintSecondary` (`pulumi.Input[str]`) - The Secondary Thumbprint of the Certificate.
          * `x509StoreName` (`pulumi.Input[str]`) - The X509 Store where the Certificate Exists, such as `My`.

        The **certificate_common_names** object supports the following:

          * `commonNames` (`pulumi.Input[list]`) - A `common_names` block as defined below.
            * `certificateCommonName` (`pulumi.Input[str]`) - The common or subject name of the certificate.
            * `certificateIssuerThumbprint` (`pulumi.Input[str]`) - The Issuer Thumbprint of the Certificate.

          * `x509StoreName` (`pulumi.Input[str]`) - The X509 Store where the Certificate Exists, such as `My`.

        The **client_certificate_thumbprints** object supports the following:

          * `isAdmin` (`pulumi.Input[bool]`) - Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
          * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint associated with the Client Certificate.

        The **diagnostics_config** object supports the following:

          * `blobEndpoint` (`pulumi.Input[str]`) - The Blob Endpoint of the Storage Account.
          * `protectedAccountKeyName` (`pulumi.Input[str]`) - The protected diagnostics storage key name, such as `StorageAccountKey1`.
          * `queueEndpoint` (`pulumi.Input[str]`) - The Queue Endpoint of the Storage Account.
          * `storage_account_name` (`pulumi.Input[str]`) - The name of the Storage Account where the Diagnostics should be sent to.
          * `tableEndpoint` (`pulumi.Input[str]`) - The Table Endpoint of the Storage Account.

        The **fabric_settings** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the Fabric Setting, such as `Security` or `Federation`.
          * `parameters` (`pulumi.Input[dict]`) - A map containing settings for the specified Fabric Setting.

        The **node_types** object supports the following:

          * `applicationPorts` (`pulumi.Input[dict]`) - A `application_ports` block as defined below.
            * `endPort` (`pulumi.Input[float]`) - The end of the Application Port Range on this Node Type.
            * `startPort` (`pulumi.Input[float]`) - The start of the Application Port Range on this Node Type.

          * `capacities` (`pulumi.Input[dict]`) - The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
          * `clientEndpointPort` (`pulumi.Input[float]`) - The Port used for the Client Endpoint for this Node Type. Changing this forces a new resource to be created.
          * `durabilityLevel` (`pulumi.Input[str]`) - The Durability Level for this Node Type. Possible values include `Bronze`, `Gold` and `Silver`. Defaults to `Bronze`. Changing this forces a new resource to be created.
          * `ephemeralPorts` (`pulumi.Input[dict]`) - A `ephemeral_ports` block as defined below.
            * `endPort` (`pulumi.Input[float]`) - The end of the Ephemeral Port Range on this Node Type.
            * `startPort` (`pulumi.Input[float]`) - The start of the Ephemeral Port Range on this Node Type.

          * `httpEndpointPort` (`pulumi.Input[float]`) - The Port used for the HTTP Endpoint for this Node Type. Changing this forces a new resource to be created.
          * `instanceCount` (`pulumi.Input[float]`) - The number of nodes for this Node Type.
          * `isPrimary` (`pulumi.Input[bool]`) - Is this the Primary Node Type? Changing this forces a new resource to be created.
          * `name` (`pulumi.Input[str]`) - The name of the Node Type. Changing this forces a new resource to be created.
          * `placementProperties` (`pulumi.Input[dict]`) - The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
          * `reverseProxyEndpointPort` (`pulumi.Input[float]`) - The Port used for the Reverse Proxy Endpoint  for this Node Type. Changing this will upgrade the cluster.

        The **reverse_proxy_certificate** object supports the following:

          * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint of the Certificate.
          * `thumbprintSecondary` (`pulumi.Input[str]`) - The Secondary Thumbprint of the Certificate.
          * `x509StoreName` (`pulumi.Input[str]`) - The X509 Store where the Certificate Exists, such as `My`.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['add_on_features'] = add_on_features
            __props__['azure_active_directory'] = azure_active_directory
            __props__['certificate'] = certificate
            __props__['certificate_common_names'] = certificate_common_names
            __props__['client_certificate_thumbprints'] = client_certificate_thumbprints
            __props__['cluster_code_version'] = cluster_code_version
            __props__['diagnostics_config'] = diagnostics_config
            __props__['fabric_settings'] = fabric_settings
            __props__['location'] = location
            if management_endpoint is None:
                raise TypeError("Missing required property 'management_endpoint'")
            __props__['management_endpoint'] = management_endpoint
            __props__['name'] = name
            if node_types is None:
                raise TypeError("Missing required property 'node_types'")
            __props__['node_types'] = node_types
            if reliability_level is None:
                raise TypeError("Missing required property 'reliability_level'")
            __props__['reliability_level'] = reliability_level
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['reverse_proxy_certificate'] = reverse_proxy_certificate
            __props__['tags'] = tags
            if upgrade_mode is None:
                raise TypeError("Missing required property 'upgrade_mode'")
            __props__['upgrade_mode'] = upgrade_mode
            if vm_image is None:
                raise TypeError("Missing required property 'vm_image'")
            __props__['vm_image'] = vm_image
            __props__['cluster_endpoint'] = None
        super(Cluster, __self__).__init__(
            'azure:servicefabric/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, add_on_features=None, azure_active_directory=None, certificate=None, certificate_common_names=None, client_certificate_thumbprints=None, cluster_code_version=None, cluster_endpoint=None, diagnostics_config=None, fabric_settings=None, location=None, management_endpoint=None, name=None, node_types=None, reliability_level=None, resource_group_name=None, reverse_proxy_certificate=None, tags=None, upgrade_mode=None, vm_image=None):
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] add_on_features: A List of one or more features which should be enabled, such as `DnsService`.
        :param pulumi.Input[dict] azure_active_directory: An `azure_active_directory` block as defined below.
        :param pulumi.Input[dict] certificate: A `certificate` block as defined below. Conflicts with `certificate_common_names`.
        :param pulumi.Input[dict] certificate_common_names: A `certificate_common_names` block as defined below. Conflicts with `certificate`.
        :param pulumi.Input[list] client_certificate_thumbprints: One or two `client_certificate_thumbprint` blocks as defined below.
        :param pulumi.Input[str] cluster_code_version: Required if Upgrade Mode set to `Manual`, Specifies the Version of the Cluster Code of the cluster.
        :param pulumi.Input[str] cluster_endpoint: The Cluster Endpoint for this Service Fabric Cluster.
        :param pulumi.Input[dict] diagnostics_config: A `diagnostics_config` block as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[list] fabric_settings: One or more `fabric_settings` blocks as defined below.
        :param pulumi.Input[str] location: Specifies the Azure Region where the Service Fabric Cluster should exist. Changing this forces a new resource to be created.
        :param pulumi.Input[str] management_endpoint: Specifies the Management Endpoint of the cluster such as `http://example.com`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the Service Fabric Cluster. Changing this forces a new resource to be created.
        :param pulumi.Input[list] node_types: One or more `node_type` blocks as defined below.
        :param pulumi.Input[str] reliability_level: Specifies the Reliability Level of the Cluster. Possible values include `None`, `Bronze`, `Silver`, `Gold` and `Platinum`.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group in which the Service Fabric Cluster exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] reverse_proxy_certificate: A `reverse_proxy_certificate` block as defined below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] upgrade_mode: Specifies the Upgrade Mode of the cluster. Possible values are `Automatic` or `Manual`.
        :param pulumi.Input[str] vm_image: Specifies the Image expected for the Service Fabric Cluster, such as `Windows`. Changing this forces a new resource to be created.

        The **azure_active_directory** object supports the following:

          * `clientApplicationId` (`pulumi.Input[str]`) - The Azure Active Directory Client ID which should be used for the Client Application.
          * `clusterApplicationId` (`pulumi.Input[str]`) - The Azure Active Directory Cluster Application ID.
          * `tenant_id` (`pulumi.Input[str]`) - The Azure Active Directory Tenant ID.

        The **certificate** object supports the following:

          * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint of the Certificate.
          * `thumbprintSecondary` (`pulumi.Input[str]`) - The Secondary Thumbprint of the Certificate.
          * `x509StoreName` (`pulumi.Input[str]`) - The X509 Store where the Certificate Exists, such as `My`.

        The **certificate_common_names** object supports the following:

          * `commonNames` (`pulumi.Input[list]`) - A `common_names` block as defined below.
            * `certificateCommonName` (`pulumi.Input[str]`) - The common or subject name of the certificate.
            * `certificateIssuerThumbprint` (`pulumi.Input[str]`) - The Issuer Thumbprint of the Certificate.

          * `x509StoreName` (`pulumi.Input[str]`) - The X509 Store where the Certificate Exists, such as `My`.

        The **client_certificate_thumbprints** object supports the following:

          * `isAdmin` (`pulumi.Input[bool]`) - Does the Client Certificate have Admin Access to the cluster? Non-admin clients can only perform read only operations on the cluster.
          * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint associated with the Client Certificate.

        The **diagnostics_config** object supports the following:

          * `blobEndpoint` (`pulumi.Input[str]`) - The Blob Endpoint of the Storage Account.
          * `protectedAccountKeyName` (`pulumi.Input[str]`) - The protected diagnostics storage key name, such as `StorageAccountKey1`.
          * `queueEndpoint` (`pulumi.Input[str]`) - The Queue Endpoint of the Storage Account.
          * `storage_account_name` (`pulumi.Input[str]`) - The name of the Storage Account where the Diagnostics should be sent to.
          * `tableEndpoint` (`pulumi.Input[str]`) - The Table Endpoint of the Storage Account.

        The **fabric_settings** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the Fabric Setting, such as `Security` or `Federation`.
          * `parameters` (`pulumi.Input[dict]`) - A map containing settings for the specified Fabric Setting.

        The **node_types** object supports the following:

          * `applicationPorts` (`pulumi.Input[dict]`) - A `application_ports` block as defined below.
            * `endPort` (`pulumi.Input[float]`) - The end of the Application Port Range on this Node Type.
            * `startPort` (`pulumi.Input[float]`) - The start of the Application Port Range on this Node Type.

          * `capacities` (`pulumi.Input[dict]`) - The capacity tags applied to the nodes in the node type, the cluster resource manager uses these tags to understand how much resource a node has.
          * `clientEndpointPort` (`pulumi.Input[float]`) - The Port used for the Client Endpoint for this Node Type. Changing this forces a new resource to be created.
          * `durabilityLevel` (`pulumi.Input[str]`) - The Durability Level for this Node Type. Possible values include `Bronze`, `Gold` and `Silver`. Defaults to `Bronze`. Changing this forces a new resource to be created.
          * `ephemeralPorts` (`pulumi.Input[dict]`) - A `ephemeral_ports` block as defined below.
            * `endPort` (`pulumi.Input[float]`) - The end of the Ephemeral Port Range on this Node Type.
            * `startPort` (`pulumi.Input[float]`) - The start of the Ephemeral Port Range on this Node Type.

          * `httpEndpointPort` (`pulumi.Input[float]`) - The Port used for the HTTP Endpoint for this Node Type. Changing this forces a new resource to be created.
          * `instanceCount` (`pulumi.Input[float]`) - The number of nodes for this Node Type.
          * `isPrimary` (`pulumi.Input[bool]`) - Is this the Primary Node Type? Changing this forces a new resource to be created.
          * `name` (`pulumi.Input[str]`) - The name of the Node Type. Changing this forces a new resource to be created.
          * `placementProperties` (`pulumi.Input[dict]`) - The placement tags applied to nodes in the node type, which can be used to indicate where certain services (workload) should run.
          * `reverseProxyEndpointPort` (`pulumi.Input[float]`) - The Port used for the Reverse Proxy Endpoint  for this Node Type. Changing this will upgrade the cluster.

        The **reverse_proxy_certificate** object supports the following:

          * `thumbprint` (`pulumi.Input[str]`) - The Thumbprint of the Certificate.
          * `thumbprintSecondary` (`pulumi.Input[str]`) - The Secondary Thumbprint of the Certificate.
          * `x509StoreName` (`pulumi.Input[str]`) - The X509 Store where the Certificate Exists, such as `My`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_on_features"] = add_on_features
        __props__["azure_active_directory"] = azure_active_directory
        __props__["certificate"] = certificate
        __props__["certificate_common_names"] = certificate_common_names
        __props__["client_certificate_thumbprints"] = client_certificate_thumbprints
        __props__["cluster_code_version"] = cluster_code_version
        __props__["cluster_endpoint"] = cluster_endpoint
        __props__["diagnostics_config"] = diagnostics_config
        __props__["fabric_settings"] = fabric_settings
        __props__["location"] = location
        __props__["management_endpoint"] = management_endpoint
        __props__["name"] = name
        __props__["node_types"] = node_types
        __props__["reliability_level"] = reliability_level
        __props__["resource_group_name"] = resource_group_name
        __props__["reverse_proxy_certificate"] = reverse_proxy_certificate
        __props__["tags"] = tags
        __props__["upgrade_mode"] = upgrade_mode
        __props__["vm_image"] = vm_image
        return Cluster(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

