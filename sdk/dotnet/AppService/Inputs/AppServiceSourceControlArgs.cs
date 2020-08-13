// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class AppServiceSourceControlArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The branch of the remote repository to use. Defaults to 'master'.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// Limits to manual integration. Defaults to `false` if not specified.
        /// </summary>
        [Input("manualIntegration")]
        public Input<bool>? ManualIntegration { get; set; }

        /// <summary>
        /// The URL of the source code repository.
        /// </summary>
        [Input("repoUrl")]
        public Input<string>? RepoUrl { get; set; }

        /// <summary>
        /// Enable roll-back for the repository. Defaults to `false` if not specified.
        /// </summary>
        [Input("rollbackEnabled")]
        public Input<bool>? RollbackEnabled { get; set; }

        /// <summary>
        /// Use Mercurial if `true`, otherwise uses Git.
        /// </summary>
        [Input("useMercurial")]
        public Input<bool>? UseMercurial { get; set; }

        public AppServiceSourceControlArgs()
        {
        }
    }
}
