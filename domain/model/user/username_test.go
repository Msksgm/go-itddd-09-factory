package user

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewUserName(t *testing.T) {
	data := []struct {
		testName string
		userName string
		want     *UserName
		errMsg   string
	}{
		{"success", "username", &UserName{value: "username"}, ""},
		{"fail because value is less than 3 characters", "us", nil, "UserName is more than 3 characters."},
	}
	for _, d := range data {
		t.Run(d.testName, func(t *testing.T) {
			got, err := NewUserName(d.userName)
			if diff := cmp.Diff(d.want, got, cmp.AllowUnexported(UserName{})); diff != "" {
				t.Errorf("mismatch (-want, +got):\n%s", diff)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
	t.Run("fail because value is more than 20 characters", func(t *testing.T) {
		_, err := NewUserName("usernameusernameusername")
		want := "UserName is less than 20 characters."
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func Test_UserNameEquals(t *testing.T) {
	userName, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("equal", func(t *testing.T) {
		otherUserName, err := NewUserName("username")
		if err != nil {
			t.Fatal(err)
		}
		if !userName.Equals(*otherUserName) {
			t.Errorf("userName %v must be equal to otherUserName %v", userName, otherUserName)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		otherUserName, err := NewUserName("otherusername")
		if err != nil {
			t.Fatal(err)
		}
		if userName.Equals(*otherUserName) {
			t.Errorf("userName %v must not be equal to otherUserName %v", userName, otherUserName)
		}
	})
}

func Test_UserNameString(t *testing.T) {
	userName, err := NewUserName("username")
	if err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("UserName: [value: %s]", userName.value)
	got := fmt.Sprint(userName)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
