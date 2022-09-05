package graph

import "gqlgen_sample2/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	pokemon []*model.Pokemon
	battle  []*model.Battle
}
