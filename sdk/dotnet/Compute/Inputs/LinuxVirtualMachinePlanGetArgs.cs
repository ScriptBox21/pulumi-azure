// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class LinuxVirtualMachinePlanGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the Name of the Marketplace Image this Virtual Machine should be created from. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies the Product of the Marketplace Image this Virtual Machine should be created from. Changing this forces a new resource to be created.
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        /// <summary>
        /// Specifies the Publisher of the Marketplace Image this Virtual Machine should be created from. Changing this forces a new resource to be created.
        /// </summary>
        [Input("publisher", required: true)]
        public Input<string> Publisher { get; set; } = null!;

        public LinuxVirtualMachinePlanGetArgs()
        {
        }
    }
}
