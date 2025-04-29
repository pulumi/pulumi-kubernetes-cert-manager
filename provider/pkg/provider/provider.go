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

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	pp "github.com/pulumi/pulumi/sdk/v3/go/pulumi/provider"
)

const (
	ProviderName  = "kubernetes-cert-manager"
	ComponentName = ProviderName + ":index:CertManager"
)

// Serve launches the gRPC server for the resource provider.
func Serve(version string, schema []byte) {
	if err := provider.ComponentMain(ProviderName, version, schema, Construct); err != nil {
		cmdutil.ExitError(err.Error())
	}
}

// Construct is the RPC call that initiates the creation of a new component resource. It
// creates, registers, and returns the resulting object.
func Construct(ctx *pulumi.Context, typ, name string, inputs pp.ConstructInputs,
	opts pulumi.ResourceOption) (*pp.ConstructResult, error) {
	args := &CertManagerArgs{}
	if err := inputs.CopyTo(args); err != nil {
		return nil, err
	}

	// Set default values for the Crds configuration
	// The cert-manager Helm chart is transitioning from using installCRDs (boolean)
	// to a structured object crds: { enabled: boolean, keep: boolean }
	//
	// This section handles both formats and ensures proper defaults are set.
	// For the structured format:
	// - crds.enabled (default: false) - Whether to install CRDs
	// - crds.keep (default: false) - Whether to keep CRDs after chart uninstall
	keepFalse := false
	enabledFalse := false

	// Initialize the Crds object if it doesn't exist
	if args.Crds == nil {
		args.Crds = &CertManagerCrds{
			Keep:    &keepFalse,    // Default: don't keep CRDs after uninstall
			Enabled: &enabledFalse, // Default: don't install CRDs
		}
	} else {
		// Ensure all fields have proper defaults set
		if args.Crds.Keep == nil {
			args.Crds.Keep = &keepFalse
		}
		if args.Crds.Enabled == nil {
			args.Crds.Enabled = &enabledFalse
		}
	}

	// Handle legacy installCRDs parameter for backward compatibility
	// For background: In the Helm chart, setting both installCRDs=true and crds.enabled=true
	// causes a conflict, so we need to handle this case specifically.
	if args.InstallCRDs != nil && *args.InstallCRDs {
		// If installCRDs is true, we set crds.enabled=true and clear installCRDs
		// to avoid sending conflicting configuration to the Helm chart
		enabledTrue := true
		args.Crds.Enabled = &enabledTrue
		args.InstallCRDs = nil
	}

	return helmbase.Construct(ctx, &CertManager{}, typ, name, args, inputs, opts)
}
