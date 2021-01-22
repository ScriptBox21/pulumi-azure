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
    public sealed class KubernetesClusterIdentity
    {
        /// <summary>
        /// The principal id of the system assigned identity which is used by master components.
        /// </summary>
        public readonly string? PrincipalId;
        /// <summary>
        /// The Tenant ID used for Azure Active Directory Application. If this isn't specified the Tenant ID of the current Subscription is used.
        /// </summary>
        public readonly string? TenantId;
        /// <summary>
        /// The type of identity used for the managed cluster. Possible values are `SystemAssigned` and `UserAssigned`. If `UserAssigned` is set, a `user_assigned_identity_id` must be set as well.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The ID of a user assigned identity.
        /// </summary>
        public readonly string? UserAssignedIdentityId;

        [OutputConstructor]
        private KubernetesClusterIdentity(
            string? principalId,

            string? tenantId,

            string type,

            string? userAssignedIdentityId)
        {
            PrincipalId = principalId;
            TenantId = tenantId;
            Type = type;
            UserAssignedIdentityId = userAssignedIdentityId;
        }
    }
}
