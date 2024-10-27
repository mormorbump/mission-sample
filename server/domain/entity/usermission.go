package entity

import "com.graffity/mission-sample/server/domain/value"

type UserMission struct {
	UserID           UserID
	MissionID        MissionID
	CurrentThreshold int64 // MissionProgressのThresholdが設定される
	CurrentProgress  int64
	StatusType       value.StatusType
}

type UserMissions []*UserMission

type UserMissionPK struct {
	UserID    UserID
	MissionID MissionID
}

type UserMissionPKs []*UserMissionPK

func NewUserMission(userID UserID, missionID MissionID, firstThreshold int64) *UserMission {
	return &UserMission{
		UserID:           userID,
		MissionID:        missionID,
		CurrentThreshold: firstThreshold,
		CurrentProgress:  0,
		StatusType:       value.StatusTypeProgress,
	}
}

// AddProgress
// 指定の閾値を上限値として、上限を超えないように加算
func (um *UserMission) AddProgress(progress, threshold int64) {
	um.CurrentProgress += progress
	if um.CurrentProgress > threshold {
		um.CurrentProgress = threshold
	}
}

func (um *UserMission) ReplaceProgress(progress, threshold int64) {
	um.CurrentProgress = progress
	if um.CurrentProgress > threshold {
		um.CurrentProgress = threshold
	}
}

// IsClear
// CurrentThresholdにはその時のMissionProgressのthresholdが入っている。
// これは報酬を受け取ったタイミングで次のMissionProgress.thresholdに更新される
func (um *UserMission) IsClear() bool {
	return um.CurrentProgress >= um.CurrentThreshold
}

// ToMapByMissionID userごとのmissionとuserMissionは1:1
func (ums UserMissions) ToMapByMissionID() map[MissionID]*UserMission {
	ret := make(map[MissionID]*UserMission)
	for _, um := range ums {
		ret[um.MissionID] = um
	}
	return ret
}
