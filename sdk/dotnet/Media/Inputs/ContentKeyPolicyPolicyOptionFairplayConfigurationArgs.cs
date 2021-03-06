// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Media.Inputs
{

    public sealed class ContentKeyPolicyPolicyOptionFairplayConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key that must be used as FairPlay Application Secret key.
        /// </summary>
        [Input("ask")]
        public Input<string>? Ask { get; set; }

        /// <summary>
        /// A `offline_rental_configuration` block as defined below.
        /// </summary>
        [Input("offlineRentalConfiguration")]
        public Input<Inputs.ContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationArgs>? OfflineRentalConfiguration { get; set; }

        /// <summary>
        /// The Base64 representation of FairPlay certificate in PKCS 12 (pfx) format (including private key).
        /// </summary>
        [Input("pfx")]
        public Input<string>? Pfx { get; set; }

        /// <summary>
        /// The password encrypting FairPlay certificate in PKCS 12 (pfx) format.
        /// </summary>
        [Input("pfxPassword")]
        public Input<string>? PfxPassword { get; set; }

        /// <summary>
        /// The rental and lease key type. Supported values are `DualExpiry`, `PersistentLimited`, `PersistentUnlimited` or `Undefined`.
        /// </summary>
        [Input("rentalAndLeaseKeyType")]
        public Input<string>? RentalAndLeaseKeyType { get; set; }

        /// <summary>
        /// The rental duration. Must be greater than 0.
        /// </summary>
        [Input("rentalDurationSeconds")]
        public Input<int>? RentalDurationSeconds { get; set; }

        public ContentKeyPolicyPolicyOptionFairplayConfigurationArgs()
        {
        }
    }
}
