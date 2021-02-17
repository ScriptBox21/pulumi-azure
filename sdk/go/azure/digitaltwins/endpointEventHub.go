// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Digital Twins Event Hub Endpoint.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/digitaltwins"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/eventhub"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleInstance, err := digitaltwins.NewInstance(ctx, "exampleInstance", &digitaltwins.InstanceArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHubNamespace, err := eventhub.NewEventHubNamespace(ctx, "exampleEventHubNamespace", &eventhub.EventHubNamespaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("Standard"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleEventHub, err := eventhub.NewEventHub(ctx, "exampleEventHub", &eventhub.EventHubArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PartitionCount:    pulumi.Int(2),
// 			MessageRetention:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleAuthorizationRule, err := eventhub.NewAuthorizationRule(ctx, "exampleAuthorizationRule", &eventhub.AuthorizationRuleArgs{
// 			NamespaceName:     exampleEventHubNamespace.Name,
// 			EventhubName:      exampleEventHub.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Listen:            pulumi.Bool(false),
// 			Send:              pulumi.Bool(true),
// 			Manage:            pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitaltwins.NewEndpointEventHub(ctx, "exampleEndpointEventHub", &digitaltwins.EndpointEventHubArgs{
// 			DigitalTwinsId:                    exampleInstance.ID(),
// 			EventhubPrimaryConnectionString:   exampleAuthorizationRule.PrimaryConnectionString,
// 			EventhubSecondaryConnectionString: exampleAuthorizationRule.SecondaryConnectionString,
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
// Digital Twins Eventhub Endpoints can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:digitaltwins/endpointEventHub:EndpointEventHub example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
// ```
type EndpointEventHub struct {
	pulumi.CustomResourceState

	// The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
	DeadLetterStorageSecret pulumi.StringPtrOutput `pulumi:"deadLetterStorageSecret"`
	// The resource ID of the Digital Twins Instance. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	DigitalTwinsId pulumi.StringOutput `pulumi:"digitalTwinsId"`
	// The primary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubPrimaryConnectionString pulumi.StringOutput `pulumi:"eventhubPrimaryConnectionString"`
	// The secondary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubSecondaryConnectionString pulumi.StringOutput `pulumi:"eventhubSecondaryConnectionString"`
	// The name which should be used for this Digital Twins Event Hub Endpoint. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewEndpointEventHub registers a new resource with the given unique name, arguments, and options.
func NewEndpointEventHub(ctx *pulumi.Context,
	name string, args *EndpointEventHubArgs, opts ...pulumi.ResourceOption) (*EndpointEventHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DigitalTwinsId == nil {
		return nil, errors.New("invalid value for required argument 'DigitalTwinsId'")
	}
	if args.EventhubPrimaryConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'EventhubPrimaryConnectionString'")
	}
	if args.EventhubSecondaryConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'EventhubSecondaryConnectionString'")
	}
	var resource EndpointEventHub
	err := ctx.RegisterResource("azure:digitaltwins/endpointEventHub:EndpointEventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointEventHub gets an existing EndpointEventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointEventHubState, opts ...pulumi.ResourceOption) (*EndpointEventHub, error) {
	var resource EndpointEventHub
	err := ctx.ReadResource("azure:digitaltwins/endpointEventHub:EndpointEventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointEventHub resources.
type endpointEventHubState struct {
	// The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
	DeadLetterStorageSecret *string `pulumi:"deadLetterStorageSecret"`
	// The resource ID of the Digital Twins Instance. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	DigitalTwinsId *string `pulumi:"digitalTwinsId"`
	// The primary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubPrimaryConnectionString *string `pulumi:"eventhubPrimaryConnectionString"`
	// The secondary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubSecondaryConnectionString *string `pulumi:"eventhubSecondaryConnectionString"`
	// The name which should be used for this Digital Twins Event Hub Endpoint. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	Name *string `pulumi:"name"`
}

type EndpointEventHubState struct {
	// The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
	DeadLetterStorageSecret pulumi.StringPtrInput
	// The resource ID of the Digital Twins Instance. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	DigitalTwinsId pulumi.StringPtrInput
	// The primary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubPrimaryConnectionString pulumi.StringPtrInput
	// The secondary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubSecondaryConnectionString pulumi.StringPtrInput
	// The name which should be used for this Digital Twins Event Hub Endpoint. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	Name pulumi.StringPtrInput
}

func (EndpointEventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointEventHubState)(nil)).Elem()
}

