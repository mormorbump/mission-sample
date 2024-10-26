package service

import (
	"com.graffity/mission-sample/server/domain/entity"
	"com.graffity/mission-sample/server/domain/value"
	"context"
)

func UpdateUserMission(ctx context.Context, um *entity.UserMission, mps entity.MissionProgresses, progress int64) bool {
	var updated bool
	switch um.StatusType {
	case value.StatusTypeProgress:
		updated = updateCountInProgress(ctx, um, mps, progress)
	case value.StatusTypeClear:
		updated = updateCountInClear(ctx, um, mps, progress)
	case value.StatusTypeReceived:
		updated = updateCountInReceived(ctx, um, mps, progress)
	}
	return updated
}

// 進捗中の回数系ミッションを更新するメソッド
// IsClearでは現在のMissionProgress.thresholdとの判定がなされる
func updateCountInProgress(ctx context.Context, um *entity.UserMission, mps entity.MissionProgresses, progress int64) bool {
	lastThreshold := mps.GetLastThreshold()
	um.AddProgress(progress, lastThreshold)
	if um.IsClear() {
		um.StatusType = value.StatusTypeClear
	}
	return true
}

// 現在のMissionProgressをクリア済みの回数系ミッションを更新するメソッド。
// すでに全てのMissionProgressをクリア済み(最終MissionProgressのthresholdに到達)なら更新しない場合があるのが違い。
func updateCountInClear(ctx context.Context, um *entity.UserMission, mps entity.MissionProgresses, progress int64) bool {
	lastThreshold := mps.GetLastThreshold()
	// 基本inClearならこの条件に入るはず
	if um.CurrentProgress >= lastThreshold {
		return false
	}
	// マスタ更新とかで整合性崩れた時用？
	um.AddProgress(progress, lastThreshold)
	return true
}

// 報酬受け取り済みの回数系ミッションを更新するメソッド
// ここでupdateされないとUserMissionのcurrentThresholdが更新(次のMissionProgress.thresholdになる)される仕組み。
// つまり報酬を受け取らないとCurrentProgressはたされてくけど、Mission自体は進捗しないよってこと
// また、一気に加算することで新しい閾値も達成した場合は、Progressの状態にはならず、Clearの状態になる(連続クリアを担保)
func updateCountInReceived(ctx context.Context, um *entity.UserMission, mps entity.MissionProgresses, progress int64) bool {
	lastThreshold := mps.GetLastThreshold()
	if um.CurrentProgress >= lastThreshold {
		return false
	}
	um.CurrentThreshold = mps.GetNextProgress(um.CurrentProgress).Threshold
	// 最大閾値じゃないなら次のthresholdをUserMissionに適用する
	um.AddProgress(progress, lastThreshold)

	// 進捗度加算後、次の閾値もクリアしてた場合はClearにする
	if um.IsClear() {
		um.StatusType = value.StatusTypeClear
	} else {
		um.StatusType = value.StatusTypeProgress
	}
	return true
}
