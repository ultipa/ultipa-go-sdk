package http

import (
	ultipa "ultipa-go-sdk/rpc"
	"ultipa-go-sdk/sdk/structs"
)

type ExplainPlan struct {
	Explain []*structs.Explain
}

func ParseExplainPlan(ex *ultipa.ExplainPlan) (*ExplainPlan, error) {
	explainPlan := ExplainPlan{
		Explain: []*structs.Explain{},
	}

	if ex == nil {
		return &explainPlan, nil
	}

	for _, planNode := range ex.PlanNodes {
		explain := structs.Explain{
			//Type:        planNode.GetType(),
			Alias:       planNode.GetAlias(),
			ChildrenNum: planNode.GetChildrenNum(),
			Uql:         planNode.GetUql(),
			Infos:       planNode.GetInfos(),
		}
		explainPlan.Explain = append(explainPlan.Explain, &explain)
	}

	return &explainPlan, nil
}
