package user

type UserApplicationService struct {
	userFactory    UserFactorier
	userRepository UserRepositorier
	userService    UserService
}

func NewUserApplicationService(userFactory UserFactorier, userRepository UserRepositorier, userService UserService) (*UserApplicationService, error) {
	return &UserApplicationService{userFactory: userFactory, userRepository: userRepository, userService: userService}, nil
}

func (uas *UserApplicationService) Register(name string) error {
	userName, err := NewUserName(name)
	if err != nil {
		return err
	}

	user, err := uas.userFactory.Create(*userName)
	if err != nil {
		return err
	}

	isUserExists, err := uas.userService.Exists(user)
	if isUserExists {
		return err
	}

	if err := uas.userRepository.Save(user); err != nil {
		return err
	}
	return nil
}
