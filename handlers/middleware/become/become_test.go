package become

import (
	"context"
	"net/http"
	"oj/db"
	"oj/models/users"
	"testing"
)

type testResponseWriter struct {
	t              *testing.T
	wantStatusCode int
}

func (w testResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w testResponseWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func (w testResponseWriter) WriteHeader(code int) {
	if code != http.StatusOK {
		if w.wantStatusCode != code {
			w.t.Errorf("got status code %d want %d", code, w.wantStatusCode)
		}
	}
	return
}

type testHandler struct {
	want string
	t    *testing.T
}

func (h testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	got := users.FromContext(r.Context())
	if got.Username != h.want {
		h.t.Errorf("want %v got %v", h.want, got.Username)
	}
}

func TestProvider(t *testing.T) {
	r := http.Request{}

	var bob users.User
	db.DB.MustExec("insert into users(username) values(?)", "bob")
	err := db.DB.Get(&bob, "select * from users where username = ?", "bob")
	if err != nil {
		panic(err)
	}

	for _, tc := range []struct {
		name           string
		requestUser    users.User
		wantUsername   string
		wantStatusCode int
	}{
		{
			name:         "alice as self",
			requestUser:  users.User{Username: "alice"},
			wantUsername: "alice",
		},
		{
			name:           "not admin alice as bob",
			requestUser:    users.User{Username: "alice", BecomeUserID: &bob.ID},
			wantStatusCode: 401,
		},
		{
			name:         "admin alice as bob",
			requestUser:  users.User{Username: "alice", BecomeUserID: &bob.ID, Admin: true},
			wantUsername: "bob",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := users.NewContext(context.Background(), tc.requestUser)
			res := Provider(testHandler{t: t, want: tc.wantUsername})
			res.ServeHTTP(testResponseWriter{t: t, wantStatusCode: tc.wantStatusCode}, r.WithContext(ctx))
		})
	}

}
