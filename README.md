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

Afterwards, import the library and instantiate it within your Pulumi program:

## Configuration

This component supports all of the configuration options of the [official Helm chart](
https://github.com/jetstack/cert-manager/tree/master/deploy/charts/cert-manager), except that these
are strongly typed so you will get IDE support and static error checking.

### CRDs Configuration

The component handles Custom Resource Definitions (CRDs) for cert-manager in two ways:

1. **Modern approach (recommended)**: Use the structured `crds` object:
   ```typescript
   const manager = new certmanager.CertManager("cert-manager", {
     crds: {
       enabled: true,  // Whether to install CRDs (default: false)
       keep: false,    // Whether to keep CRDs after uninstall (default: false)
     },
     // Other configuration...
   });
   ```

2. **Legacy approach (deprecated)**: Use the boolean `installCRDs` parameter:
   ```typescript
   const manager = new certmanager.CertManager("cert-manager", {
     installCRDs: true,
     // Other configuration...
   });
   ```

The component handles both approaches correctly, but the structured `crds` object is preferred for new deployments as it offers more fine-grained control.

### Other Configuration

The Helm deployment uses reasonable defaults, including the chart name and repo URL, however,
if you need to override them, you may do so using the `helmOptions` parameter. Refer to
[the API docs for the `kubernetes:helm/v3:Release` Pulumi type](
https://www.pulumi.com/docs/reference/pkg/kubernetes/helm/v3/release/#inputs) for a full set of choices.

For complete details, refer to the Pulumi Package details within the Pulumi Registry.
