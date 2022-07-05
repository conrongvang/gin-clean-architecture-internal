package usecase

type AuthUseCase interface {
	Login()
}

type authUsecase struct {
}

func NewAuthUseCase() AuthUseCase {
	return &authUsecase{}
}

func (l *authUsecase) Login() {

}
