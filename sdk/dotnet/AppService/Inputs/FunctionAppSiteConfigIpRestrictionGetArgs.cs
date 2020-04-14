// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Inputs
{

    public sealed class FunctionAppSiteConfigIpRestrictionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP Address CIDR notation used for this IP Restriction.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The Subnet ID used for this IP Restriction.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public FunctionAppSiteConfigIpRestrictionGetArgs()
        {
        }
    }
}
