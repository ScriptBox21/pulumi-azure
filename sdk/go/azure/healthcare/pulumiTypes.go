// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ServiceAuthenticationConfiguration struct {
	// The intended audience to receive authentication tokens for the service. The default value is https://azurehealthcareapis.com
	Audience *string `pulumi:"audience"`
	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service. The default authority is the Directory defined in the authentication scheme in use when running this provider.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	Authority *string `pulumi:"authority"`
	// Enables the 'SMART on FHIR' option for mobile and web implementations.
	SmartProxyEnabled *bool `pulumi:"smartProxyEnabled"`
}

// ServiceAuthenticationConfigurationInput is an input type that accepts ServiceAuthenticationConfigurationArgs and ServiceAuthenticationConfigurationOutput values.
// You can construct a concrete instance of `ServiceAuthenticationConfigurationInput` via:
//
// 		 ServiceAuthenticationConfigurationArgs{...}
//
type ServiceAuthenticationConfigurationInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationOutput() ServiceAuthenticationConfigurationOutput
	ToServiceAuthenticationConfigurationOutputWithContext(context.Context) ServiceAuthenticationConfigurationOutput
}

type ServiceAuthenticationConfigurationArgs struct {
	// The intended audience to receive authentication tokens for the service. The default value is https://azurehealthcareapis.com
	Audience pulumi.StringPtrInput `pulumi:"audience"`
	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service. The default authority is the Directory defined in the authentication scheme in use when running this provider.
	// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
	Authority pulumi.StringPtrInput `pulumi:"authority"`
	// Enables the 'SMART on FHIR' option for mobile and web implementations.
	SmartProxyEnabled pulumi.BoolPtrInput `pulumi:"smartProxyEnabled"`
}

func (ServiceAuthenticationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfiguration)(nil)).Elem()
}

func (i ServiceAuthenticationConfigurationArgs) ToServiceAuthenticationConfigurationOutput() ServiceAuthenticationConfigurationOutput {
	return i.ToServiceAuthenticationConfigurationOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationArgs) ToServiceAuthenticationConfigurationOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationOutput)
}

func (i ServiceAuthenticationConfigurationArgs) ToServiceAuthenticationConfigurationPtrOutput() ServiceAuthenticationConfigurationPtrOutput {
	return i.ToServiceAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationArgs) ToServiceAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationOutput).ToServiceAuthenticationConfigurationPtrOutputWithContext(ctx)
}

// ServiceAuthenticationConfigurationPtrInput is an input type that accepts ServiceAuthenticationConfigurationArgs, ServiceAuthenticationConfigurationPtr and ServiceAuthenticationConfigurationPtrOutput values.
// You can construct a concrete instance of `ServiceAuthenticationConfigurationPtrInput` via:
//
// 		 ServiceAuthenticationConfigurationArgs{...}
//
//  or:
//
// 		 nil
//
type ServiceAuthenticationConfigurationPtrInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationPtrOutput() ServiceAuthenticationConfigurationPtrOutput
	ToServiceAuthenticationConfigurationPtrOutputWithContext(context.Context) ServiceAuthenticationConfigurationPtrOutput
}

type serviceAuthenticationConfigurationPtrType ServiceAuthenticationConfigurationArgs

func ServiceAuthenticationConfigurationPtr(v *ServiceAuthenticationConfigurationArgs) ServiceAuthenticationConfigurationPtrInput {
	return (*serviceAuthenticationConfigurationPtrType)(v)
}

func (*serviceAuthenticationConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfiguration)(nil)).Elem()
}

func (i *serviceAuthenticationConfigurationPtrType) ToServiceAuthenticationConfigurationPtrOutput() ServiceAuthenticationConfigurationPtrOutput {
	return i.ToServiceAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (i *serviceAuthenticationConfigurationPtrType) ToServiceAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationPtrOutput)
}

type ServiceAuthenticationConfigurationOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfiguration)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationOutput) ToServiceAuthenticationConfigurationOutput() ServiceAuthenticationConfigurationOutput {
	return o
}

func (o ServiceAuthenticationConfigurationOutput) ToServiceAuthenticationConfigurationOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationOutput {
	return o
}

