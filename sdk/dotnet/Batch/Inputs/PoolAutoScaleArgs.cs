// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch.Inputs
{

    public sealed class PoolAutoScaleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The interval to wait before evaluating if the pool needs to be scaled. Defaults to `PT15M`.
        /// </summary>
        [Input("evaluationInterval")]
        public Input<string>? EvaluationInterval { get; set; }

        /// <summary>
        /// The autoscale formula that needs to be used for scaling the Batch pool.
        /// </summary>
        [Input("formula", required: true)]
        public Input<string> Formula { get; set; } = null!;

        public PoolAutoScaleArgs()
        {
        }
    }
}
