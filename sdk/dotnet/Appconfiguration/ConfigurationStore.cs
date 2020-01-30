// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppConfiguration
{
    /// <summary>
    /// Manages an Azure App Configuration.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/app_configuration.html.markdown.
    /// </summary>
    public partial class ConfigurationStore : Pulumi.CustomResource
    {
        /// <summary>
        /// The URL of the App Configuration.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the App Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An `access_key` block as defined below containing the primary read access key.
        /// </summary>
        [Output("primaryReadKeys")]
        public Output<ImmutableArray<Outputs.ConfigurationStorePrimaryReadKeys>> PrimaryReadKeys { get; private set; } = null!;

        /// <summary>
        /// An `access_key` block as defined below containing the primary write access key.
        /// </summary>
        [Output("primaryWriteKeys")]
        public Output<ImmutableArray<Outputs.ConfigurationStorePrimaryWriteKeys>> PrimaryWriteKeys { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// An `access_key` block as defined below containing the secondary read access key.
        /// </summary>
        [Output("secondaryReadKeys")]
        public Output<ImmutableArray<Outputs.ConfigurationStoreSecondaryReadKeys>> SecondaryReadKeys { get; private set; } = null!;

        /// <summary>
        /// An `access_key` block as defined below containing the secondary write access key.
        /// </summary>
        [Output("secondaryWriteKeys")]
        public Output<ImmutableArray<Outputs.ConfigurationStoreSecondaryWriteKeys>> SecondaryWriteKeys { get; private set; } = null!;

        /// <summary>
        /// The SKU name of the the App Configuration. Possible values are `free` and `standard`.
        /// </summary>
        [Output("sku")]
        public Output<string?> Sku { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationStore(string name, ConfigurationStoreArgs args, CustomResourceOptions? options = null)
            : base("azure:appconfiguration/configurationStore:ConfigurationStore", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationStore(string name, Input<string> id, ConfigurationStoreState? state = null, CustomResourceOptions? options = null)
            : base("azure:appconfiguration/configurationStore:ConfigurationStore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationStore Get(string name, Input<string> id, ConfigurationStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigurationStore(name, id, state, options);
        }
    }

    public sealed class ConfigurationStoreArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the App Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The SKU name of the the App Configuration. Possible values are `free` and `standard`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

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

        public ConfigurationStoreArgs()
        {
        }
    }

    public sealed class ConfigurationStoreState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the App Configuration.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Specifies the name of the App Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("primaryReadKeys")]
        private InputList<Inputs.ConfigurationStorePrimaryReadKeysGetArgs>? _primaryReadKeys;

        /// <summary>
        /// An `access_key` block as defined below containing the primary read access key.
        /// </summary>
        public InputList<Inputs.ConfigurationStorePrimaryReadKeysGetArgs> PrimaryReadKeys
        {
            get => _primaryReadKeys ?? (_primaryReadKeys = new InputList<Inputs.ConfigurationStorePrimaryReadKeysGetArgs>());
            set => _primaryReadKeys = value;
        }

        [Input("primaryWriteKeys")]
        private InputList<Inputs.ConfigurationStorePrimaryWriteKeysGetArgs>? _primaryWriteKeys;

        /// <summary>
        /// An `access_key` block as defined below containing the primary write access key.
        /// </summary>
        public InputList<Inputs.ConfigurationStorePrimaryWriteKeysGetArgs> PrimaryWriteKeys
        {
            get => _primaryWriteKeys ?? (_primaryWriteKeys = new InputList<Inputs.ConfigurationStorePrimaryWriteKeysGetArgs>());
            set => _primaryWriteKeys = value;
        }

        /// <summary>
        /// The name of the resource group in which to create the App Configuration. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        [Input("secondaryReadKeys")]
        private InputList<Inputs.ConfigurationStoreSecondaryReadKeysGetArgs>? _secondaryReadKeys;

        /// <summary>
        /// An `access_key` block as defined below containing the secondary read access key.
        /// </summary>
        public InputList<Inputs.ConfigurationStoreSecondaryReadKeysGetArgs> SecondaryReadKeys
        {
            get => _secondaryReadKeys ?? (_secondaryReadKeys = new InputList<Inputs.ConfigurationStoreSecondaryReadKeysGetArgs>());
            set => _secondaryReadKeys = value;
        }

        [Input("secondaryWriteKeys")]
        private InputList<Inputs.ConfigurationStoreSecondaryWriteKeysGetArgs>? _secondaryWriteKeys;

        /// <summary>
        /// An `access_key` block as defined below containing the secondary write access key.
        /// </summary>
        public InputList<Inputs.ConfigurationStoreSecondaryWriteKeysGetArgs> SecondaryWriteKeys
        {
            get => _secondaryWriteKeys ?? (_secondaryWriteKeys = new InputList<Inputs.ConfigurationStoreSecondaryWriteKeysGetArgs>());
            set => _secondaryWriteKeys = value;
        }

        /// <summary>
        /// The SKU name of the the App Configuration. Possible values are `free` and `standard`.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

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

        public ConfigurationStoreState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ConfigurationStorePrimaryReadKeysGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection string including the endpoint, id and secret.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The ID of the access key.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The secret of the access key.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        public ConfigurationStorePrimaryReadKeysGetArgs()
        {
        }
    }

    public sealed class ConfigurationStorePrimaryWriteKeysGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection string including the endpoint, id and secret.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The ID of the access key.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The secret of the access key.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        public ConfigurationStorePrimaryWriteKeysGetArgs()
        {
        }
    }

    public sealed class ConfigurationStoreSecondaryReadKeysGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection string including the endpoint, id and secret.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The ID of the access key.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The secret of the access key.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        public ConfigurationStoreSecondaryReadKeysGetArgs()
        {
        }
    }

    public sealed class ConfigurationStoreSecondaryWriteKeysGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The connection string including the endpoint, id and secret.
        /// </summary>
        [Input("connectionString")]
        public Input<string>? ConnectionString { get; set; }

        /// <summary>
        /// The ID of the access key.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The secret of the access key.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        public ConfigurationStoreSecondaryWriteKeysGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ConfigurationStorePrimaryReadKeys
    {
        /// <summary>
        /// The connection string including the endpoint, id and secret.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// The ID of the access key.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The secret of the access key.
        /// </summary>
        public readonly string Secret;

        [OutputConstructor]
        private ConfigurationStorePrimaryReadKeys(
            string connectionString,
            string id,
            string secret)
        {
            ConnectionString = connectionString;
            Id = id;
            Secret = secret;
        }
    }

    [OutputType]
    public sealed class ConfigurationStorePrimaryWriteKeys
    {
        /// <summary>
        /// The connection string including the endpoint, id and secret.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// The ID of the access key.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The secret of the access key.
        /// </summary>
        public readonly string Secret;

        [OutputConstructor]
        private ConfigurationStorePrimaryWriteKeys(
            string connectionString,
            string id,
            string secret)
        {
            ConnectionString = connectionString;
            Id = id;
            Secret = secret;
        }
    }

    [OutputType]
    public sealed class ConfigurationStoreSecondaryReadKeys
    {
        /// <summary>
        /// The connection string including the endpoint, id and secret.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// The ID of the access key.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The secret of the access key.
        /// </summary>
        public readonly string Secret;

        [OutputConstructor]
        private ConfigurationStoreSecondaryReadKeys(
            string connectionString,
            string id,
            string secret)
        {
            ConnectionString = connectionString;
            Id = id;
            Secret = secret;
        }
    }

    [OutputType]
    public sealed class ConfigurationStoreSecondaryWriteKeys
    {
        /// <summary>
        /// The connection string including the endpoint, id and secret.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// The ID of the access key.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The secret of the access key.
        /// </summary>
        public readonly string Secret;

        [OutputConstructor]
        private ConfigurationStoreSecondaryWriteKeys(
            string connectionString,
            string id,
            string secret)
        {
            ConnectionString = connectionString;
            Id = id;
            Secret = secret;
        }
    }
    }
}
