package router

type any interface{}
type H map[string]any

type LoginPost struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterPost struct {
	Usename   string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birth     string `json:"birth"`
}
