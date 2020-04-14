// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Redis Cache
func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	var rv LookupCacheResult
	err := ctx.Invoke("azure:redis/getCache:getCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCache.
type LookupCacheArgs struct {
	// The name of the Redis cache
	Name string `pulumi:"name"`
	// The name of the resource group the Redis cache instance is located in.
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Zones             []string `pulumi:"zones"`
}

// A collection of values returned by getCache.
type LookupCacheResult struct {
	// The size of the Redis Cache deployed.
	Capacity int `pulumi:"capacity"`
	// Whether the SSL port is enabled.
	EnableNonSslPort bool `pulumi:"enableNonSslPort"`
	// The SKU family/pricing group used. Possible values are `C` (for Basic/Standard SKU family) and `P` (for `Premium`)
	Family string `pulumi:"family"`
	// The Hostname of the Redis Instance
	Hostname string `pulumi:"hostname"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The location of the Redis Cache.
	Location string `pulumi:"location"`
	// The minimum TLS version.
	MinimumTlsVersion string `pulumi:"minimumTlsVersion"`
	Name              string `pulumi:"name"`
	// A list of `patchSchedule` blocks as defined below - only available for Premium SKU's.
	PatchSchedules []GetCachePatchSchedule `pulumi:"patchSchedules"`
	// The non-SSL Port of the Redis Instance
	Port int `pulumi:"port"`
	// The Primary Access Key for the Redis Instance
	PrimaryAccessKey string `pulumi:"primaryAccessKey"`
	// The primary connection string of the Redis Instance.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	PrivateStaticIpAddress  string `pulumi:"privateStaticIpAddress"`
	// A `redisConfiguration` block as defined below.
	RedisConfigurations []GetCacheRedisConfiguration `pulumi:"redisConfigurations"`
	ResourceGroupName   string                       `pulumi:"resourceGroupName"`
	// The Secondary Access Key for the Redis Instance
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
	// The secondary connection string of the Redis Instance.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	ShardCount                int    `pulumi:"shardCount"`
	// The SKU of Redis used. Possible values are `Basic`, `Standard` and `Premium`.
	SkuName string `pulumi:"skuName"`
	// The SSL Port of the Redis Instance
	SslPort  int               `pulumi:"sslPort"`
	SubnetId string            `pulumi:"subnetId"`
	Tags     map[string]string `pulumi:"tags"`
	Zones    []string          `pulumi:"zones"`
}
