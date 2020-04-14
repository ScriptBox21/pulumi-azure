// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Storage.Outputs
{

    [OutputType]
    public sealed class ShareAclAccessPolicy
    {
        /// <summary>
        /// The ISO8061 UTC time at which this Access Policy should be valid until.
        /// </summary>
        public readonly string Expiry;
        /// <summary>
        /// The permissions which should associated with this Shared Identifier.
        /// </summary>
        public readonly string Permissions;
        /// <summary>
        /// The ISO8061 UTC time at which this Access Policy should be valid from.
        /// </summary>
        public readonly string Start;

        [OutputConstructor]
        private ShareAclAccessPolicy(
            string expiry,

            string permissions,

            string start)
        {
            Expiry = expiry;
            Permissions = permissions;
            Start = start;
        }
    }
}
