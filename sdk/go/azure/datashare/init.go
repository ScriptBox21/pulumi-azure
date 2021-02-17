// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datashare

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
	case "azure:datashare/account:Account":
		r, err = NewAccount(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datashare/datasetBlobStorage:DatasetBlobStorage":
		r, err = NewDatasetBlobStorage(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datashare/datasetDataLakeGen1:DatasetDataLakeGen1":
		r, err = NewDatasetDataLakeGen1(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datashare/datasetDataLakeGen2:DatasetDataLakeGen2":
		r, err = NewDatasetDataLakeGen2(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datashare/datasetKustoCluster:DatasetKustoCluster":
		r, err = NewDatasetKustoCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datashare/datasetKustoDatabase:DatasetKustoDatabase":
		r, err = NewDatasetKustoDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure:datashare/share:Share":
		r, err = NewShare(ctx, name, nil, pulumi.URN_(urn))
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
		"datashare/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datashare/datasetBlobStorage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datashare/datasetDataLakeGen1",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datashare/datasetDataLakeGen2",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datashare/datasetKustoCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datashare/datasetKustoDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"datashare/share",
		&module{version},
	)
}
