// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation.Inputs
{

    public sealed class RunBookPublishContentLinkGetArgs : Pulumi.ResourceArgs
    {
        [Input("hash")]
        public Input<Inputs.RunBookPublishContentLinkHashGetArgs>? Hash { get; set; }

        /// <summary>
        /// The uri of the runbook content.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        [Input("version")]
        public Input<string>? Version { get; set; }

        public RunBookPublishContentLinkGetArgs()
        {
        }
    }
}
