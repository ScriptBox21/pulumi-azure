// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Outputs
{

    [OutputType]
    public sealed class HBaseClusterComponentVersion
    {
        /// <summary>
        /// The version of HBase which should be used for this HDInsight HBase Cluster. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Hbase;

        [OutputConstructor]
        private HBaseClusterComponentVersion(string hbase)
        {
            Hbase = hbase;
        }
    }
}
