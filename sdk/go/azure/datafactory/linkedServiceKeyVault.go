// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between Key Vault and Azure Data Factory.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/keyvault"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := core.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("eastus"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleKeyVault, err := keyvault.NewKeyVault(ctx, "exampleKeyVault", &keyvault.KeyVaultArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			TenantId:          pulumi.String(current.TenantId),
// 			SkuName:           pulumi.String("standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceKeyVault(ctx, "exampleLinkedServiceKeyVault", &datafactory.LinkedServiceKeyVaultArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			KeyVaultId:        exampleKeyVault.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Data Factory Key Vault Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceKeyVault:LinkedServiceKeyVault example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceKeyVault struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service Key Vault.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Key Vault.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Key Vault.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Key Vault.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// The ID the Azure Key Vault resource.
	KeyVaultId pulumi.StringOutput `pulumi:"keyVaultId"`
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Key Vault.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceKeyVault registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceKeyVault(ctx *pulumi.Context,
	name string, args *LinkedServiceKeyVaultArgs, opts ...pulumi.ResourceOption) (*LinkedServiceKeyVault, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.KeyVaultId == nil {
		return nil, errors.New("invalid value for required argument 'KeyVaultId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LinkedServiceKeyVault
	err := ctx.RegisterResource("azure:datafactory/linkedServiceKeyVault:LinkedServiceKeyVault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceKeyVault gets an existing LinkedServiceKeyVault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceKeyVault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceKeyVaultState, opts ...pulumi.ResourceOption) (*LinkedServiceKeyVault, error) {
	var resource LinkedServiceKeyVault
	err := ctx.ReadResource("azure:datafactory/linkedServiceKeyVault:LinkedServiceKeyVault", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceKeyVault resources.
type linkedServiceKeyVaultState struct {
	// A map of additional properties to associate with the Data Factory Linked Service Key Vault.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Key Vault.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Key Vault.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Key Vault.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// The ID the Azure Key Vault resource.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Key Vault.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceKeyVaultState struct {
	// A map of additional properties to associate with the Data Factory Linked Service Key Vault.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service Key Vault.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service Key Vault.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service Key Vault.
	IntegrationRuntimeName pulumi.StringPtrInput
	// The ID the Azure Key Vault resource.
	KeyVaultId pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service Key Vault.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceKeyVaultState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceKeyVaultState)(nil)).Elem()
}

type linkedServiceKeyVaultArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service Key Vault.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Key Vault.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Key Vault.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Key Vault.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// The ID the Azure Key Vault resource.
	KeyVaultId string `pulumi:"keyVaultId"`
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Key Vault.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceKeyVault resource.
type LinkedServiceKeyVaultArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service Key Vault.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service Key Vault.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service Key Vault.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service Key Vault.
	IntegrationRuntimeName pulumi.StringPtrInput
	// The ID the Azure Key Vault resource.
	KeyVaultId pulumi.StringInput
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service Key Vault.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceKeyVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceKeyVaultArgs)(nil)).Elem()
}

type LinkedServiceKeyVaultInput interface {
	pulumi.Input

	ToLinkedServiceKeyVaultOutput() LinkedServiceKeyVaultOutput
	ToLinkedServiceKeyVaultOutputWithContext(ctx context.Context) LinkedServiceKeyVaultOutput
}

func (*LinkedServiceKeyVault) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceKeyVault)(nil))
}

func (i *LinkedServiceKeyVault) ToLinkedServiceKeyVaultOutput() LinkedServiceKeyVaultOutput {
	return i.ToLinkedServiceKeyVaultOutputWithContext(context.Background())
}

func (i *LinkedServiceKeyVault) ToLinkedServiceKeyVaultOutputWithContext(ctx context.Context) LinkedServiceKeyVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceKeyVaultOutput)
}

func (i *LinkedServiceKeyVault) ToLinkedServiceKeyVaultPtrOutput() LinkedServiceKeyVaultPtrOutput {
	return i.ToLinkedServiceKeyVaultPtrOutputWithContext(context.Background())
}

func (i *LinkedServiceKeyVault) ToLinkedServiceKeyVaultPtrOutputWithContext(ctx context.Context) LinkedServiceKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceKeyVaultPtrOutput)
}

type LinkedServiceKeyVaultPtrInput interface {
	pulumi.Input

	ToLinkedServiceKeyVaultPtrOutput() LinkedServiceKeyVaultPtrOutput
	ToLinkedServiceKeyVaultPtrOutputWithContext(ctx context.Context) LinkedServiceKeyVaultPtrOutput
}

type linkedServiceKeyVaultPtrType LinkedServiceKeyVaultArgs

func (*linkedServiceKeyVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceKeyVault)(nil))
}

func (i *linkedServiceKeyVaultPtrType) ToLinkedServiceKeyVaultPtrOutput() LinkedServiceKeyVaultPtrOutput {
	return i.ToLinkedServiceKeyVaultPtrOutputWithContext(context.Background())
}

func (i *linkedServiceKeyVaultPtrType) ToLinkedServiceKeyVaultPtrOutputWithContext(ctx context.Context) LinkedServiceKeyVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceKeyVaultPtrOutput)
}

