// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// PlaylistsColumns holds the columns for the "playlists" table.
	PlaylistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
	}
	// PlaylistsTable holds the schema information for the "playlists" table.
	PlaylistsTable = &schema.Table{
		Name:       "playlists",
		Columns:    PlaylistsColumns,
		PrimaryKey: []*schema.Column{PlaylistsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "playlists_users_playlists",
				Columns: []*schema.Column{PlaylistsColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PlaylistVideosColumns holds the columns for the "playlist_videos" table.
	PlaylistVideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "playlist_id", Type: field.TypeInt, Nullable: true},
		{Name: "resolution_id", Type: field.TypeInt, Nullable: true},
		{Name: "User_id", Type: field.TypeInt, Nullable: true},
		{Name: "video_id", Type: field.TypeInt, Nullable: true},
	}
	// PlaylistVideosTable holds the schema information for the "playlist_videos" table.
	PlaylistVideosTable = &schema.Table{
		Name:       "playlist_videos",
		Columns:    PlaylistVideosColumns,
		PrimaryKey: []*schema.Column{PlaylistVideosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "playlist_videos_playlists_playlist_videos",
				Columns: []*schema.Column{PlaylistVideosColumns[2]},

				RefColumns: []*schema.Column{PlaylistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "playlist_videos_resolutions_playlist_videos",
				Columns: []*schema.Column{PlaylistVideosColumns[3]},

				RefColumns: []*schema.Column{ResolutionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "playlist_videos_users_playlist_videos",
				Columns: []*schema.Column{PlaylistVideosColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "playlist_videos_videos_playlist_videos",
				Columns: []*schema.Column{PlaylistVideosColumns[5]},

				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ResolutionsColumns holds the columns for the "resolutions" table.
	ResolutionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeInt},
	}
	// ResolutionsTable holds the schema information for the "resolutions" table.
	ResolutionsTable = &schema.Table{
		Name:        "resolutions",
		Columns:     ResolutionsColumns,
		PrimaryKey:  []*schema.Column{ResolutionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "student_id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "identification_number", Type: field.TypeString, Size: 13},
		{Name: "email", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// VideosColumns holds the columns for the "videos" table.
	VideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
	}
	// VideosTable holds the schema information for the "videos" table.
	VideosTable = &schema.Table{
		Name:       "videos",
		Columns:    VideosColumns,
		PrimaryKey: []*schema.Column{VideosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "videos_users_videos",
				Columns: []*schema.Column{VideosColumns[3]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PlaylistsTable,
		PlaylistVideosTable,
		ResolutionsTable,
		UsersTable,
		VideosTable,
	}
)

func init() {
	PlaylistsTable.ForeignKeys[0].RefTable = UsersTable
	PlaylistVideosTable.ForeignKeys[0].RefTable = PlaylistsTable
	PlaylistVideosTable.ForeignKeys[1].RefTable = ResolutionsTable
	PlaylistVideosTable.ForeignKeys[2].RefTable = UsersTable
	PlaylistVideosTable.ForeignKeys[3].RefTable = VideosTable
	VideosTable.ForeignKeys[0].RefTable = UsersTable
}