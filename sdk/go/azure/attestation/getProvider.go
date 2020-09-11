// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package attestation

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupProvider(ctx *pulumi.Context, args *LookupProviderArgs, opts ...pulumi.InvokeOption) (*LookupProviderResult, error) {
	var rv LookupProviderResult
	err := ctx.Invoke("azure:attestation/getProvider:getProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProvider.
type LookupProviderArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getProvider.
type LookupProviderResult struct {
	AttestationUri string `pulumi:"attestationUri"`
	// The provider-assigned unique ID for this managed resource.
	Id                string            `pulumi:"id"`
	Location          string            `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	TrustModel        string            `pulumi:"trustModel"`
}
