package entity

// User model for one user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type userDb struct {
	Data []User `json:"data"`
}

type userModel struct {
	storage
	users map[string]*User
}

var (
	// UserModel model for users
	UserModel userModel
)

func init() {
	addModel(&UserModel, "user_data")
}

// Init initialize a user model
func (model *userModel) Init(path string) {
	logger.Println("[usermodel] initializing")
	model.path = path
	model.users = make(map[string]*User)

	model.load()
	logger.Println("[usermodel] initialized")
}

// AddUser add a new user to database
func (model *userModel) AddUser(user *User) {
	logger.Println("[usermodel] try adding new user", user.Username)
	model.users[user.Username] = user
	model.dump()
	logger.Println("[usermodel] added new user", user.Username)
}

// FindUserCondition filter function to query user
type FindUserCondition func(*User) bool

// FindBy find userList with provided condition
func (model *userModel) FindBy(condition FindUserCondition) []User {
	result := []User{}
	for _, user := range model.users {
		if condition(user) {
			result = append(result, *user)
		}
	}
	return result
}

// FindByUsername find user by username
func (model *userModel) FindByUsername(username string) *User {
	return model.users[username]
}

func (model *userModel) load() {
	var userDb userDb
	model.storage.load(&userDb)
	for _, user := range userDb.Data {
		model.users[user.Username] = &user
	}
}

func (model *userModel) dump() {
	var userDb userDb
	for _, user := range model.users {
		userDb.Data = append(userDb.Data, *user)
	}
	model.storage.dump(&userDb)
}
