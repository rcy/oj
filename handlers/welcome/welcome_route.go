package welcome

import (
	cryptorand "crypto/rand"
	"database/sql"
	_ "embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	mathrand "math/rand"
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

	r.Get("/signout", signout)
}

//go:embed layout.gohtml
var layoutContent string

func mustLayout(content string) *template.Template {
	return template.Must(template.New("").Parse(layoutContent + content))
}

//go:embed welcome.gohtml
var welcomeContent string
var welcomeTemplate = mustLayout(welcomeContent)

func welcome(w http.ResponseWriter, r *http.Request) {
	err := welcomeTemplate.Execute(w, nil)
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

//go:embed welcome_kids.gohtml
var welcomeKidsContent string
var welcomeKidsTemplate = mustLayout(welcomeKidsContent)

func welcomeKids(w http.ResponseWriter, r *http.Request) {
	err := welcomeKidsTemplate.Execute(w, struct{ Error string }{""})
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

//go:embed welcome_parents.gohtml
var welcomeParentsContent string
var welcomeParentsTemplate = mustLayout(welcomeParentsContent)

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
		digit := mathrand.Intn(10)
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
		render.Error(w, "Error generating code enMJFDN8M4Z6y5p6n", 500)
		return
	}
	code := generateDigitCode()
	_, err = db.DB.Exec("insert into codes(nonce, email, code) values(?, ?, ?)", nonce, address, code)
	if err != nil {
		render.Error(w, "Error generating code YQChKPeCivnvM9P82", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Value: nonce, Path: "/", Expires: time.Now().Add(time.Hour)})

	// email code to user
	_, _, err = email.Send(
		fmt.Sprintf("Parent sign in code: %s", code),
		fmt.Sprintf("Your Octopus Jr verification code is %s", code),
		address)
	if err != nil {
		render.Error(w, "Error emailing code gYqGXoK6XfC2va3Rp", 500)
		return
	}

	// redirect to page to input code
	http.Redirect(w, r, "/welcome/parents/code", http.StatusSeeOther)
}

//go:embed welcome_parents_code.gohtml
var welcomeParentsCodeContent string
var welcomeParentsCodeTemplate = mustLayout(welcomeParentsCodeContent)

func parentsCode(w http.ResponseWriter, r *http.Request) {
	err := welcomeParentsCodeTemplate.Execute(w, struct{ Error string }{""})
	if err != nil {
		render.Error(w, err.Error(), 500)
	}
}

func kidsUsernameAction(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	user, err := users.FindByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = welcomeKidsTemplate.Execute(w, struct{ Error string }{"User not found"})
			if err != nil {
				render.Error(w, err.Error(), 500)
			}
			return
		}
		render.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store generated code in pending registrations table along with email
	nonce, err := generateSecureString(32)
	if err != nil {
		render.Error(w, "Error generating code wN6Cd9vQLHYQ2euxb", 500)
		return
	}
	code := generateDigitCode()
	_, err = db.DB.Exec("insert into kids_codes(nonce, user_id, code) values(?, ?, ?)", nonce, user.ID, code)
	if err != nil {
		render.Error(w, "Error generating code qYBJ24gqRrmFEJWAs", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "kh_nonce", Value: nonce, Path: "/", Expires: time.Now().Add(time.Hour)})

	// email code to kids parent(s)
	parents, err := users.GetParents(user.ID)
	if err != nil {
		render.Error(w, "Error getting parents wdEXqpGbDeTc69Ju3", 500)
		return
	}

	if len(parents) == 0 {
		render.Error(w, "No parents QNw5BhAWCEQxwQ4LE", 500)
		return
	}

	for _, parent := range parents {
		_, _, err = email.Send(
			fmt.Sprintf("Code for %s is %s", username, code),
			fmt.Sprintf(fmt.Sprintf("Your child, %s, is trying to login to Octopus Jr.  The verification code is %s.",
				username, code)),
			*parent.Email)
		if err != nil {
			render.Error(w, fmt.Sprintf("Error emailing code: %s", err), http.StatusInternalServerError)
		}
	}

	// redirect to page to input code
	http.Redirect(w, r, "/welcome/kids/code", http.StatusSeeOther)
}

//go:embed welcome_kids_code.gohtml
var welcomeKidsCodeContent string
var welcomeKidsCodeTemplate = mustLayout(welcomeKidsCodeContent)

func kidsCode(w http.ResponseWriter, r *http.Request) {
	err := welcomeKidsCodeTemplate.Execute(w, struct{ Error string }{""})
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
			render.Error(w, "Error retrieving code RkfeaQB4rAX7uxdY3", 500)
			return
		}
	}

	if userID != 0 {
		log.Println("code is good")
		// found email, code is good
		// create user if not exists
		user, err := users.FindById(userID)
		if err != nil {
			render.Error(w, "error getting user:"+err.Error(), 500)
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
		err = welcomeKidsCodeTemplate.Execute(w, struct{ Error string }{"bad code, try again"})
		if err != nil {
			render.Error(w, err.Error(), 500)
			return
		}
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
			render.Error(w, "Error retrieving code qmNpb3qvPM8oGwmLn", 500)
			return
		}
	}

	if email != "" {
		log.Println("code is good")
		// found email, code is good
		// create user if not exists
		user, err := users.FindOrCreateParentByEmail(email)
		if err != nil {
			render.Error(w, "error getting user: "+err.Error(), 500)
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
		// code is bad
		log.Println("code is bad")
		err := welcomeParentsCodeTemplate.Execute(w, struct{ Error string }{"bad code, try again"})
		if err != nil {
			render.Error(w, err.Error(), 500)
		}
		//http.Redirect(w, r, "/welcome/parents/code?retry", 303)

	}
}

func generateSecureString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := cryptorand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)
	return randomString, nil
}
