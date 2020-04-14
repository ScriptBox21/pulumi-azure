// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch.Outputs
{

    [OutputType]
    public sealed class PoolStartTaskUserIdentity
    {
        /// <summary>
        /// A `auto_user` block that describes the user identity under which the start task runs.
        /// </summary>
        public readonly Outputs.PoolStartTaskUserIdentityAutoUser? AutoUser;
        /// <summary>
        /// The username to be used by the Batch pool start task.
        /// </summary>
        public readonly string? UserName;

        [OutputConstructor]
        private PoolStartTaskUserIdentity(
            Outputs.PoolStartTaskUserIdentityAutoUser? autoUser,

            string? userName)
        {
            AutoUser = autoUser;
            UserName = userName;
        }
    }
}
