// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Virtual Machine Extension to provide post deployment configuration
// and run automated tasks.
// 
// > **NOTE:** Custom Script Extensions for Linux & Windows require that the `commandToExecute` returns a `0` exit code to be classified as successfully deployed. You can achieve this by appending `exit 0` to the end of your `commandToExecute`.
// 
// > **NOTE:** Custom Script Extensions require that the Azure Virtual Machine Guest Agent is running on the Virtual Machine.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_machine_extension.html.markdown.
type Extension struct {
	pulumi.CustomResourceState

	// Specifies if the platform deploys
	// the latest minor version update to the `typeHandlerVersion` specified.
	AutoUpgradeMinorVersion pulumi.BoolPtrOutput `pulumi:"autoUpgradeMinorVersion"`
	// The location where the extension is created. Changing
	// this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the virtual machine extension peering. Changing
	// this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protectedSettings passed to the
	// extension, like settings, these are specified as a JSON object in a string.
	ProtectedSettings pulumi.StringPtrOutput `pulumi:"protectedSettings"`
	// The publisher of the extension, available publishers
	// can be found by using the Azure CLI.
	Publisher pulumi.StringOutput `pulumi:"publisher"`
	// The name of the resource group in which to
	// create the virtual network. Changing this forces a new resource to be
	// created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The settings passed to the extension, these are
	// specified as a JSON object in a string.
	Settings pulumi.StringPtrOutput `pulumi:"settings"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of extension, available types for a publisher can
	// be found using the Azure CLI.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the version of the extension to
	// use, available versions can be found using the Azure CLI.
	TypeHandlerVersion pulumi.StringOutput `pulumi:"typeHandlerVersion"`
	// The resource ID of the virtual machine. This value replaces 
	// `location`, `resourceGroupName` and `virtualMachineName`. Changing this forces a new
	// resource to be created
	VirtualMachineId pulumi.StringOutput `pulumi:"virtualMachineId"`
	// The name of the virtual machine. Changing
	// this forces a new resource to be created.
	VirtualMachineName pulumi.StringOutput `pulumi:"virtualMachineName"`
}

// NewExtension registers a new resource with the given unique name, arguments, and options.
func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil || args.Publisher == nil {
		return nil, errors.New("missing required argument 'Publisher'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil || args.TypeHandlerVersion == nil {
		return nil, errors.New("missing required argument 'TypeHandlerVersion'")
	}
	if args == nil {
		args = &ExtensionArgs{}
	}
	var resource Extension
	err := ctx.RegisterResource("azure:compute/extension:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtension gets an existing Extension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("azure:compute/extension:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extension resources.
type extensionState struct {
	// Specifies if the platform deploys
	// the latest minor version update to the `typeHandlerVersion` specified.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// The location where the extension is created. Changing
	// this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the virtual machine extension peering. Changing
	// this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The protectedSettings passed to the
	// extension, like settings, these are specified as a JSON object in a string.
	ProtectedSettings *string `pulumi:"protectedSettings"`
	// The publisher of the extension, available publishers
	// can be found by using the Azure CLI.
	Publisher *string `pulumi:"publisher"`
	// The name of the resource group in which to
	// create the virtual network. Changing this forces a new resource to be
	// created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The settings passed to the extension, these are
	// specified as a JSON object in a string.
	Settings *string `pulumi:"settings"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of extension, available types for a publisher can
	// be found using the Azure CLI.
	Type *string `pulumi:"type"`
	// Specifies the version of the extension to
	// use, available versions can be found using the Azure CLI.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
	// The resource ID of the virtual machine. This value replaces 
	// `location`, `resourceGroupName` and `virtualMachineName`. Changing this forces a new
	// resource to be created
	VirtualMachineId *string `pulumi:"virtualMachineId"`
	// The name of the virtual machine. Changing
	// this forces a new resource to be created.
	VirtualMachineName *string `pulumi:"virtualMachineName"`
}

