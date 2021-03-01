// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement
{
    /// <summary>
    /// Manages an API Management Service.
    /// 
    /// ## Disclaimers
    /// 
    /// &gt; **Note:** It's possible to define Custom Domains both within the `azure.apimanagement.Service` resource via the `hostname_configurations` block and by using the `azure.apimanagement.CustomDomain` resource. However it's not possible to use both methods to manage Custom Domains within an API Management Service, since there'll be conflicts.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleService = new Azure.ApiManagement.Service("exampleService", new Azure.ApiManagement.ServiceArgs
    ///         {
    ///             Location = exampleResourceGroup.Location,
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             PublisherName = "My Company",
    ///             PublisherEmail = "company@exmaple.com",
    ///             SkuName = "Developer_1",
    ///             Policy = new Azure.ApiManagement.Inputs.ServicePolicyArgs
    ///             {
    ///                 XmlContent = @"    &lt;policies&gt;
    ///       &lt;inbound /&gt;
    ///       &lt;backend /&gt;
    ///       &lt;outbound /&gt;
    ///       &lt;on-error /&gt;
    ///     &lt;/policies&gt;
    /// ",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// API Management Services can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:apimanagement/service:Service example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1
    /// ```
    /// </summary>
    [AzureResourceType("azure:apimanagement/service:Service")]
    public partial class Service : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `additional_location` blocks as defined below.
        /// </summary>
        [Output("additionalLocations")]
        public Output<ImmutableArray<Outputs.ServiceAdditionalLocation>> AdditionalLocations { get; private set; } = null!;

        /// <summary>
        /// One or more (up to 10) `certificate` blocks as defined below.
        /// </summary>
        [Output("certificates")]
        public Output<ImmutableArray<Outputs.ServiceCertificate>> Certificates { get; private set; } = null!;

        /// <summary>
        /// The URL for the Developer Portal associated with this API Management service.
        /// </summary>
        [Output("developerPortalUrl")]
        public Output<string> DeveloperPortalUrl { get; private set; } = null!;

        /// <summary>
        /// The URL of the Regional Gateway for the API Management Service in the specified region.
        /// </summary>
        [Output("gatewayRegionalUrl")]
        public Output<string> GatewayRegionalUrl { get; private set; } = null!;

        /// <summary>
        /// The URL of the Gateway for the API Management Service.
        /// </summary>
        [Output("gatewayUrl")]
        public Output<string> GatewayUrl { get; private set; } = null!;

        /// <summary>
        /// A `hostname_configuration` block as defined below.
        /// </summary>
        [Output("hostnameConfiguration")]
        public Output<Outputs.ServiceHostnameConfiguration?> HostnameConfiguration { get; private set; } = null!;

        /// <summary>
        /// An `identity` block is documented below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ServiceIdentity?> Identity { get; private set; } = null!;

        /// <summary>
        /// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The URL for the Management API associated with this API Management service.
        /// </summary>
        [Output("managementApiUrl")]
        public Output<string> ManagementApiUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Email address from which the notification will be sent.
        /// </summary>
        [Output("notificationSenderEmail")]
        public Output<string> NotificationSenderEmail { get; private set; } = null!;

        /// <summary>
        /// A `policy` block as defined below.
        /// </summary>
        [Output("policy")]
        public Output<Outputs.ServicePolicy> Policy { get; private set; } = null!;

        /// <summary>
        /// The URL for the Publisher Portal associated with this API Management service.
        /// </summary>
        [Output("portalUrl")]
        public Output<string> PortalUrl { get; private set; } = null!;

        /// <summary>
        /// The Private IP addresses of the API Management Service.  Available only when the API Manager instance is using Virtual Network mode.
        /// </summary>
        [Output("privateIpAddresses")]
        public Output<ImmutableArray<string>> PrivateIpAddresses { get; private set; } = null!;

        /// <summary>
        /// A `protocols` block as defined below.
        /// </summary>
        [Output("protocols")]
        public Output<Outputs.ServiceProtocols> Protocols { get; private set; } = null!;

        /// <summary>
        /// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
        /// </summary>
        [Output("publicIpAddresses")]
        public Output<ImmutableArray<string>> PublicIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The email of publisher/company.
        /// </summary>
        [Output("publisherEmail")]
        public Output<string> PublisherEmail { get; private set; } = null!;

        /// <summary>
        /// The name of publisher/company.
        /// </summary>
        [Output("publisherName")]
        public Output<string> PublisherName { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
        /// </summary>
        [Output("scmUrl")]
        public Output<string> ScmUrl { get; private set; } = null!;

        /// <summary>
        /// A `security` block as defined below.
        /// </summary>
        [Output("security")]
        public Output<Outputs.ServiceSecurity> Security { get; private set; } = null!;

        /// <summary>
        /// A `sign_in` block as defined below.
        /// </summary>
        [Output("signIn")]
        public Output<Outputs.ServiceSignIn> SignIn { get; private set; } = null!;

        /// <summary>
        /// A `sign_up` block as defined below.
        /// </summary>
        [Output("signUp")]
        public Output<Outputs.ServiceSignUp> SignUp { get; private set; } = null!;

        /// <summary>
        /// `sku_name` is a string consisting of two parts separated by an underscore(\_). The first part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
        /// </summary>
        [Output("skuName")]
        public Output<string> SkuName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// A `tenant_access` block as defined below.
        /// </summary>
        [Output("tenantAccess")]
        public Output<Outputs.ServiceTenantAccess> TenantAccess { get; private set; } = null!;

        /// <summary>
        /// A `virtual_network_configuration` block as defined below. Required when `virtual_network_type` is `External` or `Internal`.
        /// </summary>
        [Output("virtualNetworkConfiguration")]
        public Output<Outputs.ServiceVirtualNetworkConfiguration?> VirtualNetworkConfiguration { get; private set; } = null!;

        /// <summary>
        /// The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`. 
        /// &gt; **NOTE:** Please ensure that in the subnet, inbound port 3443 is open when `virtual_network_type` is `Internal` or `External`. And please ensure other necessary ports are open according to [api management network configuration](https://docs.microsoft.com/en-us/azure/api-management/api-management-using-with-vnet#-common-network-configuration-issues).
        /// </summary>
        [Output("virtualNetworkType")]
        public Output<string?> VirtualNetworkType { get; private set; } = null!;


        /// <summary>
        /// Create a Service resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Service(string name, ServiceArgs args, CustomResourceOptions? options = null)
            : base("azure:apimanagement/service:Service", name, args ?? new ServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Service(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
            : base("azure:apimanagement/service:Service", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Service resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Service Get(string name, Input<string> id, ServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new Service(name, id, state, options);
        }
    }

    public sealed class ServiceArgs : Pulumi.ResourceArgs
    {
        [Input("additionalLocations")]
        private InputList<Inputs.ServiceAdditionalLocationArgs>? _additionalLocations;

        /// <summary>
        /// One or more `additional_location` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ServiceAdditionalLocationArgs> AdditionalLocations
        {
            get => _additionalLocations ?? (_additionalLocations = new InputList<Inputs.ServiceAdditionalLocationArgs>());
            set => _additionalLocations = value;
        }

        [Input("certificates")]
        private InputList<Inputs.ServiceCertificateArgs>? _certificates;

        /// <summary>
        /// One or more (up to 10) `certificate` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ServiceCertificateArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.ServiceCertificateArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// A `hostname_configuration` block as defined below.
        /// </summary>
        [Input("hostnameConfiguration")]
        public Input<Inputs.ServiceHostnameConfigurationArgs>? HostnameConfiguration { get; set; }

        /// <summary>
        /// An `identity` block is documented below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Email address from which the notification will be sent.
        /// </summary>
        [Input("notificationSenderEmail")]
        public Input<string>? NotificationSenderEmail { get; set; }

        /// <summary>
        /// A `policy` block as defined below.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.ServicePolicyArgs>? Policy { get; set; }

        /// <summary>
        /// A `protocols` block as defined below.
        /// </summary>
        [Input("protocols")]
        public Input<Inputs.ServiceProtocolsArgs>? Protocols { get; set; }

        /// <summary>
        /// The email of publisher/company.
        /// </summary>
        [Input("publisherEmail", required: true)]
        public Input<string> PublisherEmail { get; set; } = null!;

        /// <summary>
        /// The name of publisher/company.
        /// </summary>
        [Input("publisherName", required: true)]
        public Input<string> PublisherName { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A `security` block as defined below.
        /// </summary>
        [Input("security")]
        public Input<Inputs.ServiceSecurityArgs>? Security { get; set; }

        /// <summary>
        /// A `sign_in` block as defined below.
        /// </summary>
        [Input("signIn")]
        public Input<Inputs.ServiceSignInArgs>? SignIn { get; set; }

        /// <summary>
        /// A `sign_up` block as defined below.
        /// </summary>
        [Input("signUp")]
        public Input<Inputs.ServiceSignUpArgs>? SignUp { get; set; }

        /// <summary>
        /// `sku_name` is a string consisting of two parts separated by an underscore(\_). The first part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
        /// </summary>
        [Input("skuName", required: true)]
        public Input<string> SkuName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A `tenant_access` block as defined below.
        /// </summary>
        [Input("tenantAccess")]
        public Input<Inputs.ServiceTenantAccessArgs>? TenantAccess { get; set; }

        /// <summary>
        /// A `virtual_network_configuration` block as defined below. Required when `virtual_network_type` is `External` or `Internal`.
        /// </summary>
        [Input("virtualNetworkConfiguration")]
        public Input<Inputs.ServiceVirtualNetworkConfigurationArgs>? VirtualNetworkConfiguration { get; set; }

        /// <summary>
        /// The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`. 
        /// &gt; **NOTE:** Please ensure that in the subnet, inbound port 3443 is open when `virtual_network_type` is `Internal` or `External`. And please ensure other necessary ports are open according to [api management network configuration](https://docs.microsoft.com/en-us/azure/api-management/api-management-using-with-vnet#-common-network-configuration-issues).
        /// </summary>
        [Input("virtualNetworkType")]
        public Input<string>? VirtualNetworkType { get; set; }

        public ServiceArgs()
        {
        }
    }

    public sealed class ServiceState : Pulumi.ResourceArgs
    {
        [Input("additionalLocations")]
        private InputList<Inputs.ServiceAdditionalLocationGetArgs>? _additionalLocations;

        /// <summary>
        /// One or more `additional_location` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ServiceAdditionalLocationGetArgs> AdditionalLocations
        {
            get => _additionalLocations ?? (_additionalLocations = new InputList<Inputs.ServiceAdditionalLocationGetArgs>());
            set => _additionalLocations = value;
        }

        [Input("certificates")]
        private InputList<Inputs.ServiceCertificateGetArgs>? _certificates;

        /// <summary>
        /// One or more (up to 10) `certificate` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ServiceCertificateGetArgs> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<Inputs.ServiceCertificateGetArgs>());
            set => _certificates = value;
        }

        /// <summary>
        /// The URL for the Developer Portal associated with this API Management service.
        /// </summary>
        [Input("developerPortalUrl")]
        public Input<string>? DeveloperPortalUrl { get; set; }

        /// <summary>
        /// The URL of the Regional Gateway for the API Management Service in the specified region.
        /// </summary>
        [Input("gatewayRegionalUrl")]
        public Input<string>? GatewayRegionalUrl { get; set; }

        /// <summary>
        /// The URL of the Gateway for the API Management Service.
        /// </summary>
        [Input("gatewayUrl")]
        public Input<string>? GatewayUrl { get; set; }

        /// <summary>
        /// A `hostname_configuration` block as defined below.
        /// </summary>
        [Input("hostnameConfiguration")]
        public Input<Inputs.ServiceHostnameConfigurationGetArgs>? HostnameConfiguration { get; set; }

        /// <summary>
        /// An `identity` block is documented below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ServiceIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// The Azure location where the API Management Service exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The URL for the Management API associated with this API Management service.
        /// </summary>
        [Input("managementApiUrl")]
        public Input<string>? ManagementApiUrl { get; set; }

        /// <summary>
        /// The name of the API Management Service. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Email address from which the notification will be sent.
        /// </summary>
        [Input("notificationSenderEmail")]
        public Input<string>? NotificationSenderEmail { get; set; }

        /// <summary>
        /// A `policy` block as defined below.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.ServicePolicyGetArgs>? Policy { get; set; }

        /// <summary>
        /// The URL for the Publisher Portal associated with this API Management service.
        /// </summary>
        [Input("portalUrl")]
        public Input<string>? PortalUrl { get; set; }

        [Input("privateIpAddresses")]
        private InputList<string>? _privateIpAddresses;

        /// <summary>
        /// The Private IP addresses of the API Management Service.  Available only when the API Manager instance is using Virtual Network mode.
        /// </summary>
        public InputList<string> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<string>());
            set => _privateIpAddresses = value;
        }

        /// <summary>
        /// A `protocols` block as defined below.
        /// </summary>
        [Input("protocols")]
        public Input<Inputs.ServiceProtocolsGetArgs>? Protocols { get; set; }

        [Input("publicIpAddresses")]
        private InputList<string>? _publicIpAddresses;

        /// <summary>
        /// Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
        /// </summary>
        public InputList<string> PublicIpAddresses
        {
            get => _publicIpAddresses ?? (_publicIpAddresses = new InputList<string>());
            set => _publicIpAddresses = value;
        }

        /// <summary>
        /// The email of publisher/company.
        /// </summary>
        [Input("publisherEmail")]
        public Input<string>? PublisherEmail { get; set; }

        /// <summary>
        /// The name of publisher/company.
        /// </summary>
        [Input("publisherName")]
        public Input<string>? PublisherName { get; set; }

        /// <summary>
        /// The name of the Resource Group in which the API Management Service should be exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The URL for the SCM (Source Code Management) Endpoint associated with this API Management service.
        /// </summary>
        [Input("scmUrl")]
        public Input<string>? ScmUrl { get; set; }

        /// <summary>
        /// A `security` block as defined below.
        /// </summary>
        [Input("security")]
        public Input<Inputs.ServiceSecurityGetArgs>? Security { get; set; }

        /// <summary>
        /// A `sign_in` block as defined below.
        /// </summary>
        [Input("signIn")]
        public Input<Inputs.ServiceSignInGetArgs>? SignIn { get; set; }

        /// <summary>
        /// A `sign_up` block as defined below.
        /// </summary>
        [Input("signUp")]
        public Input<Inputs.ServiceSignUpGetArgs>? SignUp { get; set; }

        /// <summary>
        /// `sku_name` is a string consisting of two parts separated by an underscore(\_). The first part is the `name`, valid values include: `Consumption`, `Developer`, `Basic`, `Standard` and `Premium`. The second part is the `capacity` (e.g. the number of deployed units of the `sku`), which must be a positive `integer` (e.g. `Developer_1`).
        /// </summary>
        [Input("skuName")]
        public Input<string>? SkuName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A `tenant_access` block as defined below.
        /// </summary>
        [Input("tenantAccess")]
        public Input<Inputs.ServiceTenantAccessGetArgs>? TenantAccess { get; set; }

        /// <summary>
        /// A `virtual_network_configuration` block as defined below. Required when `virtual_network_type` is `External` or `Internal`.
        /// </summary>
        [Input("virtualNetworkConfiguration")]
        public Input<Inputs.ServiceVirtualNetworkConfigurationGetArgs>? VirtualNetworkConfiguration { get; set; }

        /// <summary>
        /// The type of virtual network you want to use, valid values include: `None`, `External`, `Internal`. 
        /// &gt; **NOTE:** Please ensure that in the subnet, inbound port 3443 is open when `virtual_network_type` is `Internal` or `External`. And please ensure other necessary ports are open according to [api management network configuration](https://docs.microsoft.com/en-us/azure/api-management/api-management-using-with-vnet#-common-network-configuration-issues).
        /// </summary>
        [Input("virtualNetworkType")]
        public Input<string>? VirtualNetworkType { get; set; }

        public ServiceState()
        {
        }
    }
}
