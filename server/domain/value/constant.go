package value

type MissionType = int32

const (
	MissionTypeLoginCount            MissionType = 1 // 通算ログイン回数
	MissionTypeQuestClearCount       MissionType = 2 // クエストクリア回数
	MissionTypeTargetQuestClearCount MissionType = 3 // 指定クエストクリア回数
	MissionTypeMiniGameClearReach    MissionType = 4 // ミニゲーム最高得点到達
	MissionTypeCharacterLevelReach   MissionType = 5 // キャラクター最大レベル到達
)

type StatusType = int32 // ユーザのミッション状況

const (
	StatusTypeProgress StatusType = 1 // 進行中
	StatusTypeClear    StatusType = 2 // 達成
	StatusTypeReceived StatusType = 3 // 報酬受取済み
)
