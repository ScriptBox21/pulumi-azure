// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.KeyVault.Inputs
{

    public sealed class CertificateIssuerAdminArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// E-mail address of the admin.
        /// </summary>
        [Input("emailAddress", required: true)]
        public Input<string> EmailAddress { get; set; } = null!;

        /// <summary>
        /// First name of the admin.
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// Last name of the admin.
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        /// <summary>
        /// Phone number of the admin.
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        public CertificateIssuerAdminArgs()
        {
        }
    }
}
