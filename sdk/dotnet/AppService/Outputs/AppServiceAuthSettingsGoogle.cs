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
    public sealed class AppServiceAuthSettingsGoogle
    {
        /// <summary>
        /// The OpenID Connect Client ID for the Google web application.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The client secret associated with the Google web application.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The OAuth 2.0 scopes that will be requested as part of Google Sign-In authentication. https://developers.google.com/identity/sign-in/web/
        /// </summary>
        public readonly ImmutableArray<string> OauthScopes;

        [OutputConstructor]
        private AppServiceAuthSettingsGoogle(
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
