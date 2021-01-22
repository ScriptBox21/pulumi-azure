// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Inputs
{

    public sealed class AutoscaleSettingProfileRuleMetricTriggerDimensionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the dimension.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The dimension operator. Possible values are `Equals` and `NotEquals`. `Equals` means being equal to any of the values. `NotEquals` means being not equal to any of the values.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// A list of dimension values.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public AutoscaleSettingProfileRuleMetricTriggerDimensionGetArgs()
        {
        }
    }
}
