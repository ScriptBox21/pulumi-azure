// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp.Inputs
{

    public sealed class VolumeDataProtectionReplicationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The endpoint type, default value is `dst` for destination.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// Primary volume's location.
        /// </summary>
        [Input("remoteVolumeLocation", required: true)]
        public Input<string> RemoteVolumeLocation { get; set; } = null!;

        /// <summary>
        /// Primary volume's resource id.
        /// </summary>
        [Input("remoteVolumeResourceId", required: true)]
        public Input<string> RemoteVolumeResourceId { get; set; } = null!;

        /// <summary>
        /// Replication frequency, supported values are '10minutes', 'hourly', 'daily', values are case sensitive.
        /// </summary>
        [Input("replicationFrequency", required: true)]
        public Input<string> ReplicationFrequency { get; set; } = null!;

        public VolumeDataProtectionReplicationGetArgs()
        {
        }
    }
}
