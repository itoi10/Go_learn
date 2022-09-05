package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen_sample2/graph/generated"
	"gqlgen_sample2/graph/model"
	"strconv"
)

var PokemonData []*model.Pokemon
var BattleData []*model.Battle
var pokemonId = 0

// CreatePokemon is the resolver for the CreatePokemon field.
func (r *mutationResolver) CreatePokemon(ctx context.Context, input model.NewPokemon) (*model.Pokemon, error) {
	pokemonId++
	newPokemon := model.Pokemon{
		ID:          strconv.Itoa(pokemonId),
		Name:        input.Name,
		Power:       input.Power,
		Description: &input.Description,
	}
	PokemonData = append(PokemonData, &newPokemon)
	return &newPokemon, nil
}

// DeletePokemon is the resolver for the DeletePokemon field.
func (r *mutationResolver) DeletePokemon(ctx context.Context, id string) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented: DeletePokemon - DeletePokemon"))
}

// UpdatePokemon is the resolver for the UpdatePokemon field.
func (r *mutationResolver) UpdatePokemon(ctx context.Context, input model.NewPokemon) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented: UpdatePokemon - UpdatePokemon"))
}

// CreateBattle is the resolver for the CreateBattle field.
func (r *mutationResolver) CreateBattle(ctx context.Context, input model.NewBattle) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented: CreateBattle - CreateBattle"))
}

// DeleteBattle is the resolver for the DeleteBattle field.
func (r *mutationResolver) DeleteBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented: DeleteBattle - DeleteBattle"))
}

// UpdateBattle is the resolver for the UpdateBattle field.
func (r *mutationResolver) UpdateBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented: UpdateBattle - UpdateBattle"))
}

// GetPokemon is the resolver for the GetPokemon field.
func (r *queryResolver) GetPokemon(ctx context.Context, id string) (*model.Pokemon, error) {
	panic(fmt.Errorf("not implemented: GetPokemon - GetPokemon"))
}

// ListPokemon is the resolver for the ListPokemon field.
func (r *queryResolver) ListPokemon(ctx context.Context) ([]*model.Pokemon, error) {
	return PokemonData, nil
}

// GetBattle is the resolver for the GetBattle field.
func (r *queryResolver) GetBattle(ctx context.Context, id string) (*model.Battle, error) {
	panic(fmt.Errorf("not implemented: GetBattle - GetBattle"))
}

// ListBattle is the resolver for the ListBattle field.
func (r *queryResolver) ListBattle(ctx context.Context) ([]*model.Battle, error) {
	panic(fmt.Errorf("not implemented: ListBattle - ListBattle"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
