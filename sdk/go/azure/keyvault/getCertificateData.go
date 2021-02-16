// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keyvault

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access data stored in an existing Key Vault Certificate.
//
// > **Note:** All arguments including the secret value will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > **Note:** This data source uses the `GetSecret` function of the Azure API, to get the key of the certificate. Therefore you need secret/get permission
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleKeyVault, err := keyvault.LookupKeyVault(ctx, &keyvault.LookupKeyVaultArgs{
// 			Name:              "examplekv",
// 			ResourceGroupName: "some-resource-group",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = keyvault.GetCertificateData(ctx, &keyvault.GetCertificateDataArgs{
// 			Name:       "secret-sauce",
// 			KeyVaultId: exampleKeyVault.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("examplePem", data.Azurerm_key_vault_certificate.Example.Pem)
// 		return nil
// 	})
// }
// ```
func GetCertificateData(ctx *pulumi.Context, args *GetCertificateDataArgs, opts ...pulumi.InvokeOption) (*GetCertificateDataResult, error) {
	var rv GetCertificateDataResult
	err := ctx.Invoke("azure:keyvault/getCertificateData:getCertificateData", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificateData.
type GetCertificateDataArgs struct {
	// Specifies the ID of the Key Vault instance where the Secret resides, available on the `keyvault.KeyVault` Data Source / Resource.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the name of the Key Vault Secret.
	Name string `pulumi:"name"`
	// Specifies the version of the certificate to look up.  (Defaults to latest)
	Version *string `pulumi:"version"`
}

// A collection of values returned by getCertificateData.
type GetCertificateDataResult struct {
	// Expiry date of certificate in RFC3339 format.
	Expires string `pulumi:"expires"`
	// The raw Key Vault Certificate data represented as a hexadecimal string.
	Hex string `pulumi:"hex"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Key Vault Certificate Key.
	Key        string `pulumi:"key"`
	KeyVaultId string `pulumi:"keyVaultId"`
	Name       string `pulumi:"name"`
	// The Key Vault Certificate in PEM format.
	Pem string `pulumi:"pem"`
	// A mapping of tags to assign to the resource.
	Tags    map[string]string `pulumi:"tags"`
	Version string            `pulumi:"version"`
}
