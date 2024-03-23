package domain

func (user_name UserName) ExportName() string {
	return user_name.name
}

func (user_pass UserPass) ExportPass() string {
	return user_pass.pass
}
