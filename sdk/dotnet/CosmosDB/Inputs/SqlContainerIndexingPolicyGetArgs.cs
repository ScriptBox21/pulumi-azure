// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.CosmosDB.Inputs
{

    public sealed class SqlContainerIndexingPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("excludedPaths")]
        private InputList<Inputs.SqlContainerIndexingPolicyExcludedPathGetArgs>? _excludedPaths;

        /// <summary>
        /// One or more `excluded_path` blocks as defined below. Either `included_path` or `excluded_path` must contain the `path` `/*`
        /// </summary>
        public InputList<Inputs.SqlContainerIndexingPolicyExcludedPathGetArgs> ExcludedPaths
        {
            get => _excludedPaths ?? (_excludedPaths = new InputList<Inputs.SqlContainerIndexingPolicyExcludedPathGetArgs>());
            set => _excludedPaths = value;
        }

        [Input("includedPaths")]
        private InputList<Inputs.SqlContainerIndexingPolicyIncludedPathGetArgs>? _includedPaths;

        /// <summary>
        /// One or more `included_path` blocks as defined below. Either `included_path` or `excluded_path` must contain the `path` `/*`
        /// </summary>
        public InputList<Inputs.SqlContainerIndexingPolicyIncludedPathGetArgs> IncludedPaths
        {
            get => _includedPaths ?? (_includedPaths = new InputList<Inputs.SqlContainerIndexingPolicyIncludedPathGetArgs>());
            set => _includedPaths = value;
        }

        /// <summary>
        /// Indicates the indexing mode. Possible values include: `Consistent` and `None`. Defaults to `Consistent`.
        /// </summary>
        [Input("indexingMode")]
        public Input<string>? IndexingMode { get; set; }

        public SqlContainerIndexingPolicyGetArgs()
        {
        }
    }
}
