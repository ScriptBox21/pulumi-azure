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
    public sealed class KubernetesClusterAddonProfileOmsAgentOmsAgentIdentity
    {
        /// <summary>
        /// The Client ID for the Service Principal.
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// The Object ID of the user-defined Managed Identity used by the OMS Agents.
        /// </summary>
        public readonly string? ObjectId;
        /// <summary>
        /// The ID of a user assigned identity.
        /// </summary>
        public readonly string? UserAssignedIdentityId;

        [OutputConstructor]
        private KubernetesClusterAddonProfileOmsAgentOmsAgentIdentity(
            string? clientId,

            string? objectId,

            string? userAssignedIdentityId)
        {
            ClientId = clientId;
            ObjectId = objectId;
            UserAssignedIdentityId = userAssignedIdentityId;
        }
    }
}
