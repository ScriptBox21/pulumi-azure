// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a API Management Service API Diagnostics Logs.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/appinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
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
// 		exampleInsights, err := appinsights.NewInsights(ctx, "exampleInsights", &appinsights.InsightsArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationType:   pulumi.String("web"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleService, err := apimanagement.NewService(ctx, "exampleService", &apimanagement.ServiceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			PublisherName:     pulumi.String("My Company"),
// 			PublisherEmail:    pulumi.String("company@mycompany.io"),
// 			SkuName:           pulumi.String("Developer_1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleApi, err := apimanagement.NewApi(ctx, "exampleApi", &apimanagement.ApiArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApiManagementName: exampleService.Name,
// 			Revision:          pulumi.String("1"),
// 			DisplayName:       pulumi.String("Example API"),
// 			Path:              pulumi.String("example"),
// 			Protocols: pulumi.StringArray{
// 				pulumi.String("https"),
// 			},
// 			Import: &apimanagement.ApiImportArgs{
// 				ContentFormat: pulumi.String("swagger-link-json"),
// 				ContentValue:  pulumi.String("http://conferenceapi.azurewebsites.net/?format=json"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleLogger, err := apimanagement.NewLogger(ctx, "exampleLogger", &apimanagement.LoggerArgs{
// 			ApiManagementName: exampleService.Name,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ApplicationInsights: &apimanagement.LoggerApplicationInsightsArgs{
// 				InstrumentationKey: exampleInsights.InstrumentationKey,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apimanagement.NewApiDiagnostic(ctx, "exampleApiDiagnostic", &apimanagement.ApiDiagnosticArgs{
// 			ResourceGroupName:       exampleResourceGroup.Name,
// 			ApiManagementName:       exampleService.Name,
// 			ApiName:                 exampleApi.Name,
// 			ApiManagementLoggerId:   exampleLogger.ID(),
// 			SamplingPercentage:      pulumi.Float64(5),
// 			AlwaysLogErrors:         pulumi.Bool(true),
// 			LogClientIp:             pulumi.Bool(true),
// 			Verbosity:               pulumi.String("Verbose"),
// 			HttpCorrelationProtocol: pulumi.String("W3C"),
// 			FrontendRequest: &apimanagement.ApiDiagnosticFrontendRequestArgs{
// 				BodyBytes: pulumi.Int(32),
// 				HeadersToLogs: pulumi.StringArray{
// 					pulumi.String("content-type"),
// 					pulumi.String("accept"),
// 					pulumi.String("origin"),
// 				},
// 			},
// 			FrontendResponse: &apimanagement.ApiDiagnosticFrontendResponseArgs{
// 				BodyBytes: pulumi.Int(32),
// 				HeadersToLogs: pulumi.StringArray{
// 					pulumi.String("content-type"),
// 					pulumi.String("content-length"),
// 					pulumi.String("origin"),
// 				},
// 			},
// 			BackendRequest: &apimanagement.ApiDiagnosticBackendRequestArgs{
// 				BodyBytes: pulumi.Int(32),
// 				HeadersToLogs: pulumi.StringArray{
// 					pulumi.String("content-type"),
// 					pulumi.String("accept"),
// 					pulumi.String("origin"),
// 				},
// 			},
// 			BackendResponse: &apimanagement.ApiDiagnosticBackendResponseArgs{
// 				BodyBytes: pulumi.Int(32),
// 				HeadersToLogs: pulumi.StringArray{
// 					pulumi.String("content-type"),
// 					pulumi.String("content-length"),
// 					pulumi.String("origin"),
// 				},
// 			},
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
// API Management Service API Diagnostics Logs can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:apimanagement/apiDiagnostic:ApiDiagnostic example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/instance1/apis/api1/diagnostics/diagnostic1/loggers/logger1
// ```
type ApiDiagnostic struct {
	pulumi.CustomResourceState

	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors pulumi.BoolOutput `pulumi:"alwaysLogErrors"`
	// The ID (name) of the Diagnostics Logger.
	ApiManagementLoggerId pulumi.StringOutput `pulumi:"apiManagementLoggerId"`
	// The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiName pulumi.StringOutput `pulumi:"apiName"`
	// A `backendRequest` block as defined below.
	BackendRequest ApiDiagnosticBackendRequestOutput `pulumi:"backendRequest"`
	// A `backendResponse` block as defined below.
	BackendResponse ApiDiagnosticBackendResponseOutput `pulumi:"backendResponse"`
	// A `frontendRequest` block as defined below.
	FrontendRequest ApiDiagnosticFrontendRequestOutput `pulumi:"frontendRequest"`
	// A `frontendResponse` block as defined below.
	FrontendResponse ApiDiagnosticFrontendResponseOutput `pulumi:"frontendResponse"`
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol pulumi.StringOutput `pulumi:"httpCorrelationProtocol"`
	// Identifier of the Diagnostics Logs. Possible values are `applicationinsights` and `azuremonitor`. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Log client IP address.
	LogClientIp pulumi.BoolOutput `pulumi:"logClientIp"`
	// The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage pulumi.Float64Output `pulumi:"samplingPercentage"`
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity pulumi.StringOutput `pulumi:"verbosity"`
}

