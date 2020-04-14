// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory
{
    public static class GetFactory
    {
        /// <summary>
        /// Use this data source to access information about an existing Azure Data Factory (Version 2).
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFactoryResult> InvokeAsync(GetFactoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFactoryResult>("azure:datafactory/getFactory:getFactory", args ?? new GetFactoryArgs(), options.WithVersion());
    }


    public sealed class GetFactoryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the Data Factory to retrieve information about. 
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the Data Factory exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetFactoryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFactoryResult
    {
        /// <summary>
        /// A `github_configuration` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFactoryGithubConfigurationResult> GithubConfigurations;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFactoryIdentityResult> Identities;
        /// <summary>
        /// The Azure location where the resource exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// A mapping of tags assigned to the resource.
        /// ---
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// A `vsts_configuration` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFactoryVstsConfigurationResult> VstsConfigurations;

        [OutputConstructor]
        private GetFactoryResult(
            ImmutableArray<Outputs.GetFactoryGithubConfigurationResult> githubConfigurations,

            string id,

            ImmutableArray<Outputs.GetFactoryIdentityResult> identities,

            string location,

            string name,

            string resourceGroupName,

            ImmutableDictionary<string, string> tags,

            ImmutableArray<Outputs.GetFactoryVstsConfigurationResult> vstsConfigurations)
        {
            GithubConfigurations = githubConfigurations;
            Id = id;
            Identities = identities;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            VstsConfigurations = vstsConfigurations;
        }
    }
}
