package domain

type RetrieveUsersRequest struct {
	Usernames []string `json:"usernames,omitempty"`
}

type GetUserResponse struct {
	UserData []FormattedUserData `json:"user_data"`
}

type UserData struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	Followers   int    `json:"followers"`
	PublicRepos int    `json:"public_repos"`
}

type FormattedUserData struct {
	Login            string `json:"login"`
	Name             string `json:"name"`
	Company          string `json:"company"`
	Followers        int    `json:"number_of_followers"`
	PublicRepos      int    `json:"number_of_public_repositories"`
	AverageFollowers int    `json:"average_followers"`
}
