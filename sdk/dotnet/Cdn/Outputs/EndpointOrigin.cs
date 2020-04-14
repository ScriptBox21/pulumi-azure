// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Cdn.Outputs
{

    [OutputType]
    public sealed class EndpointOrigin
    {
        /// <summary>
        /// A string that determines the hostname/IP address of the origin server. This string can be a domain name, Storage Account endpoint, Web App endpoint, IPv4 address or IPv6 address. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The HTTP port of the origin. Defaults to `80`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? HttpPort;
        /// <summary>
        /// The HTTPS port of the origin. Defaults to `443`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? HttpsPort;
        /// <summary>
        /// The name of the origin. This is an arbitrary value. However, this value needs to be unique under the endpoint. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private EndpointOrigin(
            string hostName,

            int? httpPort,

            int? httpsPort,

            string name)
        {
            HostName = hostName;
            HttpPort = httpPort;
            HttpsPort = httpsPort;
            Name = name;
        }
    }
}
