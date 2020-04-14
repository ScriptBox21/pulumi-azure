// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayTrustedRootCertificate
    {
        /// <summary>
        /// The contents of the Trusted Root Certificate which should be used.
        /// </summary>
        public readonly string Data;
        /// <summary>
        /// The ID of the Rewrite Rule Set
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The Name of the Trusted Root Certificate to use.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ApplicationGatewayTrustedRootCertificate(
            string data,

            string? id,

            string name)
        {
            Data = data;
            Id = id;
            Name = name;
        }
    }
}
