// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Inputs
{

    public sealed class StormClusterComponentVersionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The version of Storm which should be used for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
        /// </summary>
        [Input("storm", required: true)]
        public Input<string> Storm { get; set; } = null!;

        public StormClusterComponentVersionGetArgs()
        {
        }
    }
}
