// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Inputs
{

    public sealed class ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures Automatic Gain Control (AGC) and Color Stripe in the license. Must be between 0 and 3 inclusive.
        /// </summary>
        [Input("agcAndColorStripeRestriction")]
        public Input<int>? AgcAndColorStripeRestriction { get; set; }

        /// <summary>
        /// Configures Unknown output handling settings of the license. Supported values are `Allowed`, `AllowedWithVideoConstriction` or `NotAllowed`.
        /// </summary>
        [Input("allowPassingVideoContentToUnknownOutput")]
        public Input<string>? AllowPassingVideoContentToUnknownOutput { get; set; }

        /// <summary>
        /// Specifies the output protection level for compressed digital audio. Supported values are 100, 150 or 200.
        /// </summary>
        [Input("analogVideoOpl")]
        public Input<int>? AnalogVideoOpl { get; set; }

        /// <summary>
        /// Specifies the output protection level for compressed digital audio.Supported values are 100, 150 or 200.
        /// </summary>
        [Input("compressedDigitalAudioOpl")]
        public Input<int>? CompressedDigitalAudioOpl { get; set; }

        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        [Input("digitalVideoOnlyContentRestriction")]
        public Input<bool>? DigitalVideoOnlyContentRestriction { get; set; }

        /// <summary>
        /// The amount of time that the license is valid after the license is first used to play content.
        /// </summary>
        [Input("firstPlayExpiration")]
        public Input<string>? FirstPlayExpiration { get; set; }

        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        [Input("imageConstraintForAnalogComponentVideoRestriction")]
        public Input<bool>? ImageConstraintForAnalogComponentVideoRestriction { get; set; }

        /// <summary>
        /// Enables the Image Constraint For Analog Component Video Restriction in the license.
        /// </summary>
        [Input("imageConstraintForAnalogComputerMonitorRestriction")]
        public Input<bool>? ImageConstraintForAnalogComputerMonitorRestriction { get; set; }

        /// <summary>
        /// Configures the Serial Copy Management System (SCMS) in the license. Must be between 0 and 3 inclusive.
        /// </summary>
        [Input("scmsRestriction")]
        public Input<int>? ScmsRestriction { get; set; }

        /// <summary>
        /// Specifies the output protection level for uncompressed digital audio. Supported values are 100, 150, 250 or 300.
        /// </summary>
        [Input("uncompressedDigitalAudioOpl")]
        public Input<int>? UncompressedDigitalAudioOpl { get; set; }

        /// <summary>
        /// Specifies the output protection level for uncompressed digital video. Supported values are 100, 150, 250 or 300.
        /// </summary>
        [Input("uncompressedDigitalVideoOpl")]
        public Input<int>? UncompressedDigitalVideoOpl { get; set; }

        public ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightGetArgs()
        {
        }
    }
}
