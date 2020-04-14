// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class FunctionAppAuthSettingsActiveDirectoryGetArgs : Pulumi.ResourceArgs
    {
        [Input("allowedAudiences")]
        private InputList<string>? _allowedAudiences;

        /// <summary>
        /// Allowed audience values to consider when validating JWTs issued by Azure Active Directory.
        /// </summary>
        public InputList<string> AllowedAudiences
        {
            get => _allowedAudiences ?? (_allowedAudiences = new InputList<string>());
            set => _allowedAudiences = value;
        }

        /// <summary>
        /// The Client ID of this relying party application. Enables OpenIDConnection authentication with Azure Active Directory.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// The Client Secret of this relying party application. If no secret is provided, implicit flow will be used.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        public FunctionAppAuthSettingsActiveDirectoryGetArgs()
        {
        }
    }
}
