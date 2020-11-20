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
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Key Vault.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Key Vault. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceKeyVault registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceKeyVault(ctx *pulumi.Context,
	name string, args *LinkedServiceKeyVaultArgs, opts ...pulumi.ResourceOption) (*LinkedServiceKeyVault, error) {
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.KeyVaultId == nil {
		return nil, errors.New("missing required argument 'KeyVaultId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LinkedServiceKeyVaultArgs{}
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
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
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
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
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
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
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
	// Specifies the name of the Data Factory Linked Service Key Vault. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
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

func (LinkedServiceKeyVault) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceKeyVault)(nil)).Elem()
}

func (i LinkedServiceKeyVault) ToLinkedServiceKeyVaultOutput() LinkedServiceKeyVaultOutput {
	return i.ToLinkedServiceKeyVaultOutputWithContext(context.Background())
}

func (i LinkedServiceKeyVault) ToLinkedServiceKeyVaultOutputWithContext(ctx context.Context) LinkedServiceKeyVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceKeyVaultOutput)
}

type LinkedServiceKeyVaultOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceKeyVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceKeyVaultOutput)(nil)).Elem()
}

func (o LinkedServiceKeyVaultOutput) ToLinkedServiceKeyVaultOutput() LinkedServiceKeyVaultOutput {
	return o
}

func (o LinkedServiceKeyVaultOutput) ToLinkedServiceKeyVaultOutputWithContext(ctx context.Context) LinkedServiceKeyVaultOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceKeyVaultOutput{})
}
