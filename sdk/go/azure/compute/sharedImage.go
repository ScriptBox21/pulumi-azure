// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Shared Image within a Shared Image Gallery.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/shared_image.html.markdown.
type SharedImage struct {
	pulumi.CustomResourceState

	// A description of this Shared Image.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The End User Licence Agreement for the Shared Image.
	Eula pulumi.StringPtrOutput `pulumi:"eula"`
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName pulumi.StringOutput `pulumi:"galleryName"`
	// An `identifier` block as defined below.
	Identifier SharedImageIdentifierOutput `pulumi:"identifier"`
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri pulumi.StringPtrOutput `pulumi:"privacyStatementUri"`
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri pulumi.StringPtrOutput `pulumi:"releaseNoteUri"`
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Shared Image.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSharedImage registers a new resource with the given unique name, arguments, and options.
func NewSharedImage(ctx *pulumi.Context,
	name string, args *SharedImageArgs, opts ...pulumi.ResourceOption) (*SharedImage, error) {
	if args == nil || args.GalleryName == nil {
		return nil, errors.New("missing required argument 'GalleryName'")
	}
	if args == nil || args.Identifier == nil {
		return nil, errors.New("missing required argument 'Identifier'")
	}
	if args == nil || args.OsType == nil {
		return nil, errors.New("missing required argument 'OsType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &SharedImageArgs{}
	}
	var resource SharedImage
	err := ctx.RegisterResource("azure:compute/sharedImage:SharedImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSharedImage gets an existing SharedImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSharedImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SharedImageState, opts ...pulumi.ResourceOption) (*SharedImage, error) {
	var resource SharedImage
	err := ctx.ReadResource("azure:compute/sharedImage:SharedImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SharedImage resources.
type sharedImageState struct {
	// A description of this Shared Image.
	Description *string `pulumi:"description"`
	// The End User Licence Agreement for the Shared Image.
	Eula *string `pulumi:"eula"`
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName *string `pulumi:"galleryName"`
	// An `identifier` block as defined below.
	Identifier *SharedImageIdentifier `pulumi:"identifier"`
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
	OsType *string `pulumi:"osType"`
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Shared Image.
	Tags map[string]string `pulumi:"tags"`
}

type SharedImageState struct {
	// A description of this Shared Image.
	Description pulumi.StringPtrInput
	// The End User Licence Agreement for the Shared Image.
	Eula pulumi.StringPtrInput
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName pulumi.StringPtrInput
	// An `identifier` block as defined below.
	Identifier SharedImageIdentifierPtrInput
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
	OsType pulumi.StringPtrInput
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri pulumi.StringPtrInput
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri pulumi.StringPtrInput
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Shared Image.
	Tags pulumi.StringMapInput
}

func (SharedImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedImageState)(nil)).Elem()
}

type sharedImageArgs struct {
	// A description of this Shared Image.
	Description *string `pulumi:"description"`
	// The End User Licence Agreement for the Shared Image.
	Eula *string `pulumi:"eula"`
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName string `pulumi:"galleryName"`
	// An `identifier` block as defined below.
	Identifier SharedImageIdentifier `pulumi:"identifier"`
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
	OsType string `pulumi:"osType"`
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Shared Image.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SharedImage resource.
type SharedImageArgs struct {
	// A description of this Shared Image.
	Description pulumi.StringPtrInput
	// The End User Licence Agreement for the Shared Image.
	Eula pulumi.StringPtrInput
	// Specifies the name of the Shared Image Gallery in which this Shared Image should exist. Changing this forces a new resource to be created.
	GalleryName pulumi.StringInput
	// An `identifier` block as defined below.
	Identifier SharedImageIdentifierInput
	// Specifies the supported Azure location where the Shared Image Gallery exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Shared Image. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The type of Operating System present in this Shared Image. Possible values are `Linux` and `Windows`.
	OsType pulumi.StringInput
	// The URI containing the Privacy Statement associated with this Shared Image.
	PrivacyStatementUri pulumi.StringPtrInput
	// The URI containing the Release Notes associated with this Shared Image.
	ReleaseNoteUri pulumi.StringPtrInput
	// The name of the resource group in which the Shared Image Gallery exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Shared Image.
	Tags pulumi.StringMapInput
}

func (SharedImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sharedImageArgs)(nil)).Elem()
}
