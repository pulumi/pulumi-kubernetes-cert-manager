// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCertManager.Inputs
{

    public sealed class CertManagerCaInjectorArgs : Pulumi.ResourceArgs
    {
        [Input("affinity")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.AffinityArgs>? Affinity { get; set; }

        /// <summary>
        /// Container Security Context to be set on the cainjector component container. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
        /// </summary>
        [Input("containerSecurityContext")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.SecurityContextArgs>? ContainerSecurityContext { get; set; }

        [Input("deploymentAnnotations")]
        private InputMap<string>? _deploymentAnnotations;

        /// <summary>
        /// Optional additional annotations to add to the cainjector Deployment
        /// </summary>
        public InputMap<string> DeploymentAnnotations
        {
            get => _deploymentAnnotations ?? (_deploymentAnnotations = new InputMap<string>());
            set => _deploymentAnnotations = value;
        }

        [Input("extraArgs")]
        private InputList<string>? _extraArgs;

        /// <summary>
        /// Optional additional arguments for cainjector
        /// </summary>
        public InputList<string> ExtraArgs
        {
            get => _extraArgs ?? (_extraArgs = new InputList<string>());
            set => _extraArgs = value;
        }

        [Input("image")]
        public Input<Inputs.CertManagerImageArgs>? Image { get; set; }

        [Input("nodeSelector")]
        private InputMap<string>? _nodeSelector;
        public InputMap<string> NodeSelector
        {
            get => _nodeSelector ?? (_nodeSelector = new InputMap<string>());
            set => _nodeSelector = value;
        }

        [Input("podAnnotations")]
        private InputMap<string>? _podAnnotations;

        /// <summary>
        /// Optional additional annotations to add to the cainjector Pods
        /// </summary>
        public InputMap<string> PodAnnotations
        {
            get => _podAnnotations ?? (_podAnnotations = new InputMap<string>());
            set => _podAnnotations = value;
        }

        [Input("podLabels")]
        private InputMap<string>? _podLabels;

        /// <summary>
        /// Optional additional labels to add to the Webhook Pods
        /// </summary>
        public InputMap<string> PodLabels
        {
            get => _podLabels ?? (_podLabels = new InputMap<string>());
            set => _podLabels = value;
        }

        /// <summary>
        /// Pod Security Context to be set on the cainjector component Pod. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
        /// </summary>
        [Input("podSecurityContext")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.PodSecurityContextArgs>? PodSecurityContext { get; set; }

        [Input("replicaCount")]
        public Input<int>? ReplicaCount { get; set; }

        [Input("resources")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Core.V1.ResourceRequirementsArgs>? Resources { get; set; }

        [Input("serviceAccount")]
        public Input<Inputs.CertManagerServiceAccountArgs>? ServiceAccount { get; set; }

        [Input("strategy")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Apps.V1.DeploymentStrategyArgs>? Strategy { get; set; }

        [Input("timeoutSeconds")]
        public Input<int>? TimeoutSeconds { get; set; }

        [Input("tolerations")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs>? _tolerations;
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs> Tolerations
        {
            get => _tolerations ?? (_tolerations = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.TolerationArgs>());
            set => _tolerations = value;
        }

        public CertManagerCaInjectorArgs()
        {
        }
    }
}
