// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ProfileContainerNetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        [Input("ipConfigurations", required: true)]
        private InputList<Inputs.ProfileContainerNetworkInterfaceIpConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// One or more `ip_configuration` blocks as documented below.
        /// </summary>
        public InputList<Inputs.ProfileContainerNetworkInterfaceIpConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.ProfileContainerNetworkInterfaceIpConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// Specifies the name of the IP Configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ProfileContainerNetworkInterfaceArgs()
        {
        }
    }
}
