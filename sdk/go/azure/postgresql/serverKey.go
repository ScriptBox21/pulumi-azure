// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Customer Managed Key for a PostgreSQL Server.
type ServerKey struct {
	pulumi.CustomResourceState

	// The URL to a Key Vault Key.
	KeyVaultKeyId pulumi.StringOutput `pulumi:"keyVaultKeyId"`
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
}

// NewServerKey registers a new resource with the given unique name, arguments, and options.
func NewServerKey(ctx *pulumi.Context,
	name string, args *ServerKeyArgs, opts ...pulumi.ResourceOption) (*ServerKey, error) {
	if args == nil || args.KeyVaultKeyId == nil {
		return nil, errors.New("missing required argument 'KeyVaultKeyId'")
	}
	if args == nil || args.ServerId == nil {
		return nil, errors.New("missing required argument 'ServerId'")
	}
	if args == nil {
		args = &ServerKeyArgs{}
	}
	var resource ServerKey
	err := ctx.RegisterResource("azure:postgresql/serverKey:ServerKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerKey gets an existing ServerKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerKeyState, opts ...pulumi.ResourceOption) (*ServerKey, error) {
	var resource ServerKey
	err := ctx.ReadResource("azure:postgresql/serverKey:ServerKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerKey resources.
type serverKeyState struct {
	// The URL to a Key Vault Key.
	KeyVaultKeyId *string `pulumi:"keyVaultKeyId"`
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId *string `pulumi:"serverId"`
}

type ServerKeyState struct {
	// The URL to a Key Vault Key.
	KeyVaultKeyId pulumi.StringPtrInput
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId pulumi.StringPtrInput
}

func (ServerKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverKeyState)(nil)).Elem()
}

type serverKeyArgs struct {
	// The URL to a Key Vault Key.
	KeyVaultKeyId string `pulumi:"keyVaultKeyId"`
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId string `pulumi:"serverId"`
}

// The set of arguments for constructing a ServerKey resource.
type ServerKeyArgs struct {
	// The URL to a Key Vault Key.
	KeyVaultKeyId pulumi.StringInput
	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerId pulumi.StringInput
}

func (ServerKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverKeyArgs)(nil)).Elem()
}