func (o ServiceAuthenticationConfigurationOutput) ToServiceAuthenticationConfigurationPtrOutput() ServiceAuthenticationConfigurationPtrOutput {
	return o.ToServiceAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (o ServiceAuthenticationConfigurationOutput) ToServiceAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfiguration) *ServiceAuthenticationConfiguration {
		return &v
	}).(ServiceAuthenticationConfigurationPtrOutput)
}

// The intended audience to receive authentication tokens for the service. The default value is https://azurehealthcareapis.com
func (o ServiceAuthenticationConfigurationOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfiguration) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

// The Azure Active Directory (tenant) that serves as the authentication authority to access the service. The default authority is the Directory defined in the authentication scheme in use when running this provider.
// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
func (o ServiceAuthenticationConfigurationOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfiguration) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

// Enables the 'SMART on FHIR' option for mobile and web implementations.
func (o ServiceAuthenticationConfigurationOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfiguration) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfiguration)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationPtrOutput) ToServiceAuthenticationConfigurationPtrOutput() ServiceAuthenticationConfigurationPtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationPtrOutput) ToServiceAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationPtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationPtrOutput) Elem() ServiceAuthenticationConfigurationOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfiguration) ServiceAuthenticationConfiguration { return *v }).(ServiceAuthenticationConfigurationOutput)
}

// The intended audience to receive authentication tokens for the service. The default value is https://azurehealthcareapis.com
func (o ServiceAuthenticationConfigurationPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfiguration) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

// The Azure Active Directory (tenant) that serves as the authentication authority to access the service. The default authority is the Directory defined in the authentication scheme in use when running this provider.
// Authority must be registered to Azure AD and in the following format: https://{Azure-AD-endpoint}/{tenant-id}.
func (o ServiceAuthenticationConfigurationPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfiguration) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

// Enables the 'SMART on FHIR' option for mobile and web implementations.
func (o ServiceAuthenticationConfigurationPtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfiguration) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type ServiceCorsConfiguration struct {
	// If credentials are allowed via CORS.
	AllowCredentials *bool `pulumi:"allowCredentials"`
	// A set of headers to be allowed via CORS.
	AllowedHeaders []string `pulumi:"allowedHeaders"`
	// The methods to be allowed via CORS.
	AllowedMethods []string `pulumi:"allowedMethods"`
	// A set of origins to be allowed via CORS.
	AllowedOrigins []string `pulumi:"allowedOrigins"`
	// The max age to be allowed via CORS.
	MaxAgeInSeconds *int `pulumi:"maxAgeInSeconds"`
}

// ServiceCorsConfigurationInput is an input type that accepts ServiceCorsConfigurationArgs and ServiceCorsConfigurationOutput values.
// You can construct a concrete instance of `ServiceCorsConfigurationInput` via:
//
// 		 ServiceCorsConfigurationArgs{...}
//
type ServiceCorsConfigurationInput interface {
	pulumi.Input

	ToServiceCorsConfigurationOutput() ServiceCorsConfigurationOutput
	ToServiceCorsConfigurationOutputWithContext(context.Context) ServiceCorsConfigurationOutput
}

type ServiceCorsConfigurationArgs struct {
	// If credentials are allowed via CORS.
	AllowCredentials pulumi.BoolPtrInput `pulumi:"allowCredentials"`
	// A set of headers to be allowed via CORS.
	AllowedHeaders pulumi.StringArrayInput `pulumi:"allowedHeaders"`
	// The methods to be allowed via CORS.
	AllowedMethods pulumi.StringArrayInput `pulumi:"allowedMethods"`
	// A set of origins to be allowed via CORS.
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	// The max age to be allowed via CORS.
	MaxAgeInSeconds pulumi.IntPtrInput `pulumi:"maxAgeInSeconds"`
}

func (ServiceCorsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfiguration)(nil)).Elem()
}

func (i ServiceCorsConfigurationArgs) ToServiceCorsConfigurationOutput() ServiceCorsConfigurationOutput {
	return i.ToServiceCorsConfigurationOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationArgs) ToServiceCorsConfigurationOutputWithContext(ctx context.Context) ServiceCorsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationOutput)
}

