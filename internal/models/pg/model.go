package pg

type UserDataParams struct {
	Name         string `json:"name"`
	Mail         string `json:"mail"`
	Countries_id int    `json:"countries_id"`
}
