package user

import "testing"

func Test_Register(t *testing.T) {
	data := []struct {
		testname       string
		userName       string
		create         func(userName UserName) (*User, error)
		findByUserName func(name UserName) (*User, error)
		exists         func(user User) (bool, error)
		save           func(user User) error
		errMsg         string
	}{
		{
			"success",
			"userName",
			func(userName UserName) (*User, error) {
				return &User{name: UserName{value: "userName"}, id: UserId{value: "userId"}}, nil
			},
			func(name UserName) (*User, error) { return nil, nil },
			func(user User) (bool, error) { return false, nil },
			func(user User) error { return nil },
			"",
		},
	}
	userApplicationService := UserApplicationService{}

	for _, d := range data {
		t.Run(d.testname, func(t *testing.T) {
			userApplicationService.userFactory = &UserFactorierStub{create: d.create}
			userApplicationService.userRepository = &UserRepositorierStub{findByUserName: d.findByUserName, save: d.save}
			userApplicationService.userService = &UserServicerStub{exists: d.exists}

			err := userApplicationService.Register("userName")
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
