import * as k8s from "@pulumi/kubernetes";
import * as certmanager from "@pulumi/kubernetes-cert-manager";

// Create a sandbox namespace.
const ns = new k8s.core.v1.Namespace("sandbox-ns");

// Install a cert manager into our cluster.
const manager = new certmanager.CertManager("cert-manager", {
    installCRDs: true,
    helmOptions: {
        namespace: ns.metadata.name,
        version: "v1.15.3",
    },
    image: {
        repository: "public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-controller",
        tag: "v1.15.3-eks-a-v0.21.3-dev-build.0"
    },
    cainjector: {
        "image": {
            repository: "public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-cainjector",
            tag: "v1.15.3-eks-a-v0.21.3-dev-build.0",
        },
    },
    startupapicheck: {
        "image": {
            repository: "public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-startupapicheck",
            tag: "v1.15.3-eks-a-v0.21.3-dev-build.0",
        }
    },
    webhook: {
        image: {
            repository: "public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-webhook",
            tag: "v1.15.3-eks-a-v0.21.3-dev-build.0"
        }
    }
});

// Create a cluster issuer that uses self-signed certificates.
// This is not very secure, but has the least amount of external
// dependencies, so is simple. Please refer to
// https://cert-manager.io/docs/configuration/selfsigned/
// for additional details on other signing providers.
const issuer = new k8s.apiextensions.CustomResource(
    "issuer",
    {
        apiVersion: "cert-manager.io/v1",
        kind: "Issuer",
        metadata: {
            name: "selfsigned-issuer",
            namespace: ns.metadata.name,
        },
        spec: {
            selfSigned: {},
        },
    },
    { dependsOn: manager }
);

export const certManagerStatus = manager.status;
