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
    public sealed class SlotAuthSettingsFacebook
    {
        /// <summary>
        /// The App ID of the Facebook app used for login
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// The App Secret of the Facebook app used for Facebook Login.
        /// </summary>
        public readonly string AppSecret;
        /// <summary>
        /// The OAuth 2.0 scopes that will be requested as part of Facebook Login authentication. https://developers.facebook.com/docs/facebook-login
        /// </summary>
        public readonly ImmutableArray<string> OauthScopes;

        [OutputConstructor]
        private SlotAuthSettingsFacebook(
            string appId,

            string appSecret,

            ImmutableArray<string> oauthScopes)
        {
            AppId = appId;
            AppSecret = appSecret;
            OauthScopes = oauthScopes;
        }
    }
}
