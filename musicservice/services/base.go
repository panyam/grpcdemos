package services

import (
	"fmt"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
	"sort"
	"time"
)

type EntityService[T any] struct {
	IDCount  int
	Entities map[string]*T

	// Getters/Setters for ID
	IDSetter func(entity *T, id string)
	IDGetter func(entity *T) string

	// Getters/Setters for created timestamp
	CreatedAtSetter func(entity *T, ts *tspb.Timestamp)
	CreatedAtGetter func(entity *T) *tspb.Timestamp

	// Getters/Setters for udpated timestamp
	UpdatedAtSetter func(entity *T, ts *tspb.Timestamp)
	UpdatedAtGetter func(entity *T) *tspb.Timestamp

	// A function that returns whether entity1 < entity2 for sorting
	LTFunc func(ent1, ent2 *T) bool
}

func NewEntityService[T any]() *EntityService[T] {
	return &EntityService[T]{
		Entities: make(map[string]*T),
	}
}

func (s *EntityService[T]) Create(entity *T) *T {
	s.IDCount++
	newid := fmt.Sprintf("%d", s.IDCount)
	s.IDSetter(entity, newid)
	s.CreatedAtSetter(entity, tspb.New(time.Now()))
	s.UpdatedAtSetter(entity, tspb.New(time.Now()))
	return entity
}

func (s *EntityService[T]) Get(id string) *T {
	if entity, ok := s.Entities[id]; ok {
		return entity
	}
	return nil
}

func (s *EntityService[T]) BatchGet(ids []string) map[string]*T {
	out := make(map[string]*T)
	for _, id := range ids {
		if entity, ok := s.Entities[id]; ok {
			out[id] = entity
		}
	}
	return out
}

// Updates specific fields of an Entity
func (s *EntityService[T]) Update(entity *T) *T {
	s.UpdatedAtSetter(entity, tspb.New(time.Now()))
	return entity
}

// Deletes an song from our system.
func (s *EntityService[T]) Delete(id string) bool {
	_, ok := s.Entities[id]
	if ok {
		delete(s.Entities, id)
	}
	return ok
}

// Finds and retrieves songs matching the particular criteria.
func (s *EntityService[T]) List(ltfunc func(t1, t2 *T) bool) (out []*T) {
	for _, ent := range s.Entities {
		out = append(out, ent)
	}
	if ltfunc == nil {
		ltfunc = s.LTFunc
	}
	// Sort in reverse order of name
	sort.Slice(out, func(idx1, idx2 int) bool {
		ent1 := out[idx1]
		ent2 := out[idx2]
		return ltfunc(ent1, ent2)
	})
	return
}
