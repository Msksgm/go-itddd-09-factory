package user

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type UserFactorierStub struct {
	create func(userName UserName) (*User, error)
}

func (ufs UserFactorierStub) Create(userName UserName) (*User, error) {
	return ufs.create(userName)
}

func Test_Create(t *testing.T) {
	name, err := NewUserName("userName")
	if err != nil {
		t.Error(err)
	}

	userFactory, err := NewUserFactory()
	if err != nil {
		t.Error(err)
	}

	got, err := userFactory.Create(*name)
	if err != nil {
		t.Error(err)
	}

	want := &User{id: UserId{value: "id"}, name: *name}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(User{}, UserId{}, UserName{}), cmpopts.IgnoreFields(*got, "id")); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
