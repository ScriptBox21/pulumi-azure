// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages the association between a Network Interface and a Application Gateway's Backend Address Pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
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
// 		exampleVirtualNetwork, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
// 			AddressSpaces: pulumi.StringArray{
// 				pulumi.String("10.0.0.0/16"),
// 			},
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		frontend, err := network.NewSubnet(ctx, "frontend", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.1.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewSubnet(ctx, "backend", &network.SubnetArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			VirtualNetworkName: exampleVirtualNetwork.Name,
// 			AddressPrefixes: pulumi.StringArray{
// 				pulumi.String("10.0.2.0/24"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		examplePublicIp, err := network.NewPublicIp(ctx, "examplePublicIp", &network.PublicIpArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			AllocationMethod:  pulumi.String("Dynamic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		network, err := network.NewApplicationGateway(ctx, "network", &network.ApplicationGatewayArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			Sku: &network.ApplicationGatewaySkuArgs{
// 				Name:     pulumi.String("Standard_Small"),
// 				Tier:     pulumi.String("Standard"),
// 				Capacity: pulumi.Int(2),
// 			},
// 			GatewayIpConfigurations: network.ApplicationGatewayGatewayIpConfigurationArray{
// 				&network.ApplicationGatewayGatewayIpConfigurationArgs{
// 					Name:     pulumi.String("my-gateway-ip-configuration"),
// 					SubnetId: frontend.ID(),
// 				},
// 			},
// 			FrontendPorts: network.ApplicationGatewayFrontendPortArray{
// 				&network.ApplicationGatewayFrontendPortArgs{
// 					Name: pulumi.String(frontendPortName),
// 					Port: pulumi.Int(80),
// 				},
// 			},
// 			FrontendIpConfigurations: network.ApplicationGatewayFrontendIpConfigurationArray{
// 				&network.ApplicationGatewayFrontendIpConfigurationArgs{
// 					Name:              pulumi.String(frontendIpConfigurationName),
// 					PublicIpAddressId: examplePublicIp.ID(),
// 				},
// 			},
// 			BackendAddressPools: network.ApplicationGatewayBackendAddressPoolArray{
// 				&network.ApplicationGatewayBackendAddressPoolArgs{
// 					Name: pulumi.String(backendAddressPoolName),
// 				},
// 			},
// 			BackendHttpSettings: network.ApplicationGatewayBackendHttpSettingArray{
// 				&network.ApplicationGatewayBackendHttpSettingArgs{
// 					Name:                pulumi.String(httpSettingName),
// 					CookieBasedAffinity: pulumi.String("Disabled"),
// 					Port:                pulumi.Int(80),
// 					Protocol:            pulumi.String("Http"),
// 					RequestTimeout:      pulumi.Int(1),
// 				},
// 			},
// 			HttpListeners: network.ApplicationGatewayHttpListenerArray{
// 				&network.ApplicationGatewayHttpListenerArgs{
// 					Name:                        pulumi.String(listenerName),
// 					FrontendIpConfigurationName: pulumi.String(frontendIpConfigurationName),
// 					FrontendPortName:            pulumi.String(frontendPortName),
// 					Protocol:                    pulumi.String("Http"),
// 				},
// 			},
// 			RequestRoutingRules: network.ApplicationGatewayRequestRoutingRuleArray{
// 				&network.ApplicationGatewayRequestRoutingRuleArgs{
// 					Name:                    pulumi.String(requestRoutingRuleName),
// 					RuleType:                pulumi.String("Basic"),
// 					HttpListenerName:        pulumi.String(listenerName),
// 					BackendAddressPoolName:  pulumi.String(backendAddressPoolName),
// 					BackendHttpSettingsName: pulumi.String(httpSettingName),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleNetworkInterface, err := network.NewNetworkInterface(ctx, "exampleNetworkInterface", &network.NetworkInterfaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
// 				&network.NetworkInterfaceIpConfigurationArgs{
// 					Name:                       pulumi.String("testconfiguration1"),
// 					SubnetId:                   frontend.ID(),
// 					PrivateIpAddressAllocation: pulumi.String("Dynamic"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(ctx, "exampleNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation", &network.NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs{
// 			NetworkInterfaceId:  exampleNetworkInterface.ID(),
// 			IpConfigurationName: pulumi.String("testconfiguration1"),
// 			BackendAddressPoolId: pulumi.String(network.BackendAddressPools.ApplyT(func(backendAddressPools []network.ApplicationGatewayBackendAddressPool) (string, error) {
// 				return backendAddressPools[0].Id, nil
// 			}).(pulumi.StringOutput)),
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
// Associations between Network Interfaces and Application Gateway Backend Address Pools can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation association1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/microsoft.network/networkInterfaces/nic1/ipConfigurations/example|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/applicationGateways/gateway1/backendAddressPools/pool1
// ```
type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringOutput `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringOutput `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
}

// NewNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(ctx *pulumi.Context,
	name string, args *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, error) {
	if args == nil || args.BackendAddressPoolId == nil {
		return nil, errors.New("missing required argument 'BackendAddressPoolId'")
	}
	if args == nil || args.IpConfigurationName == nil {
		return nil, errors.New("missing required argument 'IpConfigurationName'")
	}
	if args == nil || args.NetworkInterfaceId == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceId'")
	}
	if args == nil {
		args = &NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs{}
	}
	var resource NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation
	err := ctx.RegisterResource("azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation gets an existing NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceApplicationGatewayBackendAddressPoolAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation, error) {
	var resource NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation
	err := ctx.ReadResource("azure:network/networkInterfaceApplicationGatewayBackendAddressPoolAssociation:NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation resources.
type networkInterfaceApplicationGatewayBackendAddressPoolAssociationState struct {
	// The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId *string `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName *string `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
}

type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState struct {
	// The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringPtrInput
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringPtrInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringPtrInput
}

func (NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceApplicationGatewayBackendAddressPoolAssociationState)(nil)).Elem()
}

type networkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs struct {
	// The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId string `pulumi:"backendAddressPoolId"`
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName string `pulumi:"ipConfigurationName"`
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// The set of arguments for constructing a NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation resource.
type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs struct {
	// The ID of the Application Gateway's Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	BackendAddressPoolId pulumi.StringInput
	// The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
	IpConfigurationName pulumi.StringInput
	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceId pulumi.StringInput
}

func (NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceApplicationGatewayBackendAddressPoolAssociationArgs)(nil)).Elem()
}

type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationInput interface {
	pulumi.Input

	ToNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput() NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput
	ToNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutputWithContext(ctx context.Context) NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput
}

func (NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation)(nil)).Elem()
}

func (i NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) ToNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput() NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput {
	return i.ToNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutputWithContext(context.Background())
}

func (i NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation) ToNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutputWithContext(ctx context.Context) NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput)
}

type NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput struct {
	*pulumi.OutputState
}

func (NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput)(nil)).Elem()
}

func (o NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput) ToNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput() NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput {
	return o
}

func (o NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput) ToNetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutputWithContext(ctx context.Context) NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationOutput{})
}
