// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Inputs
{

    public sealed class ScaleSetNetworkProfileDnsSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServers", required: true)]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// Specifies an array of dns servers.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        public ScaleSetNetworkProfileDnsSettingsArgs()
        {
        }
    }
}
