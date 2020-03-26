// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Configures Network Packet Capturing against a Virtual Machine using a Network Watcher.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_packet_capture.html.markdown.
type NetworkPacketCapture struct {
	pulumi.CustomResourceState

	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters NetworkPacketCaptureFilterArrayOutput `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrOutput `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrOutput `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrOutput `pulumi:"maximumCaptureDuration"`
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringOutput `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation NetworkPacketCaptureStorageLocationOutput `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewNetworkPacketCapture registers a new resource with the given unique name, arguments, and options.
func NewNetworkPacketCapture(ctx *pulumi.Context,
	name string, args *NetworkPacketCaptureArgs, opts ...pulumi.ResourceOption) (*NetworkPacketCapture, error) {
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
		args = &NetworkPacketCaptureArgs{}
	}
	var resource NetworkPacketCapture
	err := ctx.RegisterResource("azure:network/networkPacketCapture:NetworkPacketCapture", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPacketCapture gets an existing NetworkPacketCapture resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPacketCapture(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPacketCaptureState, opts ...pulumi.ResourceOption) (*NetworkPacketCapture, error) {
	var resource NetworkPacketCapture
	err := ctx.ReadResource("azure:network/networkPacketCapture:NetworkPacketCapture", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPacketCapture resources.
type networkPacketCaptureState struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters []NetworkPacketCaptureFilter `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket *int `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession *int `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration *int `pulumi:"maximumCaptureDuration"`
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName *string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation *NetworkPacketCaptureStorageLocation `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type NetworkPacketCaptureState struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters NetworkPacketCaptureFilterArrayInput
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrInput
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrInput
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrInput
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringPtrInput
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation NetworkPacketCaptureStorageLocationPtrInput
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringPtrInput
}

func (NetworkPacketCaptureState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPacketCaptureState)(nil)).Elem()
}

type networkPacketCaptureArgs struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters []NetworkPacketCaptureFilter `pulumi:"filters"`
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket *int `pulumi:"maximumBytesPerPacket"`
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession *int `pulumi:"maximumBytesPerSession"`
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration *int `pulumi:"maximumCaptureDuration"`
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation NetworkPacketCaptureStorageLocation `pulumi:"storageLocation"`
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a NetworkPacketCapture resource.
type NetworkPacketCaptureArgs struct {
	// One or more `filter` blocks as defined below. Changing this forces a new resource to be created.
	Filters NetworkPacketCaptureFilterArrayInput
	// The number of bytes captured per packet. The remaining bytes are truncated. Defaults to `0` (Entire Packet Captured). Changing this forces a new resource to be created.
	MaximumBytesPerPacket pulumi.IntPtrInput
	// Maximum size of the capture in Bytes. Defaults to `1073741824` (1GB). Changing this forces a new resource to be created.
	MaximumBytesPerSession pulumi.IntPtrInput
	// The maximum duration of the capture session in seconds. Defaults to `18000` (5 hours). Changing this forces a new resource to be created.
	MaximumCaptureDuration pulumi.IntPtrInput
	// The name to use for this Network Packet Capture. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringInput
	// The name of the resource group in which the Network Watcher exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `storageLocation` block as defined below. Changing this forces a new resource to be created.
	StorageLocation NetworkPacketCaptureStorageLocationInput
	// The ID of the Resource to capture packets from. Changing this forces a new resource to be created.
	TargetResourceId pulumi.StringInput
}

func (NetworkPacketCaptureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPacketCaptureArgs)(nil)).Elem()
}
