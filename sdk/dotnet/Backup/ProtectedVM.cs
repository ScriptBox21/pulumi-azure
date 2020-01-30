// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Backup
{
    /// <summary>
    /// Manages Azure Backup for an Azure VM
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/backup_protected_vm.html.markdown.
    /// </summary>
    public partial class ProtectedVM : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the id of the backup policy to use.
        /// </summary>
        [Output("backupPolicyId")]
        public Output<string> BackupPolicyId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Output("recoveryVaultName")]
        public Output<string> RecoveryVaultName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
        /// </summary>
        [Output("sourceVmId")]
        public Output<string> SourceVmId { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ProtectedVM resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProtectedVM(string name, ProtectedVMArgs args, CustomResourceOptions? options = null)
            : base("azure:backup/protectedVM:ProtectedVM", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ProtectedVM(string name, Input<string> id, ProtectedVMState? state = null, CustomResourceOptions? options = null)
            : base("azure:backup/protectedVM:ProtectedVM", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProtectedVM resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProtectedVM Get(string name, Input<string> id, ProtectedVMState? state = null, CustomResourceOptions? options = null)
        {
            return new ProtectedVM(name, id, state, options);
        }
    }

    public sealed class ProtectedVMArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the backup policy to use.
        /// </summary>
        [Input("backupPolicyId", required: true)]
        public Input<string> BackupPolicyId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public Input<string> RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceVmId", required: true)]
        public Input<string> SourceVmId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ProtectedVMArgs()
        {
        }
    }

    public sealed class ProtectedVMState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the id of the backup policy to use.
        /// </summary>
        [Input("backupPolicyId")]
        public Input<string>? BackupPolicyId { get; set; }

        /// <summary>
        /// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
        /// </summary>
        [Input("recoveryVaultName")]
        public Input<string>? RecoveryVaultName { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the Recovery Services Protected VM. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
        /// </summary>
        [Input("sourceVmId")]
        public Input<string>? SourceVmId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ProtectedVMState()
        {
        }
    }
}
