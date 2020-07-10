// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayHttpListenerArgs : Pulumi.ResourceArgs
    {
        [Input("customErrorConfigurations")]
        private InputList<Inputs.ApplicationGatewayHttpListenerCustomErrorConfigurationArgs>? _customErrorConfigurations;

        /// <summary>
        /// One or more `custom_error_configuration` blocks as defined below.
        /// </summary>
        public InputList<Inputs.ApplicationGatewayHttpListenerCustomErrorConfigurationArgs> CustomErrorConfigurations
        {
            get => _customErrorConfigurations ?? (_customErrorConfigurations = new InputList<Inputs.ApplicationGatewayHttpListenerCustomErrorConfigurationArgs>());
            set => _customErrorConfigurations = value;
        }

        /// <summary>
        /// The ID of the Web Application Firewall Policy which should be used as a HTTP Listener.
        /// </summary>
        [Input("firewallPolicyId")]
        public Input<string>? FirewallPolicyId { get; set; }

        /// <summary>
        /// The ID of the associated Frontend Configuration.
        /// </summary>
        [Input("frontendIpConfigurationId")]
        public Input<string>? FrontendIpConfigurationId { get; set; }

        /// <summary>
        /// The Name of the Frontend IP Configuration used for this HTTP Listener.
        /// </summary>
        [Input("frontendIpConfigurationName", required: true)]
        public Input<string> FrontendIpConfigurationName { get; set; } = null!;

        /// <summary>
        /// The ID of the associated Frontend Port.
        /// </summary>
        [Input("frontendPortId")]
        public Input<string>? FrontendPortId { get; set; }

        /// <summary>
        /// The Name of the Frontend Port use for this HTTP Listener.
        /// </summary>
        [Input("frontendPortName", required: true)]
        public Input<string> FrontendPortName { get; set; } = null!;

        /// <summary>
        /// The Hostname which should be used for this HTTP Listener.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        [Input("hostNames")]
        private InputList<string>? _hostNames;

        /// <summary>
        /// A list of Hostname(s) should be used for this HTTP Listener. It allows special wildcard characters.
        /// </summary>
        public InputList<string> HostNames
        {
            get => _hostNames ?? (_hostNames = new InputList<string>());
            set => _hostNames = value;
        }

        /// <summary>
        /// The ID of the Rewrite Rule Set
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The Name of the HTTP Listener.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The Protocol to use for this HTTP Listener. Possible values are `Http` and `Https`.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Should Server Name Indication be Required? Defaults to `false`.
        /// </summary>
        [Input("requireSni")]
        public Input<bool>? RequireSni { get; set; }

        /// <summary>
        /// The ID of the associated SSL Certificate.
        /// </summary>
        [Input("sslCertificateId")]
        public Input<string>? SslCertificateId { get; set; }

        /// <summary>
        /// The name of the associated SSL Certificate which should be used for this HTTP Listener.
        /// </summary>
        [Input("sslCertificateName")]
        public Input<string>? SslCertificateName { get; set; }

        public ApplicationGatewayHttpListenerArgs()
        {
        }
    }
}
