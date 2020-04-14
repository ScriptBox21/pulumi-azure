// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp.Inputs
{

    public sealed class AccountActiveDirectoryGetArgs : Pulumi.ResourceArgs
    {
        [Input("dnsServers", required: true)]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// A list of DNS server IP addresses for the Active Directory domain. Only allows `IPv4` address.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// The name of the Active Directory domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// The Organizational Unit (OU) within the Active Directory Domain.
        /// </summary>
        [Input("organizationalUnit")]
        public Input<string>? OrganizationalUnit { get; set; }

        /// <summary>
        /// The password associated with the `username`.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The NetBIOS name which should be used for the NetApp SMB Server, which will be registered as a computer account in the AD and used to mount volumes.
        /// </summary>
        [Input("smbServerName", required: true)]
        public Input<string> SmbServerName { get; set; } = null!;

        /// <summary>
        /// The Username of Active Directory Domain Administrator.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public AccountActiveDirectoryGetArgs()
        {
        }
    }
}