func (i ServiceCorsConfigurationArgs) ToServiceCorsConfigurationPtrOutput() ServiceCorsConfigurationPtrOutput {
	return i.ToServiceCorsConfigurationPtrOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationArgs) ToServiceCorsConfigurationPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationOutput).ToServiceCorsConfigurationPtrOutputWithContext(ctx)
}

// ServiceCorsConfigurationPtrInput is an input type that accepts ServiceCorsConfigurationArgs, ServiceCorsConfigurationPtr and ServiceCorsConfigurationPtrOutput values.
// You can construct a concrete instance of `ServiceCorsConfigurationPtrInput` via:
//
// 		 ServiceCorsConfigurationArgs{...}
//
//  or:
//
// 		 nil
//
type ServiceCorsConfigurationPtrInput interface {
	pulumi.Input

	ToServiceCorsConfigurationPtrOutput() ServiceCorsConfigurationPtrOutput
	ToServiceCorsConfigurationPtrOutputWithContext(context.Context) ServiceCorsConfigurationPtrOutput
}

type serviceCorsConfigurationPtrType ServiceCorsConfigurationArgs

func ServiceCorsConfigurationPtr(v *ServiceCorsConfigurationArgs) ServiceCorsConfigurationPtrInput {
	return (*serviceCorsConfigurationPtrType)(v)
}

func (*serviceCorsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfiguration)(nil)).Elem()
}

func (i *serviceCorsConfigurationPtrType) ToServiceCorsConfigurationPtrOutput() ServiceCorsConfigurationPtrOutput {
	return i.ToServiceCorsConfigurationPtrOutputWithContext(context.Background())
}

func (i *serviceCorsConfigurationPtrType) ToServiceCorsConfigurationPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationPtrOutput)
}

type ServiceCorsConfigurationOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfiguration)(nil)).Elem()
}

func (o ServiceCorsConfigurationOutput) ToServiceCorsConfigurationOutput() ServiceCorsConfigurationOutput {
	return o
}

func (o ServiceCorsConfigurationOutput) ToServiceCorsConfigurationOutputWithContext(ctx context.Context) ServiceCorsConfigurationOutput {
	return o
}

func (o ServiceCorsConfigurationOutput) ToServiceCorsConfigurationPtrOutput() ServiceCorsConfigurationPtrOutput {
	return o.ToServiceCorsConfigurationPtrOutputWithContext(context.Background())
}

func (o ServiceCorsConfigurationOutput) ToServiceCorsConfigurationPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) *ServiceCorsConfiguration {
		return &v
	}).(ServiceCorsConfigurationPtrOutput)
}

// If credentials are allowed via CORS.
func (o ServiceCorsConfigurationOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

// A set of headers to be allowed via CORS.
func (o ServiceCorsConfigurationOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

// The methods to be allowed via CORS.
func (o ServiceCorsConfigurationOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

// A set of origins to be allowed via CORS.
func (o ServiceCorsConfigurationOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

// The max age to be allowed via CORS.
func (o ServiceCorsConfigurationOutput) MaxAgeInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) *int { return v.MaxAgeInSeconds }).(pulumi.IntPtrOutput)
}

type ServiceCorsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfiguration)(nil)).Elem()
}

func (o ServiceCorsConfigurationPtrOutput) ToServiceCorsConfigurationPtrOutput() ServiceCorsConfigurationPtrOutput {
	return o
}

func (o ServiceCorsConfigurationPtrOutput) ToServiceCorsConfigurationPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationPtrOutput {
	return o
}

func (o ServiceCorsConfigurationPtrOutput) Elem() ServiceCorsConfigurationOutput {
	return o.ApplyT(func(v *ServiceCorsConfiguration) ServiceCorsConfiguration { return *v }).(ServiceCorsConfigurationOutput)
}

