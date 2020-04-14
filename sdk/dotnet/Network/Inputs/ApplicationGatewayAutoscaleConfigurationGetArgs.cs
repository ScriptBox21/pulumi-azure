// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayAutoscaleConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum capacity for autoscaling. Accepted values are in the range `2` to `125`.
        /// </summary>
        [Input("maxCapacity")]
        public Input<int>? MaxCapacity { get; set; }

        /// <summary>
        /// Minimum capacity for autoscaling. Accepted values are in the range `0` to `100`.
        /// </summary>
        [Input("minCapacity", required: true)]
        public Input<int> MinCapacity { get; set; } = null!;

        public ApplicationGatewayAutoscaleConfigurationGetArgs()
        {
        }
    }
}
