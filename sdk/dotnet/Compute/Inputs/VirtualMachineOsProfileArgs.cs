// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class VirtualMachineOsProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password associated with the local administrator account.
        /// </summary>
        [Input("adminPassword")]
        public Input<string>? AdminPassword { get; set; }

        /// <summary>
        /// Specifies the name of the local administrator account.
        /// </summary>
        [Input("adminUsername", required: true)]
        public Input<string> AdminUsername { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Virtual Machine.
        /// </summary>
        [Input("computerName", required: true)]
        public Input<string> ComputerName { get; set; } = null!;

        /// <summary>
        /// Specifies custom data to supply to the machine. On Linux-based systems, this can be used as a cloud-init script. On other systems, this will be copied as a file on disk. Internally, this provider will base64 encode this value before sending it to the API. The maximum length of the binary array is 65535 bytes.
        /// </summary>
        [Input("customData")]
        public Input<string>? CustomData { get; set; }

        public VirtualMachineOsProfileArgs()
        {
        }
    }
}
