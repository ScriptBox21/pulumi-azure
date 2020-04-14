// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class WindowsVirtualMachineScaleSetAdditionalCapabilitiesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should the capacity to enable Data Disks of the `UltraSSD_LRS` storage account type be supported on this Virtual Machine Scale Set? Defaults to `false`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("ultraSsdEnabled")]
        public Input<bool>? UltraSsdEnabled { get; set; }

        public WindowsVirtualMachineScaleSetAdditionalCapabilitiesGetArgs()
        {
        }
    }
}
