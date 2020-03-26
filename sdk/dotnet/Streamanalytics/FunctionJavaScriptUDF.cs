// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.StreamAnalytics
{
    /// <summary>
    /// Manages a JavaScript UDF Function within Stream Analytics Streaming Job.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_function_javascript_udf.html.markdown.
    /// </summary>
    public partial class FunctionJavaScriptUDF : Pulumi.CustomResource
    {
        /// <summary>
        /// One or more `input` blocks as defined below.
        /// </summary>
        [Output("inputs")]
        public Output<ImmutableArray<Outputs.FunctionJavaScriptUDFInputs>> Inputs { get; private set; } = null!;

        /// <summary>
        /// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An `output` blocks as defined below.
        /// </summary>
        [Output("output")]
        public Output<Outputs.FunctionJavaScriptUDFOutput> Output { get; private set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Output("resourceGroupName")]
        public Output<string> ResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The JavaScript of this UDF Function.
        /// </summary>
        [Output("script")]
        public Output<string> Script { get; private set; } = null!;

        /// <summary>
        /// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Output("streamAnalyticsJobName")]
        public Output<string> StreamAnalyticsJobName { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionJavaScriptUDF resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionJavaScriptUDF(string name, FunctionJavaScriptUDFArgs args, CustomResourceOptions? options = null)
            : base("azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private FunctionJavaScriptUDF(string name, Input<string> id, FunctionJavaScriptUDFState? state = null, CustomResourceOptions? options = null)
            : base("azure:streamanalytics/functionJavaScriptUDF:FunctionJavaScriptUDF", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FunctionJavaScriptUDF resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionJavaScriptUDF Get(string name, Input<string> id, FunctionJavaScriptUDFState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionJavaScriptUDF(name, id, state, options);
        }
    }

    public sealed class FunctionJavaScriptUDFArgs : Pulumi.ResourceArgs
    {
        [Input("inputs", required: true)]
        private InputList<Inputs.FunctionJavaScriptUDFInputsArgs>? _inputs;

        /// <summary>
        /// One or more `input` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FunctionJavaScriptUDFInputsArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.FunctionJavaScriptUDFInputsArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An `output` blocks as defined below.
        /// </summary>
        [Input("output", required: true)]
        public Input<Inputs.FunctionJavaScriptUDFOutputArgs> Output { get; set; } = null!;

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The JavaScript of this UDF Function.
        /// </summary>
        [Input("script", required: true)]
        public Input<string> Script { get; set; } = null!;

        /// <summary>
        /// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("streamAnalyticsJobName", required: true)]
        public Input<string> StreamAnalyticsJobName { get; set; } = null!;

        public FunctionJavaScriptUDFArgs()
        {
        }
    }

    public sealed class FunctionJavaScriptUDFState : Pulumi.ResourceArgs
    {
        [Input("inputs")]
        private InputList<Inputs.FunctionJavaScriptUDFInputsGetArgs>? _inputs;

        /// <summary>
        /// One or more `input` blocks as defined below.
        /// </summary>
        public InputList<Inputs.FunctionJavaScriptUDFInputsGetArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.FunctionJavaScriptUDFInputsGetArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// The name of the JavaScript UDF Function. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// An `output` blocks as defined below.
        /// </summary>
        [Input("output")]
        public Input<Inputs.FunctionJavaScriptUDFOutputGetArgs>? Output { get; set; }

        /// <summary>
        /// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        /// </summary>
        [Input("resourceGroupName")]
        public Input<string>? ResourceGroupName { get; set; }

        /// <summary>
        /// The JavaScript of this UDF Function.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// The name of the Stream Analytics Job where this Function should be created. Changing this forces a new resource to be created.
        /// </summary>
        [Input("streamAnalyticsJobName")]
        public Input<string>? StreamAnalyticsJobName { get; set; }

        public FunctionJavaScriptUDFState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class FunctionJavaScriptUDFInputsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Data Type for the Input Argument of this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FunctionJavaScriptUDFInputsArgs()
        {
        }
    }

    public sealed class FunctionJavaScriptUDFInputsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Data Type for the Input Argument of this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FunctionJavaScriptUDFInputsGetArgs()
        {
        }
    }

    public sealed class FunctionJavaScriptUDFOutputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Data Type output from this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FunctionJavaScriptUDFOutputArgs()
        {
        }
    }

    public sealed class FunctionJavaScriptUDFOutputGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Data Type output from this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FunctionJavaScriptUDFOutputGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class FunctionJavaScriptUDFInputs
    {
        /// <summary>
        /// The Data Type for the Input Argument of this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FunctionJavaScriptUDFInputs(string type)
        {
            Type = type;
        }
    }

    [OutputType]
    public sealed class FunctionJavaScriptUDFOutput
    {
        /// <summary>
        /// The Data Type output from this JavaScript Function. Possible values include `array`, `any`, `bigint`, `datetime`, `float`, `nvarchar(max)` and `record`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FunctionJavaScriptUDFOutput(string type)
        {
            Type = type;
        }
    }
    }
}
