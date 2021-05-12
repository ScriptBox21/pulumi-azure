// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class GetEnvironmentV3ClusterSettingResult
    {
        /// <summary>
        /// The name of this v3 App Service Environment.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value for the Cluster Setting.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetEnvironmentV3ClusterSettingResult(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
