// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Inputs
{

    public sealed class GremlinGraphIndexPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the indexing policy is automatic. Defaults to `true`.
        /// </summary>
        [Input("automatic")]
        public Input<bool>? Automatic { get; set; }

        [Input("excludedPaths")]
        private InputList<string>? _excludedPaths;

        /// <summary>
        /// List of paths to exclude from indexing. Required if `indexing_mode` is `Consistent` or `Lazy`.
        /// </summary>
        public InputList<string> ExcludedPaths
        {
            get => _excludedPaths ?? (_excludedPaths = new InputList<string>());
            set => _excludedPaths = value;
        }

        [Input("includedPaths")]
        private InputList<string>? _includedPaths;

        /// <summary>
        /// List of paths to include in the indexing. Required if `indexing_mode` is `Consistent` or `Lazy`.
        /// </summary>
        public InputList<string> IncludedPaths
        {
            get => _includedPaths ?? (_includedPaths = new InputList<string>());
            set => _includedPaths = value;
        }

        /// <summary>
        /// Indicates the indexing mode. Possible values include: `Consistent`, `Lazy`, `None`.
        /// </summary>
        [Input("indexingMode", required: true)]
        public Input<string> IndexingMode { get; set; } = null!;

        public GremlinGraphIndexPolicyArgs()
        {
        }
    }
}
