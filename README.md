# Pulumi Cert Manager Component

This repo contains the Pulumi Cert Manager component for Kubernetes. This add-on automates the
management and issuance of TLS certificates from various issuing sources. It ensures certificates
are valid and up to date periodically, and attempts to renew certificates at an appropriate time
before expiry.

This component wraps [the Jetstack Cert Manager Helm Chart](https://github.com/jetstack/cert-manager),
and offers a Pulumi-friendly and strongly-typed way to manage Cert Manager installations.

For examples of usage, see [the official documentation](https://cert-manager.io/docs/),
or refer to [the examples](/examples) in this repo.

## To Use

To use this component, first install the Pulumi Package:

```bash
# Node.js (JavaScript/TypeScript)
npm install @pulumi/kubernetes-cert-manager

# Python
pip install pulumi_kubernetes_cert_manager

# Go
go get github.com/pulumi/pulumi-kubernetes-cert-manager/sdk/go/kubernetes-cert-manager
```

Afterwards, import the library and instantiate it within your Pulumi program:

### TypeScript
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as k8s from "@pulumi/kubernetes";
import * as certmanager from "@pulumi/kubernetes-cert-manager";

// Create a namespace for cert-manager
const ns = new k8s.core.v1.Namespace("cert-manager-namespace", {
    metadata: {
        name: "cert-system",
    }
});

// Install cert-manager with CRDs
const cm = new certmanager.CertManager("cert-manager-deployment", {
    // Option 1: Using the new recommended approach
    crds: {
        enabled: true,
        keep: true, // Set to true to keep CRDs after uninstall
    },
    
    // Option 2: Using deprecated option (not recommended)
    // installCRDs: true,
    
    helmOptions: {
        namespace: ns.metadata.name,
    },
}, { parent: ns });
```

### Python
```python
import pulumi
import pulumi_kubernetes as k8s
import pulumi_kubernetes_cert_manager as certmanager

# Create a namespace for cert-manager
ns = k8s.core.v1.Namespace("cert-manager-namespace",
    metadata={
        "name": "cert-system",
    })

# Install cert-manager with CRDs
cm = certmanager.CertManager("cert-manager-deployment",
    # Option 1: Using the new recommended approach
    crds={
        "enabled": True,
        "keep": True,  # Set to true to keep CRDs after uninstall
    },
    
    # Option 2: Using deprecated option (not recommended)
    # install_crds=True,
    
    helm_options={
        "namespace": ns.metadata["name"],
    })
```

### Go
```go
package main

import (
	kubernetes_cert_manager "github.com/pulumi/pulumi-kubernetes-cert-manager/sdk/go/kubernetes-cert-manager"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a namespace for cert-manager
		ns, err := corev1.NewNamespace(ctx, "cert-manager-namespace", &corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("cert-system"),
			},
		})
		if err != nil {
			return err
		}

		// Option 1: Using the new recommended approach
		enabled := true
		keep := true
		
		// Install cert-manager with CRDs
		_, err = kubernetes_cert_manager.NewCertManager(ctx, "cert-manager-deployment", &kubernetes_cert_manager.CertManagerArgs{
			Crds: &kubernetes_cert_manager.CertManagerCrdsArgs{
				Enabled: pulumi.BoolPtr(enabled),
				Keep:    pulumi.BoolPtr(keep), // Set to true to keep CRDs after uninstall
			},
			
			// Option 2: Using deprecated option (not recommended)
			// InstallCRDs: pulumi.BoolPtr(enabled),
			
			HelmOptions: &helmv3.ReleaseArgs{
				Namespace: ns.Metadata.Name(),
			},
		})
		
		return err
	})
}
```

## Configuration

This component supports all of the configuration options of the [official Helm chart](
https://github.com/jetstack/cert-manager/tree/master/deploy/charts/cert-manager), except that these
are strongly typed so you will get IDE support and static error checking.

The Helm deployment uses reasonable defaults, including the chart name and repo URL, however,
if you need to override them, you may do so using the `helmOptions` parameter. Refer to
[the API docs for the `kubernetes:helm/v3:Release` Pulumi type](
https://www.pulumi.com/docs/reference/pkg/kubernetes/helm/v3/release/#inputs) for a full set of choices.

For complete details, refer to the Pulumi Package details within the Pulumi Registry.
