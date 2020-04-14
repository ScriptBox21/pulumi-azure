// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.FrontDoor.Inputs
{

    public sealed class FrontdoorRoutingRuleRedirectConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination fragment in the portion of URL after '#'. Set this to add a fragment to the redirect URL.
        /// </summary>
        [Input("customFragment")]
        public Input<string>? CustomFragment { get; set; }

        /// <summary>
        /// Set this to change the URL for the redirection.
        /// </summary>
        [Input("customHost")]
        public Input<string>? CustomHost { get; set; }

        /// <summary>
        /// The path to retain as per the incoming request, or update in the URL for the redirection.
        /// </summary>
        [Input("customPath")]
        public Input<string>? CustomPath { get; set; }

        /// <summary>
        /// Replace any existing query string from the incoming request URL.
        /// </summary>
        [Input("customQueryString")]
        public Input<string>? CustomQueryString { get; set; }

        /// <summary>
        /// Protocol to use when redirecting. Valid options are `HttpOnly`, `HttpsOnly`, or `MatchRequest`. Defaults to `MatchRequest`
        /// </summary>
        [Input("redirectProtocol", required: true)]
        public Input<string> RedirectProtocol { get; set; } = null!;

        /// <summary>
        /// Status code for the redirect. Valida options are `Moved`, `Found`, `TemporaryRedirect`, `PermanentRedirect`. Defaults to `Found`
        /// </summary>
        [Input("redirectType", required: true)]
        public Input<string> RedirectType { get; set; } = null!;

        public FrontdoorRoutingRuleRedirectConfigurationArgs()
        {
        }
    }
}
