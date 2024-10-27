package registry

// アプリケーションの依存性管理の効率化、依存性注入を容易にするためRegistryパターンを導入
// UI, Infrastructureのさらに外側にいるイメージ
import (
	"com.graffity/mission-sample/server/applicationservice/component"
	"com.graffity/mission-sample/server/applicationservice/component/mission"
	"com.graffity/mission-sample/server/applicationservice/usecase"
	"com.graffity/mission-sample/server/domain/repository"
	"com.graffity/mission-sample/server/domain/value"
	"com.graffity/mission-sample/server/user_interface/handler"
)

type UserRegistry interface {
	UserUsecase() *usecase.UserUsecase
	UserHandler() *handler.UserHandler
	UserProcessor() *handler.UserHandler
	UserRepository() *handler.UserHandler
}

type UserRegistryImpl struct {
	mRepo  repository.MissionRepository
	mpRepo repository.MissionProgressRepository
	umRepo repository.UserMissionRepository
}

func NewUserRegistryImpl(
	missionRepository repository.MissionRepository,
	missionProgressRepository repository.MissionProgressRepository,
	userMissionRepository repository.UserMissionRepository,
) *UserRegistryImpl {
	return &UserRegistryImpl{
		mRepo:  missionRepository,
		mpRepo: missionProgressRepository,
		umRepo: userMissionRepository,
	}
}

func (r *UserRegistryImpl) UserHandler() *handler.UserHandler {
	return handler.NewUserHandler(r.UserUsecase())
}

func (r *UserRegistryImpl) UserUsecase() *usecase.UserUsecase {
	return usecase.NewUserUsecase(r.UserRepository(), r.missionProcessor())
}

// TODO AddReporterの呼び出し場所がここで適切か
func (r *UserRegistryImpl) missionProcessor() *component.MissionProcessor {
	processor := component.NewMissionProcessor(r.mRepo, r.mpRepo, r.umRepo)
	countReporter := mission.NewCountReporter(r.mRepo, r.mpRepo, r.umRepo)
	reachReporter := mission.NewReachReporter(r.mRepo, r.mpRepo, r.umRepo)
	processor.AddReporter(
		mission.Info{
			MissionType: value.MissionTypeLoginCount,
			Reporter:    countReporter,
		},
		mission.Info{
			MissionType: value.MissionTypeUserCreateReach,
			Reporter:    reachReporter,
		},
	)
	return processor
}

func (r *UserRegistryImpl) UserRepository() repository.UserRepository {
	return nil
}
