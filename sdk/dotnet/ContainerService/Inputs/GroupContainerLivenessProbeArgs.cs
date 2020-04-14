// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class GroupContainerLivenessProbeArgs : Pulumi.ResourceArgs
    {
        [Input("execs")]
        private InputList<string>? _execs;

        /// <summary>
        /// Commands to be run to validate container readiness. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<string> Execs
        {
            get => _execs ?? (_execs = new InputList<string>());
            set => _execs = value;
        }

        /// <summary>
        /// How many times to try the probe before restarting the container (liveness probe) or marking the container as unhealthy (readiness probe). The default value is `3` and the minimum value is `1`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("failureThreshold")]
        public Input<int>? FailureThreshold { get; set; }

        [Input("httpGets")]
        private InputList<Inputs.GroupContainerLivenessProbeHttpGetArgs>? _httpGets;

        /// <summary>
        /// The definition of the httpget for this container as documented in the `httpget` block below. Changing this forces a new resource to be created.
        /// </summary>
        public InputList<Inputs.GroupContainerLivenessProbeHttpGetArgs> HttpGets
        {
            get => _httpGets ?? (_httpGets = new InputList<Inputs.GroupContainerLivenessProbeHttpGetArgs>());
            set => _httpGets = value;
        }

        /// <summary>
        /// Number of seconds after the container has started before liveness or readiness probes are initiated. Changing this forces a new resource to be created.
        /// </summary>
        [Input("initialDelaySeconds")]
        public Input<int>? InitialDelaySeconds { get; set; }

        /// <summary>
        /// How often (in seconds) to perform the probe. The default value is `10` and the minimum value is `1`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("periodSeconds")]
        public Input<int>? PeriodSeconds { get; set; }

        /// <summary>
        /// Minimum consecutive successes for the probe to be considered successful after having failed. The default value is `1` and the minimum value is `1`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("successThreshold")]
        public Input<int>? SuccessThreshold { get; set; }

        /// <summary>
        /// Number of seconds after which the probe times out. The default value is `1` and the minimum value is `1`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        public GroupContainerLivenessProbeArgs()
        {
        }
    }
}
