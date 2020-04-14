// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ContainerService.Inputs
{

    public sealed class KubernetesClusterAddonProfileAciConnectorLinuxArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is the virtual node addon enabled?
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// The subnet name for the virtual nodes to run. This is required when `aci_connector_linux` `enabled` argument is set to `true`.
        /// </summary>
        [Input("subnetName")]
        public Input<string>? SubnetName { get; set; }

        public KubernetesClusterAddonProfileAciConnectorLinuxArgs()
        {
        }
    }
}
