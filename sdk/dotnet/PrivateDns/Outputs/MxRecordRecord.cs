// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.PrivateDns.Outputs
{

    [OutputType]
    public sealed class MxRecordRecord
    {
        /// <summary>
        /// The FQDN of the exchange to MX record points to.
        /// </summary>
        public readonly string Exchange;
        /// <summary>
        /// The preference of the MX record.
        /// </summary>
        public readonly int Preference;

        [OutputConstructor]
        private MxRecordRecord(
            string exchange,

            int preference)
        {
            Exchange = exchange;
            Preference = preference;
        }
    }
}
