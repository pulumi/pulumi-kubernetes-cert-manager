// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.KubernetesCertManager.Inputs
{

    public sealed class CertManagerGlobalRbacArgs : global::Pulumi.ResourceArgs
    {
        [Input("create")]
        public Input<bool>? Create { get; set; }

        public CertManagerGlobalRbacArgs()
        {
        }
        public static new CertManagerGlobalRbacArgs Empty => new CertManagerGlobalRbacArgs();
    }
}
