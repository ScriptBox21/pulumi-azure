// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class ScaleSetOsProfileWindowsConfigAdditionalUnattendConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of the component to configure with the added content. The only allowable value is `Microsoft-Windows-Shell-Setup`.
        /// </summary>
        [Input("component", required: true)]
        public Input<string> Component { get; set; } = null!;

        /// <summary>
        /// Specifies the base-64 encoded XML formatted content that is added to the unattend.xml file for the specified path and component.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the pass that the content applies to. The only allowable value is `oobeSystem`.
        /// </summary>
        [Input("pass", required: true)]
        public Input<string> Pass { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the setting to which the content applies. Possible values are: `FirstLogonCommands` and `AutoLogon`.
        /// </summary>
        [Input("settingName", required: true)]
        public Input<string> SettingName { get; set; } = null!;

        public ScaleSetOsProfileWindowsConfigAdditionalUnattendConfigArgs()
        {
        }
    }
}
