package entity

import (
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

type MissionID string
type Missions []*Mission

type MissionPK struct {
	MissionID string
}

type MissionPKs []*MissionPK

func (m *Mission) IsTarget(targetID string) bool {
	return m.TargetID == "" || m.TargetID == targetID
}
