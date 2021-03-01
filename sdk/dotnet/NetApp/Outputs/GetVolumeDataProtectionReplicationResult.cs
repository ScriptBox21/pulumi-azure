// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp.Outputs
{

    [OutputType]
    public sealed class GetVolumeDataProtectionReplicationResult
    {
        public readonly string EndpointType;
        public readonly string RemoteVolumeLocation;
        public readonly string RemoteVolumeResourceId;
        public readonly string ReplicationSchedule;

        [OutputConstructor]
        private GetVolumeDataProtectionReplicationResult(
            string endpointType,

            string remoteVolumeLocation,

            string remoteVolumeResourceId,

            string replicationSchedule)
        {
            EndpointType = endpointType;
            RemoteVolumeLocation = remoteVolumeLocation;
            RemoteVolumeResourceId = remoteVolumeResourceId;
            ReplicationSchedule = replicationSchedule;
        }
    }
}
