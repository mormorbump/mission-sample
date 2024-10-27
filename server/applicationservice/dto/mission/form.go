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
	Targets     entity.Targets
}

type Forms []*Form

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
