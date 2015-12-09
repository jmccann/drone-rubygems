package main

type Params struct {
	APIKey      string `json:"api_key"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Gemspec     string `json:"gemspec"`
	Host        string `json:"host"`
	SkipCleanup bool   `json:"skip_cleanup"`
}
