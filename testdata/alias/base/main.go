package base

// WithMapAny は map value として any（= *types.Alias）を持つ
type WithMapAny struct {
	Extra map[string]any `json:"extra"`
}

// WithSliceAny は slice element として any を持つ
type WithSliceAny struct {
	Items []any `json:"items"`
}

// UserID は named-basic alias（alias 名喪失挙動の固定）
type UserID = string

// WithUserID は named-basic alias 経由のフィールドを持つ
type WithUserID struct {
	ID UserID `json:"id"`
}

// WithHandler は *types.Signature を踏む（Any fallback 検証）
type WithHandler struct {
	H func(int) error `json:"handler"`
}

// Box は *types.TypeParam を踏む generic struct
type Box[T any] struct {
	V T `json:"v"`
}

// VecAny は Go 1.24+ generic alias（type alias with type parameters）
type VecAny[T any] = []T

// WithVecAny は generic alias を field に持つ
type WithVecAny struct {
	Items VecAny[int] `json:"items"`
}
