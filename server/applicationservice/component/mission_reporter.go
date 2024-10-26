package component

import (
	"com.graffity/mission-sample/server/applicationservice/dto"
	"com.graffity/mission-sample/server/domain/entity"
	"context"
)

// MissionReporter どのようなミッションでも、Documentを受け取って、Resultsを返すような実装を強制する
// componentの役割としては、具体的なビジネスロジックを持たないものの、ユースケースの流れを制御し、エンティティやリポジトリと協調して目的の処理を実現する層なのでserviceにはいる。
// 具体的には、ミッションの更新ロジックを提供するが、具体的なビジネスロジックそのもの（つまり、ドメイン層のビジネスロジック）ではなく、外部のデータやDTO（DocumentとResults）を使用して処理を行っている
type MissionReporter interface {
	Report(ctx context.Context, userID entity.UserID, document *dto.Document) (dto.Results, error)
}
