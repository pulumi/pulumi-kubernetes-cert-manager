package main

import (
	certmanager "github.com/pulumi/pulumi-kubernetes-cert-manager/sdk/go/kubernetes-cert-manager"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/apiextensions"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Create a namespace
		ns, err := corev1.NewNamespace(ctx, "sandbox-ns", &corev1.NamespaceArgs{})
		if err != nil {
			return err
		}

		// Install cert-manager
		manager, err := certmanager.NewCertManager(ctx, "cert-manager", &certmanager.CertManagerArgs{
			InstallCRDs: pulumi.Bool(true),
			HelmOptions: &certmanager.ReleaseArgs{
				Namespace: ns.Metadata.Name(),
				Version:   pulumi.String("v1.15.3"),
			},
			Image: &certmanager.CertManagerImageArgs{
				Repository: pulumi.String("public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-controller"),
				Tag:        pulumi.String("v1.15.3-eks-a-v0.21.3-dev-build.0"),
			},
			Cainjector: &certmanager.CertManagerCaInjectorArgs{
				Image: &certmanager.CertManagerImageArgs{
					Repository: pulumi.String("public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-cainjector"),
					Tag:        pulumi.String("v1.15.3-eks-a-v0.21.3-dev-build.0"),
				},
			},
			Startupapicheck: &certmanager.CertManagerStartupAPICheckArgs{
				Image: certmanager.CertManagerImageArgs{
					Repository: pulumi.String("public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-startupapicheck"),
					Tag:        pulumi.String("v1.15.3-eks-a-v0.21.3-dev-build.0"),
				},
			},
			Webhook: &certmanager.CertManagerWebhookArgs{
				Image: certmanager.CertManagerImageArgs{
					Repository: pulumi.String("public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-webhook"),
					Tag:        pulumi.String("v1.15.3-eks-a-v0.21.3-dev-build.0"),
				},
			},
		})

		if err != nil {
			return err
		}

		// Create a self-signed Issuer
		apiextensions.NewCustomResource(ctx, "issuer", &apiextensions.CustomResourceArgs{
			ApiVersion: pulumi.String("cert-manager.io/v1"),
			Kind:       pulumi.String("Issuer"),
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("selfsigned-issuer"),
				Namespace: ns.Metadata.Name(),
			},
			OtherFields: kubernetes.UntypedArgs{
				"spec": map[string]interface{}{
					"selfSigned": map[string]interface{}{},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{manager}))

		// Export the status of cert-manager
		ctx.Export("certManagerStatus", manager.Status)

		return nil
	})
}
