package planphysical

import (
	"context"
	"tiny_planner/pkg/c_sql/b_planner/plancore"
	"tiny_planner/pkg/c_sql/c_exec_engine"
)

type PhysicalPlan interface {
	plancore.Plan

	// ToPB converts the physical plan to a protobuf message.
	ToPB(ctx context.Context) (execution.Executor, error)
}

type basePhysicalPlan struct {
	plancore.BasePlan
}

func (p *basePhysicalPlan) ToPB(ctx context.Context) (execution.Executor, error) {
	executorBuilder := execution.NewExecutorBuilder(ctx, p.Schema())
	return executorBuilder.Build(p)
}

var _ PhysicalPlan = &basePhysicalPlan{}
var _ PhysicalPlan = &PhysicalSelection{}
var _ PhysicalPlan = &PhysicalProjection{}
var _ PhysicalPlan = &PhysicalTableReader{}

type PhysicalSelection struct {
	basePhysicalPlan
}

type PhysicalProjection struct {
	basePhysicalPlan
}

type PhysicalTableReader struct {
	basePhysicalPlan
}
