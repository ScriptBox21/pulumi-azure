// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp.Inputs
{

    public sealed class VolumeExportPolicyRuleGetArgs : Pulumi.ResourceArgs
    {
        [Input("allowedClients", required: true)]
        private InputList<string>? _allowedClients;

        /// <summary>
        /// A list of allowed clients IPv4 addresses.
        /// </summary>
        public InputList<string> AllowedClients
        {
            get => _allowedClients ?? (_allowedClients = new InputList<string>());
            set => _allowedClients = value;
        }

        /// <summary>
        /// Is the CIFS protocol allowed?
        /// </summary>
        [Input("cifsEnabled")]
        public Input<bool>? CifsEnabled { get; set; }

        /// <summary>
        /// Is the NFSv3 protocol allowed?
        /// </summary>
        [Input("nfsv3Enabled")]
        public Input<bool>? Nfsv3Enabled { get; set; }

        /// <summary>
        /// Is the NFSv4 protocol allowed?
        /// </summary>
        [Input("nfsv4Enabled")]
        public Input<bool>? Nfsv4Enabled { get; set; }

        /// <summary>
        /// A list of allowed protocols. Valid values include `CIFS`, `NFSv3`, or `NFSv4.1`. Only one value is supported at this time. This replaces the previous arguments: `cifs_enabled`, `nfsv3_enabled` and `nfsv4_enabled`.
        /// </summary>
        [Input("protocolsEnabled")]
        public Input<string>? ProtocolsEnabled { get; set; }

        /// <summary>
        /// The index number of the rule.
        /// </summary>
        [Input("ruleIndex", required: true)]
        public Input<int> RuleIndex { get; set; } = null!;

        /// <summary>
        /// Is the file system on unix read only?
        /// </summary>
        [Input("unixReadOnly")]
        public Input<bool>? UnixReadOnly { get; set; }

        /// <summary>
        /// Is the file system on unix read and write?
        /// </summary>
        [Input("unixReadWrite")]
        public Input<bool>? UnixReadWrite { get; set; }

        public VolumeExportPolicyRuleGetArgs()
        {
        }
    }
}
