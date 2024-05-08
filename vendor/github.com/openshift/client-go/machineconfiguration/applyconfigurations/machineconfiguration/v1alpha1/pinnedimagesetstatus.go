// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PinnedImageSetStatusApplyConfiguration represents an declarative configuration of the PinnedImageSetStatus type for use
// with apply.
type PinnedImageSetStatusApplyConfiguration struct {
	Conditions []v1.Condition `json:"conditions,omitempty"`
}

// PinnedImageSetStatusApplyConfiguration constructs an declarative configuration of the PinnedImageSetStatus type for use with
// apply.
func PinnedImageSetStatus() *PinnedImageSetStatusApplyConfiguration {
	return &PinnedImageSetStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *PinnedImageSetStatusApplyConfiguration) WithConditions(values ...v1.Condition) *PinnedImageSetStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}