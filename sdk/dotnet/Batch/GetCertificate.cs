// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Batch
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to access information about an existing certificate in a Batch Account.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/batch_certificate.html.markdown.
        /// </summary>
        [Obsolete("Use GetCertificate.InvokeAsync() instead")]
        public static Task<GetCertificateResult> GetCertificate(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("azure:batch/getCertificate:getCertificate", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetCertificate
    {
        /// <summary>
        /// Use this data source to access information about an existing certificate in a Batch Account.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/batch_certificate.html.markdown.
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("azure:batch/getCertificate:getCertificate", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the Batch certificate.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where this Batch account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCertificateArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetCertificateResult
    {
        public readonly string AccountName;
        /// <summary>
        /// The format of the certificate, such as `Cer` or `Pfx`.
        /// </summary>
        public readonly string Format;
        public readonly string Name;
        /// <summary>
        /// The public key of the certificate.
        /// </summary>
        public readonly string PublicData;
        public readonly string ResourceGroupName;
        /// <summary>
        /// The thumbprint of the certificate.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// The algorithm of the certificate thumbprint.
        /// </summary>
        public readonly string ThumbprintAlgorithm;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetCertificateResult(
            string accountName,
            string format,
            string name,
            string publicData,
            string resourceGroupName,
            string thumbprint,
            string thumbprintAlgorithm,
            string id)
        {
            AccountName = accountName;
            Format = format;
            Name = name;
            PublicData = publicData;
            ResourceGroupName = resourceGroupName;
            Thumbprint = thumbprint;
            ThumbprintAlgorithm = thumbprintAlgorithm;
            Id = id;
        }
    }
}
