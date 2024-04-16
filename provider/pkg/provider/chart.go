// Copyright 2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	helmbase "github.com/pulumi/pulumi-go-helmbase"
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CertManager installs a fully-configured CertManager stack in Kubernetes.
type CertManager struct {
	pulumi.ResourceState
	Status helmv3.ReleaseStatusOutput `pulumi:"status" pschema:"out"`
}

func (c *CertManager) SetOutputs(out helmv3.ReleaseStatusOutput) { c.Status = out }
func (c *CertManager) Type() string                              { return ComponentName }
func (c *CertManager) DefaultChartName() string                  { return "cert-manager" }
func (c *CertManager) DefaultRepoURL() string                    { return "https://charts.jetstack.io" }

// CertManager contains the set of arguments for creating a CertManager component resource.
type CertManagerArgs struct {
	Global       *CertManagerGlobal         `pulumi:"global"`
	InstallCRDs  *bool                      `pulumi:"installCRDs"`
	ReplicaCount *int                       `pulumi:"replicaCount"`
	Strategy     *appsv1.DeploymentStrategy `pulumi:"strategy" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:apps/v1:DeploymentStrategy"`
	// Comma separated list of feature gates that should be enabled on the controller pod.
	FeatureGates *string           `pulumi:"featureGates"`
	Image        *CertManagerImage `pulumi:"image"`
	// Override the namespace used to store DNS provider credentials etc. for ClusterIssuer
	// resources. By default, the same namespace as cert-manager is deployed within is
	// used. This namespace will not be automatically created by the Helm chart.
	ClusterResourceNamespace *string                    `pulumi:"clusterResourceNamespace"`
	ServiceAccount           *CertManagerServiceAccount `pulumi:"serviceAccount"`
	// Optional additional arguments.
	ExtraArgs *[]string                    `pulumi:"extraArgs"`
	ExtraEnv  *[]corev1.EnvVar             `pulumi:"extraEnv" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:EnvVar"`
	Resources *corev1.ResourceRequirements `pulumi:"resources" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements"`
	// Pod Security Context. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SecurityContext *corev1.PodSecurityContext `pulumi:"securityContext" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:PodSecurityContext"`
	// Container Security Context to be set on the controller component container.
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	ContainerSecurityContext *corev1.SecurityContext `pulumi:"containerSecurityContext" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:SecurityContext"`
	Volumes                  *[]corev1.Volume        `pulumi:"extraVolumes" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Volume"`
	VolumeMounts             *[]corev1.VolumeMount   `pulumi:"extraVolumeMounts" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:VolumeMount"`
	// Optional additional annotations to add to the controller Deployment
	DeploymentAnnotations *map[string]string `pulumi:"deploymentAnnotations"`
	// Optional additional annotations to add to the controller Pods
	PodAnnotations *map[string]string `pulumi:"podAnnotations"`
	PodLabels      *map[string]string `pulumi:"podLabels"`
	// Optional additional labels to add to the controller Service
	ServiceLabels *map[string]string `pulumi:"serviceLabels"`
	// Optional additional annotations to add to the controller service
	ServiceAnnotations *map[string]string `pulumi:"serviceAnnotations"`
	// Optional DNS settings, useful if you have a public and private DNS zone for
	// the same domain on Route 53. What follows is an example of ensuring
	// cert-manager can access an ingress or DNS TXT records at all times.
	// NOTE: This requires Kubernetes 1.10 or `CustomPodDNS` feature gate enabled for
	// the cluster to work.
	PodDnsPolicy    *string                     `pulumi:"podDnsPolicy"`
	PodDnsConfig    *corev1.PodDNSConfig        `pulumi:"podDnsConfig" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:PodDNSConfig"`
	NodeSelector    *corev1.NodeSelector        `pulumi:"nodeSelector" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:NodeSelector"`
	IngressShim     *CertManagerIngressShim     `pulumi:"ingressShim"`
	Prometheus      *CertManagerPrometheus      `pulumi:"prometheus"`
	HttpProxy       *string                     `pulumi:"http_proxy"`
	HttpsProxy      *string                     `pulumi:"https_proxy"`
	NoProxy         *[]string                   `pulumi:"no_proxy"`
	Affinity        *corev1.Affinity            `pulumi:"affinity" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Affinity"`
	Tolerations     *[]corev1.Toleration        `pulumi:"tolerations" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Toleration"`
	Webhook         *CertManagerWebhook         `pulumi:"webhook"`
	CAInjector      *CertManagerCaInjector      `pulumi:"cainjector"`
	StartupAPICheck *CertManagerStartupAPICheck `pulumi:"startupapicheck"`

	// HelmOptions is an escape hatch that lets the end user control any aspect of the
	// Helm deployment. This exposes the entirety of the underlying Helm Release component args.
	HelmOptions *helmbase.ReleaseType `pulumi:"helmOptions" pschema:"ref=#/types/chart-cert-manager:index:Release" json:"-"`
}

