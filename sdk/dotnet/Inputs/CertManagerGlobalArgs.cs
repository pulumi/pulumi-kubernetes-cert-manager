// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCertManager.Inputs
{

    public sealed class CertManagerGlobalArgs : global::Pulumi.ResourceArgs
    {
        [Input("imagePullSecrets")]
        private InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.LocalObjectReferenceArgs>? _imagePullSecrets;

        /// <summary>
        /// Reference to one or more secrets to be used when pulling images. ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
        /// </summary>
        public InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.LocalObjectReferenceArgs> ImagePullSecrets
        {
            get => _imagePullSecrets ?? (_imagePullSecrets = new InputList<Pulumi.Kubernetes.Types.Inputs.Core.V1.LocalObjectReferenceArgs>());
            set => _imagePullSecrets = value;
        }

        [Input("leaderElection")]
        public Input<Inputs.CertManagerGlobalLeaderElectionArgs>? LeaderElection { get; set; }

        /// <summary>
        /// Set the verbosity of cert-manager. Range of 0 - 6 with 6 being the most verbose.
        /// </summary>
        [Input("logLevel")]
        public Input<int>? LogLevel { get; set; }

        [Input("podSecurityPolicy")]
        public Input<Inputs.CertManagerGlobalPodSecurityPolicyArgs>? PodSecurityPolicy { get; set; }

        /// <summary>
        /// Optional priority class to be used for the cert-manager pods.
        /// </summary>
        [Input("priorityClassName")]
        public Input<string>? PriorityClassName { get; set; }

        [Input("rbac")]
        public Input<Inputs.CertManagerGlobalRbacArgs>? Rbac { get; set; }

        public CertManagerGlobalArgs()
        {
        }
        public static new CertManagerGlobalArgs Empty => new CertManagerGlobalArgs();
    }
}
