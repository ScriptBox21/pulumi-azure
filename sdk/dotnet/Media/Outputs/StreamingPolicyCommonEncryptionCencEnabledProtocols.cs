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
    public sealed class StreamingPolicyCommonEncryptionCencEnabledProtocols
    {
        /// <summary>
        /// Enable DASH protocol or not. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        public readonly bool? Dash;
        /// <summary>
        /// Enable Download protocol or not. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        public readonly bool? Download;
        /// <summary>
        /// Enable HLS protocol or not. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        public readonly bool? Hls;
        /// <summary>
        /// Enable SmoothStreaming protocol or not. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        public readonly bool? SmoothStreaming;

        [OutputConstructor]
        private StreamingPolicyCommonEncryptionCencEnabledProtocols(
            bool? dash,

            bool? download,

            bool? hls,

            bool? smoothStreaming)
        {
            Dash = dash;
            Download = download;
            Hls = hls;
            SmoothStreaming = smoothStreaming;
        }
    }
}
