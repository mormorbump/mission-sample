package entity

// MissionProgress ミッションの進捗管理用
// MissionIDによってThresholdの意味が変わる
type MissionProgress struct {
	MissionID string
	Threshold int64
	ItemID    string
}

type MissionProgresses []*MissionProgress

type MissionProgressPK struct {
	MissionID string
	Threshold string
}

type MissionProgressPKs []*MissionProgressPK

// GetFirstThreshold　MissionProgressesのうち、最初のthresholdを取得
// 初めてやるミッションの時に使用
// リストが0だったら0
func (mps MissionProgresses) GetFirstThreshold() int64 {
	if len(mps) == 0 {
		return 0
	}
	return mps[0].Threshold
}

// GetLastThreshold MissionProgressesのうち、最後のthresholdを取得
// ミッションが完全に終了するかどうかの判定に利用
func (mps MissionProgresses) GetLastThreshold() int64 {
	if len(mps) == 0 {
		return 0
	}
	return mps[len(mps)-1].Threshold
}

func (mps MissionProgresses) GetNextProgress(threshold int64) *MissionProgress {
	for _, mp := range mps {
		if mp.Threshold > threshold {
			return mp
		}
	}
	return nil
}