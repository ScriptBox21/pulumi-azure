// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Image.
func LookupImage(ctx *pulumi.Context, args *LookupImageArgs, opts ...pulumi.InvokeOption) (*LookupImageResult, error) {
	var rv LookupImageResult
	err := ctx.Invoke("azure:compute/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type LookupImageArgs struct {
	// The name of the Image.
	Name *string `pulumi:"name"`
	// Regex pattern of the image to match.
	NameRegex *string `pulumi:"nameRegex"`
	// The Name of the Resource Group where this Image exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// By default when matching by regex, images are sorted by name in ascending order and the first match is chosen, to sort descending, set this flag.
	SortDescending *bool `pulumi:"sortDescending"`
}

// A collection of values returned by getImage.
type LookupImageResult struct {
	// a collection of `dataDisk` blocks as defined below.
	DataDisks []GetImageDataDisk `pulumi:"dataDisks"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// the Azure Location where this Image exists.
	Location string `pulumi:"location"`
	// the name of the Image.
	Name      *string `pulumi:"name"`
	NameRegex *string `pulumi:"nameRegex"`
	// a `osDisk` block as defined below.
	OsDisks           []GetImageOsDisk `pulumi:"osDisks"`
	ResourceGroupName string           `pulumi:"resourceGroupName"`
	SortDescending    *bool            `pulumi:"sortDescending"`
	// a mapping of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// is zone resiliency enabled?
	ZoneResilient bool `pulumi:"zoneResilient"`
}
