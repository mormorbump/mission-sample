package entity

// Target ミッション対象か(クエストIDやキャラIDなど)を判別するIDと、更新する進捗値の情報
type Target struct {
	// ミッション対象のIDリスト
	// クエストIDやキャラIDなど
	ID TargetID

	// IDに対して加算や最大値更新するための進捗値
	// クエストクリア回数や、最大到達レベルなど
	Progress int64
}

// Targets
// IDと進捗情報のリスト
type Targets []*Target

type TargetID string
type TargetIDs []TargetID
