// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateDns.Inputs
{

    public sealed class MxRecordRecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FQDN of the exchange to MX record points to.
        /// </summary>
        [Input("exchange", required: true)]
        public Input<string> Exchange { get; set; } = null!;

        /// <summary>
        /// The preference of the MX record.
        /// </summary>
        [Input("preference", required: true)]
        public Input<int> Preference { get; set; } = null!;

        public MxRecordRecordArgs()
        {
        }
    }
}
