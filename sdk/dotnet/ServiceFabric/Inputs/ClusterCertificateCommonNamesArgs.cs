// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ServiceFabric.Inputs
{

    public sealed class ClusterCertificateCommonNamesArgs : Pulumi.ResourceArgs
    {
        [Input("commonNames", required: true)]
        private InputList<Inputs.ClusterCertificateCommonNamesCommonNameArgs>? _commonNames;

        /// <summary>
        /// A `common_names` block as defined below.
        /// </summary>
        public InputList<Inputs.ClusterCertificateCommonNamesCommonNameArgs> CommonNames
        {
            get => _commonNames ?? (_commonNames = new InputList<Inputs.ClusterCertificateCommonNamesCommonNameArgs>());
            set => _commonNames = value;
        }

        /// <summary>
        /// The X509 Store where the Certificate Exists, such as `My`.
        /// </summary>
        [Input("x509StoreName", required: true)]
        public Input<string> X509StoreName { get; set; } = null!;

        public ClusterCertificateCommonNamesArgs()
        {
        }
    }
}
