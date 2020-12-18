package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"gofriday.dev/awesome/ent/confirmstatus"
)

// ConfirmMixin implements the ent.Mixin for sharing
// confirm with package schemas.
type ConfirmMixin struct {
	mixin.Schema
}

func (ConfirmMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("confirm_status").
			GoType(confirmstatus.StatusUnknown).
			Default(string(confirmstatus.StatusUnknown)).
			Optional(),
		field.Time("result_at").
			Optional(),
	}
}
