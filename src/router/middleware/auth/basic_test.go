package auth

import (
	"context"
	"net/http"
	"testing"

	"github.com/Hack-Portal/backend/src/adapters/gateways"
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/utils/password"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var testUser = &models.User{}

func AddTestData(t *testing.T) {
	t.Helper()
	ug := gateways.NewUserGateway(dbconn)

	hashedPass, err := password.HashPassword("test_password")
	if err != nil {
		t.Fatal(err)
	}

	testUser = &models.User{
		UserID:   uuid.New().String(),
		Name:     "test_user",
		Password: hashedPass,
		Role:     1,
	}
	_, err = ug.Create(context.TODO(), testUser)
	if err != nil {
		t.Fatal(err)
	}
}

func RunTestServe(t *testing.T) http.Handler {
	t.Helper()
	e := echo.New()
	e.Use(NewBasicAuth(
		gateways.NewUserGateway(dbconn),
	).AuthN())
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	e.GET("/secret", func(c echo.Context) error {
		if c.Get(IsGuest).(bool) {
			return echo.ErrUnauthorized
		}
		return c.String(200, "secret")
	})
	return e
}

type basicAuthNTestCase struct {
	name           string
	uid            string
	pass           string
	wantRootStatus int
	wantSecStatus  int
}

func TestBasicAuthN(t *testing.T) {
	AddTestData(t)

	testCases := []basicAuthNTestCase{
		{
			name:           "Success",
			uid:            testUser.UserID,
			pass:           "test_password",
			wantRootStatus: 200,
			wantSecStatus:  200,
		},
		{
			name:           "UserID not found",
			uid:            testUser.UserID,
			pass:           "test_password",
			wantRootStatus: 200,
			wantSecStatus:  401,
		},
		{
			name:           "Psasword not found",
			uid:            testUser.UserID,
			pass:           "test_password",
			wantRootStatus: 200,
			wantSecStatus:  401,
		},
		{
			name:           "UserID and Password not found",
			uid:            testUser.UserID,
			pass:           "test_password",
			wantRootStatus: 200,
			wantSecStatus:  401,
		},
		{
			name:           "user not found",
			uid:            "not_found_user",
			pass:           "test_password",
			wantRootStatus: 401,
			wantSecStatus:  401,
		},
		{
			name:           "password is wrong",
			uid:            testUser.UserID,
			pass:           "wrong_password",
			wantRootStatus: 401,
			wantSecStatus:  401,
		},
	}
	go http.ListenAndServe(":8080", RunTestServe(t))

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			CheckRoot(t, tc)
			CheckSecret(t, tc)
		})
	}
}

func CheckRoot(t *testing.T, tc basicAuthNTestCase) {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth(tc.uid, tc.pass)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != tc.wantRootStatus {
		b := resp.Body
		defer b.Close()
		buf := make([]byte, 1024)
		n, _ := b.Read(buf)
		t.Log(string(buf[:n]))

		t.Errorf("want %d, but got %d", tc.wantRootStatus, resp.StatusCode)
	}
}

func CheckSecret(t *testing.T, tc basicAuthNTestCase) {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/secret", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.SetBasicAuth(tc.uid, tc.pass)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != tc.wantRootStatus {
		t.Errorf("want %d, but got %d", tc.wantRootStatus, resp.StatusCode)
	}
}
