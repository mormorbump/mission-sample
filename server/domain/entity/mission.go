package entity

import (
	"com.graffity/mission-sample/server/applicationservice/dto/mission"
	"com.graffity/mission-sample/server/domain/value"
)

// Mission ミッションの種類やミッションの判定対象を格納
type Mission struct {
	ID          MissionID
	MissionType value.MissionType
	// ミッション対象のIDリスト。null許可
	// クエストIDやキャラクターIDなどが入り、それはMissionTypeにより対象が異なる
	TargetID string
	Name     string
}

type Missions []*Mission
type MissionID string
type MissionIDs []MissionID

type MissionPK struct {
	MissionID MissionID
}

type MissionPKs []*MissionPK

func (m *Mission) IsTarget(targetID string) bool {
	return m.TargetID == "" || m.TargetID == targetID
}

func (ms Missions) FilterByMissionType(missionType value.MissionType) Missions {
	ret := make(Missions, 0)
	for _, m := range ms {
		if m.MissionType == missionType {
			ret = append(ret, m)
		}
	}
	return ret
}

// FilterByTargets MissionsからtargetIDがnilまたは一致するものをFilter
// ここのレシーバのMissionsはMissionTypeで絞った後を想定しているので、Targetsの中のTargetIDの型は一致する
func (ms Missions) FilterByTargets(targets mission.Targets) Missions {
	ret := make(Missions, 0)
	for _, t := range targets {
		ret = append(ret, ms.FilterByTargetID(t.ID)...)
	}
	return ret
}

func (ms Missions) FilterByTargetID(targetID string) Missions {
	ret := make(Missions, 0, len(ms))
	for _, m := range ms {
		if m.IsTarget(targetID) {
			ret = append(ret, m)
		}
	}
	return ret
}

func (ms Missions) GetIDs() []string {
	ret := make([]string, 0, len(ms))
	for _, m := range ms {
		ret = append(ret, string(m.ID))
	}
	return ret
}

func (ms Missions) ToMapByMissionType() map[value.MissionType]Missions {
	ret := make(map[value.MissionType]Missions)
	for _, m := range ms {
		missions, ok := ret[m.MissionType]
		// なければMissionTypeの箱を新規作成
		if !ok {
			missions = make(Missions, 0)
		}
		ret[m.MissionType] = append(missions, m)
	}
	return ret
}
