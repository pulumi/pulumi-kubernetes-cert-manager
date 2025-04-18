// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCertManager.Inputs
{

    public sealed class CertManagerPrometheusServiceMonitorArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("interval")]
        public Input<string>? Interval { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("prometheusInstance")]
        public Input<string>? PrometheusInstance { get; set; }

        [Input("string")]
        public Input<string>? String { get; set; }

        [Input("targetPort")]
        public Input<int>? TargetPort { get; set; }

        public CertManagerPrometheusServiceMonitorArgs()
        {
        }
        public static new CertManagerPrometheusServiceMonitorArgs Empty => new CertManagerPrometheusServiceMonitorArgs();
    }
}
