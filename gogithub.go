// Go implementation of the GitHub API v3 (http://developer.github.com/v3/)
package gogithub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GitHub API URL
const githubApiUrl = "https://api.github.com"

// Authenticated GitHub user.
// Only supports Basic Authentication at this time.
type AuthenticatedUser struct {
	username string
	password string
}

// GitHub User Plan (http://developer.github.com/v3/users/)
type GithubUserPlan struct {
	Name          string `json:"name"`
	Space         int    `json:"space"`
	Collaborators int    `json:"collaborators"`
	PrivateRepos  int    `json:"private_repos"`
}

// GitHub User Email (http://developer.github.com/v3/users/emails/)
type GithubUserEmails []string

// GitHub user interface (http://developer.github.com/v3/users/)
type GithubUser struct {
	Login       string         `json:"login"`
	Id          int            `json:"id"`
	AvatarUrl   string         `json:"avatar_url"`
	GravatarId  string         `json:"gravatar_id"`
	Url         string         `json:"url"`
	Name        string         `json:"name"`
	Company     string         `json:"company"`
	Blog        string         `json:"blog"`
	Location    string         `json:"location"`
	Email       string         `json:"email"`
	Hireable    bool           `json:"hirable"`
	Bio         string         `json:"bio"`
	PublicRepos int            `json:"public_repos"`
	PublicGists int            `json:"public_gists"`
	Followers   int            `json:"followers"`
	Following   int            `json:"following"`
	HtmlUrl     string         `json:"html_url"`
	CreatedAt   string         `json:"created_at"`
	Type        string         `json:"type"`
	Plan        GithubUserPlan `json:"plan"`
}

// Creates an AuthenticatedUser object that stores the username and password.
func Client(credentials map[string]string) (AuthenticatedUser, error) {
	user := AuthenticatedUser{}

	// Empty map -- no credentials
	if len(credentials) == 0 {
		return user, nil
	}

	// Map has something, check what it is.
	username := credentials["username"]
	password := credentials["password"]
	if username != "" && password != "" {
		user.username = username
		user.password = password
	} else {
		return AuthenticatedUser{}, fmt.Errorf("Unsupported authentication type. Only username/password supported at this time.")
	}

	return user, nil
}

func (u *AuthenticatedUser) get(path string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Could not create new request: %s", err)
	}

	if u.username != "" && u.password != "" {
		req.SetBasicAuth(u.username, u.password)
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("Could not process request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Received non-OK status from Github: %s", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Could not read body: %s", err)
	}

	return body, nil
}

// Fetches user details and returns an object of type GithubUser.
func (u *AuthenticatedUser) GetUser(username string) (GithubUser, error) {
	var path string
	// If you pass in your own username, or you pass in an empty username and you have credentials
	// then get "yourself".
	// Otherwise, fetch the user details unauthenticated.
	if u.username == username || (u.username != "" && username == "") {
		path = fmt.Sprintf("%s/user", githubApiUrl)
	} else {
		path = fmt.Sprintf("%s/users/%s", githubApiUrl, username)
	}

	body, err := u.get(path)
	if err != nil {
		return GithubUser{}, fmt.Errorf("Could not GET: %s", err)
	}

	githubUser := GithubUser{}
	if err = json.Unmarshal(body, &githubUser); err != nil {
		return GithubUser{}, fmt.Errorf("Could not parse Github user response: %s", err)
	}

	return githubUser, nil
}

// Fetches the list of user emails and returns GithubUserEmails.
func (u *AuthenticatedUser) GetEmails() (GithubUserEmails, error) {
	// Precondition check since this only works when you're authenticated.
	if u.username == "" || u.password == "" {
		return GithubUserEmails{}, fmt.Errorf("You must be authenticated to fetch emails.")
	}

	path := fmt.Sprintf("%s/user/emails", githubApiUrl)

	body, err := u.get(path)
	if err != nil {
		return GithubUserEmails{}, fmt.Errorf("Could not GET: %s", err)
	}

	githubUserEmails := GithubUserEmails{}
	if err = json.Unmarshal(body, &githubUserEmails); err != nil {
		return GithubUserEmails{}, fmt.Errorf("Could not parse Github user email resposne: %s", err)
	}

	return githubUserEmails, nil
}
