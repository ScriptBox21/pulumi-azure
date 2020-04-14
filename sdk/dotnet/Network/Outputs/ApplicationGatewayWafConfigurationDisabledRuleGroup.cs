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
    public sealed class ApplicationGatewayWafConfigurationDisabledRuleGroup
    {
        /// <summary>
        /// The rule group where specific rules should be disabled. Accepted values are:  `crs_20_protocol_violations`, `crs_21_protocol_anomalies`, `crs_23_request_limits`, `crs_30_http_policy`, `crs_35_bad_robots`, `crs_40_generic_attacks`, `crs_41_sql_injection_attacks`, `crs_41_xss_attacks`, `crs_42_tight_security`, `crs_45_trojans`, `General`, `REQUEST-911-METHOD-ENFORCEMENT`, `REQUEST-913-SCANNER-DETECTION`, `REQUEST-920-PROTOCOL-ENFORCEMENT`, `REQUEST-921-PROTOCOL-ATTACK`, `REQUEST-930-APPLICATION-ATTACK-LFI`, `REQUEST-931-APPLICATION-ATTACK-RFI`, `REQUEST-932-APPLICATION-ATTACK-RCE`, `REQUEST-933-APPLICATION-ATTACK-PHP`, `REQUEST-941-APPLICATION-ATTACK-XSS`, `REQUEST-942-APPLICATION-ATTACK-SQLI`, `REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION`
        /// </summary>
        public readonly string RuleGroupName;
        /// <summary>
        /// A list of rules which should be disabled in that group. Disables all rules in the specified group if `rules` is not specified.
        /// </summary>
        public readonly ImmutableArray<int> Rules;

        [OutputConstructor]
        private ApplicationGatewayWafConfigurationDisabledRuleGroup(
            string ruleGroupName,

            ImmutableArray<int> rules)
        {
            RuleGroupName = ruleGroupName;
            Rules = rules;
        }
    }
}
