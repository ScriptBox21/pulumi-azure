// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class WindowsVirtualMachineAdditionalUnattendContentGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The XML formatted content that is added to the unattend.xml file for the specified path and component. Changing this forces a new resource to be created.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The name of the setting to which the content applies. Possible values are `AutoLogon` and `FirstLogonCommands`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("setting", required: true)]
        public Input<string> Setting { get; set; } = null!;

        public WindowsVirtualMachineAdditionalUnattendContentGetArgs()
        {
        }
    }
}
