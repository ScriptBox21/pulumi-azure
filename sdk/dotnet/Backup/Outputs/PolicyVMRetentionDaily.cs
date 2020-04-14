// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Backup.Outputs
{

    [OutputType]
    public sealed class PolicyVMRetentionDaily
    {
        /// <summary>
        /// The number of yearly backups to keep. Must be between `1` and `9999`
        /// </summary>
        public readonly int Count;

        [OutputConstructor]
        private PolicyVMRetentionDaily(int count)
        {
            Count = count;
        }
    }
}
