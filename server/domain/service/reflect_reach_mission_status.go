package service

import (
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/value"
	"context"
)

type ReflectReachMissionStatus struct {
	um                 *entity.UserMission
	mps                entity.MissionProgresses
	replaceProgressNum int64
}

func NewReflectReachMissionStatus(um *entity.UserMission, mps entity.MissionProgresses, progress int64) *ReflectReachMissionStatus {
	return &ReflectReachMissionStatus{
		um:                 um,
		mps:                mps,
		replaceProgressNum: progress,
	}
}

func (r *ReflectReachMissionStatus) UpdateUserMission(ctx context.Context) bool {
	var updated bool
	switch r.um.StatusType {
	case value.StatusTypeProgress:
		updated = r.updateReachInProgress(ctx)
	case value.StatusTypeClear:
		updated = r.updateReachInClear(ctx)
	case value.StatusTypeReceived:
		updated = r.updateReachInReceived(ctx)
	}
	return updated
}

// 進捗中の到達系ミッションを更新するメソッド
// IsClearでは現在のMissionProgress.thresholdとの判定がなされる
func (r *ReflectReachMissionStatus) updateReachInProgress(ctx context.Context) bool {
	if r.um.CurrentProgress >= r.replaceProgressNum {
		return false
	}
	lastThreshold := r.mps.GetLastThreshold()
	r.um.ReplaceProgress(r.replaceProgressNum, lastThreshold)
	if r.um.IsClear() {
		r.um.StatusType = value.StatusTypeClear
	}
	return true
}

// 「現在のMissionProgressをクリア済みの到達系ミッション」を更新するメソッド。
// すでに全てのMissionProgressをクリア済み(最終MissionProgressのthresholdに到達)なら更新しない場合があるのが違い。
func (r *ReflectReachMissionStatus) updateReachInClear(ctx context.Context) bool {
	// 最大値が更新されない場合はそのままreturn
	// inClearであるので、まずcurrentProgress >= CurrentThresholdである。
	// つまり、r.um.CurrentProgress >= r.replaceProgressNum > r.um.CurrentThresholdの時も更新したい(Thresholdより上ならCurrentProgressが下がっても更新したい)かどうか
	// TODO r.um.CurrentProgress >= r.replaceProgressNumでもいい気はする
	if r.um.CurrentThreshold >= r.replaceProgressNum {
		return false
	}
	lastThreshold := r.mps.GetLastThreshold()
	// 最終MissionProgressじゃなければlastは超えない
	if r.um.CurrentProgress >= lastThreshold {
		return false
	}
	r.um.ReplaceProgress(r.replaceProgressNum, lastThreshold)
	return true
}

// 報酬受け取り済みの到達系ミッションを更新するメソッド
// ここでupdateされたら(つまり報酬を受け取ったら)UserMissionのcurrentThresholdが更新(次のMissionProgress.thresholdになる)される仕組み。
// つまり報酬を受け取らないとCurrentProgressはたされてくけど、Mission自体は進捗しないよってこと
// また、一気に加算することで新しい閾値も達成した場合は、Progressの状態にはならず、Clearの状態になる(連続クリアを担保)
func (r *ReflectReachMissionStatus) updateReachInReceived(ctx context.Context) bool {
	if r.um.CurrentThreshold >= r.replaceProgressNum {
		return false
	}
	lastThreshold := r.mps.GetLastThreshold()
	if r.um.CurrentProgress >= lastThreshold {
		return false
	}
	r.um.CurrentThreshold = r.mps.GetNextProgress(r.um.CurrentProgress).Threshold
	r.um.ReplaceProgress(r.replaceProgressNum, lastThreshold)

	// 進捗度加算後、次の閾値もクリアしてた場合はClearにする
	if r.um.IsClear() {
		r.um.StatusType = value.StatusTypeClear
	} else {
		r.um.StatusType = value.StatusTypeProgress
	}
	return true
}
