// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Inputs
{

    public sealed class ProviderFeaturesVirtualMachineArgs : Pulumi.ResourceArgs
    {
        [Input("deleteOsDiskOnDeletion", required: true)]
        public Input<bool> DeleteOsDiskOnDeletion { get; set; } = null!;

        public ProviderFeaturesVirtualMachineArgs()
        {
        }
    }
}