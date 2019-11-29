package controller

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/jonwzh/blog-sys/vm"
)

// template

// PopulateTemplates : Map template name to template.Template
func PopulateTemplates() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + fi.Name() + "'")
			}
			defer f.Close()
			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}
			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}
			result[fi.Name()] = tmpl
		}() // QUESTION: WHY DEFINE A SEPARATE FUNCTION
	}
	return result
}

// session

func getSessionUser(r *http.Request) (string, error) {
	var username string
	session, err := store.Get(r, sessionName)
	if err != nil {
		return "", err
	}

	val := session.Values["user"]
	fmt.Println("val: ", val)
	username, ok := val.(string)
	if !ok {
		return "", errors.New("Failed to get user from the session")
	}
	fmt.Println("username: ", username)
	return username, nil
}

func setSessionUser(w http.ResponseWriter, r *http.Request, username string) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values["user"] = username
	err = session.Save(r, w)
	return err
}

func clearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1 // clear the session

	err = session.Save(r, w)
	return err
}


// Login values check
func checkLen(fieldname, fieldvalue string, minLen, maxLen int) string {
	lenField := len(fieldname)
	if lenField < minLen || lenField > maxLen {
		return fmt.Sprintf("%s is not valid, the length must be between %d and %d", fieldname, minLen, maxLen)
	}
	return ""
}

func checkUsername(username string) string {
	return checkLen("Username", username, 3, 15)
}

func checkPassword(pwd string) string {
	return checkLen("Password", pwd, 6, 25)
}

func checkEmail(email string) string {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
        return fmt.Sprintf("Email is not valid")
    }
    return ""
}

func checkUsernamePassword(username, pwd string) string {
	if !vm.CheckLogin(username, pwd) {
		return fmt.Sprintf("Username and password are not correct.")
	}
	return ""
}

func checkUserExists(username string) string {
	if !vm.CheckUserExists(username) {
        return fmt.Sprintf("Username already exists, please choose another username")
    }
    return ""
}

func checkLoginDetails(username, pwd string) []string {
	var errs []string
    // if errCheck := checkUsername(username); len(errCheck) > 0 {
    //     errs = append(errs, errCheck)
    // }
    // if errCheck := checkPassword(pwd); len(errCheck) > 0 {
    //     errs = append(errs, errCheck)
    // }
    if errCheck := checkUsernamePassword(username, pwd); len(errCheck) > 0 {
        errs = append(errs, errCheck)
    }
    return errs
}

func checkRegisterDetails(username, email, pwd1, pwd2 string) []string {
    var errs []string
    if pwd1 != pwd2 {
        errs = append(errs, "The two passwords provided do not match")
    }
    if errCheck := checkUsername(username); len(errCheck) > 0 {
        errs = append(errs, errCheck)
    }
    if errCheck := checkPassword(pwd1); len(errCheck) > 0 {
        errs = append(errs, errCheck)
    }
    if errCheck := checkEmail(email); len(errCheck) > 0 {
        errs = append(errs, errCheck)
    }
    if errCheck := checkUserExists(username); len(errCheck) > 0 {
        errs = append(errs, errCheck)
    }
    return errs
}

func addUser(username, pwd, email string) error {
	return vm.AddUser(username, pwd, email)
}
