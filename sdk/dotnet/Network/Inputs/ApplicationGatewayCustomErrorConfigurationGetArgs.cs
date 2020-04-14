// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Network.Inputs
{

    public sealed class ApplicationGatewayCustomErrorConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Error page URL of the application gateway customer error.
        /// </summary>
        [Input("customErrorPageUrl", required: true)]
        public Input<string> CustomErrorPageUrl { get; set; } = null!;

        /// <summary>
        /// The ID of the Rewrite Rule Set
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Status code of the application gateway customer error. Possible values are `HttpStatus403` and `HttpStatus502`
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<string> StatusCode { get; set; } = null!;

        public ApplicationGatewayCustomErrorConfigurationGetArgs()
        {
        }
    }
}
