package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"vetPartnersGraphQL/db/mongo"
	"vetPartnersGraphQL/graph/generated"
	"vetPartnersGraphQL/graph/model"
)

var db = mongo.Connect()

func (r *mutationResolver) CreateAnimal(ctx context.Context, input *model.NewAnimal) (*model.Animals, error) {
	return db.Save(input), nil
}

func (r *queryResolver) SingleAnimal(ctx context.Context, id string) (*model.Animals, error) {
	return db.FetchAnimal(id), nil
}

func (r *queryResolver) AllAnimals(ctx context.Context) ([]*model.Animals, error) {
	return db.FetchAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
