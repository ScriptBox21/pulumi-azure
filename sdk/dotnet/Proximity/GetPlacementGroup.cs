// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Proximity
{
    public static class GetPlacementGroup
    {
        /// <summary>
        /// Use this data source to access information about an existing Proximity Placement Group.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPlacementGroupResult> InvokeAsync(GetPlacementGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPlacementGroupResult>("azure:proximity/getPlacementGroup:getPlacementGroup", args ?? new GetPlacementGroupArgs(), options.WithVersion());
    }


    public sealed class GetPlacementGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Proximity Placement Group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group in which the Proximity Placement Group exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPlacementGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPlacementGroupResult
    {
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetPlacementGroupResult(
            string id,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags)
        {
            Id = id;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
        }
    }
}
