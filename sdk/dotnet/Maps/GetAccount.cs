// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Maps
{
    public static class GetAccount
    {
        /// <summary>
        /// Use this data source to access information about an existing Azure Maps Account.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azure:maps/getAccount:getAccount", args ?? new GetAccountArgs(), options.WithVersion());
    }


    public sealed class GetAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Maps Account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Resource Group in which the Maps Account is located.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetAccountArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// The primary key used to authenticate and authorize access to the Maps REST APIs.
        /// </summary>
        public readonly string PrimaryAccessKey;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The primary key used to authenticate and authorize access to the Maps REST APIs. The second key is given to provide seamless key regeneration.
        /// </summary>
        public readonly string SecondaryAccessKey;
        /// <summary>
        /// The sku of the Azure Maps Account.
        /// </summary>
        public readonly string SkuName;
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// A unique identifier for the Maps Account.
        /// </summary>
        public readonly string XMsClientId;

        [OutputConstructor]
        private GetAccountResult(
            string id,

            string name,

            string primaryAccessKey,

            string resourceGroupName,

            string secondaryAccessKey,

            string skuName,

            ImmutableDictionary<string, string>? tags,

            string xMsClientId)
        {
            Id = id;
            Name = name;
            PrimaryAccessKey = primaryAccessKey;
            ResourceGroupName = resourceGroupName;
            SecondaryAccessKey = secondaryAccessKey;
            SkuName = skuName;
            Tags = tags;
            XMsClientId = xMsClientId;
        }
    }
}
