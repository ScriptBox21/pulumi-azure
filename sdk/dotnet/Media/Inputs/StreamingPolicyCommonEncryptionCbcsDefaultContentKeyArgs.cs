// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Inputs
{

    public sealed class StreamingPolicyCommonEncryptionCbcsDefaultContentKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Label can be used to specify Content Key when creating a Streaming Locator. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Policy used by Default Key. Changing this forces a new Streaming Policy to be created.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        public StreamingPolicyCommonEncryptionCbcsDefaultContentKeyArgs()
        {
        }
    }
}