// If credentials are allowed via CORS.
func (o ServiceCorsConfigurationPtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

// A set of headers to be allowed via CORS.
func (o ServiceCorsConfigurationPtrOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

// The methods to be allowed via CORS.
func (o ServiceCorsConfigurationPtrOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

// A set of origins to be allowed via CORS.
func (o ServiceCorsConfigurationPtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

// The max age to be allowed via CORS.
func (o ServiceCorsConfigurationPtrOutput) MaxAgeInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfiguration) *int { return v.MaxAgeInSeconds }).(pulumi.IntPtrOutput)
}

type GetServiceAuthenticationConfiguration struct {
	// The intended audience to receive authentication tokens for the service.
	Audience string `pulumi:"audience"`
	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	Authority string `pulumi:"authority"`
	// Is the 'SMART on FHIR' option for mobile and web implementations enbled?
	SmartProxyEnabled bool `pulumi:"smartProxyEnabled"`
}

// GetServiceAuthenticationConfigurationInput is an input type that accepts GetServiceAuthenticationConfigurationArgs and GetServiceAuthenticationConfigurationOutput values.
// You can construct a concrete instance of `GetServiceAuthenticationConfigurationInput` via:
//
// 		 GetServiceAuthenticationConfigurationArgs{...}
//
type GetServiceAuthenticationConfigurationInput interface {
	pulumi.Input

	ToGetServiceAuthenticationConfigurationOutput() GetServiceAuthenticationConfigurationOutput
	ToGetServiceAuthenticationConfigurationOutputWithContext(context.Context) GetServiceAuthenticationConfigurationOutput
}

type GetServiceAuthenticationConfigurationArgs struct {
	// The intended audience to receive authentication tokens for the service.
	Audience pulumi.StringInput `pulumi:"audience"`
	// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
	Authority pulumi.StringInput `pulumi:"authority"`
	// Is the 'SMART on FHIR' option for mobile and web implementations enbled?
	SmartProxyEnabled pulumi.BoolInput `pulumi:"smartProxyEnabled"`
}

func (GetServiceAuthenticationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAuthenticationConfiguration)(nil)).Elem()
}

func (i GetServiceAuthenticationConfigurationArgs) ToGetServiceAuthenticationConfigurationOutput() GetServiceAuthenticationConfigurationOutput {
	return i.ToGetServiceAuthenticationConfigurationOutputWithContext(context.Background())
}

func (i GetServiceAuthenticationConfigurationArgs) ToGetServiceAuthenticationConfigurationOutputWithContext(ctx context.Context) GetServiceAuthenticationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceAuthenticationConfigurationOutput)
}

// GetServiceAuthenticationConfigurationArrayInput is an input type that accepts GetServiceAuthenticationConfigurationArray and GetServiceAuthenticationConfigurationArrayOutput values.
// You can construct a concrete instance of `GetServiceAuthenticationConfigurationArrayInput` via:
//
// 		 GetServiceAuthenticationConfigurationArray{ GetServiceAuthenticationConfigurationArgs{...} }
//
type GetServiceAuthenticationConfigurationArrayInput interface {
	pulumi.Input

	ToGetServiceAuthenticationConfigurationArrayOutput() GetServiceAuthenticationConfigurationArrayOutput
	ToGetServiceAuthenticationConfigurationArrayOutputWithContext(context.Context) GetServiceAuthenticationConfigurationArrayOutput
}

type GetServiceAuthenticationConfigurationArray []GetServiceAuthenticationConfigurationInput

func (GetServiceAuthenticationConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceAuthenticationConfiguration)(nil)).Elem()
}

func (i GetServiceAuthenticationConfigurationArray) ToGetServiceAuthenticationConfigurationArrayOutput() GetServiceAuthenticationConfigurationArrayOutput {
	return i.ToGetServiceAuthenticationConfigurationArrayOutputWithContext(context.Background())
}

func (i GetServiceAuthenticationConfigurationArray) ToGetServiceAuthenticationConfigurationArrayOutputWithContext(ctx context.Context) GetServiceAuthenticationConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceAuthenticationConfigurationArrayOutput)
}

type GetServiceAuthenticationConfigurationOutput struct{ *pulumi.OutputState }

func (GetServiceAuthenticationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceAuthenticationConfiguration)(nil)).Elem()
}

func (o GetServiceAuthenticationConfigurationOutput) ToGetServiceAuthenticationConfigurationOutput() GetServiceAuthenticationConfigurationOutput {
	return o
}

