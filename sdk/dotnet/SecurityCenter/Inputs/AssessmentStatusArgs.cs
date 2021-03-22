// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.SecurityCenter.Inputs
{

    public sealed class AssessmentStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the cause of the assessment status.
        /// </summary>
        [Input("cause")]
        public Input<string>? Cause { get; set; }

        /// <summary>
        /// Specifies the programmatic code of the assessment status. Possible values are `Healthy`, `Unhealthy` and `NotApplicable`.
        /// </summary>
        [Input("code", required: true)]
        public Input<string> Code { get; set; } = null!;

        /// <summary>
        /// Specifies the human readable description of the assessment status.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public AssessmentStatusArgs()
        {
        }
    }
}