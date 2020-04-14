// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class WindowsVirtualMachineScaleSetPlanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Windows Virtual Machine Scale Set. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        [Input("publisher", required: true)]
        public Input<string> Publisher { get; set; } = null!;

        public WindowsVirtualMachineScaleSetPlanArgs()
        {
        }
    }
}
