// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class WindowsVirtualMachineScaleSetWinrmListenerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Secret URL of a Key Vault Certificate, which must be specified when `protocol` is set to `Https`.
        /// </summary>
        [Input("certificateUrl")]
        public Input<string>? CertificateUrl { get; set; }

        /// <summary>
        /// The Protocol of the WinRM Listener. Possible values are `Http` and `Https`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public WindowsVirtualMachineScaleSetWinrmListenerArgs()
        {
        }
    }
}
