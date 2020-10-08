// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MSSql
{
    /// <summary>
    /// Manages a Ms Sql Server Extended Auditing Policy.
    /// 
    /// &gt; **NOTE:** The Server Extended Auditing Policy Can be set inline here as well as with the mssql_server_extended_auditing_policy resource resource. You can only use one or the other and using both will cause a conflict.
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
    ///         var exampleServer = new Azure.MSSql.Server("exampleServer", new Azure.MSSql.ServerArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Version = "12.0",
    ///             AdministratorLogin = "missadministrator",
    ///             AdministratorLoginPassword = "AdminPassword123!",
    ///         });
    ///         var exampleAccount = new Azure.Storage.Account("exampleAccount", new Azure.Storage.AccountArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             AccountTier = "Standard",
    ///             AccountReplicationType = "LRS",
    ///         });
    ///         var exampleServerExtendedAuditingPolicy = new Azure.MSSql.ServerExtendedAuditingPolicy("exampleServerExtendedAuditingPolicy", new Azure.MSSql.ServerExtendedAuditingPolicyArgs
    ///         {
    ///             ServerId = exampleServer.Id,
    ///             StorageEndpoint = exampleAccount.PrimaryBlobEndpoint,
    ///             StorageAccountAccessKey = exampleAccount.PrimaryAccessKey,
    ///             StorageAccountAccessKeyIsSecondary = false,
    ///             RetentionInDays = 6,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ServerExtendedAuditingPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The number of days to retain logs for in the storage account.
        /// </summary>
        [Output("retentionInDays")]
        public Output<int?> RetentionInDays { get; private set; } = null!;

        /// <summary>
        /// The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// The access key to use for the auditing storage account.
        /// </summary>
        [Output("storageAccountAccessKey")]
        public Output<string?> StorageAccountAccessKey { get; private set; } = null!;

        /// <summary>
        /// Is `storage_account_access_key` value the storage's secondary key?
        /// </summary>
        [Output("storageAccountAccessKeyIsSecondary")]
        public Output<bool?> StorageAccountAccessKeyIsSecondary { get; private set; } = null!;

        /// <summary>
        /// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
        /// </summary>
        [Output("storageEndpoint")]
        public Output<string> StorageEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a ServerExtendedAuditingPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerExtendedAuditingPolicy(string name, ServerExtendedAuditingPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy", name, args ?? new ServerExtendedAuditingPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerExtendedAuditingPolicy(string name, Input<string> id, ServerExtendedAuditingPolicyState? state = null, CustomResourceOptions? options = null)
            : base("azure:mssql/serverExtendedAuditingPolicy:ServerExtendedAuditingPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerExtendedAuditingPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerExtendedAuditingPolicy Get(string name, Input<string> id, ServerExtendedAuditingPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerExtendedAuditingPolicy(name, id, state, options);
        }
    }

    public sealed class ServerExtendedAuditingPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days to retain logs for in the storage account.
        /// </summary>
        [Input("retentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        /// <summary>
        /// The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        /// <summary>
        /// The access key to use for the auditing storage account.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// Is `storage_account_access_key` value the storage's secondary key?
        /// </summary>
        [Input("storageAccountAccessKeyIsSecondary")]
        public Input<bool>? StorageAccountAccessKeyIsSecondary { get; set; }

        /// <summary>
        /// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
        /// </summary>
        [Input("storageEndpoint", required: true)]
        public Input<string> StorageEndpoint { get; set; } = null!;

        public ServerExtendedAuditingPolicyArgs()
        {
        }
    }

    public sealed class ServerExtendedAuditingPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of days to retain logs for in the storage account.
        /// </summary>
        [Input("retentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        /// <summary>
        /// The ID of the sql server to set the extended auditing policy. Changing this forces a new resource to be created.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// The access key to use for the auditing storage account.
        /// </summary>
        [Input("storageAccountAccessKey")]
        public Input<string>? StorageAccountAccessKey { get; set; }

        /// <summary>
        /// Is `storage_account_access_key` value the storage's secondary key?
        /// </summary>
        [Input("storageAccountAccessKeyIsSecondary")]
        public Input<bool>? StorageAccountAccessKeyIsSecondary { get; set; }

        /// <summary>
        /// The blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all extended auditing logs.
        /// </summary>
        [Input("storageEndpoint")]
        public Input<string>? StorageEndpoint { get; set; }

        public ServerExtendedAuditingPolicyState()
        {
        }
    }
}