package domain

func (user_name UserName) ExportName() string {
	return user_name.name
}
