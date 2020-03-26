// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing Azure Data Factory (Version 2).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/data_factory.html.markdown.
        /// </summary>
        [Obsolete("Use GetFactory.InvokeAsync() instead")]
        public static Task<GetFactoryResult> GetFactory(GetFactoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFactoryResult>("azure:datafactory/getFactory:getFactory", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetFactory
    {
        /// <summary>
        /// Use this data source to access information about an existing Azure Data Factory (Version 2).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/data_factory.html.markdown.
        /// </summary>
        public static Task<GetFactoryResult> InvokeAsync(GetFactoryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFactoryResult>("azure:datafactory/getFactory:getFactory", args ?? InvokeArgs.Empty, options.WithVersion());
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
        public readonly ImmutableArray<Outputs.GetFactoryGithubConfigurationsResult> GithubConfigurations;
        /// <summary>
        /// An `identity` block as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFactoryIdentitiesResult> Identities;
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
        public readonly ImmutableArray<Outputs.GetFactoryVstsConfigurationsResult> VstsConfigurations;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFactoryResult(
            ImmutableArray<Outputs.GetFactoryGithubConfigurationsResult> githubConfigurations,
            ImmutableArray<Outputs.GetFactoryIdentitiesResult> identities,
            string location,
            string name,
            string resourceGroupName,
            ImmutableDictionary<string, string> tags,
            ImmutableArray<Outputs.GetFactoryVstsConfigurationsResult> vstsConfigurations,
            string id)
        {
            GithubConfigurations = githubConfigurations;
            Identities = identities;
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Tags = tags;
            VstsConfigurations = vstsConfigurations;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetFactoryGithubConfigurationsResult
    {
        /// <summary>
        /// The VSTS account name.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// The branch of the repository to get code from.
        /// </summary>
        public readonly string BranchName;
        /// <summary>
        /// The GitHub Enterprise host name. 
        /// </summary>
        public readonly string GitUrl;
        /// <summary>
        /// The name of the git repository.
        /// </summary>
        public readonly string RepositoryName;
        /// <summary>
        /// The root folder within the repository.
        /// </summary>
        public readonly string RootFolder;

        [OutputConstructor]
        private GetFactoryGithubConfigurationsResult(
            string accountName,
            string branchName,
            string gitUrl,
            string repositoryName,
            string rootFolder)
        {
            AccountName = accountName;
            BranchName = branchName;
            GitUrl = gitUrl;
            RepositoryName = repositoryName;
            RootFolder = rootFolder;
        }
    }

    [OutputType]
    public sealed class GetFactoryIdentitiesResult
    {
        /// <summary>
        /// The ID of the Principal (Client) in Azure Active Directory.
        /// </summary>
        public readonly string PrincipalId;
        /// <summary>
        /// The Tenant ID associated with the VSTS account.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The identity type of the Data Factory.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetFactoryIdentitiesResult(
            string principalId,
            string tenantId,
            string type)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetFactoryVstsConfigurationsResult
    {
        /// <summary>
        /// The VSTS account name.
        /// </summary>
        public readonly string AccountName;
        /// <summary>
        /// The branch of the repository to get code from.
        /// </summary>
        public readonly string BranchName;
        /// <summary>
        /// The name of the VSTS project.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The name of the git repository.
        /// </summary>
        public readonly string RepositoryName;
        /// <summary>
        /// The root folder within the repository.
        /// </summary>
        public readonly string RootFolder;
        /// <summary>
        /// The Tenant ID associated with the VSTS account.
        /// </summary>
        public readonly string TenantId;

        [OutputConstructor]
        private GetFactoryVstsConfigurationsResult(
            string accountName,
            string branchName,
            string projectName,
            string repositoryName,
            string rootFolder,
            string tenantId)
        {
            AccountName = accountName;
            BranchName = branchName;
            ProjectName = projectName;
            RepositoryName = repositoryName;
            RootFolder = rootFolder;
            TenantId = tenantId;
        }
    }
    }
}
