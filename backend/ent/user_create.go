// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/playlist"
	"github.com/tanapon395/playlist-video/ent/playlist_video"
	"github.com/tanapon395/playlist-video/ent/user"
	"github.com/tanapon395/playlist-video/ent/video"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetStudentID sets the student_id field.
func (uc *UserCreate) SetStudentID(s string) *UserCreate {
	uc.mutation.SetStudentID(s)
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetIdentificationNumber sets the identification_number field.
func (uc *UserCreate) SetIdentificationNumber(s string) *UserCreate {
	uc.mutation.SetIdentificationNumber(s)
	return uc
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.mutation.SetAge(i)
	return uc
}

// AddPlaylistIDs adds the playlists edge to Playlist by ids.
func (uc *UserCreate) AddPlaylistIDs(ids ...int) *UserCreate {
	uc.mutation.AddPlaylistIDs(ids...)
	return uc
}

// AddPlaylists adds the playlists edges to Playlist.
func (uc *UserCreate) AddPlaylists(p ...*Playlist) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPlaylistIDs(ids...)
}

// AddVideoIDs adds the videos edge to Video by ids.
func (uc *UserCreate) AddVideoIDs(ids ...int) *UserCreate {
	uc.mutation.AddVideoIDs(ids...)
	return uc
}

// AddVideos adds the videos edges to Video.
func (uc *UserCreate) AddVideos(v ...*Video) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uc.AddVideoIDs(ids...)
}

// AddPlaylistVideoIDs adds the playlist_videos edge to Playlist_Video by ids.
func (uc *UserCreate) AddPlaylistVideoIDs(ids ...int) *UserCreate {
	uc.mutation.AddPlaylistVideoIDs(ids...)
	return uc
}

// AddPlaylistVideos adds the playlist_videos edges to Playlist_Video.
func (uc *UserCreate) AddPlaylistVideos(p ...*Playlist_Video) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPlaylistVideoIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if err := uc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) preSave() error {
	if _, ok := uc.mutation.StudentID(); !ok {
		return &ValidationError{Name: "student_id", err: errors.New("ent: missing required field \"student_id\"")}
	}
	if v, ok := uc.mutation.StudentID(); ok {
		if err := user.StudentIDValidator(v); err != nil {
			return &ValidationError{Name: "student_id", err: fmt.Errorf("ent: validator failed for field \"student_id\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := uc.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := uc.mutation.IdentificationNumber(); !ok {
		return &ValidationError{Name: "identification_number", err: errors.New("ent: missing required field \"identification_number\"")}
	}
	if v, ok := uc.mutation.IdentificationNumber(); ok {
		if err := user.IdentificationNumberValidator(v); err != nil {
			return &ValidationError{Name: "identification_number", err: fmt.Errorf("ent: validator failed for field \"identification_number\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	if v, ok := uc.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	u, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.StudentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldStudentID,
		})
		u.StudentID = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
		u.Name = value
	}
	if value, ok := uc.mutation.IdentificationNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldIdentificationNumber,
		})
		u.IdentificationNumber = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
		u.Email = value
	}
	if value, ok := uc.mutation.Age(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: user.FieldAge,
		})
		u.Age = value
	}
	if nodes := uc.mutation.PlaylistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PlaylistsTable,
			Columns: []string{user.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.VideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VideosTable,
			Columns: []string{user.VideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PlaylistVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PlaylistVideosTable,
			Columns: []string{user.PlaylistVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist_video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return u, _spec
}

// UserCreateBulk is the builder for creating a bulk of User entities.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
