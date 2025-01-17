using System;
using System.Collections.Generic;
using Pulumi;
using Pulumi.Kubernetes.Core.V1;
using Pulumi.Kubernetes.Types.Inputs.Core.V1;
using Pulumi.Kubernetes.Types.Inputs.Meta.V1;
using Pulumi.KubernetesCertManager;
using Pulumi.KubernetesCertManager.Inputs;
using Pulumi.Random;
using Pulumi.Kubernetes.ApiExtensions;

return await Deployment.RunAsync(() =>
{
  var randomString = new RandomString("random", new RandomStringArgs
  {
    Length = 16,
    Special = false
  });

  var config = new Config();
  var confRepo = config.Get("repository");
  var repository = confRepo != null ? Output.Create(confRepo) : randomString.Result;

  // Create a sandbox namespace
  var ns = new Namespace("sandbox-ns");

  // Install cert manager into the cluster
  
  var manager = new CertManager("cert-manager", new CertManagerArgs
  {
    InstallCRDs = true,
    HelmOptions = new ReleaseArgs
    {
      Namespace = ns.Metadata.Apply(m => m.Name),
      Version = "v1.15.3"
    },
    Cainjector = new CertManagerCaInjectorArgs
    {
      Image = new CertManagerImageArgs
      {
        Repository = "public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-cainjector",
        Tag = "v1.15.3-eks-a-v0.21.3-dev-build.0"
      }
    },
    Startupapicheck = new CertManagerStartupAPICheckArgs
    {
      Image = new CertManagerImageArgs
      {
        Repository = "public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-startupapicheck",
        Tag = "v1.15.3-eks-a-v0.21.3-dev-build.0"
      }
    },
    Webhook = new CertManagerWebhookArgs
    {
      Image = new CertManagerImageArgs
      {
        Repository = "public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-webhook",
        Tag = "v1.15.3-eks-a-v0.21.3-dev-build.0"
      }
    }
  });

  // Convert status to a serializable format
  return new Dictionary<string, object?>
  {
    ["certManagerNamespace"] = ns.Metadata.Apply(m => m.Name),
    ["certManagerName"] = manager.Urn
  };
});
