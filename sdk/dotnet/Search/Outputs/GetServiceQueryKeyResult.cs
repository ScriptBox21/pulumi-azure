// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Search.Outputs
{

    [OutputType]
    public sealed class GetServiceQueryKeyResult
    {
        /// <summary>
        /// The value of this Query Key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The Name of the Search Service.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetServiceQueryKeyResult(
            string key,

            string name)
        {
            Key = key;
            Name = name;
        }
    }
}
