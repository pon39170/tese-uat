package schema

import (
	"errors"
	"regexp"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("student_id").Validate(func(s string) error {
			match, _ := regexp.MatchString("[BMD]\\d{7}", s)
			if !match {
				return errors.New("รูปแบบรหัสนักศึกษาไม่ถูกต้อง")
			}
			return nil
		}),
		// field.String("student_id").Match(regexp.MustCompile("[BMD]\\d{7}")),
		field.String("name").
		Validate(func(s string) error {
			
			if len(s) == 0 {
				return errors.New("โปรดกรอกชื่อและนามสกุล")
			}
			return nil
		}).NotEmpty(),
		field.String("identification_number").
		Validate(func(s string) error {
			
			if len(s) != 13 {
				return errors.New("รหัสประจำตัวประชาชนไม่ครบ")
			}
			return nil
		}).MaxLen(13).MinLen(13),
		field.String("email").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$",s)
			if !match {
				return errors.New("Email ไม่ถูกต้อง")
			}
			return nil
		}),
		field.Int("age").
		Validate(func(s int) error {
			
			if s <= 0 {
				return errors.New("อายุไม่ถูกต้อง")
			}
			return nil
		}).Min(1),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("playlists", Playlist.Type).StorageKey(edge.Column("owner_id")),
		edge.To("videos", Video.Type).StorageKey(edge.Column("owner_id")),
		edge.To("playlist_videos", Playlist_Video.Type).StorageKey(edge.Column("User_id")),
	}
}
