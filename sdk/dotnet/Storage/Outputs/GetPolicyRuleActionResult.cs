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
    public sealed class GetPolicyRuleActionResult
    {
        /// <summary>
        /// A `base_blob` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPolicyRuleActionBaseBlobResult> BaseBlobs;
        /// <summary>
        /// A `snapshot` block as documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetPolicyRuleActionSnapshotResult> Snapshots;

        [OutputConstructor]
        private GetPolicyRuleActionResult(
            ImmutableArray<Outputs.GetPolicyRuleActionBaseBlobResult> baseBlobs,

            ImmutableArray<Outputs.GetPolicyRuleActionSnapshotResult> snapshots)
        {
            BaseBlobs = baseBlobs;
            Snapshots = snapshots;
        }
    }
}
