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
    public sealed class KubernetesClusterAddonProfileAzurePolicy
    {
        /// <summary>
        /// Is the Azure Policy for Kubernetes Add On enabled?
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private KubernetesClusterAddonProfileAzurePolicy(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
