// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot
{
    public static class GetIotHub
    {
        /// <summary>
        /// Use this data source to access information about an existing IoTHub.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Azure = Pulumi.Azure;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Azure.Iot.GetIotHub.InvokeAsync(new Azure.Iot.GetIotHubArgs
        ///         {
        ///             Name = "existing",
        ///             ResourceGroupName = "existing",
        ///         }));
        ///         this.Id = example.Apply(example =&gt; example.Id);
        ///     }
        /// 
        ///     [Output("id")]
        ///     public Output&lt;string&gt; Id { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIotHubResult> InvokeAsync(GetIotHubArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIotHubResult>("azure:iot/getIotHub:getIotHub", args ?? new GetIotHubArgs(), options.WithVersion());
    }


    public sealed class GetIotHubArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of this IoTHub.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the IoTHub exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// A mapping of tags which should be assigned to the IoTHub.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetIotHubArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIotHubResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ResourceGroupName;
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetIotHubResult(
            string id,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string>? tags)
        {
            Id = id;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
