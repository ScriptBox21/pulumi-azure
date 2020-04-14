// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewaySslPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("cipherSuites")]
        private InputList<string>? _cipherSuites;

        /// <summary>
        /// A List of accepted cipher suites. Possible values are: `TLS_DHE_DSS_WITH_AES_128_CBC_SHA`, `TLS_DHE_DSS_WITH_AES_128_CBC_SHA256`, `TLS_DHE_DSS_WITH_AES_256_CBC_SHA`, `TLS_DHE_DSS_WITH_AES_256_CBC_SHA256`, `TLS_DHE_RSA_WITH_AES_128_CBC_SHA`, `TLS_DHE_RSA_WITH_AES_128_GCM_SHA256`, `TLS_DHE_RSA_WITH_AES_256_CBC_SHA`, `TLS_DHE_RSA_WITH_AES_256_GCM_SHA384`, `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA`, `TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256`, `TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256`, `TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA`, `TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384`, `TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384`, `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA`, `TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256`, `TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA`, `TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384`, `TLS_RSA_WITH_3DES_EDE_CBC_SHA`, `TLS_RSA_WITH_AES_128_CBC_SHA`, `TLS_RSA_WITH_AES_128_CBC_SHA256`, `TLS_RSA_WITH_AES_128_GCM_SHA256`, `TLS_RSA_WITH_AES_256_CBC_SHA`, `TLS_RSA_WITH_AES_256_CBC_SHA256` and `TLS_RSA_WITH_AES_256_GCM_SHA384`.
        /// </summary>
        public InputList<string> CipherSuites
        {
            get => _cipherSuites ?? (_cipherSuites = new InputList<string>());
            set => _cipherSuites = value;
        }

        [Input("disabledProtocols")]
        private InputList<string>? _disabledProtocols;

        /// <summary>
        /// A list of SSL Protocols which should be disabled on this Application Gateway. Possible values are `TLSv1_0`, `TLSv1_1` and `TLSv1_2`.
        /// </summary>
        public InputList<string> DisabledProtocols
        {
            get => _disabledProtocols ?? (_disabledProtocols = new InputList<string>());
            set => _disabledProtocols = value;
        }

        /// <summary>
        /// The minimal TLS version. Possible values are `TLSv1_0`, `TLSv1_1` and `TLSv1_2`.
        /// </summary>
        [Input("minProtocolVersion")]
        public Input<string>? MinProtocolVersion { get; set; }

        /// <summary>
        /// The Name of the Policy e.g AppGwSslPolicy20170401S. Required if `policy_type` is set to `Predefined`. Possible values can change over time and
        /// are published here https://docs.microsoft.com/en-us/azure/application-gateway/application-gateway-ssl-policy-overview. Not compatible with `disabled_protocols`.
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The Type of the Policy. Possible values are `Predefined` and `Custom`.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        public ApplicationGatewaySslPolicyGetArgs()
        {
        }
    }
}
