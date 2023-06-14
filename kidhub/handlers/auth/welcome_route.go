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
	"oj/handlers/render"
	"oj/models/users"
	"oj/services/email"
	"time"

	"github.com/go-chi/chi/v5"
)

func Route(r chi.Router) {
	r.Get("/", welcome)

	r.Get("/parents", welcomeParents)
	r.Post("/parents/email", emailRegisterAction)
	r.Get("/parents/code", parentsCode)
	r.Post("/parents/code", parentsCodeAction)

	r.Get("/kids", welcomeKids)
	r.Post("/kids/username", kidsUsernameAction)
	r.Get("/kids/code", kidsCode)
	r.Post("/kids/code", kidsCodeAction)

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
		render.Error(w, err.Error(), 500)
	}
}

var welcomeKidsTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome_kids.html",
))

func welcomeKids(w http.ResponseWriter, r *http.Request) {
	err := welcomeKidsTemplate.Execute(w, nil)
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

var welcomeParentsTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome_parents.html",
))

func welcomeParents(w http.ResponseWriter, r *http.Request) {
	err := welcomeParentsTemplate.Execute(w, nil)
	if err != nil {
		render.Error(w, err.Error(), 500)
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
	address := r.FormValue("email")
	if address == "" {
		http.Redirect(w, r, "/welcome/parents", http.StatusSeeOther)
		return
	}

	// store generated code in pending registrations table along with email
	nonce, err := generateSecureString(32)
	if err != nil {
		log.Printf("Error generating secure string enMJFDN8M4Z6y5p6n: %s", err)
		http.Error(w, "Error generating code enMJFDN8M4Z6y5p6n", 500)
		return
	}
	code := generateDigitCode()
	_, err = db.DB.Exec("insert into codes(nonce, email, code) values(?, ?, ?)", nonce, address, code)
	if err != nil {
		log.Printf("Error generating code YQChKPeCivnvM9P82: %s", err)
		http.Error(w, "Error generating code YQChKPeCivnvM9P82", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Value: nonce, Path: "/", Expires: time.Now().Add(time.Hour)})

	// email code to user
	_, _, err = email.Send("hello", code, address)
	if err != nil {
		log.Printf("Error emailing code gYqGXoK6XfC2va3Rp: %s", err)
		http.Error(w, "Error emailing code gYqGXoK6XfC2va3Rp", 500)
		return
	}

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
		render.Error(w, err.Error(), 500)
	}
}

func kidsUsernameAction(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		http.Redirect(w, r, "/welcome/kids", http.StatusSeeOther)
		return
	}

	user, err := users.FindByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Redirect(w, r, "/welcome/kids?badusername", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/welcome/kids?weirderror", http.StatusSeeOther)
		return
	}

	// store generated code in pending registrations table along with email
	nonce, err := generateSecureString(32)
	if err != nil {
		log.Printf("Error generating secure string wN6Cd9vQLHYQ2euxb: %s", err)
		http.Error(w, "Error generating code wN6Cd9vQLHYQ2euxb", 500)
		return
	}
	code := generateDigitCode()
	_, err = db.DB.Exec("insert into kids_codes(nonce, user_id, code) values(?, ?, ?)", nonce, user.ID, code)
	if err != nil {
		log.Printf("Error generating code qYBJ24gqRrmFEJWAs: %s", err)
		http.Error(w, "Error generating code qYBJ24gqRrmFEJWAs", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Value: nonce, Path: "/", Expires: time.Now().Add(time.Hour)})

	// email code to kids parent(s)
	parents, err := users.GetParents(user.ID)
	if err != nil {
		log.Printf("Error getting parents wdEXqpGbDeTc69Ju3: %s", err)
		http.Error(w, "Error getting parents wdEXqpGbDeTc69Ju3", 500)
		return
	}

	if len(parents) == 0 {
		log.Printf("No parents QNw5BhAWCEQxwQ4LE: %s", err)
		http.Error(w, "No parents QNw5BhAWCEQxwQ4LE", 500)
		return
	}

	for _, parent := range parents {
		_, _, err = email.Send("code for your kid", code, *parent.Email)
		if err != nil {
			log.Printf("Error sending email BohZie4YoPfrTHwj4: %s", err)
		}
	}

	// redirect to page to input code
	http.Redirect(w, r, "/welcome/kids/code", http.StatusSeeOther)
}

