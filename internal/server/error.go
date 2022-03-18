package server

import "forum/internal"

func (c *Dbs) CheckUniq(m *internal.User) bool {
	notUn := 0
	notUniq := c.Db.QueryRow((`SELECT user.id FROM user WHERE user.name = ?`), m.Name)
	notUniq.Scan(&notUn)
	if notUn != 0 {
		m.ErrorL = true
		return false
	}
	notUnE := 0
	notUniqE := c.Db.QueryRow((`SELECT user.id FROM user WHERE user.email = ?`), m.Email)
	notUniqE.Scan(&notUnE)
	if notUnE != 0 {
		m.ErrorEm = true
		return false
	}
	return true
}
func (c *Dbs) CheckEmail(m *internal.User) bool {
	for _, i := range m.Email {
		if i == '@' {
			for _, j := range m.Email {
				if j == '.' {

					return true
				}
			}

		}
	}
	m.ErrorE = true
	return false
}
func (c *Dbs) CheckEmpty(m *internal.User) bool {
	if len(m.Name) == 0 {
		m.ErrorEmpty = true
		return false
	}
	for _, i := range m.Name {
		if i == ' ' {
			m.ErrorEmpty = true
			return false
		}
	}
	if len(m.Email) == 0 {
		m.ErrorEmpty = true
		return false
	}
	for _, i := range m.Email {
		if i == ' ' {
			m.ErrorEmpty = true
			return false
		}
	}
	if len(m.Password) == 0 {
		m.ErrorEmpty = true
		return false
	}
	for _, i := range m.Password {
		if i == ' ' {
			m.ErrorEmpty = true
			return false
		}
	}
	return true
}
