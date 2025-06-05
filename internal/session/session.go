package session

import "github.com/gin-contrib/sessions"

const UserSession string = "sid"

func CreateUserSession(
	session sessions.Session,
	sessionName, id string,
) error {
	session.Set(sessionName, id)

	return session.Save()
}
