// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NotificationHub.Outputs
{

    [OutputType]
    public sealed class GetHubGcmCredentialResult
    {
        /// <summary>
        /// The API Key associated with the Google Cloud Messaging service.
        /// </summary>
        public readonly string ApiKey;

        [OutputConstructor]
        private GetHubGcmCredentialResult(string apiKey)
        {
            ApiKey = apiKey;
        }
    }
}
