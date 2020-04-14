// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Enables you to manage DNS AAAA Records within Azure DNS.
type AaaaRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS AAAA Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS AAAA Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of IPv4 Addresses. Conflicts with `targetResourceId`.
	Records pulumi.StringArrayOutput `pulumi:"records"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId pulumi.StringPtrOutput `pulumi:"targetResourceId"`
	Ttl              pulumi.IntOutput       `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewAaaaRecord registers a new resource with the given unique name, arguments, and options.
func NewAaaaRecord(ctx *pulumi.Context,
	name string, args *AaaaRecordArgs, opts ...pulumi.ResourceOption) (*AaaaRecord, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Ttl == nil {
		return nil, errors.New("missing required argument 'Ttl'")
	}
	if args == nil || args.ZoneName == nil {
		return nil, errors.New("missing required argument 'ZoneName'")
	}
	if args == nil {
		args = &AaaaRecordArgs{}
	}
	var resource AaaaRecord
	err := ctx.RegisterResource("azure:dns/aaaaRecord:AaaaRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAaaaRecord gets an existing AaaaRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAaaaRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AaaaRecordState, opts ...pulumi.ResourceOption) (*AaaaRecord, error) {
	var resource AaaaRecord
	err := ctx.ReadResource("azure:dns/aaaaRecord:AaaaRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AaaaRecord resources.
type aaaaRecordState struct {
	// The FQDN of the DNS AAAA Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS AAAA Record.
	Name *string `pulumi:"name"`
	// List of IPv4 Addresses. Conflicts with `targetResourceId`.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId *string `pulumi:"targetResourceId"`
	Ttl              *int    `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName *string `pulumi:"zoneName"`
}

type AaaaRecordState struct {
	// The FQDN of the DNS AAAA Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS AAAA Record.
	Name pulumi.StringPtrInput
	// List of IPv4 Addresses. Conflicts with `targetResourceId`.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId pulumi.StringPtrInput
	Ttl              pulumi.IntPtrInput
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringPtrInput
}

func (AaaaRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*aaaaRecordState)(nil)).Elem()
}

type aaaaRecordArgs struct {
	// The name of the DNS AAAA Record.
	Name *string `pulumi:"name"`
	// List of IPv4 Addresses. Conflicts with `targetResourceId`.
	Records []string `pulumi:"records"`
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId *string `pulumi:"targetResourceId"`
	Ttl              int     `pulumi:"ttl"`
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a AaaaRecord resource.
type AaaaRecordArgs struct {
	// The name of the DNS AAAA Record.
	Name pulumi.StringPtrInput
	// List of IPv4 Addresses. Conflicts with `targetResourceId`.
	Records pulumi.StringArrayInput
	// Specifies the resource group where the DNS Zone (parent resource) exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The Azure resource id of the target object. Conflicts with `records`
	TargetResourceId pulumi.StringPtrInput
	Ttl              pulumi.IntInput
	// Specifies the DNS Zone where the resource exists. Changing this forces a new resource to be created.
	ZoneName pulumi.StringInput
}

func (AaaaRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aaaaRecordArgs)(nil)).Elem()
}
