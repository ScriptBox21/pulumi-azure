// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Outputs
{

    [OutputType]
    public sealed class GetActionGroupAzureAppPushReceiverResult
    {
        /// <summary>
        /// The email address of this receiver.
        /// </summary>
        public readonly string EmailAddress;
        /// <summary>
        /// Specifies the name of the Action Group.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetActionGroupAzureAppPushReceiverResult(
            string emailAddress,

            string name)
        {
            EmailAddress = emailAddress;
            Name = name;
        }
    }
}
