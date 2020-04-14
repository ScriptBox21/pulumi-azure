// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Inputs
{

    public sealed class BackendCredentialsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An `authorization` block as defined below.
        /// </summary>
        [Input("authorization")]
        public Input<Inputs.BackendCredentialsAuthorizationGetArgs>? Authorization { get; set; }

        [Input("certificates")]
        private InputList<string>? _certificates;

        /// <summary>
        /// A list of client certificate thumbprints to present to the backend host. The certificates must exist within the API Management Service.
        /// </summary>
        public InputList<string> Certificates
        {
            get => _certificates ?? (_certificates = new InputList<string>());
            set => _certificates = value;
        }

        [Input("header")]
        private InputMap<string>? _header;

        /// <summary>
        /// A mapping of header parameters to pass to the backend host. The keys are the header names and the values are a comma separated string of header values. This is converted to a list before being passed to the API.
        /// </summary>
        public InputMap<string> Header
        {
            get => _header ?? (_header = new InputMap<string>());
            set => _header = value;
        }

        [Input("query")]
        private InputMap<string>? _query;

        /// <summary>
        /// A mapping of query parameters to pass to the backend host. The keys are the query names and the values are a comma separated string of query values. This is converted to a list before being passed to the API.
        /// </summary>
        public InputMap<string> Query
        {
            get => _query ?? (_query = new InputMap<string>());
            set => _query = value;
        }

        public BackendCredentialsGetArgs()
        {
        }
    }
}
