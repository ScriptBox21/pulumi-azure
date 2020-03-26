// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package batch

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing certificate in a Batch Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/batch_certificate.html.markdown.
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azure:batch/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// The name of the Batch certificate.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where this Batch account exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	AccountName string `pulumi:"accountName"`
	// The format of the certificate, such as `Cer` or `Pfx`.
	Format string `pulumi:"format"`
	// id is the provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The public key of the certificate.
	PublicData        string `pulumi:"publicData"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The thumbprint of the certificate.
	Thumbprint string `pulumi:"thumbprint"`
	// The algorithm of the certificate thumbprint.
	ThumbprintAlgorithm string `pulumi:"thumbprintAlgorithm"`
}
