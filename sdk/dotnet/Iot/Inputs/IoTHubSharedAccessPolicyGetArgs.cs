// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot.Inputs
{

    public sealed class IoTHubSharedAccessPolicyGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the shared access policy.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// The permissions assigned to the shared access policy.
        /// </summary>
        [Input("permissions")]
        public Input<string>? Permissions { get; set; }

        /// <summary>
        /// The primary key.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        /// <summary>
        /// The secondary key.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        public IoTHubSharedAccessPolicyGetArgs()
        {
        }
    }
}
