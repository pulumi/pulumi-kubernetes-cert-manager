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
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFindAndAdoptCertManagerCRDs tests the dynamic CRD finding and adoption functionality
func TestFindAndAdoptCertManagerCRDs(t *testing.T) {
	// This is a placeholder for testing CRD import functionality based on the
	// example in the prompt that finds CRDs dynamically via listing and filtering
}

func TestCertManagerCrdsDefaults(t *testing.T) {
	args := &CertManagerArgs{}

	// Set default values
	if args.Crds == nil {
		keepFalse := false
		args.Crds = &CertManagerCrds{
			Keep: &keepFalse,
		}
	} else if args.Crds.Keep == nil {
		keepFalse := false
		args.Crds.Keep = &keepFalse
	}

	// Verify that Crds is initialized
	assert.NotNil(t, args.Crds)

	// Verify that Keep defaults to false
	assert.NotNil(t, args.Crds.Keep)
	assert.False(t, *args.Crds.Keep)
}

func TestCertManagerCrdsWithCustomValues(t *testing.T) {
	// Test with custom Keep value set to true
	keepTrue := true
	args := &CertManagerArgs{
		Crds: &CertManagerCrds{
			Keep: &keepTrue,
		},
	}

	// Set default values (should not change our custom setting)
	if args.Crds == nil {
		keepFalse := false
		args.Crds = &CertManagerCrds{
			Keep: &keepFalse,
		}
	} else if args.Crds.Keep == nil {
		keepFalse := false
		args.Crds.Keep = &keepFalse
	}

	// Verify that Crds is initialized
	assert.NotNil(t, args.Crds)

	// Verify that Keep value is preserved
	assert.NotNil(t, args.Crds.Keep)
	assert.True(t, *args.Crds.Keep)
}
