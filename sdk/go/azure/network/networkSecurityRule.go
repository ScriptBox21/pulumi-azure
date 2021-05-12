// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Network Security Rule.
//
// > **NOTE on Network Security Groups and Network Security Rules:** This provider currently
// provides both a standalone Network Security Rule resource, and allows for Network Security Rules to be defined in-line within the Network Security Group resource.
// At this time you cannot use a Network Security Group with in-line Network Security Rules in conjunction with any Network Security Rule resources. Doing so will cause a conflict of rule settings and will overwrite rules.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/network"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		exampleNetworkSecurityGroup, err := network.NewNetworkSecurityGroup(ctx, "exampleNetworkSecurityGroup", &network.NetworkSecurityGroupArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = network.NewNetworkSecurityRule(ctx, "exampleNetworkSecurityRule", &network.NetworkSecurityRuleArgs{
// 			Priority:                 pulumi.Int(100),
// 			Direction:                pulumi.String("Outbound"),
// 			Access:                   pulumi.String("Allow"),
// 			Protocol:                 pulumi.String("Tcp"),
// 			SourcePortRange:          pulumi.String("*"),
// 			DestinationPortRange:     pulumi.String("*"),
// 			SourceAddressPrefix:      pulumi.String("*"),
// 			DestinationAddressPrefix: pulumi.String("*"),
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			NetworkSecurityGroupName: exampleNetworkSecurityGroup.Name,
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
// Network Security Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:network/networkSecurityRule:NetworkSecurityRule rule1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/networkSecurityGroups/mySecurityGroup/securityRules/rule1
// ```
type NetworkSecurityRule struct {
	pulumi.CustomResourceState

	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access pulumi.StringOutput `pulumi:"access"`
	// A description for this rule. Restricted to 140 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destinationAddressPrefixes` is not specified.
	DestinationAddressPrefix pulumi.StringPtrOutput `pulumi:"destinationAddressPrefix"`
	// List of destination address prefixes. Tags may not be used. This is required if `destinationAddressPrefix` is not specified.
	DestinationAddressPrefixes pulumi.StringArrayOutput `pulumi:"destinationAddressPrefixes"`
	// A List of destination Application Security Group ID's
	DestinationApplicationSecurityGroupIds pulumi.StringPtrOutput `pulumi:"destinationApplicationSecurityGroupIds"`
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destinationPortRanges` is not specified.
	DestinationPortRange pulumi.StringPtrOutput `pulumi:"destinationPortRange"`
	// List of destination ports or port ranges. This is required if `destinationPortRange` is not specified.
	DestinationPortRanges pulumi.StringArrayOutput `pulumi:"destinationPortRanges"`
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
	NetworkSecurityGroupName pulumi.StringOutput `pulumi:"networkSecurityGroupName"`
	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, `Esp`, `Ah` or `*` (which matches all).
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `sourceAddressPrefixes` is not specified.
	SourceAddressPrefix pulumi.StringPtrOutput `pulumi:"sourceAddressPrefix"`
	// List of source address prefixes. Tags may not be used. This is required if `sourceAddressPrefix` is not specified.
	SourceAddressPrefixes pulumi.StringArrayOutput `pulumi:"sourceAddressPrefixes"`
	// A List of source Application Security Group ID's
	SourceApplicationSecurityGroupIds pulumi.StringPtrOutput `pulumi:"sourceApplicationSecurityGroupIds"`
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `sourcePortRanges` is not specified.
	SourcePortRange pulumi.StringPtrOutput `pulumi:"sourcePortRange"`
	// List of source ports or port ranges. This is required if `sourcePortRange` is not specified.
	SourcePortRanges pulumi.StringArrayOutput `pulumi:"sourcePortRanges"`
}

// NewNetworkSecurityRule registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityRule(ctx *pulumi.Context,
	name string, args *NetworkSecurityRuleArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Access == nil {
		return nil, errors.New("invalid value for required argument 'Access'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.NetworkSecurityGroupName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSecurityGroupName'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource NetworkSecurityRule
	err := ctx.RegisterResource("azure:network/networkSecurityRule:NetworkSecurityRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityRule gets an existing NetworkSecurityRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityRuleState, opts ...pulumi.ResourceOption) (*NetworkSecurityRule, error) {
	var resource NetworkSecurityRule
	err := ctx.ReadResource("azure:network/networkSecurityRule:NetworkSecurityRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityRule resources.
type networkSecurityRuleState struct {
	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access *string `pulumi:"access"`
	// A description for this rule. Restricted to 140 characters.
	Description *string `pulumi:"description"`
	// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destinationAddressPrefixes` is not specified.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// List of destination address prefixes. Tags may not be used. This is required if `destinationAddressPrefix` is not specified.
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	// A List of destination Application Security Group ID's
	DestinationApplicationSecurityGroupIds *string `pulumi:"destinationApplicationSecurityGroupIds"`
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destinationPortRanges` is not specified.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// List of destination ports or port ranges. This is required if `destinationPortRange` is not specified.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction *string `pulumi:"direction"`
	// The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
	NetworkSecurityGroupName *string `pulumi:"networkSecurityGroupName"`
	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `pulumi:"priority"`
	// Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, `Esp`, `Ah` or `*` (which matches all).
	Protocol *string `pulumi:"protocol"`
	// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `sourceAddressPrefixes` is not specified.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// List of source address prefixes. Tags may not be used. This is required if `sourceAddressPrefix` is not specified.
	SourceAddressPrefixes []string `pulumi:"sourceAddressPrefixes"`
	// A List of source Application Security Group ID's
	SourceApplicationSecurityGroupIds *string `pulumi:"sourceApplicationSecurityGroupIds"`
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `sourcePortRanges` is not specified.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// List of source ports or port ranges. This is required if `sourcePortRange` is not specified.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
}

type NetworkSecurityRuleState struct {
	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access pulumi.StringPtrInput
	// A description for this rule. Restricted to 140 characters.
	Description pulumi.StringPtrInput
	// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destinationAddressPrefixes` is not specified.
	DestinationAddressPrefix pulumi.StringPtrInput
	// List of destination address prefixes. Tags may not be used. This is required if `destinationAddressPrefix` is not specified.
	DestinationAddressPrefixes pulumi.StringArrayInput
	// A List of destination Application Security Group ID's
	DestinationApplicationSecurityGroupIds pulumi.StringPtrInput
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destinationPortRanges` is not specified.
	DestinationPortRange pulumi.StringPtrInput
	// List of destination ports or port ranges. This is required if `destinationPortRange` is not specified.
	DestinationPortRanges pulumi.StringArrayInput
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction pulumi.StringPtrInput
	// The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
	NetworkSecurityGroupName pulumi.StringPtrInput
	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntPtrInput
	// Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, `Esp`, `Ah` or `*` (which matches all).
	Protocol pulumi.StringPtrInput
	// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `sourceAddressPrefixes` is not specified.
	SourceAddressPrefix pulumi.StringPtrInput
	// List of source address prefixes. Tags may not be used. This is required if `sourceAddressPrefix` is not specified.
	SourceAddressPrefixes pulumi.StringArrayInput
	// A List of source Application Security Group ID's
	SourceApplicationSecurityGroupIds pulumi.StringPtrInput
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `sourcePortRanges` is not specified.
	SourcePortRange pulumi.StringPtrInput
	// List of source ports or port ranges. This is required if `sourcePortRange` is not specified.
	SourcePortRanges pulumi.StringArrayInput
}

func (NetworkSecurityRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityRuleState)(nil)).Elem()
}

type networkSecurityRuleArgs struct {
	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access string `pulumi:"access"`
	// A description for this rule. Restricted to 140 characters.
	Description *string `pulumi:"description"`
	// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destinationAddressPrefixes` is not specified.
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	// List of destination address prefixes. Tags may not be used. This is required if `destinationAddressPrefix` is not specified.
	DestinationAddressPrefixes []string `pulumi:"destinationAddressPrefixes"`
	// A List of destination Application Security Group ID's
	DestinationApplicationSecurityGroupIds *string `pulumi:"destinationApplicationSecurityGroupIds"`
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destinationPortRanges` is not specified.
	DestinationPortRange *string `pulumi:"destinationPortRange"`
	// List of destination ports or port ranges. This is required if `destinationPortRange` is not specified.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction string `pulumi:"direction"`
	// The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
	NetworkSecurityGroupName string `pulumi:"networkSecurityGroupName"`
	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority int `pulumi:"priority"`
	// Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, `Esp`, `Ah` or `*` (which matches all).
	Protocol string `pulumi:"protocol"`
	// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `sourceAddressPrefixes` is not specified.
	SourceAddressPrefix *string `pulumi:"sourceAddressPrefix"`
	// List of source address prefixes. Tags may not be used. This is required if `sourceAddressPrefix` is not specified.
	SourceAddressPrefixes []string `pulumi:"sourceAddressPrefixes"`
	// A List of source Application Security Group ID's
	SourceApplicationSecurityGroupIds *string `pulumi:"sourceApplicationSecurityGroupIds"`
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `sourcePortRanges` is not specified.
	SourcePortRange *string `pulumi:"sourcePortRange"`
	// List of source ports or port ranges. This is required if `sourcePortRange` is not specified.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
}

// The set of arguments for constructing a NetworkSecurityRule resource.
type NetworkSecurityRuleArgs struct {
	// Specifies whether network traffic is allowed or denied. Possible values are `Allow` and `Deny`.
	Access pulumi.StringInput
	// A description for this rule. Restricted to 140 characters.
	Description pulumi.StringPtrInput
	// CIDR or destination IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. Besides, it also supports all available Service Tags like ‘Sql.WestEurope‘, ‘Storage.EastUS‘, etc. You can list the available service tags with the cli: ```shell az network list-service-tags --location westcentralus```. For further information please see [Azure CLI - az network list-service-tags](https://docs.microsoft.com/en-us/cli/azure/network?view=azure-cli-latest#az-network-list-service-tags). This is required if `destinationAddressPrefixes` is not specified.
	DestinationAddressPrefix pulumi.StringPtrInput
	// List of destination address prefixes. Tags may not be used. This is required if `destinationAddressPrefix` is not specified.
	DestinationAddressPrefixes pulumi.StringArrayInput
	// A List of destination Application Security Group ID's
	DestinationApplicationSecurityGroupIds pulumi.StringPtrInput
	// Destination Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `destinationPortRanges` is not specified.
	DestinationPortRange pulumi.StringPtrInput
	// List of destination ports or port ranges. This is required if `destinationPortRange` is not specified.
	DestinationPortRanges pulumi.StringArrayInput
	// The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are `Inbound` and `Outbound`.
	Direction pulumi.StringInput
	// The name of the security rule. This needs to be unique across all Rules in the Network Security Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Security Group that we want to attach the rule to. Changing this forces a new resource to be created.
	NetworkSecurityGroupName pulumi.StringInput
	// Specifies the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority pulumi.IntInput
	// Network protocol this rule applies to. Possible values include `Tcp`, `Udp`, `Icmp`, `Esp`, `Ah` or `*` (which matches all).
	Protocol pulumi.StringInput
	// The name of the resource group in which to create the Network Security Rule. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// CIDR or source IP range or * to match any IP. Tags such as ‘VirtualNetwork’, ‘AzureLoadBalancer’ and ‘Internet’ can also be used. This is required if `sourceAddressPrefixes` is not specified.
	SourceAddressPrefix pulumi.StringPtrInput
	// List of source address prefixes. Tags may not be used. This is required if `sourceAddressPrefix` is not specified.
	SourceAddressPrefixes pulumi.StringArrayInput
	// A List of source Application Security Group ID's
	SourceApplicationSecurityGroupIds pulumi.StringPtrInput
	// Source Port or Range. Integer or range between `0` and `65535` or `*` to match any. This is required if `sourcePortRanges` is not specified.
	SourcePortRange pulumi.StringPtrInput
	// List of source ports or port ranges. This is required if `sourcePortRange` is not specified.
	SourcePortRanges pulumi.StringArrayInput
}

func (NetworkSecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityRuleArgs)(nil)).Elem()
}

type NetworkSecurityRuleInput interface {
	pulumi.Input

	ToNetworkSecurityRuleOutput() NetworkSecurityRuleOutput
	ToNetworkSecurityRuleOutputWithContext(ctx context.Context) NetworkSecurityRuleOutput
}

func (*NetworkSecurityRule) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityRule)(nil))
}

func (i *NetworkSecurityRule) ToNetworkSecurityRuleOutput() NetworkSecurityRuleOutput {
	return i.ToNetworkSecurityRuleOutputWithContext(context.Background())
}

func (i *NetworkSecurityRule) ToNetworkSecurityRuleOutputWithContext(ctx context.Context) NetworkSecurityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRuleOutput)
}

func (i *NetworkSecurityRule) ToNetworkSecurityRulePtrOutput() NetworkSecurityRulePtrOutput {
	return i.ToNetworkSecurityRulePtrOutputWithContext(context.Background())
}

func (i *NetworkSecurityRule) ToNetworkSecurityRulePtrOutputWithContext(ctx context.Context) NetworkSecurityRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRulePtrOutput)
}

type NetworkSecurityRulePtrInput interface {
	pulumi.Input

	ToNetworkSecurityRulePtrOutput() NetworkSecurityRulePtrOutput
	ToNetworkSecurityRulePtrOutputWithContext(ctx context.Context) NetworkSecurityRulePtrOutput
}

type networkSecurityRulePtrType NetworkSecurityRuleArgs

func (*networkSecurityRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityRule)(nil))
}

func (i *networkSecurityRulePtrType) ToNetworkSecurityRulePtrOutput() NetworkSecurityRulePtrOutput {
	return i.ToNetworkSecurityRulePtrOutputWithContext(context.Background())
}

func (i *networkSecurityRulePtrType) ToNetworkSecurityRulePtrOutputWithContext(ctx context.Context) NetworkSecurityRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRulePtrOutput)
}

// NetworkSecurityRuleArrayInput is an input type that accepts NetworkSecurityRuleArray and NetworkSecurityRuleArrayOutput values.
// You can construct a concrete instance of `NetworkSecurityRuleArrayInput` via:
//
//          NetworkSecurityRuleArray{ NetworkSecurityRuleArgs{...} }
type NetworkSecurityRuleArrayInput interface {
	pulumi.Input

	ToNetworkSecurityRuleArrayOutput() NetworkSecurityRuleArrayOutput
	ToNetworkSecurityRuleArrayOutputWithContext(context.Context) NetworkSecurityRuleArrayOutput
}

type NetworkSecurityRuleArray []NetworkSecurityRuleInput

func (NetworkSecurityRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NetworkSecurityRule)(nil))
}

func (i NetworkSecurityRuleArray) ToNetworkSecurityRuleArrayOutput() NetworkSecurityRuleArrayOutput {
	return i.ToNetworkSecurityRuleArrayOutputWithContext(context.Background())
}

func (i NetworkSecurityRuleArray) ToNetworkSecurityRuleArrayOutputWithContext(ctx context.Context) NetworkSecurityRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRuleArrayOutput)
}

// NetworkSecurityRuleMapInput is an input type that accepts NetworkSecurityRuleMap and NetworkSecurityRuleMapOutput values.
// You can construct a concrete instance of `NetworkSecurityRuleMapInput` via:
//
//          NetworkSecurityRuleMap{ "key": NetworkSecurityRuleArgs{...} }
type NetworkSecurityRuleMapInput interface {
	pulumi.Input

	ToNetworkSecurityRuleMapOutput() NetworkSecurityRuleMapOutput
	ToNetworkSecurityRuleMapOutputWithContext(context.Context) NetworkSecurityRuleMapOutput
}

type NetworkSecurityRuleMap map[string]NetworkSecurityRuleInput

func (NetworkSecurityRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NetworkSecurityRule)(nil))
}

func (i NetworkSecurityRuleMap) ToNetworkSecurityRuleMapOutput() NetworkSecurityRuleMapOutput {
	return i.ToNetworkSecurityRuleMapOutputWithContext(context.Background())
}

func (i NetworkSecurityRuleMap) ToNetworkSecurityRuleMapOutputWithContext(ctx context.Context) NetworkSecurityRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityRuleMapOutput)
}

type NetworkSecurityRuleOutput struct {
	*pulumi.OutputState
}

func (NetworkSecurityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityRule)(nil))
}

func (o NetworkSecurityRuleOutput) ToNetworkSecurityRuleOutput() NetworkSecurityRuleOutput {
	return o
}

func (o NetworkSecurityRuleOutput) ToNetworkSecurityRuleOutputWithContext(ctx context.Context) NetworkSecurityRuleOutput {
	return o
}

func (o NetworkSecurityRuleOutput) ToNetworkSecurityRulePtrOutput() NetworkSecurityRulePtrOutput {
	return o.ToNetworkSecurityRulePtrOutputWithContext(context.Background())
}

func (o NetworkSecurityRuleOutput) ToNetworkSecurityRulePtrOutputWithContext(ctx context.Context) NetworkSecurityRulePtrOutput {
	return o.ApplyT(func(v NetworkSecurityRule) *NetworkSecurityRule {
		return &v
	}).(NetworkSecurityRulePtrOutput)
}

type NetworkSecurityRulePtrOutput struct {
	*pulumi.OutputState
}

func (NetworkSecurityRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityRule)(nil))
}

func (o NetworkSecurityRulePtrOutput) ToNetworkSecurityRulePtrOutput() NetworkSecurityRulePtrOutput {
	return o
}

func (o NetworkSecurityRulePtrOutput) ToNetworkSecurityRulePtrOutputWithContext(ctx context.Context) NetworkSecurityRulePtrOutput {
	return o
}

type NetworkSecurityRuleArrayOutput struct{ *pulumi.OutputState }

func (NetworkSecurityRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkSecurityRule)(nil))
}

func (o NetworkSecurityRuleArrayOutput) ToNetworkSecurityRuleArrayOutput() NetworkSecurityRuleArrayOutput {
	return o
}

func (o NetworkSecurityRuleArrayOutput) ToNetworkSecurityRuleArrayOutputWithContext(ctx context.Context) NetworkSecurityRuleArrayOutput {
	return o
}

func (o NetworkSecurityRuleArrayOutput) Index(i pulumi.IntInput) NetworkSecurityRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkSecurityRule {
		return vs[0].([]NetworkSecurityRule)[vs[1].(int)]
	}).(NetworkSecurityRuleOutput)
}

type NetworkSecurityRuleMapOutput struct{ *pulumi.OutputState }

func (NetworkSecurityRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NetworkSecurityRule)(nil))
}

func (o NetworkSecurityRuleMapOutput) ToNetworkSecurityRuleMapOutput() NetworkSecurityRuleMapOutput {
	return o
}

func (o NetworkSecurityRuleMapOutput) ToNetworkSecurityRuleMapOutputWithContext(ctx context.Context) NetworkSecurityRuleMapOutput {
	return o
}

func (o NetworkSecurityRuleMapOutput) MapIndex(k pulumi.StringInput) NetworkSecurityRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NetworkSecurityRule {
		return vs[0].(map[string]NetworkSecurityRule)[vs[1].(string)]
	}).(NetworkSecurityRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkSecurityRuleOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRulePtrOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleArrayOutput{})
	pulumi.RegisterOutputType(NetworkSecurityRuleMapOutput{})
}
