// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Inputs
{

    public sealed class ContentKeyPolicyPolicyOptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable a configuration for non-DRM keys.
        /// </summary>
        [Input("clearKeyConfigurationEnabled")]
        public Input<bool>? ClearKeyConfigurationEnabled { get; set; }

        /// <summary>
        /// A `fairplay_configuration` block as defined above. Check license requirements here https://docs.microsoft.com/en-us/azure/media-services/latest/fairplay-license-overview.
        /// </summary>
        [Input("fairplayConfiguration")]
        public Input<Inputs.ContentKeyPolicyPolicyOptionFairplayConfigurationArgs>? FairplayConfiguration { get; set; }

        /// <summary>
        /// The name which should be used for this Policy Option.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Enable an open restriction. License or key will be delivered on every request.
        /// </summary>
        [Input("openRestrictionEnabled")]
        public Input<bool>? OpenRestrictionEnabled { get; set; }

        [Input("playreadyConfigurationLicenses")]
        private InputList<Inputs.ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseArgs>? _playreadyConfigurationLicenses;

        /// <summary>
        /// One or more `playready_configuration_license` blocks as defined above.
        /// </summary>
        public InputList<Inputs.ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseArgs> PlayreadyConfigurationLicenses
        {
            get => _playreadyConfigurationLicenses ?? (_playreadyConfigurationLicenses = new InputList<Inputs.ContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseArgs>());
            set => _playreadyConfigurationLicenses = value;
        }

        /// <summary>
        /// A `token_restriction` block as defined below.
        /// </summary>
        [Input("tokenRestriction")]
        public Input<Inputs.ContentKeyPolicyPolicyOptionTokenRestrictionArgs>? TokenRestriction { get; set; }

        /// <summary>
        /// The Widevine template.
        /// </summary>
        [Input("widevineConfigurationTemplate")]
        public Input<string>? WidevineConfigurationTemplate { get; set; }

        public ContentKeyPolicyPolicyOptionArgs()
        {
        }
    }
}
