// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot.Outputs
{

    [OutputType]
    public sealed class SecuritySolutionRecommendationsEnabled
    {
        /// <summary>
        /// Is Principal Authentication enabled for the ACR repository? Defaults to `true`.
        /// </summary>
        public readonly bool? AcrAuthentication;
        /// <summary>
        /// Is Agent send underutilized messages enabled? Defaults to `true`.
        /// </summary>
        public readonly bool? AgentSendUnutilizedMsg;
        /// <summary>
        /// Is Security related system configuration issues identified? Defaults to `true`.
        /// </summary>
        public readonly bool? Baseline;
        /// <summary>
        /// Is IoT Edge Hub memory optimized? Defaults to `true`.
        /// </summary>
        public readonly bool? EdgeHubMemOptimize;
        /// <summary>
        /// Is logging configured for IoT Edge module? Defaults to `true`.
        /// </summary>
        public readonly bool? EdgeLoggingOption;
        /// <summary>
        /// Is inconsistent module settings enabled for SecurityGroup? Defaults to `true`.
        /// </summary>
        public readonly bool? InconsistentModuleSettings;
        /// <summary>
        /// is Azure IoT Security agent installed? Defaults to `true`.
        /// </summary>
        public readonly bool? InstallAgent;
        /// <summary>
        /// Is Default IP filter policy denied? Defaults to `true`.
        /// </summary>
        public readonly bool? IpFilterDenyAll;
        /// <summary>
        /// Is IP filter rule source allowable IP range too large? Defaults to `true`.
        /// </summary>
        public readonly bool? IpFilterPermissiveRule;
        /// <summary>
        /// Is any ports open on the device? Defaults to `true`.
        /// </summary>
        public readonly bool? OpenPorts;
        /// <summary>
        /// Does firewall policy exist which allow necessary communication to/from the device? Defaults to `true`.
        /// </summary>
        public readonly bool? PermissiveFirewallPolicy;
        /// <summary>
        /// Is only necessary addresses or ports are permitted in? Defaults to `true`.
        /// </summary>
        public readonly bool? PermissiveInputFirewallRules;
        /// <summary>
        /// Is only necessary addresses or ports are permitted out? Defaults to `true`.
        /// </summary>
        public readonly bool? PermissiveOutputFirewallRules;
        /// <summary>
        /// Is high level permissions are needed for the module? Defaults to `true`.
        /// </summary>
        public readonly bool? PrivilegedDockerOptions;
        /// <summary>
        /// Is any credentials shared among devices? Defaults to `true`.
        /// </summary>
        public readonly bool? SharedCredentials;
        /// <summary>
        /// Does TLS cipher suite need to be updated? Defaults to `true`.
        /// </summary>
        public readonly bool? VulnerableTlsCipherSuite;

        [OutputConstructor]
        private SecuritySolutionRecommendationsEnabled(
            bool? acrAuthentication,

            bool? agentSendUnutilizedMsg,

            bool? baseline,

            bool? edgeHubMemOptimize,

            bool? edgeLoggingOption,

            bool? inconsistentModuleSettings,

            bool? installAgent,

            bool? ipFilterDenyAll,

            bool? ipFilterPermissiveRule,

            bool? openPorts,

            bool? permissiveFirewallPolicy,

            bool? permissiveInputFirewallRules,

            bool? permissiveOutputFirewallRules,

            bool? privilegedDockerOptions,

            bool? sharedCredentials,

            bool? vulnerableTlsCipherSuite)
        {
            AcrAuthentication = acrAuthentication;
            AgentSendUnutilizedMsg = agentSendUnutilizedMsg;
            Baseline = baseline;
            EdgeHubMemOptimize = edgeHubMemOptimize;
            EdgeLoggingOption = edgeLoggingOption;
            InconsistentModuleSettings = inconsistentModuleSettings;
            InstallAgent = installAgent;
            IpFilterDenyAll = ipFilterDenyAll;
            IpFilterPermissiveRule = ipFilterPermissiveRule;
            OpenPorts = openPorts;
            PermissiveFirewallPolicy = permissiveFirewallPolicy;
            PermissiveInputFirewallRules = permissiveInputFirewallRules;
            PermissiveOutputFirewallRules = permissiveOutputFirewallRules;
            PrivilegedDockerOptions = privilegedDockerOptions;
            SharedCredentials = sharedCredentials;
            VulnerableTlsCipherSuite = vulnerableTlsCipherSuite;
        }
    }
}
