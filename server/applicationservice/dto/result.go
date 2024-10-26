package dto

// Result クリアしたミッション情報の結果を返すための構造体。
// ミッションマスターやユーザー情報、クリアした閾値の情報が含まれる
type Result struct {
	MissionData *MissionData
}

type Results []*Result
