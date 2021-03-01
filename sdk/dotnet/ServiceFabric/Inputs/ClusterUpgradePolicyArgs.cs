// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceFabric.Inputs
{

    public sealed class ClusterUpgradePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `delta_health_policy` block as defined below
        /// </summary>
        [Input("deltaHealthPolicy")]
        public Input<Inputs.ClusterUpgradePolicyDeltaHealthPolicyArgs>? DeltaHealthPolicy { get; set; }

        [Input("forceRestartEnabled")]
        public Input<bool>? ForceRestartEnabled { get; set; }

        /// <summary>
        /// Specifies the duration, in "hh:mm:ss" string format, after which Service Fabric retries the health check if the previous health check fails. Defaults to `00:45:00`.
        /// </summary>
        [Input("healthCheckRetryTimeout")]
        public Input<string>? HealthCheckRetryTimeout { get; set; }

        /// <summary>
        /// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric waits in order to verify that the cluster is stable before it continues to the next upgrade domain or completes the upgrade. This wait duration prevents undetected changes of health right after the health check is performed. Defaults to `00:01:00`.
        /// </summary>
        [Input("healthCheckStableDuration")]
        public Input<string>? HealthCheckStableDuration { get; set; }

        /// <summary>
        /// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric waits before it performs the initial health check after it finishes the upgrade on the upgrade domain. Defaults to `00:00:30`.
        /// </summary>
        [Input("healthCheckWaitDuration")]
        public Input<string>? HealthCheckWaitDuration { get; set; }

        /// <summary>
        /// A `health_policy` block as defined below
        /// </summary>
        [Input("healthPolicy")]
        public Input<Inputs.ClusterUpgradePolicyHealthPolicyArgs>? HealthPolicy { get; set; }

        /// <summary>
        /// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric takes to upgrade a single upgrade domain. After this period, the upgrade fails. Defaults to `02:00:00`.
        /// </summary>
        [Input("upgradeDomainTimeout")]
        public Input<string>? UpgradeDomainTimeout { get; set; }

        /// <summary>
        /// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric waits for a replica set to reconfigure into a safe state, if it is not already in a safe state, before Service Fabric proceeds with the upgrade. Defaults to `10675199.02:48:05.4775807`.
        /// </summary>
        [Input("upgradeReplicaSetCheckTimeout")]
        public Input<string>? UpgradeReplicaSetCheckTimeout { get; set; }

        /// <summary>
        /// Specifies the duration, in "hh:mm:ss" string format, that Service Fabric takes for the entire upgrade. After this period, the upgrade fails. Defaults to `12:00:00`.
        /// </summary>
        [Input("upgradeTimeout")]
        public Input<string>? UpgradeTimeout { get; set; }

        public ClusterUpgradePolicyArgs()
        {
        }
    }
}