func (o GetServiceAuthenticationConfigurationOutput) ToGetServiceAuthenticationConfigurationOutputWithContext(ctx context.Context) GetServiceAuthenticationConfigurationOutput {
	return o
}

// The intended audience to receive authentication tokens for the service.
func (o GetServiceAuthenticationConfigurationOutput) Audience() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAuthenticationConfiguration) string { return v.Audience }).(pulumi.StringOutput)
}

// The Azure Active Directory (tenant) that serves as the authentication authority to access the service.
func (o GetServiceAuthenticationConfigurationOutput) Authority() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceAuthenticationConfiguration) string { return v.Authority }).(pulumi.StringOutput)
}

// Is the 'SMART on FHIR' option for mobile and web implementations enbled?
func (o GetServiceAuthenticationConfigurationOutput) SmartProxyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceAuthenticationConfiguration) bool { return v.SmartProxyEnabled }).(pulumi.BoolOutput)
}

type GetServiceAuthenticationConfigurationArrayOutput struct{ *pulumi.OutputState }

func (GetServiceAuthenticationConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceAuthenticationConfiguration)(nil)).Elem()
}

func (o GetServiceAuthenticationConfigurationArrayOutput) ToGetServiceAuthenticationConfigurationArrayOutput() GetServiceAuthenticationConfigurationArrayOutput {
	return o
}

func (o GetServiceAuthenticationConfigurationArrayOutput) ToGetServiceAuthenticationConfigurationArrayOutputWithContext(ctx context.Context) GetServiceAuthenticationConfigurationArrayOutput {
	return o
}

func (o GetServiceAuthenticationConfigurationArrayOutput) Index(i pulumi.IntInput) GetServiceAuthenticationConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetServiceAuthenticationConfiguration {
		return vs[0].([]GetServiceAuthenticationConfiguration)[vs[1].(int)]
	}).(GetServiceAuthenticationConfigurationOutput)
}

type GetServiceCorsConfiguration struct {
	// Are credentials are allowed via CORS?
	AllowCredentials bool `pulumi:"allowCredentials"`
	// The set of headers to be allowed via CORS.
	AllowedHeaders []string `pulumi:"allowedHeaders"`
	// The methods to be allowed via CORS.
	AllowedMethods []string `pulumi:"allowedMethods"`
	// The set of origins to be allowed via CORS.
	AllowedOrigins []string `pulumi:"allowedOrigins"`
	// The max age to be allowed via CORS.
	MaxAgeInSeconds int `pulumi:"maxAgeInSeconds"`
}

// GetServiceCorsConfigurationInput is an input type that accepts GetServiceCorsConfigurationArgs and GetServiceCorsConfigurationOutput values.
// You can construct a concrete instance of `GetServiceCorsConfigurationInput` via:
//
// 		 GetServiceCorsConfigurationArgs{...}
//
type GetServiceCorsConfigurationInput interface {
	pulumi.Input

	ToGetServiceCorsConfigurationOutput() GetServiceCorsConfigurationOutput
	ToGetServiceCorsConfigurationOutputWithContext(context.Context) GetServiceCorsConfigurationOutput
}

type GetServiceCorsConfigurationArgs struct {
	// Are credentials are allowed via CORS?
	AllowCredentials pulumi.BoolInput `pulumi:"allowCredentials"`
	// The set of headers to be allowed via CORS.
	AllowedHeaders pulumi.StringArrayInput `pulumi:"allowedHeaders"`
	// The methods to be allowed via CORS.
	AllowedMethods pulumi.StringArrayInput `pulumi:"allowedMethods"`
	// The set of origins to be allowed via CORS.
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	// The max age to be allowed via CORS.
	MaxAgeInSeconds pulumi.IntInput `pulumi:"maxAgeInSeconds"`
}

func (GetServiceCorsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceCorsConfiguration)(nil)).Elem()
}

func (i GetServiceCorsConfigurationArgs) ToGetServiceCorsConfigurationOutput() GetServiceCorsConfigurationOutput {
	return i.ToGetServiceCorsConfigurationOutputWithContext(context.Background())
}

func (i GetServiceCorsConfigurationArgs) ToGetServiceCorsConfigurationOutputWithContext(ctx context.Context) GetServiceCorsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceCorsConfigurationOutput)
}

