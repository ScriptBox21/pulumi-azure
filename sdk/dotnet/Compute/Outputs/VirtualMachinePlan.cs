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
    public sealed class VirtualMachinePlan
    {
        /// <summary>
        /// Specifies the name of the image from the marketplace.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the product of the image from the marketplace.
        /// </summary>
        public readonly string Product;
        /// <summary>
        /// Specifies the publisher of the image.
        /// </summary>
        public readonly string Publisher;

        [OutputConstructor]
        private VirtualMachinePlan(
            string name,

            string product,

            string publisher)
        {
            Name = name;
            Product = product;
            Publisher = publisher;
        }
    }
}
