// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dns

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Enables you to manage DNS CNAME Records within Azure DNS.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/dns_cname_record.html.markdown.
type CNameRecord struct {
	pulumi.CustomResourceState

	// The FQDN of the DNS CName Record.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The name of the DNS CNAME Record.
	Name pulumi.StringOutput `pulumi:"name"`
	// The target of the CNAME.
	Record pulumi.StringPtrOutput `pulumi:"record"`
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

// NewCNameRecord registers a new resource with the given unique name, arguments, and options.
func NewCNameRecord(ctx *pulumi.Context,
	name string, args *CNameRecordArgs, opts ...pulumi.ResourceOption) (*CNameRecord, error) {
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
		args = &CNameRecordArgs{}
	}
	var resource CNameRecord
	err := ctx.RegisterResource("azure:dns/cNameRecord:CNameRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCNameRecord gets an existing CNameRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCNameRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CNameRecordState, opts ...pulumi.ResourceOption) (*CNameRecord, error) {
	var resource CNameRecord
	err := ctx.ReadResource("azure:dns/cNameRecord:CNameRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CNameRecord resources.
type cnameRecordState struct {
	// The FQDN of the DNS CName Record.
	Fqdn *string `pulumi:"fqdn"`
	// The name of the DNS CNAME Record.
	Name *string `pulumi:"name"`
	// The target of the CNAME.
	Record *string `pulumi:"record"`
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

type CNameRecordState struct {
	// The FQDN of the DNS CName Record.
	Fqdn pulumi.StringPtrInput
	// The name of the DNS CNAME Record.
	Name pulumi.StringPtrInput
	// The target of the CNAME.
	Record pulumi.StringPtrInput
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

func (CNameRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameRecordState)(nil)).Elem()
}

type cnameRecordArgs struct {
	// The name of the DNS CNAME Record.
	Name *string `pulumi:"name"`
	// The target of the CNAME.
	Record *string `pulumi:"record"`
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

// The set of arguments for constructing a CNameRecord resource.
type CNameRecordArgs struct {
	// The name of the DNS CNAME Record.
	Name pulumi.StringPtrInput
	// The target of the CNAME.
	Record pulumi.StringPtrInput
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

func (CNameRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameRecordArgs)(nil)).Elem()
}