var welcomeKidsCodeTemplate = template.Must(template.ParseFiles(
	"handlers/auth/layout.html",
	"handlers/auth/welcome_kids_code.html",
))

func kidsCode(w http.ResponseWriter, r *http.Request) {
	retry := r.URL.Query().Has("retry")
	log.Printf("retry(%v)", retry)

	err := welcomeKidsCodeTemplate.Execute(w, struct{ Retry bool }{Retry: retry})
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

func kidsCodeAction(w http.ResponseWriter, r *http.Request) {
	var userID int64

	cookie, err := r.Cookie("kh_nonce")
	if err != nil {
		if err != http.ErrNoCookie {
			log.Printf("weird error 792pR3LQagv5ej3Xi %s", err)
		}
		http.Redirect(w, r, "/welcome/parents", 303)
		return
	}

	nonce := cookie.Value
	code := r.FormValue("code")

	// look up code
	// XXX fetch by id alone, compare code, and add retry count
	err = db.DB.Get(&userID, "select user_id from kids_codes where nonce = ? and code = ?", nonce, code)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Printf("Error retrieving code RkfeaQB4rAX7uxdY3: %s", err)
			http.Error(w, "Error retrieving code RkfeaQB4rAX7uxdY3", 500)
			return
		}
	}

	if userID != 0 {
		log.Println("code is good")
		// found email, code is good
		// create user if not exists
		user, err := users.FindById(userID)
		if err != nil {
			render.Error(w, "error getting user", 500)
			return
		}
		log.Printf("user %v", user)
		// create a new session
		key, err := generateSecureString(32)
		if err != nil {
			render.Error(w, "error creating session", 500)
			return
		}
		_, err = db.DB.Exec("insert into sessions(key, user_id) values(?, ?)", key, user.ID)
		if err != nil {
			render.Error(w, "error creating session", 500)
			return
		}
		// set session cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_session", Value: key, Path: "/", Expires: time.Now().Add(30 * 24 * time.Hour)})
		// clear nonce cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Path: "/", Expires: time.Now().Add(-time.Hour)})

		// redirect
		http.Redirect(w, r, "/", 303)
	} else {
		log.Println("code is bad")
		// code is bad
		http.Redirect(w, r, "/welcome/kids/code?retry", 303)
	}
}

func parentsCodeAction(w http.ResponseWriter, r *http.Request) {
	var email string

	cookie, err := r.Cookie("kh_nonce")
	if err != nil {
		if err != http.ErrNoCookie {
			log.Printf("weird error 792pR3LQagv5ej3Xi %s", err)
		}
		http.Redirect(w, r, "/welcome/parents", 303)
		return
	}

	nonce := cookie.Value
	code := r.FormValue("code")

	// look up code
	// XXX fetch by id alone, compare code, and add retry count
	err = db.DB.Get(&email, "select email from codes where nonce = ? and code = ?", nonce, code)
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
			render.Error(w, "error getting user", 500)
			return
		}
		log.Printf("user %v", user)
		// create a new session
		key, err := generateSecureString(32)
		if err != nil {
			render.Error(w, "error creating session", 500)
			return
		}
		_, err = db.DB.Exec("insert into sessions(key, user_id) values(?, ?)", key, user.ID)
		if err != nil {
			render.Error(w, "error creating session", 500)
			return
		}
		// set session cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_session", Value: key, Path: "/", Expires: time.Now().Add(30 * 24 * time.Hour)})
		// clear nonce cookie
		http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Path: "/", Expires: time.Now().Add(-time.Hour)})

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
