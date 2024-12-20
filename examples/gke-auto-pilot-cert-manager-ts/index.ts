import * as k8s from "@pulumi/kubernetes";
import * as certmanager from "@pulumi/kubernetes-cert-manager";

// Create a sandbox namespace.
const ns = new k8s.core.v1.Namespace("gke-ap-cert-mgr");

// Install a cert manager into our cluster.
const manager = new certmanager.CertManager("cert-manager", {
    installCRDs: true,
    helmOptions: {
        namespace: ns.metadata.name,
        version: "v1.15.3"
    },
    global: {
        leaderElection: {
            namespace: ns.metadata.name,
        }
    },
    prometheus: {
        enabled: false,
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