type ExtensionState struct {
	// Specifies if the platform deploys
	// the latest minor version update to the `typeHandlerVersion` specified.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// The location where the extension is created. Changing
	// this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the virtual machine extension peering. Changing
	// this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The protectedSettings passed to the
	// extension, like settings, these are specified as a JSON object in a string.
	ProtectedSettings pulumi.StringPtrInput
	// The publisher of the extension, available publishers
	// can be found by using the Azure CLI.
	Publisher pulumi.StringPtrInput
	// The name of the resource group in which to
	// create the virtual network. Changing this forces a new resource to be
	// created.
	ResourceGroupName pulumi.StringPtrInput
	// The settings passed to the extension, these are
	// specified as a JSON object in a string.
	Settings pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of extension, available types for a publisher can
	// be found using the Azure CLI.
	Type pulumi.StringPtrInput
	// Specifies the version of the extension to
	// use, available versions can be found using the Azure CLI.
	TypeHandlerVersion pulumi.StringPtrInput
	// The resource ID of the virtual machine. This value replaces 
	// `location`, `resourceGroupName` and `virtualMachineName`. Changing this forces a new
	// resource to be created
	VirtualMachineId pulumi.StringPtrInput
	// The name of the virtual machine. Changing
	// this forces a new resource to be created.
	VirtualMachineName pulumi.StringPtrInput
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	// Specifies if the platform deploys
	// the latest minor version update to the `typeHandlerVersion` specified.
	AutoUpgradeMinorVersion *bool `pulumi:"autoUpgradeMinorVersion"`
	// The location where the extension is created. Changing
	// this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the virtual machine extension peering. Changing
	// this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The protectedSettings passed to the
	// extension, like settings, these are specified as a JSON object in a string.
	ProtectedSettings *string `pulumi:"protectedSettings"`
	// The publisher of the extension, available publishers
	// can be found by using the Azure CLI.
	Publisher string `pulumi:"publisher"`
	// The name of the resource group in which to
	// create the virtual network. Changing this forces a new resource to be
	// created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The settings passed to the extension, these are
	// specified as a JSON object in a string.
	Settings *string `pulumi:"settings"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of extension, available types for a publisher can
	// be found using the Azure CLI.
	Type string `pulumi:"type"`
	// Specifies the version of the extension to
	// use, available versions can be found using the Azure CLI.
	TypeHandlerVersion string `pulumi:"typeHandlerVersion"`
	// The resource ID of the virtual machine. This value replaces 
	// `location`, `resourceGroupName` and `virtualMachineName`. Changing this forces a new
	// resource to be created
	VirtualMachineId *string `pulumi:"virtualMachineId"`
	// The name of the virtual machine. Changing
	// this forces a new resource to be created.
	VirtualMachineName *string `pulumi:"virtualMachineName"`
}

// The set of arguments for constructing a Extension resource.
type ExtensionArgs struct {
	// Specifies if the platform deploys
	// the latest minor version update to the `typeHandlerVersion` specified.
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	// The location where the extension is created. Changing
	// this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the virtual machine extension peering. Changing
	// this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The protectedSettings passed to the
	// extension, like settings, these are specified as a JSON object in a string.
	ProtectedSettings pulumi.StringPtrInput
	// The publisher of the extension, available publishers
	// can be found by using the Azure CLI.
	Publisher pulumi.StringInput
	// The name of the resource group in which to
	// create the virtual network. Changing this forces a new resource to be
	// created.
	ResourceGroupName pulumi.StringPtrInput
	// The settings passed to the extension, these are
	// specified as a JSON object in a string.
	Settings pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The type of extension, available types for a publisher can
	// be found using the Azure CLI.
	Type pulumi.StringInput
	// Specifies the version of the extension to
	// use, available versions can be found using the Azure CLI.
	TypeHandlerVersion pulumi.StringInput
	// The resource ID of the virtual machine. This value replaces 
	// `location`, `resourceGroupName` and `virtualMachineName`. Changing this forces a new
	// resource to be created
	VirtualMachineId pulumi.StringPtrInput
	// The name of the virtual machine. Changing
	// this forces a new resource to be created.
	VirtualMachineName pulumi.StringPtrInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

