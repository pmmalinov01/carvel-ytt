package texttemplate

import (
	"fmt"

	"github.com/k14s/ytt/pkg/template"
)

const (
	EvaluationCtxDialectName template.EvaluationCtxDialectName = "text"
)

type EvaluationCtx struct{}

var _ template.EvaluationCtxDialect = EvaluationCtx{}

func (e EvaluationCtx) Replace(
	parentNodes []template.EvaluationNode, val interface{}) error {

	return fmt.Errorf("unsupported operation")
}

func (e EvaluationCtx) ShouldWrapRootValue(nodeVal interface{}) bool {
	_, root := nodeVal.(*NodeRoot)
	return !root
}

func (e EvaluationCtx) WrapRootValue(val interface{}) interface{} {
	if typedVal, ok := val.(*NodeRoot); ok {
		return typedVal.AsString()
	}
	panic(fmt.Sprintf("Unexpected root value %T", val))
}

func (e EvaluationCtx) PrepareNode(
	parentNode template.EvaluationNode, node template.EvaluationNode) error {

	return nil
}
