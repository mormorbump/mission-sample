package mission

import "com.graffity/mission-sample/server/domain/entity"

// Document 依頼書(Form)の情報を元に、必要なミッションの情報をまとめた資料にあたる構造体
// ミッション更新に必要なミッションマスターやユーザー情報、送られてきたフォーム情報が含まれる
type Document struct {
	Form     *Form
	DataList DataList
}

type Documents []*Document

type Data struct {
	Mission     *entity.Mission
	UserMission *entity.UserMission
}

type DataList []*Data

func CreateMissionDataList(missions entity.Missions, userMissionMap map[entity.MissionID]*entity.UserMission) DataList {
	ret := make(DataList, 0, len(missions))
	for _, m := range missions {
		um := userMissionMap[m.ID]
		ret = append(ret, &Data{
			Mission:     m,
			UserMission: um,
		})
	}
	return ret
}
