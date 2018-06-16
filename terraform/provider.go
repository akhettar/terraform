package terraform

import (
	"github.com/hashicorp/terraform/tfdiags"
	"github.com/zclconf/go-cty/cty"
)

// Provider represents the set of methods required for a complete resource
// provider plugin.
type Provider interface {
	// GetSchema returns the complete schema for the provider.
	GetSchema() (*ProviderSchema, error)

	// ValidateProviderConfig allows the provider to validate the provider
	// configuration values.
	ValidateProviderConfig(config cty.Value) tfdiags.Diagnostics

	// ValidateResourceTypeConfig allows the provider to validate the resource
	// configuration values.
	ValidateResourceTypeConfig(name string, config cty.Value) tfdiags.Diagnostics

	// ValidateDataSource allows the provider to validate the data source
	// configuration values.
	ValidateDataSourceConfig(name string, config cty.Value) tfdiags.Diagnostics

	// UpgradeResourceState is called when the state loader encounters an
	// instance state whose schema version is less than the one reported by the
	// currently-used version of the corresponding provider, and the upgraded
	// result is used for any further processing.
	UpgradeResourceState(name string, version int, state cty.Value) (cty.Value, tfdiags.Diagnostics)

	// Configure configures and initialized the provider.
	Configure(config cty.Value) tfdiags.Diagnostics

	// Stop is called when the provider should halt any in-flight actions.
	//
	// Stop should not block waiting for in-flight actions to complete. It
	// should take any action it wants and return immediately acknowledging it
	// has received the stop request. Terraform will not make any further API
	// calls to the provider after Stop is called.
	//
	// The error returned, if non-nil, is assumed to mean that signaling the
	// stop somehow failed and that the user should expect potentially waiting
	// a longer period of time.
	Stop() error

	// ReadResource refreshes a resource and returns its current state.
	ReadResource(name string, prior cty.Value) (state cty.Value, diags tfdiags.Diagnostics)

	// PlanResourceChange takes the current state and proposed state of a
	// resource, and returns the planned final state.
	//
	// The private argument is an opaque blob returned by the provider from a
	// previous apply of the same resource.
	PlanResourceChange(name string, prior, proposed cty.Value, private []byte) (state cty.Value, diags tfdiags.Diagnostics)

	// ApplyResourceChange takes the planned state for a resource, which may
	// yet contain unknown computed values, and applies the changes returning
	// the final state.
	//
	// The private argument is an opaque blob which was returned from a prior
	// apply execution, and will match the value provided to
	// PlanResourceChange. The returned private value will be saved unchanged
	// in the state, replacing any prior values.
	ApplyResourceChange(name string, prior, planned cty.Value, privatePrior []byte) (state cty.Value, private []byte, diags tfdiags.Diagnostics)

	// ImportResourceState requests that the given resource be imported.
	ImportResourceState(name, id string) (state []cty.Value, diags tfdiags.Diagnostics)

	// ReadDataSource returns the data source's current state.
	ReadDataSource(name string) (state cty.Value, diags tfdiags.Diagnostics)
}
