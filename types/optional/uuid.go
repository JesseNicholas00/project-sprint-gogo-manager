package optional

import (
	"encoding/json"
	"github.com/google/uuid"
)

type OptionalUUID struct {
	Defined bool
	V       *uuid.UUID
}

func (o *OptionalUUID) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return json.Unmarshal(data, &o.V)
}
