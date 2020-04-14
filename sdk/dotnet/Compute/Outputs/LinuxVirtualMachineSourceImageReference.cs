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
    public sealed class LinuxVirtualMachineSourceImageReference
    {
        public readonly string Offer;
        /// <summary>
        /// Specifies the Publisher of the Marketplace Image this Virtual Machine should be created from. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Publisher;
        public readonly string Sku;
        public readonly string Version;

        [OutputConstructor]
        private LinuxVirtualMachineSourceImageReference(
            string offer,

            string publisher,

            string sku,

            string version)
        {
            Offer = offer;
            Publisher = publisher;
            Sku = sku;
            Version = version;
        }
    }
}
