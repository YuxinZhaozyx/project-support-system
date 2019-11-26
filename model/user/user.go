package user

type User struct {
	Id			int			`json:"user_id"`
	Name		string		`json:"name"`
	Pwd			string		`json:"pwd"`
	Phone		string		`json:"phone"`
}

type Group struct{
	Id			int			`json:"group_id"`
	Description	string		`json:"description"`
	IdLeader	int 		`json:"leader_id"`
}

func NewUser(id int, name string, pwd string, phone string) (*User, error){
	user := User {id, name, pwd, phone}
	
	return &user, nil
}

func NewGroup(id int, description string, idLeader int)(*Group, error){
	group := Group{id, description, idLeader}

	return &group, nil
}