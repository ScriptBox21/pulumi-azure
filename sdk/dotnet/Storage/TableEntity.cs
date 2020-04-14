// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage
{
    /// <summary>
    /// Manages an Entity within a Table in an Azure Storage Account.
    /// </summary>
    public partial class TableEntity : Pulumi.CustomResource
    {
        /// <summary>
        /// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
        /// </summary>
        [Output("entity")]
        public Output<ImmutableDictionary<string, string>> Entity { get; private set; } = null!;

        /// <summary>
        /// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
        /// </summary>
        [Output("partitionKey")]
        public Output<string> PartitionKey { get; private set; } = null!;

        /// <summary>
        /// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
        /// </summary>
        [Output("rowKey")]
        public Output<string> RowKey { get; private set; } = null!;

        /// <summary>
        /// Specifies the storage account in which to create the storage table entity.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Output("storageAccountName")]
        public Output<string> StorageAccountName { get; private set; } = null!;

        /// <summary>
        /// The name of the storage table in which to create the storage table entity.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Output("tableName")]
        public Output<string> TableName { get; private set; } = null!;


        /// <summary>
        /// Create a TableEntity resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TableEntity(string name, TableEntityArgs args, CustomResourceOptions? options = null)
            : base("azure:storage/tableEntity:TableEntity", name, args ?? new TableEntityArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TableEntity(string name, Input<string> id, TableEntityState? state = null, CustomResourceOptions? options = null)
            : base("azure:storage/tableEntity:TableEntity", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TableEntity resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TableEntity Get(string name, Input<string> id, TableEntityState? state = null, CustomResourceOptions? options = null)
        {
            return new TableEntity(name, id, state, options);
        }
    }

    public sealed class TableEntityArgs : Pulumi.ResourceArgs
    {
        [Input("entity", required: true)]
        private InputMap<string>? _entity;

        /// <summary>
        /// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
        /// </summary>
        public InputMap<string> Entity
        {
            get => _entity ?? (_entity = new InputMap<string>());
            set => _entity = value;
        }

        /// <summary>
        /// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
        /// </summary>
        [Input("partitionKey", required: true)]
        public Input<string> PartitionKey { get; set; } = null!;

        /// <summary>
        /// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
        /// </summary>
        [Input("rowKey", required: true)]
        public Input<string> RowKey { get; set; } = null!;

        /// <summary>
        /// Specifies the storage account in which to create the storage table entity.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountName", required: true)]
        public Input<string> StorageAccountName { get; set; } = null!;

        /// <summary>
        /// The name of the storage table in which to create the storage table entity.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("tableName", required: true)]
        public Input<string> TableName { get; set; } = null!;

        public TableEntityArgs()
        {
        }
    }

    public sealed class TableEntityState : Pulumi.ResourceArgs
    {
        [Input("entity")]
        private InputMap<string>? _entity;

        /// <summary>
        /// A map of key/value pairs that describe the entity to be inserted/merged in to the storage table.
        /// </summary>
        public InputMap<string> Entity
        {
            get => _entity ?? (_entity = new InputMap<string>());
            set => _entity = value;
        }

        /// <summary>
        /// The key for the partition where the entity will be inserted/merged. Changing this forces a new resource.
        /// </summary>
        [Input("partitionKey")]
        public Input<string>? PartitionKey { get; set; }

        /// <summary>
        /// The key for the row where the entity will be inserted/merged. Changing this forces a new resource.
        /// </summary>
        [Input("rowKey")]
        public Input<string>? RowKey { get; set; }

        /// <summary>
        /// Specifies the storage account in which to create the storage table entity.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        /// <summary>
        /// The name of the storage table in which to create the storage table entity.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        public TableEntityState()
        {
        }
    }
}
