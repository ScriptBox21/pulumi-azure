// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SignalR.Outputs
{

    [OutputType]
    public sealed class ServiceSku
    {
        /// <summary>
        /// Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ServiceSku(
            int capacity,

            string name)
        {
            Capacity = capacity;
            Name = name;
        }
    }
}
