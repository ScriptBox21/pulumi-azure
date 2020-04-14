// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Configures Packet Capturing against a Virtual Machine using a Network Watcher.
//
// > **NOTE:** This resource has been deprecated in favour of the `network.NetworkConnectionMonitor` resource and will be removed in the next major version of the AzureRM Provider. The new resource shares the same fields as this one, and information on migrating across can be found in this guide.
type PacketCapture struct {
	pulumi.CustomResourceState

	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters PacketCaptureFilterArrayOutput `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrOutput `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrOutput `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrOutput `pulumi:"maximumCaptureDuration"`
	// The name to use for this Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringOutput `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation PacketCaptureStorageLocationOutput `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewPacketCapture registers a new resource with the given unique name, arguments, and options.
func NewPacketCapture(ctx *pulumi.Context,
	name string, args *PacketCaptureArgs, opts ...pulumi.ResourceOption) (*PacketCapture, error) {
	if args == nil || args.NetworkWatcherName == nil {
		return nil, errors.New("missing required argument 'NetworkWatcherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageLocation == nil {
		return nil, errors.New("missing required argument 'StorageLocation'")
	}
	if args == nil || args.TargetResourceId == nil {
		return nil, errors.New("missing required argument 'TargetResourceId'")
	}
	if args == nil {
		args = &PacketCaptureArgs{}
	}
	var resource PacketCapture
	err := ctx.RegisterResource("azure:network/packetCapture:PacketCapture", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPacketCapture gets an existing PacketCapture resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPacketCapture(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketCaptureState, opts ...pulumi.ResourceOption) (*PacketCapture, error) {
	var resource PacketCapture
	err := ctx.ReadResource("azure:network/packetCapture:PacketCapture", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PacketCapture resources.
type packetCaptureState struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters []PacketCaptureFilter `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket *int `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession *int `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration *int `pulumi:"maximumCaptureDuration"`
	// The name to use for this Packet Capture. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName *string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation *PacketCaptureStorageLocation `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type PacketCaptureState struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters PacketCaptureFilterArrayInput
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrInput
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrInput
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrInput
	// The name to use for this Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringPtrInput
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation PacketCaptureStorageLocationPtrInput
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringPtrInput
}

func (PacketCaptureState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCaptureState)(nil)).Elem()
}

type packetCaptureArgs struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters []PacketCaptureFilter `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket *int `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession *int `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration *int `pulumi:"maximumCaptureDuration"`
	// The name to use for this Packet Capture. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation PacketCaptureStorageLocation `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a PacketCapture resource.
type PacketCaptureArgs struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters PacketCaptureFilterArrayInput
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrInput
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrInput
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrInput
	// The name to use for this Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringInput
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation PacketCaptureStorageLocationInput
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringInput
}

func (PacketCaptureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packetCaptureArgs)(nil)).Elem()
}
