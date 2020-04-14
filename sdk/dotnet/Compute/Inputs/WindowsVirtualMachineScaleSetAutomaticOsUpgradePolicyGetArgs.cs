// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should automatic rollbacks be disabled? Changing this forces a new resource to be created.
        /// </summary>
        [Input("disableAutomaticRollback", required: true)]
        public Input<bool> DisableAutomaticRollback { get; set; } = null!;

        /// <summary>
        /// Should OS Upgrades automatically be applied to Scale Set instances in a rolling fashion when a newer version of the OS Image becomes available? Changing this forces a new resource to be created.
        /// </summary>
        [Input("enableAutomaticOsUpgrade", required: true)]
        public Input<bool> EnableAutomaticOsUpgrade { get; set; } = null!;

        public WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyGetArgs()
        {
        }
    }
}
