// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.EventGrid.Outputs
{

    [OutputType]
    public sealed class DomainInputMappingFields
    {
        /// <summary>
        /// Specifies the data version of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? DataVersion;
        /// <summary>
        /// Specifies the event time of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? EventTime;
        /// <summary>
        /// Specifies the event type of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? EventType;
        /// <summary>
        /// Specifies the id of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Specifies the subject of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Subject;
        /// <summary>
        /// Specifies the topic of the EventGrid Event to associate with the domain. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string? Topic;

        [OutputConstructor]
        private DomainInputMappingFields(
            string? dataVersion,

            string? eventTime,

            string? eventType,

            string? id,

            string? subject,

            string? topic)
        {
            DataVersion = dataVersion;
            EventTime = eventTime;
            EventType = eventType;
            Id = id;
            Subject = subject;
            Topic = topic;
        }
    }
}