type endpointEventHubArgs struct {
	// The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
	DeadLetterStorageSecret *string `pulumi:"deadLetterStorageSecret"`
	// The resource ID of the Digital Twins Instance. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	DigitalTwinsId string `pulumi:"digitalTwinsId"`
	// The primary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubPrimaryConnectionString string `pulumi:"eventhubPrimaryConnectionString"`
	// The secondary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubSecondaryConnectionString string `pulumi:"eventhubSecondaryConnectionString"`
	// The name which should be used for this Digital Twins Event Hub Endpoint. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a EndpointEventHub resource.
type EndpointEventHubArgs struct {
	// The storage secret of the dead-lettering, whose format is `https://<storageAccountname>.blob.core.windows.net/<containerName>?<SASToken>`. When an endpoint can't deliver an event within a certain time period or after trying to deliver the event a certain number of times, it can send the undelivered event to a storage account.
	DeadLetterStorageSecret pulumi.StringPtrInput
	// The resource ID of the Digital Twins Instance. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	DigitalTwinsId pulumi.StringInput
	// The primary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubPrimaryConnectionString pulumi.StringInput
	// The secondary connection string of the Event Hub Authorization Rule with a minimum of `send` permission.
	EventhubSecondaryConnectionString pulumi.StringInput
	// The name which should be used for this Digital Twins Event Hub Endpoint. Changing this forces a new Digital Twins Event Hub Endpoint to be created.
	Name pulumi.StringPtrInput
}

func (EndpointEventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointEventHubArgs)(nil)).Elem()
}

type EndpointEventHubInput interface {
	pulumi.Input

	ToEndpointEventHubOutput() EndpointEventHubOutput
	ToEndpointEventHubOutputWithContext(ctx context.Context) EndpointEventHubOutput
}

func (*EndpointEventHub) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointEventHub)(nil))
}

func (i *EndpointEventHub) ToEndpointEventHubOutput() EndpointEventHubOutput {
	return i.ToEndpointEventHubOutputWithContext(context.Background())
}

func (i *EndpointEventHub) ToEndpointEventHubOutputWithContext(ctx context.Context) EndpointEventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointEventHubOutput)
}

func (i *EndpointEventHub) ToEndpointEventHubPtrOutput() EndpointEventHubPtrOutput {
	return i.ToEndpointEventHubPtrOutputWithContext(context.Background())
}

func (i *EndpointEventHub) ToEndpointEventHubPtrOutputWithContext(ctx context.Context) EndpointEventHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointEventHubPtrOutput)
}

type EndpointEventHubPtrInput interface {
	pulumi.Input

	ToEndpointEventHubPtrOutput() EndpointEventHubPtrOutput
	ToEndpointEventHubPtrOutputWithContext(ctx context.Context) EndpointEventHubPtrOutput
}

type endpointEventHubPtrType EndpointEventHubArgs

func (*endpointEventHubPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointEventHub)(nil))
}

func (i *endpointEventHubPtrType) ToEndpointEventHubPtrOutput() EndpointEventHubPtrOutput {
	return i.ToEndpointEventHubPtrOutputWithContext(context.Background())
}

func (i *endpointEventHubPtrType) ToEndpointEventHubPtrOutputWithContext(ctx context.Context) EndpointEventHubPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointEventHubPtrOutput)
}

// EndpointEventHubArrayInput is an input type that accepts EndpointEventHubArray and EndpointEventHubArrayOutput values.
// You can construct a concrete instance of `EndpointEventHubArrayInput` via:
//
//          EndpointEventHubArray{ EndpointEventHubArgs{...} }
type EndpointEventHubArrayInput interface {
	pulumi.Input

	ToEndpointEventHubArrayOutput() EndpointEventHubArrayOutput
	ToEndpointEventHubArrayOutputWithContext(context.Context) EndpointEventHubArrayOutput
}

type EndpointEventHubArray []EndpointEventHubInput

func (EndpointEventHubArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*EndpointEventHub)(nil))
}

func (i EndpointEventHubArray) ToEndpointEventHubArrayOutput() EndpointEventHubArrayOutput {
	return i.ToEndpointEventHubArrayOutputWithContext(context.Background())
}

