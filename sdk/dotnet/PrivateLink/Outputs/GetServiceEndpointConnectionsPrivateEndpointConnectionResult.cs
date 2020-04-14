// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateLink.Outputs
{

    [OutputType]
    public sealed class GetServiceEndpointConnectionsPrivateEndpointConnectionResult
    {
        /// <summary>
        /// A message indicating if changes on the service provider require any updates or not.
        /// </summary>
        public readonly string ActionRequired;
        /// <summary>
        /// The resource id of the private link service connection between the private link service and the private link endpoint.
        /// </summary>
        public readonly string ConnectionId;
        /// <summary>
        /// The name of the connection between the private link service and the private link endpoint.
        /// </summary>
        public readonly string ConnectionName;
        /// <summary>
        /// The request for approval message or the reason for rejection message.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The resource id of the private link endpoint.
        /// </summary>
        public readonly string PrivateEndpointId;
        /// <summary>
        /// The name of the private link endpoint.
        /// </summary>
        public readonly string PrivateEndpointName;
        /// <summary>
        /// Indicates the state of the connection between the private link service and the private link endpoint, possible values are `Pending`, `Approved` or `Rejected`.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetServiceEndpointConnectionsPrivateEndpointConnectionResult(
            string actionRequired,

            string connectionId,

            string connectionName,

            string description,

            string privateEndpointId,

            string privateEndpointName,

            string status)
        {
            ActionRequired = actionRequired;
            ConnectionId = connectionId;
            ConnectionName = connectionName;
            Description = description;
            PrivateEndpointId = privateEndpointId;
            PrivateEndpointName = privateEndpointName;
            Status = status;
        }
    }
}
