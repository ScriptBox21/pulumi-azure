// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class ScaleSetOsProfileWindowsConfigWinrm
    {
        /// <summary>
        /// Specifies URL of the certificate with which new Virtual Machines is provisioned.
        /// </summary>
        public readonly string? CertificateUrl;
        /// <summary>
        /// Specifies the protocol of listener
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private ScaleSetOsProfileWindowsConfigWinrm(
            string? certificateUrl,

            string protocol)
        {
            CertificateUrl = certificateUrl;
            Protocol = protocol;
        }
    }
}