func (args *CertManagerArgs) R() **helmbase.ReleaseType { return &args.HelmOptions }

type CertManagerGlobal struct {
	// Reference to one or more secrets to be used when pulling images.
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
	ImagePullSecrets *[]corev1.LocalObjectReference `pulumi:"imagePullSecrets" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:LocalObjectReference"`
	// Optional priority class to be used for the cert-manager pods.
	PriorityClassName *string                             `pulumi:"priorityClassName"`
	Rbac              *CertManagerGlobalRbac              `pulumi:"rbac"`
	PodSecurityPolicy *CertManagerGlobalPodSecurityPolicy `pulumi:"podSecurityPolicy"`
	// Set the verbosity of cert-manager. Range of 0 - 6 with 6 being the most verbose.
	LogLevel       *int                             `pulumi:"logLevel"`
	LeaderElection *CertManagerGlobalLeaderElection `pulumi:"leaderElection"`
}

type CertManagerGlobalRbac struct {
	Create *bool `pulumi:"create"`
}

type CertManagerGlobalPodSecurityPolicy struct {
	Enabled     *bool `pulumi:"enabled"`
	UseAppArmor *bool `pulumi:"useAppArmor"`
}

type CertManagerGlobalLeaderElection struct {
	// Override the namespace used to store the ConfigMap for leader election.
	Namespace *string `pulumi:"namespace"`
	// The duration that non-leader candidates will wait after observing a
	// leadership renewal until attempting to acquire leadership of a led but
	// unrenewed leader slot. This is effectively the maximum duration that a
	// leader can be stopped before it is replaced by another candidate.
	LeaseDuration *string `pulumi:"leaseDuration"`
	// The interval between attempts by the acting master to renew a leadership
	// slot before it stops leading. This must be less than or equal to the
	// lease duration.
	RenewDeadline *string `pulumi:"renewDeadline"`
	// The duration the clients should wait between attempting acquisition and
	// renewal of a leadership.
	RetryPeriod *string `pulumi:"retryPeriod"`
}

type CertManagerImage struct {
	// You can manage a registry with `registry: quay.io`.
	Registry *string `pulumi:"registry"`
	// You can manage a registry with `repository: jetstack/cert-manager-controller`.
	Repository *string `pulumi:"repository"`
	// Override the image tag to deploy by setting this variable.
	// If no value is set, the chart's appVersion will be used.
	Tag *string `pulumi:"tag"`
	// Setting a digest will override any tag, e.g.
	// `digest: sha256:0e072dddd1f7f8fc8909a2ca6f65e76c5f0d2fcfb8be47935ae3457e8bbceb20`.
	Digest     *string `pulumi:"digest"`
	PullPolicy *string `pulumi:"pullPolicy"`
}

type CertManagerServiceAccount struct {
	// Specifies whether a service account should be created
	Create *bool `pulumi:"create"`
	// The name of the service account to use.
	// If not set and create is true, a name is generated using the fullname template.
	Name *string `pulumi:"name"`
	// Optional additional annotations to add to the controller's ServiceAccount.
	Annotations *map[string]string `pulumi:"annotations"`
	// Automount API credentials for a Service Account.
	AutomountServiceAccountToken *bool `pulumi:"automountServiceAccountToken"`
}

type CertManagerIngressShim struct {
	DefaultIssuerName  *string `pulumi:"defaultIssuerName"`
	DefaultIssuerKind  *string `pulumi:"defaultIssuerKind"`
	DefaultIssuerGroup *string `pulumi:"defaultIssuerGroup"`
}

