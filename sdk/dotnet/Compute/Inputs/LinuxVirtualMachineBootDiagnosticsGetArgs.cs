// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class LinuxVirtualMachineBootDiagnosticsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Primary/Secondary Endpoint for the Azure Storage Account which should be used to store Boot Diagnostics, including Console Output and Screenshots from the Hypervisor.
        /// </summary>
        [Input("storageAccountUri", required: true)]
        public Input<string> StorageAccountUri { get; set; } = null!;

        public LinuxVirtualMachineBootDiagnosticsGetArgs()
        {
        }
    }
}
