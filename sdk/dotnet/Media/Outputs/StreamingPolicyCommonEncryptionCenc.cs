// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Outputs
{

    [OutputType]
    public sealed class StreamingPolicyCommonEncryptionCenc
    {
        /// <summary>
        /// A `default_content_key` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        public readonly Outputs.StreamingPolicyCommonEncryptionCencDefaultContentKey? DefaultContentKey;
        /// <summary>
        /// A `drm_playready` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        public readonly Outputs.StreamingPolicyCommonEncryptionCencDrmPlayready? DrmPlayready;
        /// <summary>
        /// Template for the URL of the custom service delivering licenses to end user players. Not required when using Azure Media Services for issuing licenses. The template supports replaceable tokens that the service will update at runtime with the value specific to the request. The currently supported token values are `{AlternativeMediaId}`, which is replaced with the value of `StreamingLocatorId.AlternativeMediaId`, and `{ContentKeyId}`, which is replaced with the value of identifier of the key being requested. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        public readonly string? DrmWidevineCustomLicenseAcquisitionUrlTemplate;
        /// <summary>
        /// A `enabled_protocols` block as defined below. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        public readonly Outputs.StreamingPolicyCommonEncryptionCencEnabledProtocols? EnabledProtocols;

        [OutputConstructor]
        private StreamingPolicyCommonEncryptionCenc(
            Outputs.StreamingPolicyCommonEncryptionCencDefaultContentKey? defaultContentKey,

            Outputs.StreamingPolicyCommonEncryptionCencDrmPlayready? drmPlayready,

            string? drmWidevineCustomLicenseAcquisitionUrlTemplate,

            Outputs.StreamingPolicyCommonEncryptionCencEnabledProtocols? enabledProtocols)
        {
            DefaultContentKey = defaultContentKey;
            DrmPlayready = drmPlayready;
            DrmWidevineCustomLicenseAcquisitionUrlTemplate = drmWidevineCustomLicenseAcquisitionUrlTemplate;
            EnabledProtocols = enabledProtocols;
        }
    }
}
