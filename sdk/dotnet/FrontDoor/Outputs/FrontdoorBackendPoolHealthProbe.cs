// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor.Outputs
{

    [OutputType]
    public sealed class FrontdoorBackendPoolHealthProbe
    {
        /// <summary>
        /// Is this health probe enabled? Dafaults to `true`.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The ID of the FrontDoor.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The number of seconds between each Health Probe. Defaults to `120`.
        /// </summary>
        public readonly int? IntervalInSeconds;
        /// <summary>
        /// Specifies the name of the Health Probe.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The path to use for the Health Probe. Default is `/`.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// Specifies HTTP method the health probe uses when querying the backend pool instances. Possible values include: `Get` and `Head`. Defaults to `Get`.
        /// </summary>
        public readonly string? ProbeMethod;
        /// <summary>
        /// Protocol scheme to use for the Health Probe. Defaults to `Http`.
        /// </summary>
        public readonly string? Protocol;

        [OutputConstructor]
        private FrontdoorBackendPoolHealthProbe(
            bool? enabled,

            string? id,

            int? intervalInSeconds,

            string name,

            string? path,

            string? probeMethod,

            string? protocol)
        {
            Enabled = enabled;
            Id = id;
            IntervalInSeconds = intervalInSeconds;
            Name = name;
            Path = path;
            ProbeMethod = probeMethod;
            Protocol = protocol;
        }
    }
}
