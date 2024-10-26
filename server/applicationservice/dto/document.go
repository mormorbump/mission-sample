package dto

import "com.graffity/mission-sample/server/domain/entity"

// Document 依頼書(Form)の情報を元に、必要なミッションの情報をまとめた資料にあたる構造体
// ミッション更新に必要なミッションマスターやユーザー情報、送られてき たフォーム情報が含まれる
type Document struct {
	Form            *Form
	MissionDataList MissionDataList
}

type Documents []*Document

type MissionData struct {
	Mission     *entity.Mission
	UserMission *entity.UserMission
}

type MissionDataList []*MissionData
