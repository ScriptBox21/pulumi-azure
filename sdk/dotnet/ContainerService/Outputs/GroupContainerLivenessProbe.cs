// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Outputs
{

    [OutputType]
    public sealed class GroupContainerLivenessProbe
    {
        /// <summary>
        /// Commands to be run to validate container readiness. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<string> Execs;
        /// <summary>
        /// How many times to try the probe before restarting the container (liveness probe) or marking the container as unhealthy (readiness probe). The default value is `3` and the minimum value is `1`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? FailureThreshold;
        /// <summary>
        /// The definition of the httpget for this container as documented in the `httpget` block below. Changing this forces a new resource to be created.
        /// </summary>
        public readonly ImmutableArray<Outputs.GroupContainerLivenessProbeHttpGet> HttpGets;
        /// <summary>
        /// Number of seconds after the container has started before liveness or readiness probes are initiated. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? InitialDelaySeconds;
        /// <summary>
        /// How often (in seconds) to perform the probe. The default value is `10` and the minimum value is `1`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? PeriodSeconds;
        /// <summary>
        /// Minimum consecutive successes for the probe to be considered successful after having failed. The default value is `1` and the minimum value is `1`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? SuccessThreshold;
        /// <summary>
        /// Number of seconds after which the probe times out. The default value is `1` and the minimum value is `1`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? TimeoutSeconds;

        [OutputConstructor]
        private GroupContainerLivenessProbe(
            ImmutableArray<string> execs,

            int? failureThreshold,

            ImmutableArray<Outputs.GroupContainerLivenessProbeHttpGet> httpGets,

            int? initialDelaySeconds,

            int? periodSeconds,

            int? successThreshold,

            int? timeoutSeconds)
        {
            Execs = execs;
            FailureThreshold = failureThreshold;
            HttpGets = httpGets;
            InitialDelaySeconds = initialDelaySeconds;
            PeriodSeconds = periodSeconds;
            SuccessThreshold = successThreshold;
            TimeoutSeconds = timeoutSeconds;
        }
    }
}
