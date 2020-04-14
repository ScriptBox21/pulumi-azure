// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class TrafficManagerProfileDnsConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The relative domain name, this is combined with the domain name used by Traffic Manager to form the FQDN which is exported as documented below. Changing this forces a new resource to be created.
        /// </summary>
        [Input("relativeName", required: true)]
        public Input<string> RelativeName { get; set; } = null!;

        /// <summary>
        /// The TTL value of the Profile used by Local DNS resolvers and clients.
        /// </summary>
        [Input("ttl", required: true)]
        public Input<int> Ttl { get; set; } = null!;

        public TrafficManagerProfileDnsConfigArgs()
        {
        }
    }
}
