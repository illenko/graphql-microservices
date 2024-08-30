package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"card-service/graph/generated"
	"card-service/graph/model"
	"context"
)

// Cards is the resolver for the cards field.
func (r *queryResolver) Cards(ctx context.Context) ([]*model.Card, error) {
	return cards, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
