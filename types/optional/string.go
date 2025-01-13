package optional

import "encoding/json"

type OptionalStr struct {
	Defined bool
	V       *string
}

func (o *OptionalStr) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return json.Unmarshal(data, &o.V)
}
