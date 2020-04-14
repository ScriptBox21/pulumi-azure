// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayWafConfigurationGetArgs : Pulumi.ResourceArgs
    {
        [Input("disabledRuleGroups")]
        private InputList<Inputs.ApplicationGatewayWafConfigurationDisabledRuleGroupGetArgs>? _disabledRuleGroups;

        /// <summary>
        /// one or more `disabled_rule_group` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayWafConfigurationDisabledRuleGroupGetArgs> DisabledRuleGroups
        {
            get => _disabledRuleGroups ?? (_disabledRuleGroups = new InputList<Inputs.ApplicationGatewayWafConfigurationDisabledRuleGroupGetArgs>());
            set => _disabledRuleGroups = value;
        }

        /// <summary>
        /// Is the Web Application Firewall be enabled?
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("exclusions")]
        private InputList<Inputs.ApplicationGatewayWafConfigurationExclusionGetArgs>? _exclusions;

        /// <summary>
        /// one or more `exclusion` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayWafConfigurationExclusionGetArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.ApplicationGatewayWafConfigurationExclusionGetArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// The File Upload Limit in MB. Accepted values are in the range `1`MB to `500`MB. Defaults to `100`MB.
        /// </summary>
        [Input("fileUploadLimitMb")]
        public Input<int>? FileUploadLimitMb { get; set; }

        /// <summary>
        /// The Web Application Firewall Mode. Possible values are `Detection` and `Prevention`.
        /// </summary>
        [Input("firewallMode", required: true)]
        public Input<string> FirewallMode { get; set; } = null!;

        /// <summary>
        /// The Maximum Request Body Size in KB.  Accepted values are in the range `1`KB to `128`KB.  Defaults to `128`KB.
        /// </summary>
        [Input("maxRequestBodySizeKb")]
        public Input<int>? MaxRequestBodySizeKb { get; set; }

        /// <summary>
        /// Is Request Body Inspection enabled?  Defaults to `true`.
        /// </summary>
        [Input("requestBodyCheck")]
        public Input<bool>? RequestBodyCheck { get; set; }

        /// <summary>
        /// The Type of the Rule Set used for this Web Application Firewall.
        /// </summary>
        [Input("ruleSetType")]
        public Input<string>? RuleSetType { get; set; }

        /// <summary>
        /// The Version of the Rule Set used for this Web Application Firewall. Possible values are `2.2.9`, `3.0`, and `3.1`.
        /// </summary>
        [Input("ruleSetVersion", required: true)]
        public Input<string> RuleSetVersion { get; set; } = null!;

        public ApplicationGatewayWafConfigurationGetArgs()
        {
        }
    }
}