// LinkedServiceKeyVaultArrayInput is an input type that accepts LinkedServiceKeyVaultArray and LinkedServiceKeyVaultArrayOutput values.
// You can construct a concrete instance of `LinkedServiceKeyVaultArrayInput` via:
//
//          LinkedServiceKeyVaultArray{ LinkedServiceKeyVaultArgs{...} }
type LinkedServiceKeyVaultArrayInput interface {
	pulumi.Input

	ToLinkedServiceKeyVaultArrayOutput() LinkedServiceKeyVaultArrayOutput
	ToLinkedServiceKeyVaultArrayOutputWithContext(context.Context) LinkedServiceKeyVaultArrayOutput
}

type LinkedServiceKeyVaultArray []LinkedServiceKeyVaultInput

func (LinkedServiceKeyVaultArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LinkedServiceKeyVault)(nil))
}

func (i LinkedServiceKeyVaultArray) ToLinkedServiceKeyVaultArrayOutput() LinkedServiceKeyVaultArrayOutput {
	return i.ToLinkedServiceKeyVaultArrayOutputWithContext(context.Background())
}

func (i LinkedServiceKeyVaultArray) ToLinkedServiceKeyVaultArrayOutputWithContext(ctx context.Context) LinkedServiceKeyVaultArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceKeyVaultArrayOutput)
}

// LinkedServiceKeyVaultMapInput is an input type that accepts LinkedServiceKeyVaultMap and LinkedServiceKeyVaultMapOutput values.
// You can construct a concrete instance of `LinkedServiceKeyVaultMapInput` via:
//
//          LinkedServiceKeyVaultMap{ "key": LinkedServiceKeyVaultArgs{...} }
type LinkedServiceKeyVaultMapInput interface {
	pulumi.Input

	ToLinkedServiceKeyVaultMapOutput() LinkedServiceKeyVaultMapOutput
	ToLinkedServiceKeyVaultMapOutputWithContext(context.Context) LinkedServiceKeyVaultMapOutput
}

type LinkedServiceKeyVaultMap map[string]LinkedServiceKeyVaultInput

func (LinkedServiceKeyVaultMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LinkedServiceKeyVault)(nil))
}

func (i LinkedServiceKeyVaultMap) ToLinkedServiceKeyVaultMapOutput() LinkedServiceKeyVaultMapOutput {
	return i.ToLinkedServiceKeyVaultMapOutputWithContext(context.Background())
}

func (i LinkedServiceKeyVaultMap) ToLinkedServiceKeyVaultMapOutputWithContext(ctx context.Context) LinkedServiceKeyVaultMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceKeyVaultMapOutput)
}

type LinkedServiceKeyVaultOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceKeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceKeyVault)(nil))
}

func (o LinkedServiceKeyVaultOutput) ToLinkedServiceKeyVaultOutput() LinkedServiceKeyVaultOutput {
	return o
}

func (o LinkedServiceKeyVaultOutput) ToLinkedServiceKeyVaultOutputWithContext(ctx context.Context) LinkedServiceKeyVaultOutput {
	return o
}

func (o LinkedServiceKeyVaultOutput) ToLinkedServiceKeyVaultPtrOutput() LinkedServiceKeyVaultPtrOutput {
	return o.ToLinkedServiceKeyVaultPtrOutputWithContext(context.Background())
}

func (o LinkedServiceKeyVaultOutput) ToLinkedServiceKeyVaultPtrOutputWithContext(ctx context.Context) LinkedServiceKeyVaultPtrOutput {
	return o.ApplyT(func(v LinkedServiceKeyVault) *LinkedServiceKeyVault {
		return &v
	}).(LinkedServiceKeyVaultPtrOutput)
}

type LinkedServiceKeyVaultPtrOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceKeyVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceKeyVault)(nil))
}

func (o LinkedServiceKeyVaultPtrOutput) ToLinkedServiceKeyVaultPtrOutput() LinkedServiceKeyVaultPtrOutput {
	return o
}

func (o LinkedServiceKeyVaultPtrOutput) ToLinkedServiceKeyVaultPtrOutputWithContext(ctx context.Context) LinkedServiceKeyVaultPtrOutput {
	return o
}

type LinkedServiceKeyVaultArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceKeyVaultArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServiceKeyVault)(nil))
}

func (o LinkedServiceKeyVaultArrayOutput) ToLinkedServiceKeyVaultArrayOutput() LinkedServiceKeyVaultArrayOutput {
	return o
}

func (o LinkedServiceKeyVaultArrayOutput) ToLinkedServiceKeyVaultArrayOutputWithContext(ctx context.Context) LinkedServiceKeyVaultArrayOutput {
	return o
}

func (o LinkedServiceKeyVaultArrayOutput) Index(i pulumi.IntInput) LinkedServiceKeyVaultOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServiceKeyVault {
		return vs[0].([]LinkedServiceKeyVault)[vs[1].(int)]
	}).(LinkedServiceKeyVaultOutput)
}

type LinkedServiceKeyVaultMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceKeyVaultMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServiceKeyVault)(nil))
}

func (o LinkedServiceKeyVaultMapOutput) ToLinkedServiceKeyVaultMapOutput() LinkedServiceKeyVaultMapOutput {
	return o
}

func (o LinkedServiceKeyVaultMapOutput) ToLinkedServiceKeyVaultMapOutputWithContext(ctx context.Context) LinkedServiceKeyVaultMapOutput {
	return o
}

func (o LinkedServiceKeyVaultMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceKeyVaultOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServiceKeyVault {
		return vs[0].(map[string]LinkedServiceKeyVault)[vs[1].(string)]
	}).(LinkedServiceKeyVaultOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceKeyVaultOutput{})
	pulumi.RegisterOutputType(LinkedServiceKeyVaultPtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceKeyVaultArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceKeyVaultMapOutput{})
}