// NewApiDiagnostic registers a new resource with the given unique name, arguments, and options.
func NewApiDiagnostic(ctx *pulumi.Context,
	name string, args *ApiDiagnosticArgs, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiManagementLoggerId == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementLoggerId'")
	}
	if args.ApiManagementName == nil {
		return nil, errors.New("invalid value for required argument 'ApiManagementName'")
	}
	if args.ApiName == nil {
		return nil, errors.New("invalid value for required argument 'ApiName'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource ApiDiagnostic
	err := ctx.RegisterResource("azure:apimanagement/apiDiagnostic:ApiDiagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiDiagnostic gets an existing ApiDiagnostic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiDiagnosticState, opts ...pulumi.ResourceOption) (*ApiDiagnostic, error) {
	var resource ApiDiagnostic
	err := ctx.ReadResource("azure:apimanagement/apiDiagnostic:ApiDiagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiDiagnostic resources.
type apiDiagnosticState struct {
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors *bool `pulumi:"alwaysLogErrors"`
	// The ID (name) of the Diagnostics Logger.
	ApiManagementLoggerId *string `pulumi:"apiManagementLoggerId"`
	// The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiName *string `pulumi:"apiName"`
	// A `backendRequest` block as defined below.
	BackendRequest *ApiDiagnosticBackendRequest `pulumi:"backendRequest"`
	// A `backendResponse` block as defined below.
	BackendResponse *ApiDiagnosticBackendResponse `pulumi:"backendResponse"`
	// A `frontendRequest` block as defined below.
	FrontendRequest *ApiDiagnosticFrontendRequest `pulumi:"frontendRequest"`
	// A `frontendResponse` block as defined below.
	FrontendResponse *ApiDiagnosticFrontendResponse `pulumi:"frontendResponse"`
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// Identifier of the Diagnostics Logs. Possible values are `applicationinsights` and `azuremonitor`. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	Identifier *string `pulumi:"identifier"`
	// Log client IP address.
	LogClientIp *bool `pulumi:"logClientIp"`
	// The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage *float64 `pulumi:"samplingPercentage"`
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity *string `pulumi:"verbosity"`
}

type ApiDiagnosticState struct {
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors pulumi.BoolPtrInput
	// The ID (name) of the Diagnostics Logger.
	ApiManagementLoggerId pulumi.StringPtrInput
	// The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiManagementName pulumi.StringPtrInput
	// The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiName pulumi.StringPtrInput
	// A `backendRequest` block as defined below.
	BackendRequest ApiDiagnosticBackendRequestPtrInput
	// A `backendResponse` block as defined below.
	BackendResponse ApiDiagnosticBackendResponsePtrInput
	// A `frontendRequest` block as defined below.
	FrontendRequest ApiDiagnosticFrontendRequestPtrInput
	// A `frontendResponse` block as defined below.
	FrontendResponse ApiDiagnosticFrontendResponsePtrInput
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol pulumi.StringPtrInput
	// Identifier of the Diagnostics Logs. Possible values are `applicationinsights` and `azuremonitor`. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	Identifier pulumi.StringPtrInput
	// Log client IP address.
	LogClientIp pulumi.BoolPtrInput
	// The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage pulumi.Float64PtrInput
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity pulumi.StringPtrInput
}

func (ApiDiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticState)(nil)).Elem()
}

type apiDiagnosticArgs struct {
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors *bool `pulumi:"alwaysLogErrors"`
	// The ID (name) of the Diagnostics Logger.
	ApiManagementLoggerId string `pulumi:"apiManagementLoggerId"`
	// The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiName string `pulumi:"apiName"`
	// A `backendRequest` block as defined below.
	BackendRequest *ApiDiagnosticBackendRequest `pulumi:"backendRequest"`
	// A `backendResponse` block as defined below.
	BackendResponse *ApiDiagnosticBackendResponse `pulumi:"backendResponse"`
	// A `frontendRequest` block as defined below.
	FrontendRequest *ApiDiagnosticFrontendRequest `pulumi:"frontendRequest"`
	// A `frontendResponse` block as defined below.
	FrontendResponse *ApiDiagnosticFrontendResponse `pulumi:"frontendResponse"`
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol *string `pulumi:"httpCorrelationProtocol"`
	// Identifier of the Diagnostics Logs. Possible values are `applicationinsights` and `azuremonitor`. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	Identifier string `pulumi:"identifier"`
	// Log client IP address.
	LogClientIp *bool `pulumi:"logClientIp"`
	// The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage *float64 `pulumi:"samplingPercentage"`
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity *string `pulumi:"verbosity"`
}

// The set of arguments for constructing a ApiDiagnostic resource.
type ApiDiagnosticArgs struct {
	// Always log errors. Send telemetry if there is an erroneous condition, regardless of sampling settings.
	AlwaysLogErrors pulumi.BoolPtrInput
	// The ID (name) of the Diagnostics Logger.
	ApiManagementLoggerId pulumi.StringInput
	// The name of the API Management Service instance. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiManagementName pulumi.StringInput
	// The name of the API on which to configure the Diagnostics Logs. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ApiName pulumi.StringInput
	// A `backendRequest` block as defined below.
	BackendRequest ApiDiagnosticBackendRequestPtrInput
	// A `backendResponse` block as defined below.
	BackendResponse ApiDiagnosticBackendResponsePtrInput
	// A `frontendRequest` block as defined below.
	FrontendRequest ApiDiagnosticFrontendRequestPtrInput
	// A `frontendResponse` block as defined below.
	FrontendResponse ApiDiagnosticFrontendResponsePtrInput
	// The HTTP Correlation Protocol to use. Possible values are `None`, `Legacy` or `W3C`.
	HttpCorrelationProtocol pulumi.StringPtrInput
	// Identifier of the Diagnostics Logs. Possible values are `applicationinsights` and `azuremonitor`. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	Identifier pulumi.StringInput
	// Log client IP address.
	LogClientIp pulumi.BoolPtrInput
	// The name of the Resource Group where the API Management Service API Diagnostics Logs should exist. Changing this forces a new API Management Service API Diagnostics Logs to be created.
	ResourceGroupName pulumi.StringInput
	// Sampling (%). For high traffic APIs, please read this [documentation](https://docs.microsoft.com/azure/api-management/api-management-howto-app-insights#performance-implications-and-log-sampling) to understand performance implications and log sampling. Valid values are between `0.0` and `100.0`.
	SamplingPercentage pulumi.Float64PtrInput
	// Logging verbosity. Possible values are `verbose`, `information` or `error`.
	Verbosity pulumi.StringPtrInput
}

func (ApiDiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiDiagnosticArgs)(nil)).Elem()
}

