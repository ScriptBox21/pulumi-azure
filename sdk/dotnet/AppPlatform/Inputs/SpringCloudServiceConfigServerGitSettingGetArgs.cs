// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppPlatform.Inputs
{

    public sealed class SpringCloudServiceConfigServerGitSettingGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `http_basic_auth` block as defined below.
        /// </summary>
        [Input("httpBasicAuth")]
        public Input<Inputs.SpringCloudServiceConfigServerGitSettingHttpBasicAuthGetArgs>? HttpBasicAuth { get; set; }

        /// <summary>
        /// The default label of the Git repository, should be the branch name, tag name, or commit-id of the repository.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        [Input("repositories")]
        private InputList<Inputs.SpringCloudServiceConfigServerGitSettingRepositoryGetArgs>? _repositories;

        /// <summary>
        /// One or more `repository` blocks as defined below.
        /// </summary>
        public InputList<Inputs.SpringCloudServiceConfigServerGitSettingRepositoryGetArgs> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<Inputs.SpringCloudServiceConfigServerGitSettingRepositoryGetArgs>());
            set => _repositories = value;
        }

        [Input("searchPaths")]
        private InputList<string>? _searchPaths;

        /// <summary>
        /// An array of strings used to search subdirectories of the Git repository.
        /// </summary>
        public InputList<string> SearchPaths
        {
            get => _searchPaths ?? (_searchPaths = new InputList<string>());
            set => _searchPaths = value;
        }

        /// <summary>
        /// A `ssh_auth` block as defined below.
        /// </summary>
        [Input("sshAuth")]
        public Input<Inputs.SpringCloudServiceConfigServerGitSettingSshAuthGetArgs>? SshAuth { get; set; }

        /// <summary>
        /// The URI of the default Git repository used as the Config Server back end, should be started with `http://`, `https://`, `git@`, or `ssh://`.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public SpringCloudServiceConfigServerGitSettingGetArgs()
        {
        }
    }
}
