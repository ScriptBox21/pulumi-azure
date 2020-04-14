// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Automation.Inputs
{

    public sealed class ModuleModuleLinkGetArgs : Pulumi.ResourceArgs
    {
        [Input("hash")]
        public Input<Inputs.ModuleModuleLinkHashGetArgs>? Hash { get; set; }

        /// <summary>
        /// The uri of the module content (zip or nupkg).
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public ModuleModuleLinkGetArgs()
        {
        }
    }
}
