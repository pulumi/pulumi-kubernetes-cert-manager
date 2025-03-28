// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetescertmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.kubernetes.apps.v1.inputs.DeploymentStrategyArgs;
import com.pulumi.kubernetes.core.v1.inputs.AffinityArgs;
import com.pulumi.kubernetes.core.v1.inputs.PodSecurityContextArgs;
import com.pulumi.kubernetes.core.v1.inputs.ResourceRequirementsArgs;
import com.pulumi.kubernetes.core.v1.inputs.SecurityContextArgs;
import com.pulumi.kubernetes.core.v1.inputs.TolerationArgs;
import com.pulumi.kubernetescertmanager.inputs.CertManagerImageArgs;
import com.pulumi.kubernetescertmanager.inputs.CertManagerServiceAccountArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertManagerCaInjectorArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertManagerCaInjectorArgs Empty = new CertManagerCaInjectorArgs();

    @Import(name="affinity")
    private @Nullable Output<AffinityArgs> affinity;

    public Optional<Output<AffinityArgs>> affinity() {
        return Optional.ofNullable(this.affinity);
    }

    /**
     * Container Security Context to be set on the cainjector component container. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
     * 
     */
    @Import(name="containerSecurityContext")
    private @Nullable Output<SecurityContextArgs> containerSecurityContext;

    /**
     * @return Container Security Context to be set on the cainjector component container. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
     * 
     */
    public Optional<Output<SecurityContextArgs>> containerSecurityContext() {
        return Optional.ofNullable(this.containerSecurityContext);
    }

    /**
     * Optional additional annotations to add to the cainjector Deployment
     * 
     */
    @Import(name="deploymentAnnotations")
    private @Nullable Output<Map<String,String>> deploymentAnnotations;

    /**
     * @return Optional additional annotations to add to the cainjector Deployment
     * 
     */
    public Optional<Output<Map<String,String>>> deploymentAnnotations() {
        return Optional.ofNullable(this.deploymentAnnotations);
    }

    /**
     * Optional additional arguments for cainjector
     * 
     */
    @Import(name="extraArgs")
    private @Nullable Output<List<String>> extraArgs;

    /**
     * @return Optional additional arguments for cainjector
     * 
     */
    public Optional<Output<List<String>>> extraArgs() {
        return Optional.ofNullable(this.extraArgs);
    }

    @Import(name="image")
    private @Nullable Output<CertManagerImageArgs> image;

    public Optional<Output<CertManagerImageArgs>> image() {
        return Optional.ofNullable(this.image);
    }

    @Import(name="nodeSelector")
    private @Nullable Output<Map<String,String>> nodeSelector;

    public Optional<Output<Map<String,String>>> nodeSelector() {
        return Optional.ofNullable(this.nodeSelector);
    }

    /**
     * Optional additional annotations to add to the cainjector Pods
     * 
     */
    @Import(name="podAnnotations")
    private @Nullable Output<Map<String,String>> podAnnotations;

    /**
     * @return Optional additional annotations to add to the cainjector Pods
     * 
     */
    public Optional<Output<Map<String,String>>> podAnnotations() {
        return Optional.ofNullable(this.podAnnotations);
    }

    /**
     * Optional additional labels to add to the Webhook Pods
     * 
     */
    @Import(name="podLabels")
    private @Nullable Output<Map<String,String>> podLabels;

    /**
     * @return Optional additional labels to add to the Webhook Pods
     * 
     */
    public Optional<Output<Map<String,String>>> podLabels() {
        return Optional.ofNullable(this.podLabels);
    }

    /**
     * Pod Security Context to be set on the cainjector component Pod. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
     * 
     */
    @Import(name="podSecurityContext")
    private @Nullable Output<PodSecurityContextArgs> podSecurityContext;

    /**
     * @return Pod Security Context to be set on the cainjector component Pod. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
     * 
     */
    public Optional<Output<PodSecurityContextArgs>> podSecurityContext() {
        return Optional.ofNullable(this.podSecurityContext);
    }

    @Import(name="replicaCount")
    private @Nullable Output<Integer> replicaCount;

    public Optional<Output<Integer>> replicaCount() {
        return Optional.ofNullable(this.replicaCount);
    }

    @Import(name="resources")
    private @Nullable Output<ResourceRequirementsArgs> resources;

    public Optional<Output<ResourceRequirementsArgs>> resources() {
        return Optional.ofNullable(this.resources);
    }

    @Import(name="serviceAccount")
    private @Nullable Output<CertManagerServiceAccountArgs> serviceAccount;

    public Optional<Output<CertManagerServiceAccountArgs>> serviceAccount() {
        return Optional.ofNullable(this.serviceAccount);
    }

    @Import(name="strategy")
    private @Nullable Output<DeploymentStrategyArgs> strategy;

    public Optional<Output<DeploymentStrategyArgs>> strategy() {
        return Optional.ofNullable(this.strategy);
    }

    @Import(name="timeoutSeconds")
    private @Nullable Output<Integer> timeoutSeconds;

    public Optional<Output<Integer>> timeoutSeconds() {
        return Optional.ofNullable(this.timeoutSeconds);
    }

    @Import(name="tolerations")
    private @Nullable Output<List<TolerationArgs>> tolerations;

    public Optional<Output<List<TolerationArgs>>> tolerations() {
        return Optional.ofNullable(this.tolerations);
    }

    private CertManagerCaInjectorArgs() {}

    private CertManagerCaInjectorArgs(CertManagerCaInjectorArgs $) {
        this.affinity = $.affinity;
        this.containerSecurityContext = $.containerSecurityContext;
        this.deploymentAnnotations = $.deploymentAnnotations;
        this.extraArgs = $.extraArgs;
        this.image = $.image;
        this.nodeSelector = $.nodeSelector;
        this.podAnnotations = $.podAnnotations;
        this.podLabels = $.podLabels;
        this.podSecurityContext = $.podSecurityContext;
        this.replicaCount = $.replicaCount;
        this.resources = $.resources;
        this.serviceAccount = $.serviceAccount;
        this.strategy = $.strategy;
        this.timeoutSeconds = $.timeoutSeconds;
        this.tolerations = $.tolerations;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertManagerCaInjectorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertManagerCaInjectorArgs $;

        public Builder() {
            $ = new CertManagerCaInjectorArgs();
        }

        public Builder(CertManagerCaInjectorArgs defaults) {
            $ = new CertManagerCaInjectorArgs(Objects.requireNonNull(defaults));
        }

        public Builder affinity(@Nullable Output<AffinityArgs> affinity) {
            $.affinity = affinity;
            return this;
        }

        public Builder affinity(AffinityArgs affinity) {
            return affinity(Output.of(affinity));
        }

        /**
         * @param containerSecurityContext Container Security Context to be set on the cainjector component container. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
         * 
         * @return builder
         * 
         */
        public Builder containerSecurityContext(@Nullable Output<SecurityContextArgs> containerSecurityContext) {
            $.containerSecurityContext = containerSecurityContext;
            return this;
        }

        /**
         * @param containerSecurityContext Container Security Context to be set on the cainjector component container. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
         * 
         * @return builder
         * 
         */
        public Builder containerSecurityContext(SecurityContextArgs containerSecurityContext) {
            return containerSecurityContext(Output.of(containerSecurityContext));
        }

        /**
         * @param deploymentAnnotations Optional additional annotations to add to the cainjector Deployment
         * 
         * @return builder
         * 
         */
        public Builder deploymentAnnotations(@Nullable Output<Map<String,String>> deploymentAnnotations) {
            $.deploymentAnnotations = deploymentAnnotations;
            return this;
        }

        /**
         * @param deploymentAnnotations Optional additional annotations to add to the cainjector Deployment
         * 
         * @return builder
         * 
         */
        public Builder deploymentAnnotations(Map<String,String> deploymentAnnotations) {
            return deploymentAnnotations(Output.of(deploymentAnnotations));
        }

        /**
         * @param extraArgs Optional additional arguments for cainjector
         * 
         * @return builder
         * 
         */
        public Builder extraArgs(@Nullable Output<List<String>> extraArgs) {
            $.extraArgs = extraArgs;
            return this;
        }

        /**
         * @param extraArgs Optional additional arguments for cainjector
         * 
         * @return builder
         * 
         */
        public Builder extraArgs(List<String> extraArgs) {
            return extraArgs(Output.of(extraArgs));
        }

        /**
         * @param extraArgs Optional additional arguments for cainjector
         * 
         * @return builder
         * 
         */
        public Builder extraArgs(String... extraArgs) {
            return extraArgs(List.of(extraArgs));
        }

        public Builder image(@Nullable Output<CertManagerImageArgs> image) {
            $.image = image;
            return this;
        }

        public Builder image(CertManagerImageArgs image) {
            return image(Output.of(image));
        }

        public Builder nodeSelector(@Nullable Output<Map<String,String>> nodeSelector) {
            $.nodeSelector = nodeSelector;
            return this;
        }

        public Builder nodeSelector(Map<String,String> nodeSelector) {
            return nodeSelector(Output.of(nodeSelector));
        }

        /**
         * @param podAnnotations Optional additional annotations to add to the cainjector Pods
         * 
         * @return builder
         * 
         */
        public Builder podAnnotations(@Nullable Output<Map<String,String>> podAnnotations) {
            $.podAnnotations = podAnnotations;
            return this;
        }

        /**
         * @param podAnnotations Optional additional annotations to add to the cainjector Pods
         * 
         * @return builder
         * 
         */
        public Builder podAnnotations(Map<String,String> podAnnotations) {
            return podAnnotations(Output.of(podAnnotations));
        }

        /**
         * @param podLabels Optional additional labels to add to the Webhook Pods
         * 
         * @return builder
         * 
         */
        public Builder podLabels(@Nullable Output<Map<String,String>> podLabels) {
            $.podLabels = podLabels;
            return this;
        }

        /**
         * @param podLabels Optional additional labels to add to the Webhook Pods
         * 
         * @return builder
         * 
         */
        public Builder podLabels(Map<String,String> podLabels) {
            return podLabels(Output.of(podLabels));
        }

        /**
         * @param podSecurityContext Pod Security Context to be set on the cainjector component Pod. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
         * 
         * @return builder
         * 
         */
        public Builder podSecurityContext(@Nullable Output<PodSecurityContextArgs> podSecurityContext) {
            $.podSecurityContext = podSecurityContext;
            return this;
        }

        /**
         * @param podSecurityContext Pod Security Context to be set on the cainjector component Pod. ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
         * 
         * @return builder
         * 
         */
        public Builder podSecurityContext(PodSecurityContextArgs podSecurityContext) {
            return podSecurityContext(Output.of(podSecurityContext));
        }

        public Builder replicaCount(@Nullable Output<Integer> replicaCount) {
            $.replicaCount = replicaCount;
            return this;
        }

        public Builder replicaCount(Integer replicaCount) {
            return replicaCount(Output.of(replicaCount));
        }

        public Builder resources(@Nullable Output<ResourceRequirementsArgs> resources) {
            $.resources = resources;
            return this;
        }

        public Builder resources(ResourceRequirementsArgs resources) {
            return resources(Output.of(resources));
        }

        public Builder serviceAccount(@Nullable Output<CertManagerServiceAccountArgs> serviceAccount) {
            $.serviceAccount = serviceAccount;
            return this;
        }

        public Builder serviceAccount(CertManagerServiceAccountArgs serviceAccount) {
            return serviceAccount(Output.of(serviceAccount));
        }

        public Builder strategy(@Nullable Output<DeploymentStrategyArgs> strategy) {
            $.strategy = strategy;
            return this;
        }

        public Builder strategy(DeploymentStrategyArgs strategy) {
            return strategy(Output.of(strategy));
        }

        public Builder timeoutSeconds(@Nullable Output<Integer> timeoutSeconds) {
            $.timeoutSeconds = timeoutSeconds;
            return this;
        }

        public Builder timeoutSeconds(Integer timeoutSeconds) {
            return timeoutSeconds(Output.of(timeoutSeconds));
        }

        public Builder tolerations(@Nullable Output<List<TolerationArgs>> tolerations) {
            $.tolerations = tolerations;
            return this;
        }

        public Builder tolerations(List<TolerationArgs> tolerations) {
            return tolerations(Output.of(tolerations));
        }

        public Builder tolerations(TolerationArgs... tolerations) {
            return tolerations(List.of(tolerations));
        }

        public CertManagerCaInjectorArgs build() {
            return $;
        }
    }

}
