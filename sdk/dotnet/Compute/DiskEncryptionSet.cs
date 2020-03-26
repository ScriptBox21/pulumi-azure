// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    /// <summary>
    /// Manages a Disk Encryption Set.
    /// 
    /// &gt; **NOTE**: Disk Encryption Sets are in Public Preview and at this time is only available in `Canada Central`, `North Europe` and `West Central US` regions - [more information can be found in the preview documentation](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/disk-encryption).
    /// 
    /// &gt; **NOTE:** At this time the Key Vault used to store the Active Key for this Disk Encryption Set must have both Soft Delete &amp; Purge Protection enabled - which are not yet supported by this provider.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/disk_encryption_set.html.markdown.
    /// </summary>
    public partial class DiskEncryptionSet : Pulumi.CustomResource
    {
        /// <summary>
        /// A `identity` block defined below.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.DiskEncryptionSetIdentity> Identity { get; private set; } = null!;

        /// <summary>
        /// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
        /// </summary>
        [Output("keyVaultKeyId")]
        public Output<string> KeyVaultKeyId { get; private set; } = null!;

        /// <summary>
        /// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the Disk Encryption Set.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DiskEncryptionSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DiskEncryptionSet(string name, DiskEncryptionSetArgs args, CustomResourceOptions? options = null)
            : base("azure:compute/diskEncryptionSet:DiskEncryptionSet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DiskEncryptionSet(string name, Input<string> id, DiskEncryptionSetState? state = null, CustomResourceOptions? options = null)
            : base("azure:compute/diskEncryptionSet:DiskEncryptionSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DiskEncryptionSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DiskEncryptionSet Get(string name, Input<string> id, DiskEncryptionSetState? state = null, CustomResourceOptions? options = null)
        {
            return new DiskEncryptionSet(name, id, state, options);
        }
    }

    public sealed class DiskEncryptionSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `identity` block defined below.
        /// </summary>
        [Input("identity", required: true)]
        public Input<Inputs.DiskEncryptionSetIdentityArgs> Identity { get; set; } = null!;

        /// <summary>
        /// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
        /// </summary>
        [Input("keyVaultKeyId", required: true)]
        public Input<string> KeyVaultKeyId { get; set; } = null!;

        /// <summary>
        /// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Disk Encryption Set.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DiskEncryptionSetArgs()
        {
        }
    }

    public sealed class DiskEncryptionSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `identity` block defined below.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.DiskEncryptionSetIdentityGetArgs>? Identity { get; set; }

        /// <summary>
        /// Specifies the URL to a Key Vault Key (either from a Key Vault Key, or the Key URL for the Key Vault Secret).
        /// </summary>
        [Input("keyVaultKeyId")]
        public Input<string>? KeyVaultKeyId { get; set; }

        /// <summary>
        /// Specifies the Azure Region where the Disk Encryption Set exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Disk Encryption Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies the name of the Resource Group where the Disk Encryption Set should exist. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the Disk Encryption Set.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public DiskEncryptionSetState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DiskEncryptionSetIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The (Client) ID of the Service Principal.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The ID of the Tenant the Service Principal is assigned in.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DiskEncryptionSetIdentityArgs()
        {
        }
    }

    public sealed class DiskEncryptionSetIdentityGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The (Client) ID of the Service Principal.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The ID of the Tenant the Service Principal is assigned in.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DiskEncryptionSetIdentityGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DiskEncryptionSetIdentity
    {
        /// <summary>
        /// The (Client) ID of the Service Principal.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The ID of the Tenant the Service Principal is assigned in.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The Type of Identity which should be used for this Disk Encryption Set. At this time the only possible value is `SystemAssigned`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DiskEncryptionSetIdentity(
            string principalId,
            string tenantId,
            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }
    }
}
