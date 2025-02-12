package container

import (
	"grapql-to-do/graph/resolver"
	"grapql-to-do/graph/schema"
)

// Resolver は全リゾルバを統合する
type Resolver struct {
	*resolver.UserMutationResolver
	*resolver.TodoMutationResolver
	*resolver.TodoQueryResolver
}

func NewResolver() *Resolver {
	return &Resolver{
		UserMutationResolver: &resolver.UserMutationResolver{},
		TodoMutationResolver: &resolver.TodoMutationResolver{},
		TodoQueryResolver:    &resolver.TodoQueryResolver{},
	}
}

// Mutation を実装（schema.ResolverRoot を満たす）
func (r *Resolver) Mutation() schema.MutationResolver {
	return r
}

// Query を実装（schema.ResolverRoot を満たす）
func (r *Resolver) Query() schema.QueryResolver {
	return r
}
