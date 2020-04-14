// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid.Inputs
{

    public sealed class DomainInputMappingDefaultValuesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the default data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("dataVersion")]
        public Input<string>? DataVersion { get; set; }

        /// <summary>
        /// Specifies the default event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// Specifies the default subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        public DomainInputMappingDefaultValuesGetArgs()
        {
        }
    }
}
