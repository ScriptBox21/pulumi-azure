// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppPlatform.Inputs
{

    public sealed class SpringCloudServiceConfigServerGitSettingHttpBasicAuthGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The username that's used to access the Git repository server, required when the Git repository server supports Http Basic Authentication.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public SpringCloudServiceConfigServerGitSettingHttpBasicAuthGetArgs()
        {
        }
    }
}
