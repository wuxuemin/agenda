package entity

// Session model for one login session
type Session struct {
	Username string `json:"username"`
}

type sessionDb struct {
	CurUser Session `json:"curUser"`
}

type sessionModel struct {
	storage
	session *Session
}

var (
	// CurSessionModel model for current session
	CurSessionModel sessionModel
)

func init() {
	addModel(&CurSessionModel, "curUser")
}

// Init initialize a session model
func (model *sessionModel) Init(path string) {
	logger.Println("[sessionmodel] initializing")
	model.path = path
	model.session = &Session{}

	model.load()
	logger.Println("[sessionmodel] initialized")
}

// SetCurUser sets current user in the session
func (model *sessionModel) SetCurUser(user *User) {
	logger.Printf("[sessionmodel] try setting user '%s' to current session\n", user.Username)
	model.session.Username = user.Username
	model.dump()
	logger.Printf("[sessionmodel] set user '%s' to current session\n", user.Username)
}

// GetCurUser get username of current user
func (model *sessionModel) GetCurUser() string {
	return model.session.Username
}

func (model *sessionModel) load() {
	var sessionDb sessionDb
	model.storage.load(&sessionDb)
	model.session = &sessionDb.CurUser
}

func (model *sessionModel) dump() {
	var sessionDb sessionDb
	sessionDb.CurUser = *model.session
	model.storage.dump(&sessionDb)
}
