// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class FunctionAppAuthSettingsMicrosoft
    {
        /// <summary>
        /// The OAuth 2.0 client ID that was created for the app used for authentication.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The OAuth 2.0 client secret that was created for the app used for authentication.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication. https://msdn.microsoft.com/en-us/library/dn631845.aspx
        /// </summary>
        public readonly ImmutableArray<string> OauthScopes;

        [OutputConstructor]
        private FunctionAppAuthSettingsMicrosoft(
            string clientId,

            string clientSecret,

            ImmutableArray<string> oauthScopes)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            OauthScopes = oauthScopes;
        }
    }
}