// GetServiceCorsConfigurationArrayInput is an input type that accepts GetServiceCorsConfigurationArray and GetServiceCorsConfigurationArrayOutput values.
// You can construct a concrete instance of `GetServiceCorsConfigurationArrayInput` via:
//
// 		 GetServiceCorsConfigurationArray{ GetServiceCorsConfigurationArgs{...} }
//
type GetServiceCorsConfigurationArrayInput interface {
	pulumi.Input

	ToGetServiceCorsConfigurationArrayOutput() GetServiceCorsConfigurationArrayOutput
	ToGetServiceCorsConfigurationArrayOutputWithContext(context.Context) GetServiceCorsConfigurationArrayOutput
}

type GetServiceCorsConfigurationArray []GetServiceCorsConfigurationInput

func (GetServiceCorsConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceCorsConfiguration)(nil)).Elem()
}

func (i GetServiceCorsConfigurationArray) ToGetServiceCorsConfigurationArrayOutput() GetServiceCorsConfigurationArrayOutput {
	return i.ToGetServiceCorsConfigurationArrayOutputWithContext(context.Background())
}

func (i GetServiceCorsConfigurationArray) ToGetServiceCorsConfigurationArrayOutputWithContext(ctx context.Context) GetServiceCorsConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetServiceCorsConfigurationArrayOutput)
}

type GetServiceCorsConfigurationOutput struct{ *pulumi.OutputState }

func (GetServiceCorsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceCorsConfiguration)(nil)).Elem()
}

func (o GetServiceCorsConfigurationOutput) ToGetServiceCorsConfigurationOutput() GetServiceCorsConfigurationOutput {
	return o
}

func (o GetServiceCorsConfigurationOutput) ToGetServiceCorsConfigurationOutputWithContext(ctx context.Context) GetServiceCorsConfigurationOutput {
	return o
}

// Are credentials are allowed via CORS?
func (o GetServiceCorsConfigurationOutput) AllowCredentials() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceCorsConfiguration) bool { return v.AllowCredentials }).(pulumi.BoolOutput)
}

// The set of headers to be allowed via CORS.
func (o GetServiceCorsConfigurationOutput) AllowedHeaders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceCorsConfiguration) []string { return v.AllowedHeaders }).(pulumi.StringArrayOutput)
}

// The methods to be allowed via CORS.
func (o GetServiceCorsConfigurationOutput) AllowedMethods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceCorsConfiguration) []string { return v.AllowedMethods }).(pulumi.StringArrayOutput)
}

// The set of origins to be allowed via CORS.
func (o GetServiceCorsConfigurationOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServiceCorsConfiguration) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

// The max age to be allowed via CORS.
func (o GetServiceCorsConfigurationOutput) MaxAgeInSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v GetServiceCorsConfiguration) int { return v.MaxAgeInSeconds }).(pulumi.IntOutput)
}

type GetServiceCorsConfigurationArrayOutput struct{ *pulumi.OutputState }

func (GetServiceCorsConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetServiceCorsConfiguration)(nil)).Elem()
}

func (o GetServiceCorsConfigurationArrayOutput) ToGetServiceCorsConfigurationArrayOutput() GetServiceCorsConfigurationArrayOutput {
	return o
}

func (o GetServiceCorsConfigurationArrayOutput) ToGetServiceCorsConfigurationArrayOutputWithContext(ctx context.Context) GetServiceCorsConfigurationArrayOutput {
	return o
}

func (o GetServiceCorsConfigurationArrayOutput) Index(i pulumi.IntInput) GetServiceCorsConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetServiceCorsConfiguration {
		return vs[0].([]GetServiceCorsConfiguration)[vs[1].(int)]
	}).(GetServiceCorsConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GetServiceAuthenticationConfigurationOutput{})
	pulumi.RegisterOutputType(GetServiceAuthenticationConfigurationArrayOutput{})
	pulumi.RegisterOutputType(GetServiceCorsConfigurationOutput{})
	pulumi.RegisterOutputType(GetServiceCorsConfigurationArrayOutput{})
}
