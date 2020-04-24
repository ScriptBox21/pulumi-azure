// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayWafConfiguration
    {
        /// <summary>
        /// one or more `disabled_rule_group` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayWafConfigurationDisabledRuleGroup> DisabledRuleGroups;
        /// <summary>
        /// Is the Web Application Firewall be enabled?
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// one or more `exclusion` blocks as defined below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayWafConfigurationExclusion> Exclusions;
        /// <summary>
        /// The File Upload Limit in MB. Accepted values are in the range `1`MB to `500`MB. Defaults to `100`MB.
        /// </summary>
        public readonly int? FileUploadLimitMb;
        /// <summary>
        /// The Web Application Firewall Mode. Possible values are `Detection` and `Prevention`.
        /// </summary>
        public readonly string FirewallMode;
        /// <summary>
        /// The Maximum Request Body Size in KB.  Accepted values are in the range `1`KB to `128`KB.  Defaults to `128`KB.
        /// </summary>
        public readonly int? MaxRequestBodySizeKb;
        /// <summary>
        /// Is Request Body Inspection enabled?  Defaults to `true`.
        /// </summary>
        public readonly bool? RequestBodyCheck;
        /// <summary>
        /// The Type of the Rule Set used for this Web Application Firewall.
        /// </summary>
        public readonly string? RuleSetType;
        /// <summary>
        /// The Version of the Rule Set used for this Web Application Firewall. Possible values are `2.2.9`, `3.0`, and `3.1`.
        /// </summary>
        public readonly string RuleSetVersion;

        [OutputConstructor]
        private ApplicationGatewayWafConfiguration(
            ImmutableArray<Outputs.ApplicationGatewayWafConfigurationDisabledRuleGroup> disabledRuleGroups,

            bool enabled,

            ImmutableArray<Outputs.ApplicationGatewayWafConfigurationExclusion> exclusions,

            int? fileUploadLimitMb,

            string firewallMode,

            int? maxRequestBodySizeKb,

            bool? requestBodyCheck,

            string? ruleSetType,

            string ruleSetVersion)
        {
            DisabledRuleGroups = disabledRuleGroups;
            Enabled = enabled;
            Exclusions = exclusions;
            FileUploadLimitMb = fileUploadLimitMb;
            FirewallMode = firewallMode;
            MaxRequestBodySizeKb = maxRequestBodySizeKb;
            RequestBodyCheck = requestBodyCheck;
            RuleSetType = ruleSetType;
            RuleSetVersion = ruleSetVersion;
        }
    }
}