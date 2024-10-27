package service

import (
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/value"
	"context"
)

type ReflectCounterMissionStatus struct {
	um             *entity.UserMission
	mps            entity.MissionProgresses
	addProgressNum int64
}

func NewReflectCounterMissionStatus(um *entity.UserMission, mps entity.MissionProgresses, progress int64) *ReflectCounterMissionStatus {
	return &ReflectCounterMissionStatus{
		um:             um,
		mps:            mps,
		addProgressNum: progress,
	}
}

func (r *ReflectCounterMissionStatus) UpdateUserMission(ctx context.Context) bool {
	var updated bool
	switch r.um.StatusType {
	case value.StatusTypeProgress:
		updated = r.updateCountInProgress(ctx)
	case value.StatusTypeClear:
		updated = r.updateCountInClear(ctx)
	case value.StatusTypeReceived:
		updated = r.updateCountInReceived(ctx)
	}
	return updated
}

// 進捗中の回数系ミッションを更新するメソッド
// IsClearでは現在のMissionProgress.thresholdとの判定がなされる
func (r *ReflectCounterMissionStatus) updateCountInProgress(ctx context.Context) bool {
	lastThreshold := r.mps.GetLastThreshold()
	r.um.AddProgress(r.addProgressNum, lastThreshold)
	if r.um.IsClear() {
		r.um.StatusType = value.StatusTypeClear
	}
	return true
}

// 「現在のMissionProgressをクリア済みの回数系ミッション」を更新するメソッド。
// すでに全てのMissionProgressをクリア済み(最終MissionProgressのthresholdに到達)なら更新しない場合があるのが違い。
func (r *ReflectCounterMissionStatus) updateCountInClear(ctx context.Context) bool {
	lastThreshold := r.mps.GetLastThreshold()
	// 最終MissionProgressじゃなければlastは超えない
	if r.um.CurrentProgress >= lastThreshold {
		return false
	}
	r.um.AddProgress(r.addProgressNum, lastThreshold)
	return true
}

// 報酬受け取り済みの回数系ミッションを更新するメソッド
// ここでupdateされたら(つまり報酬を受け取ったら)UserMissionのcurrentThresholdが更新(次のMissionProgress.thresholdになる)される仕組み。
// つまり報酬を受け取らないとCurrentProgressはたされてくけど、Mission自体は進捗しないよってこと
// また、一気に加算することで新しい閾値も達成した場合は、Progressの状態にはならず、Clearの状態になる(連続クリアを担保)
func (r *ReflectCounterMissionStatus) updateCountInReceived(ctx context.Context) bool {
	lastThreshold := r.mps.GetLastThreshold()
	if r.um.CurrentProgress >= lastThreshold {
		return false
	}
	r.um.CurrentThreshold = r.mps.GetNextProgress(r.um.CurrentProgress).Threshold
	r.um.AddProgress(r.addProgressNum, lastThreshold)

	// 進捗度加算後、次の閾値もクリアしてた場合はClearにする
	if r.um.IsClear() {
		r.um.StatusType = value.StatusTypeClear
	} else {
		r.um.StatusType = value.StatusTypeProgress
	}
	return true
}
