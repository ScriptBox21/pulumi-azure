// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Core.Inputs
{

    public sealed class CustomProviderActionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the endpoint of the action. 
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the action. 
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public CustomProviderActionGetArgs()
        {
        }
    }
}
