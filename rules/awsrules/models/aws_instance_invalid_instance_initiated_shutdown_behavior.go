// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule checks the pattern is valid
type AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule returns new rule with default attributes
func NewAwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule() *AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule {
	return &AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule{
		resourceType:  "aws_instance",
		attributeName: "instance_initiated_shutdown_behavior",
		enum: []string{
			"stop",
			"terminate",
		},
	}
}

// Name returns the rule name
func (r *AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule) Name() string {
	return "aws_instance_invalid_instance_initiated_shutdown_behavior"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsInstanceInvalidInstanceInitiatedShutdownBehaviorRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`instance_initiated_shutdown_behavior is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}