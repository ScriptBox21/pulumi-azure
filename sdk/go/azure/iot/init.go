// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure:iot/consumerGroup:ConsumerGroup":
		r, err = NewConsumerGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/dpsSharedAccessPolicy:DpsSharedAccessPolicy":
		r, err = NewDpsSharedAccessPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/endpointEventhub:EndpointEventhub":
		r, err = NewEndpointEventhub(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/endpointServicebusQueue:EndpointServicebusQueue":
		r, err = NewEndpointServicebusQueue(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/endpointServicebusTopic:EndpointServicebusTopic":
		r, err = NewEndpointServicebusTopic(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/endpointStorageContainer:EndpointStorageContainer":
		r, err = NewEndpointStorageContainer(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/enrichment:Enrichment":
		r, err = NewEnrichment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/fallbackRoute:FallbackRoute":
		r, err = NewFallbackRoute(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/ioTHub:IoTHub":
		r, err = NewIoTHub(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/iotHubCertificate:IotHubCertificate":
		r, err = NewIotHubCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/iotHubDps:IotHubDps":
		r, err = NewIotHubDps(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/route:Route":
		r, err = NewRoute(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/securityDeviceGroup:SecurityDeviceGroup":
		r, err = NewSecurityDeviceGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/securitySolution:SecuritySolution":
		r, err = NewSecuritySolution(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/sharedAccessPolicy:SharedAccessPolicy":
		r, err = NewSharedAccessPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/timeSeriesInsightsAccessPolicy:TimeSeriesInsightsAccessPolicy":
		r, err = NewTimeSeriesInsightsAccessPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/timeSeriesInsightsGen2Environment:TimeSeriesInsightsGen2Environment":
		r, err = NewTimeSeriesInsightsGen2Environment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/timeSeriesInsightsReferenceDataSet:TimeSeriesInsightsReferenceDataSet":
		r, err = NewTimeSeriesInsightsReferenceDataSet(ctx, name, nil, pulumi.URN_(urn))
	case "azure:iot/timeSeriesInsightsStandardEnvironment:TimeSeriesInsightsStandardEnvironment":
		r, err = NewTimeSeriesInsightsStandardEnvironment(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"iot/consumerGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/dpsSharedAccessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/endpointEventhub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/endpointServicebusQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/endpointServicebusTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/endpointStorageContainer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/enrichment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/fallbackRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/ioTHub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/iotHubCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/iotHubDps",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/securityDeviceGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/securitySolution",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/sharedAccessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsAccessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsGen2Environment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsReferenceDataSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"iot/timeSeriesInsightsStandardEnvironment",
		&module{version},
	)
}
