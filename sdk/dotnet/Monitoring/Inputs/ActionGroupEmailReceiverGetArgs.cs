// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Inputs
{

    public sealed class ActionGroupEmailReceiverGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The email address of this receiver.
        /// </summary>
        [Input("emailAddress", required: true)]
        public Input<string> EmailAddress { get; set; } = null!;

        /// <summary>
        /// The name of the email receiver. Names must be unique (case-insensitive) across all receivers within an action group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Enables or disables the common alert schema.
        /// </summary>
        [Input("useCommonAlertSchema")]
        public Input<bool>? UseCommonAlertSchema { get; set; }

        public ActionGroupEmailReceiverGetArgs()
        {
        }
    }
}