type CertManagerPrometheus struct {
	Enabled        *bool                                `pulumi:"enabled"`
	ServiceMonitor *CertManagerPrometheusServiceMonitor `pulumi:"serviceMonitor"`
}

type CertManagerPrometheusServiceMonitor struct {
	Enabled            *bool              `pulumi:"enabled"`
	PrometheusInstance *string            `pulumi:"prometheusInstance"`
	TargetPort         *int               `pulumi:"targetPort"`
	Path               *string            `pulumi:"path"`
	Interval           *string            `pulumi:"interval"`
	ScrapeTimeout      *string            `pulumi:"string"`
	Labels             *map[string]string `pulumi:"labels"`
}

type CertManagerWebhook struct {
	ReplicaCount   *int                       `pulumi:"replicaCount"`
	TimeoutSeconds *int                       `pulumi:"timeoutSeconds"`
	Strategy       *appsv1.DeploymentStrategy `pulumi:"strategy" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:apps/v1:DeploymentStrategy"`
	// Pod Security Context to be set on the webhook component Pod.
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SecurityContext *corev1.PodSecurityContext `pulumi:"securityContext" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:PodSecurityContext"`
	// Container Security Context to be set on the webhook component container.
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	ContainerSecurityContext *corev1.SecurityContext `pulumi:"containerSecurityContext" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:SecurityContext"`
	// Optional additional annotations to add to the webhook Deployment
	DeploymentAnnotations *map[string]string `pulumi:"deploymentAnnotations"`
	// Optional additional annotations to add to the webhook Pods
	PodAnnotations *map[string]string `pulumi:"podAnnotations"`
	// Optional additional annotations to add to the webhook MutatingWebhookConfiguration
	MutatingWebhookConfigurationAnnotations *map[string]string `pulumi:"mutatingWebhookConfigurationAnnotations"`
	// Optional additional annotations to add to the webhook ValidatingWebhookConfiguration
	ValidatingWebhookConfigurationAnnotations *map[string]string `pulumi:"validatingWebhookConfigurationAnnotations"`
	// Optional additional annotations to add to the webhook service
	ServiceAnnotations *map[string]string `pulumi:"serviceAnnotations"`
	// Optional additional arguments for webhook
	ExtraArgs *[]string                    `pulumi:"extraArgs"`
	Resources *corev1.ResourceRequirements `pulumi:"resources" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements"`
	// Liveness probe values.
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
	LivenessProbe *corev1.Probe `pulumi:"livenessProbe" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Probe"`
	// Readiness probe values.
	// Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
	ReadinessProbe *corev1.Probe        `pulumi:"readinessProbe" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Probe"`
	NodeSelector   *map[string]string   `pulumi:"nodeSelector"`
	Affinity       *corev1.Affinity     `pulumi:"affinity" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Affinity"`
	Tolerations    *[]corev1.Toleration `pulumi:"tolerations" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Toleration"`
	// Optional additional labels to add to the Webhook Pods
	PodLabels *map[string]string `pulumi:"podLabels"`
	// Optional additional labels to add to the Webhook Service
	ServiceLabels  *map[string]string         `pulumi:"serviceLabels"`
	Image          *CertManagerImage          `pulumi:"image"`
	ServiceAccount *CertManagerServiceAccount `pulumi:"serviceAccount"`
	// The port that the webhook should listen on for requests.
	// In GKE private clusters, by default kubernetes apiservers are allowed to
	// talk to the cluster nodes only on 443 and 10250. so configuring
	// securePort: 10250, will work out of the box without needing to add firewall
	// rules or requiring NET_BIND_SERVICE capabilities to bind port numbers <1000
	SecurePort *int `pulumi:"securePort"`
	// Specifies if the webhook should be started in hostNetwork mode.
	// Required for use in some managed kubernetes clusters (such as AWS EKS) with custom
	// CNI (such as calico), because control-plane managed by AWS cannot communicate
	// with pods' IP CIDR and admission webhooks are not working
	// Since the default port for the webhook conflicts with kubelet on the host
	// network, `webhook.securePort` should be changed to an available port if
	// running in hostNetwork mode.
	HostNetwork *bool `pulumi:"hostNetwork"`
	// Specifies how the service should be handled. Useful if you want to expose the
	// webhook to outside of the cluster. In some cases, the control plane cannot
	// reach internal services.
	ServiceType    *string `pulumi:"serviceType"`
	LoadBalancerIP *string `pulumi:"loadBalancerIP"`
	// Overrides the mutating webhook and validating webhook so they reach the webhook
	// service using the `url` field instead of a service.
	URL *CertManagerWebhookURL `pulumi:"url"`
}

