// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Version of a Shared Image within a Shared Image Gallery.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/shared_image_version.html.markdown.
        /// </summary>
        [Obsolete("Use GetSharedImageVersion.InvokeAsync() instead")]
        public static Task<GetSharedImageVersionResult> GetSharedImageVersion(GetSharedImageVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSharedImageVersionResult>("azure:compute/getSharedImageVersion:getSharedImageVersion", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetSharedImageVersion
    {
        /// <summary>
        /// Use this data source to access information about an existing Version of a Shared Image within a Shared Image Gallery.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/shared_image_version.html.markdown.
        /// </summary>
        public static Task<GetSharedImageVersionResult> InvokeAsync(GetSharedImageVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSharedImageVersionResult>("azure:compute/getSharedImageVersion:getSharedImageVersion", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetSharedImageVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Shared Image in which the Shared Image exists.
        /// </summary>
        [Input("galleryName", required: true)]
        public string GalleryName { get; set; } = null!;

        /// <summary>
        /// The name of the Shared Image in which this Version exists.
        /// </summary>
        [Input("imageName", required: true)]
        public string ImageName { get; set; } = null!;

        /// <summary>
        /// The name of the Image Version.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group in which the Shared Image Gallery exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetSharedImageVersionArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetSharedImageVersionResult
    {
        /// <summary>
        /// Is this Image Version excluded from the `latest` filter?
        /// </summary>
        public readonly bool ExcludeFromLatest;
        public readonly string GalleryName;
        public readonly string ImageName;
        /// <summary>
        /// The supported Azure location where the Shared Image Gallery exists.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The ID of the Managed Image which was the source of this Shared Image Version.
        /// </summary>
        public readonly string ManagedImageId;
        /// <summary>
        /// The Azure Region in which this Image Version exists.
        /// </summary>
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the Shared Image.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// One or more `target_region` blocks as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSharedImageVersionTargetRegionsResult> TargetRegions;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetSharedImageVersionResult(
            bool excludeFromLatest,
            string galleryName,
            string imageName,
            string location,
            string managedImageId,
            string name,
            string resourceGroupName,
            ImmutableDictionary<string, string> tags,
            ImmutableArray<Outputs.GetSharedImageVersionTargetRegionsResult> targetRegions,
            string id)
        {
            ExcludeFromLatest = excludeFromLatest;
            GalleryName = galleryName;
            ImageName = imageName;
            Location = location;
            ManagedImageId = managedImageId;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            TargetRegions = targetRegions;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetSharedImageVersionTargetRegionsResult
    {
        /// <summary>
        /// The name of the Image Version.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The number of replicas of the Image Version to be created per region.
        /// </summary>
        public readonly int RegionalReplicaCount;
        /// <summary>
        /// The storage account type for the image version.
        /// </summary>
        public readonly string StorageAccountType;

        [OutputConstructor]
        private GetSharedImageVersionTargetRegionsResult(
            string name,
            int regionalReplicaCount,
            string storageAccountType)
        {
            Name = name;
            RegionalReplicaCount = regionalReplicaCount;
            StorageAccountType = storageAccountType;
        }
    }
    }
}
