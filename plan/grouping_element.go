package plan

import (
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/parser"
	"github.com/xitongsys/guery/row"
)

type GroupingElementNode struct {
	Expression *ExpressionNode
}

func NewGroupingElementNode(runtime *config.ConfigRuntime, t parser.IGroupingElementContext) *GroupingElementNode {
	res := &GroupingElementNode{}
	tt := t.(*parser.GroupingElementContext).Expression()
	res.Expression = NewExpressionNode(runtime, tt)
	return res
}

func (self *GroupingElementNode) Init(md *metadata.Metadata) error {
	return self.Expression.Init(md)
}

func (self *GroupingElementNode) Result(input *row.RowsGroup) (interface{}, error) {
	return self.Expression.Result(input)
}

func (self *GroupingElementNode) GetColumns() ([]string, error) {
	return self.Expression.GetColumns()
}

func (self *GroupingElementNode) GetType(md *metadata.Metadata) (gtype.Type, error) {
	return self.Expression.GetType(md)
}

func (self *GroupingElementNode) IsAggregate() bool {
	return self.Expression.IsAggregate()
}