func (i EndpointEventHubArray) ToEndpointEventHubArrayOutputWithContext(ctx context.Context) EndpointEventHubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointEventHubArrayOutput)
}

// EndpointEventHubMapInput is an input type that accepts EndpointEventHubMap and EndpointEventHubMapOutput values.
// You can construct a concrete instance of `EndpointEventHubMapInput` via:
//
//          EndpointEventHubMap{ "key": EndpointEventHubArgs{...} }
type EndpointEventHubMapInput interface {
	pulumi.Input

	ToEndpointEventHubMapOutput() EndpointEventHubMapOutput
	ToEndpointEventHubMapOutputWithContext(context.Context) EndpointEventHubMapOutput
}

type EndpointEventHubMap map[string]EndpointEventHubInput

func (EndpointEventHubMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*EndpointEventHub)(nil))
}

func (i EndpointEventHubMap) ToEndpointEventHubMapOutput() EndpointEventHubMapOutput {
	return i.ToEndpointEventHubMapOutputWithContext(context.Background())
}

func (i EndpointEventHubMap) ToEndpointEventHubMapOutputWithContext(ctx context.Context) EndpointEventHubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointEventHubMapOutput)
}

type EndpointEventHubOutput struct {
	*pulumi.OutputState
}

func (EndpointEventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointEventHub)(nil))
}

func (o EndpointEventHubOutput) ToEndpointEventHubOutput() EndpointEventHubOutput {
	return o
}

func (o EndpointEventHubOutput) ToEndpointEventHubOutputWithContext(ctx context.Context) EndpointEventHubOutput {
	return o
}

func (o EndpointEventHubOutput) ToEndpointEventHubPtrOutput() EndpointEventHubPtrOutput {
	return o.ToEndpointEventHubPtrOutputWithContext(context.Background())
}

func (o EndpointEventHubOutput) ToEndpointEventHubPtrOutputWithContext(ctx context.Context) EndpointEventHubPtrOutput {
	return o.ApplyT(func(v EndpointEventHub) *EndpointEventHub {
		return &v
	}).(EndpointEventHubPtrOutput)
}

type EndpointEventHubPtrOutput struct {
	*pulumi.OutputState
}

func (EndpointEventHubPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointEventHub)(nil))
}

func (o EndpointEventHubPtrOutput) ToEndpointEventHubPtrOutput() EndpointEventHubPtrOutput {
	return o
}

func (o EndpointEventHubPtrOutput) ToEndpointEventHubPtrOutputWithContext(ctx context.Context) EndpointEventHubPtrOutput {
	return o
}

type EndpointEventHubArrayOutput struct{ *pulumi.OutputState }

func (EndpointEventHubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointEventHub)(nil))
}

func (o EndpointEventHubArrayOutput) ToEndpointEventHubArrayOutput() EndpointEventHubArrayOutput {
	return o
}

func (o EndpointEventHubArrayOutput) ToEndpointEventHubArrayOutputWithContext(ctx context.Context) EndpointEventHubArrayOutput {
	return o
}

func (o EndpointEventHubArrayOutput) Index(i pulumi.IntInput) EndpointEventHubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointEventHub {
		return vs[0].([]EndpointEventHub)[vs[1].(int)]
	}).(EndpointEventHubOutput)
}

type EndpointEventHubMapOutput struct{ *pulumi.OutputState }

func (EndpointEventHubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EndpointEventHub)(nil))
}

func (o EndpointEventHubMapOutput) ToEndpointEventHubMapOutput() EndpointEventHubMapOutput {
	return o
}

func (o EndpointEventHubMapOutput) ToEndpointEventHubMapOutputWithContext(ctx context.Context) EndpointEventHubMapOutput {
	return o
}

func (o EndpointEventHubMapOutput) MapIndex(k pulumi.StringInput) EndpointEventHubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EndpointEventHub {
		return vs[0].(map[string]EndpointEventHub)[vs[1].(string)]
	}).(EndpointEventHubOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointEventHubOutput{})
	pulumi.RegisterOutputType(EndpointEventHubPtrOutput{})
	pulumi.RegisterOutputType(EndpointEventHubArrayOutput{})
	pulumi.RegisterOutputType(EndpointEventHubMapOutput{})
}
