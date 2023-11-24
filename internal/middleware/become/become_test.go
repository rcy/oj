package become

import (
	"context"
	"database/sql"
	"net/http"
	"oj/api"
	"oj/db"
	"oj/internal/middleware/auth"
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
	got := auth.FromContext(r.Context())
	if got.Username != h.want {
		h.t.Errorf("want %v got %v", h.want, got.Username)
	}
}

func TestProvider(t *testing.T) {
	r := http.Request{}
	queries := api.New(db.DB)

	bob, err := queries.CreateUser(context.Background(), "bob")
	if err != nil {
		panic(err)
	}
	bobID := sql.NullInt64{Int64: bob.ID, Valid: true}

	for _, tc := range []struct {
		name           string
		requestUser    api.User
		wantUsername   string
		wantStatusCode int
	}{
		{
			name:         "alice as self",
			requestUser:  api.User{Username: "alice"},
			wantUsername: "alice",
		},
		{
			name:           "not admin alice as bob",
			requestUser:    api.User{Username: "alice", BecomeUserID: bobID},
			wantStatusCode: 401,
		},
		{
			name:         "admin alice as bob",
			requestUser:  api.User{Username: "alice", BecomeUserID: bobID, Admin: true},
			wantUsername: "bob",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := auth.NewContext(context.Background(), tc.requestUser)
			res := Provider(testHandler{t: t, want: tc.wantUsername})
			res.ServeHTTP(testResponseWriter{t: t, wantStatusCode: tc.wantStatusCode}, r.WithContext(ctx))
		})
	}

}