type ApiDiagnosticInput interface {
	pulumi.Input

	ToApiDiagnosticOutput() ApiDiagnosticOutput
	ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput
}

func (*ApiDiagnostic) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDiagnostic)(nil))
}

func (i *ApiDiagnostic) ToApiDiagnosticOutput() ApiDiagnosticOutput {
	return i.ToApiDiagnosticOutputWithContext(context.Background())
}

func (i *ApiDiagnostic) ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticOutput)
}

func (i *ApiDiagnostic) ToApiDiagnosticPtrOutput() ApiDiagnosticPtrOutput {
	return i.ToApiDiagnosticPtrOutputWithContext(context.Background())
}

func (i *ApiDiagnostic) ToApiDiagnosticPtrOutputWithContext(ctx context.Context) ApiDiagnosticPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticPtrOutput)
}

type ApiDiagnosticPtrInput interface {
	pulumi.Input

	ToApiDiagnosticPtrOutput() ApiDiagnosticPtrOutput
	ToApiDiagnosticPtrOutputWithContext(ctx context.Context) ApiDiagnosticPtrOutput
}

type apiDiagnosticPtrType ApiDiagnosticArgs

func (*apiDiagnosticPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDiagnostic)(nil))
}

func (i *apiDiagnosticPtrType) ToApiDiagnosticPtrOutput() ApiDiagnosticPtrOutput {
	return i.ToApiDiagnosticPtrOutputWithContext(context.Background())
}

func (i *apiDiagnosticPtrType) ToApiDiagnosticPtrOutputWithContext(ctx context.Context) ApiDiagnosticPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticPtrOutput)
}

// ApiDiagnosticArrayInput is an input type that accepts ApiDiagnosticArray and ApiDiagnosticArrayOutput values.
// You can construct a concrete instance of `ApiDiagnosticArrayInput` via:
//
//          ApiDiagnosticArray{ ApiDiagnosticArgs{...} }
type ApiDiagnosticArrayInput interface {
	pulumi.Input

	ToApiDiagnosticArrayOutput() ApiDiagnosticArrayOutput
	ToApiDiagnosticArrayOutputWithContext(context.Context) ApiDiagnosticArrayOutput
}

type ApiDiagnosticArray []ApiDiagnosticInput