type CertManagerWebhookURL struct {
	Host *string `pulumi:"host"`
}

type CertManagerCaInjector struct {
	ReplicaCount   *int                       `pulumi:"replicaCount"`
	TimeoutSeconds *int                       `pulumi:"timeoutSeconds"`
	Strategy       *appsv1.DeploymentStrategy `pulumi:"strategy" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:apps/v1:DeploymentStrategy"`
	// Pod Security Context to be set on the cainjector component Pod.
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SecurityContext *corev1.PodSecurityContext `pulumi:"podSecurityContext" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:PodSecurityContext"`
	// Container Security Context to be set on the cainjector component container.
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	ContainerSecurityContext *corev1.SecurityContext `pulumi:"containerSecurityContext" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:SecurityContext"`
	// Optional additional annotations to add to the cainjector Deployment
	DeploymentAnnotations *map[string]string `pulumi:"deploymentAnnotations"`
	// Optional additional annotations to add to the cainjector Pods
	PodAnnotations *map[string]string `pulumi:"podAnnotations"`
	// Optional additional arguments for cainjector
	ExtraArgs    *[]string                    `pulumi:"extraArgs"`
	Resources    *corev1.ResourceRequirements `pulumi:"resources" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements"`
	NodeSelector *map[string]string           `pulumi:"nodeSelector"`
	Affinity     *corev1.Affinity             `pulumi:"affinity" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Affinity"`
	Tolerations  *[]corev1.Toleration         `pulumi:"tolerations" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Toleration"`
	// Optional additional labels to add to the Webhook Pods
	PodLabels      *map[string]string         `pulumi:"podLabels"`
	Image          *CertManagerImage          `pulumi:"image"`
	ServiceAccount *CertManagerServiceAccount `pulumi:"serviceAccount"`
}

type CertManagerStartupAPICheck struct {
	Enabled *bool `pulumi:"enabled"`
	// Pod Security Context to be set on the startupapicheck component Pod.
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SecurityContext *corev1.PodSecurityContext `pulumi:"securityContext" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:PodSecurityContext"`
	// Timeout for 'kubectl check api' command
	Timeout *string `pulumi:"timeout"`
	// Job backoffLimit
	BackoffLimit *int `pulumi:"backoffLimit"`
	// Optional additional annotations to add to the startupapicheck Job
	JobAnnotations *map[string]string `pulumi:"jobAnnotations"`
	// Optional additional annotations to add to the startupapicheck Pods
	PodAnnotations *map[string]string `pulumi:"podAnnotations"`
	// Optional additional arguments for startupapicheck
	ExtraArgs    *[]string                    `pulumi:"extraArgs"`
	Resources    *corev1.ResourceRequirements `pulumi:"resources" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:ResourceRequirements"`
	NodeSelector *map[string]string           `pulumi:"nodeSelector"`
	Affinity     *corev1.Affinity             `pulumi:"affinity" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Affinity"`
	Tolerations  *[]corev1.Toleration         `pulumi:"tolerations" pschema:"ref=/kubernetes/v4.7.1/schema.json#/types/kubernetes:core/v1:Toleration"`
	// Optional additional labels to add to the startupapicheck Pods
	PodLabels      *map[string]string              `pulumi:"podLabels"`
	Image          *CertManagerImage               `pulumi:"image"`
	RBAC           *CertManagerStartupAPICheckRBAC `pulumi:"rbac"`
	ServiceAccount *CertManagerServiceAccount      `pulumi:"serviceAccount"`
}

type CertManagerStartupAPICheckRBAC struct {
	// annotations for the startup API Check job RBAC and PSP resources
	Annotations *map[string]string `pulumi:"annotations"`
}
