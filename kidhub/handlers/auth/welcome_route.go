package auth

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"oj/db"
	"oj/handlers"
	"oj/models/users"
	"time"

	"github.com/go-chi/chi/v5"
)

func Route(r chi.Router) {
	r.Get("/", welcome)
	r.Get("/kids", welcomeKids)

	r.Get("/parents", welcomeParents)
	r.Post("/parents/email", emailRegisterAction)

	r.Get("/parents/code", parentsCode)
	r.Post("/parents/code", parentsCodeAction)

	// r.Get("/signup", getSignup)
	// r.Post("/signup", postSignup)
	r.Get("/signout", signout)
}

var welcomeTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome.html",
))

func welcome(w http.ResponseWriter, r *http.Request) {
	err := welcomeTemplate.Execute(w, nil)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

var welcomeKidsTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome_kids.html",
))

func welcomeKids(w http.ResponseWriter, r *http.Request) {
	err := welcomeKidsTemplate.Execute(w, nil)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

var welcomeParentsTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome_parents.html",
))

func welcomeParents(w http.ResponseWriter, r *http.Request) {
	err := welcomeParentsTemplate.Execute(w, nil)
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func signout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "kh_session", Path: "/", Expires: time.Now().Add(-time.Hour)})
	http.Redirect(w, r, "/", http.StatusFound)
}

func generateDigitCode() string {
	code := ""
	for i := 0; i < 4; i++ {
		digit := rand.Intn(10)
		code += fmt.Sprint(digit)
	}

	return code
}

func emailRegisterAction(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email == "" {
		http.Redirect(w, r, "/welcome/parents", http.StatusSeeOther)
		return
	}

	// store generated code in pending registrations table along with email
	result, err := db.DB.Exec("insert into codes(email, code) values(?, ?)", email, generateDigitCode())
	if err != nil {
		log.Printf("Error generating code: %s", err)
		http.Error(w, "Error generating code YQChKPeCivnvM9P82", 500)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert id", err)
		http.Error(w, "Error generating code L55gtEm4JuuozWJgZ", 500)
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "regcodeid", Value: fmt.Sprintf("%d", id), Path: "/", Expires: time.Now().Add(time.Hour)})

	// XXX: email code to user

	// redirect to page to input code
	http.Redirect(w, r, "/welcome/parents/code", http.StatusSeeOther)
}

var welcomeParentsCodeTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome_parents_code.html",
))

func parentsCode(w http.ResponseWriter, r *http.Request) {
	retry := r.URL.Query().Has("retry")
	log.Printf("retry(%v)", retry)

	err := welcomeParentsCodeTemplate.Execute(w, struct{ Retry bool }{Retry: retry})
	if err != nil {
		handlers.Error(w, err.Error(), 500)
	}
}

func parentsCodeAction(w http.ResponseWriter, r *http.Request) {
	var email string

	cookie, err := r.Cookie("regcodeid")
	if err != nil {
		if err != http.ErrNoCookie {
			log.Printf("weird error 792pR3LQagv5ej3Xi %s", err)
		}
		http.Redirect(w, r, "/welcome/parents", 303)
		return
	}

	id := cookie.Value
	code := r.FormValue("code")

	// look up code
	// XXX fetch by id alone, compare code, and add retry count
	err = db.DB.Get(&email, "select email from codes where id = ? and code = ?", id, code)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error retrieving code qmNpb3qvPM8oGwmLn: %s", err)
			http.Error(w, "Error retrieving code qmNpb3qvPM8oGwmLn", 500)
			return
		}
	}

	if email != "" {
		log.Println("code is good")
		// found email, code is good
		// create user if not exists
		user, err := users.FindOrCreateByEmail(email)
		if err != nil {
			handlers.Error(w, "error getting user", 500)
			return
		}
		log.Printf("user %v", user)
		// create a new session
		key, err := generateSecureString(32)
		_, err = db.DB.Exec("insert into sessions(key, user_id) values(?, ?)", key, user.ID)
		if err != nil {
			log.Print(err)
			handlers.Error(w, "error creating session", 500)
			return
		}
		// set cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_session", Value: key, Path: "/", Expires: time.Now().Add(30 * 24 * time.Hour)})

		// redirect
		http.Redirect(w, r, "/", 303)
	} else {
		log.Println("code is bad")
		// code is bad
		http.Redirect(w, r, "/welcome/parents/code?retry", 303)
	}
}

func generateSecureString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)
	return randomString, nil
}
