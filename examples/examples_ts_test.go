//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optpreview"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optrefresh"
)

func TestTsExamples(t *testing.T) {
	tests := map[string]struct {
		directoryName    string
		additionalConfig map[string]string
	}{
		"TestSimpleCertManagerTs": {directoryName: "simple-cert-manager-ts"},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			p := pulumitest.NewPulumiTest(t, test.directoryName,
				opttest.LocalProviderPath("pulumi-kubernetes-cert-manager", filepath.Join(getCwd(t), "..", "bin")),
				opttest.YarnLink("@pulumi/kubernetes-cert-manager"),
			)
			if test.additionalConfig != nil {
				for key, value := range test.additionalConfig {
					p.SetConfig(t, key, value)
				}
			}
			p.SetConfig(t, "repository", "public.ecr.aws/eks-anywhere-dev/cert-manager/cert-manager-controller")
			p.Up(t)
			p.Preview(t, optpreview.ExpectNoChanges())
			p.Refresh(t, optrefresh.ExpectNoChanges())
		})
	}
}

// This tests the Output being passed to repository to fix #133
func TestTsCertManagerPreview(t *testing.T) {
	t.Run("TestSimpleCertManagerTsPreview", func(t *testing.T) {
		p := pulumitest.NewPulumiTest(t, "simple-cert-manager-ts",
			opttest.LocalProviderPath("pulumi-kubernetes-cert-manager", filepath.Join(getCwd(t), "..", "bin")),
			opttest.YarnLink("@pulumi/kubernetes-cert-manager"),
		)
		p.Preview(t, optpreview.ExpectNoChanges())
	})
}
