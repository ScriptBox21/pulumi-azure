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
    public sealed class GroupContainerPort
    {
        /// <summary>
        /// The port number the container will expose. Changing this forces a new resource to be created.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The network protocol associated with port. Possible values are `TCP` &amp; `UDP`. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Protocol;

        [OutputConstructor]
        private GroupContainerPort(
            int? port,

            string? protocol)
        {
            Port = port;
            Protocol = protocol;
        }
    }
}
