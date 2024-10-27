package mission

import (
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/value"
)

// Form ミッションの更新をする際に呼び出し元から送られてくる依頼書。
// ミッションの種類、対象となる ID、更新に必要な進捗情報などが含まれる。
// これを元にUserMissionが更新されるイメージ
type Form struct {
	MissionType value.MissionType
	Targets     Targets
}

type Forms []*Form

// Target ミッション対象か(クエストIDやキャラIDなど)を判別するIDと、更新する進捗値の情報
type Target struct {
	// ミッション対象のIDリスト
	// クエストIDやキャラIDなど
	ID string

	// IDに対して加算や最大値更新するための進捗値
	// クエストクリア回数や、最大到達レベルなど
	Progress int64
}

// Targets
// IDと進捗情報のリスト
type Targets []*Target

func (f *Form) GetAggregateProgress(m *entity.Mission) int64 {
	var ret int64
	for _, t := range f.Targets {
		if m.IsTarget(t.ID) {
			ret += t.Progress
		}
	}
	return ret
}

func (f *Form) GetMaxProgress(m *entity.Mission) int64 {
	var ret int64
	for _, t := range f.Targets {
		if m.IsTarget(t.ID) && ret < t.Progress {
			ret = t.Progress
		}
	}
	return ret
}