func (ApiDiagnosticArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ApiDiagnostic)(nil))
}

func (i ApiDiagnosticArray) ToApiDiagnosticArrayOutput() ApiDiagnosticArrayOutput {
	return i.ToApiDiagnosticArrayOutputWithContext(context.Background())
}

func (i ApiDiagnosticArray) ToApiDiagnosticArrayOutputWithContext(ctx context.Context) ApiDiagnosticArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticArrayOutput)
}

// ApiDiagnosticMapInput is an input type that accepts ApiDiagnosticMap and ApiDiagnosticMapOutput values.
// You can construct a concrete instance of `ApiDiagnosticMapInput` via:
//
//          ApiDiagnosticMap{ "key": ApiDiagnosticArgs{...} }
type ApiDiagnosticMapInput interface {
	pulumi.Input

	ToApiDiagnosticMapOutput() ApiDiagnosticMapOutput
	ToApiDiagnosticMapOutputWithContext(context.Context) ApiDiagnosticMapOutput
}

type ApiDiagnosticMap map[string]ApiDiagnosticInput

func (ApiDiagnosticMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ApiDiagnostic)(nil))
}

func (i ApiDiagnosticMap) ToApiDiagnosticMapOutput() ApiDiagnosticMapOutput {
	return i.ToApiDiagnosticMapOutputWithContext(context.Background())
}

func (i ApiDiagnosticMap) ToApiDiagnosticMapOutputWithContext(ctx context.Context) ApiDiagnosticMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDiagnosticMapOutput)
}

type ApiDiagnosticOutput struct {
	*pulumi.OutputState
}

func (ApiDiagnosticOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDiagnostic)(nil))
}

func (o ApiDiagnosticOutput) ToApiDiagnosticOutput() ApiDiagnosticOutput {
	return o
}

func (o ApiDiagnosticOutput) ToApiDiagnosticOutputWithContext(ctx context.Context) ApiDiagnosticOutput {
	return o
}

func (o ApiDiagnosticOutput) ToApiDiagnosticPtrOutput() ApiDiagnosticPtrOutput {
	return o.ToApiDiagnosticPtrOutputWithContext(context.Background())
}

func (o ApiDiagnosticOutput) ToApiDiagnosticPtrOutputWithContext(ctx context.Context) ApiDiagnosticPtrOutput {
	return o.ApplyT(func(v ApiDiagnostic) *ApiDiagnostic {
		return &v
	}).(ApiDiagnosticPtrOutput)
}

type ApiDiagnosticPtrOutput struct {
	*pulumi.OutputState
}

func (ApiDiagnosticPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDiagnostic)(nil))
}

func (o ApiDiagnosticPtrOutput) ToApiDiagnosticPtrOutput() ApiDiagnosticPtrOutput {
	return o
}

func (o ApiDiagnosticPtrOutput) ToApiDiagnosticPtrOutputWithContext(ctx context.Context) ApiDiagnosticPtrOutput {
	return o
}

type ApiDiagnosticArrayOutput struct{ *pulumi.OutputState }

func (ApiDiagnosticArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiDiagnostic)(nil))
}

func (o ApiDiagnosticArrayOutput) ToApiDiagnosticArrayOutput() ApiDiagnosticArrayOutput {
	return o
}

func (o ApiDiagnosticArrayOutput) ToApiDiagnosticArrayOutputWithContext(ctx context.Context) ApiDiagnosticArrayOutput {
	return o
}

func (o ApiDiagnosticArrayOutput) Index(i pulumi.IntInput) ApiDiagnosticOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiDiagnostic {
		return vs[0].([]ApiDiagnostic)[vs[1].(int)]
	}).(ApiDiagnosticOutput)
}

type ApiDiagnosticMapOutput struct{ *pulumi.OutputState }

func (ApiDiagnosticMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApiDiagnostic)(nil))
}

func (o ApiDiagnosticMapOutput) ToApiDiagnosticMapOutput() ApiDiagnosticMapOutput {
	return o
}

func (o ApiDiagnosticMapOutput) ToApiDiagnosticMapOutputWithContext(ctx context.Context) ApiDiagnosticMapOutput {
	return o
}

func (o ApiDiagnosticMapOutput) MapIndex(k pulumi.StringInput) ApiDiagnosticOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApiDiagnostic {
		return vs[0].(map[string]ApiDiagnostic)[vs[1].(string)]
	}).(ApiDiagnosticOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiDiagnosticOutput{})
	pulumi.RegisterOutputType(ApiDiagnosticPtrOutput{})
	pulumi.RegisterOutputType(ApiDiagnosticArrayOutput{})
	pulumi.RegisterOutputType(ApiDiagnosticMapOutput{})
}
