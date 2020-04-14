// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Backup
{
    /// <summary>
    /// Manages registration of a storage account with Azure Backup. Storage accounts must be registered with an Azure Recovery Vault in order to backup file shares within the storage account. Registering a storage account with a vault creates what is known as a protection container within Azure Recovery Services. Once the container is created, Azure file shares within the storage account can be backed up using the `azure.backup.ProtectedFileShare` resource.
    /// 
    /// &gt; **NOTE:** Azure Backup for Azure File Shares is currently in public preview. During the preview, the service is subject to additional limitations and unsupported backup scenarios. [Read More](https://docs.microsoft.com/en-us/azure/backup/backup-azure-files#limitations-for-azure-file-share-backup-during-preview)
    /// </summary>
    public partial class ContainerStorageAccount : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the vault where the storage account will be registered.
        /// </summary>
        [Output("recoveryVaultName")]
        public Output<string> RecoveryVaultName { get; private set; } = null!;

        /// <summary>
        /// Name of the resource group where the vault is located.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// Azure Resource ID of the storage account to be registered
        /// </summary>
        [Output("storageAccountId")]
        public Output<string> StorageAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerStorageAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerStorageAccount(string name, ContainerStorageAccountArgs args, CustomResourceOptions? options = null)
            : base("azure:backup/containerStorageAccount:ContainerStorageAccount", name, args ?? new ContainerStorageAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerStorageAccount(string name, Input<string> id, ContainerStorageAccountState? state = null, CustomResourceOptions? options = null)
            : base("azure:backup/containerStorageAccount:ContainerStorageAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContainerStorageAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerStorageAccount Get(string name, Input<string> id, ContainerStorageAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerStorageAccount(name, id, state, options);
        }
    }

    public sealed class ContainerStorageAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the vault where the storage account will be registered.
        /// </summary>
        [Input("recoveryVaultName", required: true)]
        public Input<string> RecoveryVaultName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group where the vault is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Azure Resource ID of the storage account to be registered
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public ContainerStorageAccountArgs()
        {
        }
    }

    public sealed class ContainerStorageAccountState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the vault where the storage account will be registered.
        /// </summary>
        [Input("recoveryVaultName")]
        public Input<string>? RecoveryVaultName { get; set; }

        /// <summary>
        /// Name of the resource group where the vault is located.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// Azure Resource ID of the storage account to be registered
        /// </summary>
        [Input("storageAccountId")]
        public Input<string>? StorageAccountId { get; set; }

        public ContainerStorageAccountState()
        {
        }
    }
}
