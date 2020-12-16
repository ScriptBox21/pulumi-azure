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
    public sealed class TransformOutputBuiltinPreset
    {
        /// <summary>
        /// The built-in preset to be used for encoding videos. The allowed values are `AACGoodQualityAudio`, `AdaptiveStreaming`,`ContentAwareEncoding`, `ContentAwareEncodingExperimental`,`CopyAllBitrateNonInterleaved`, `H264MultipleBitrate1080p`,`H264MultipleBitrate720p`, `H264MultipleBitrateSD`,`H264SingleBitrate1080p`, `H264SingleBitrate720p` and `H264SingleBitrateSD`.
        /// </summary>
        public readonly string? PresetName;

        [OutputConstructor]
        private TransformOutputBuiltinPreset(string? presetName)
        {
            PresetName = presetName;
        }
    }
}
