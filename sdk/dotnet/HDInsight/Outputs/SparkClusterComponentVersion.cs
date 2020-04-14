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
    public sealed class SparkClusterComponentVersion
    {
        /// <summary>
        /// The version of Spark which should be used for this HDInsight Spark Cluster. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Spark;

        [OutputConstructor]
        private SparkClusterComponentVersion(string spark)
        {
            Spark = spark;
        }
    }
}
