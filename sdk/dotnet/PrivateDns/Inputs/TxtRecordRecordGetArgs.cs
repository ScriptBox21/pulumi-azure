// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateDns.Inputs
{

    public sealed class TxtRecordRecordGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value of the TXT record. Max length: 1024 characters
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public TxtRecordRecordGetArgs()
        {
        }
    }
}
