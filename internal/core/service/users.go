package service

import (
	"context"
	"machship/internal/core/domain"
	"machship/internal/core/port"
	"machship/internal/core/util"
	"sort"
)

type UserService struct {
	apiClient port.ApiRequest
}

func NewUserServiceImpl(apiClient port.ApiRequest) port.UserService {
	userService := new(UserService)
	userService.apiClient = apiClient
	return userService
}

func (u *UserService) GetUsers(ctx context.Context, request domain.RetrieveUsersRequest) (domain.GetUserResponse, error) {
	util.Infoln(ctx, "### Start GetUsers() service")
	var response domain.GetUserResponse
	var userData []domain.FormattedUserData

	userDataChan := make(chan *domain.UserData)
	errorChan := make(chan error)

	usernames := sanitizePayload(request.Usernames)
	for _, user := range usernames {
		go u.getUserInfo(ctx, user, userDataChan, errorChan)
	}

	for _, _ = range usernames {
		u := <-userDataChan
		e := <-errorChan
		if e != nil {
			return response, e
		}
		if u.Login != "" {
			formmatedUserData := formatUserData(*u)
			userData = append(userData, formmatedUserData)
		}
	}

	response.UserData = sortUser(userData)

	util.Infoln(ctx, "### End GetUsers() service")
	return response, nil
}

func (u *UserService) getUserInfo(ctx context.Context, username string, userDataChan chan *domain.UserData, errorChan chan error) {
	userInfo, err := u.apiClient.GetUserInfo(ctx, username)

	userDataChan <- &userInfo
	errorChan <- err
}

func sortUser(userData []domain.FormattedUserData) []domain.FormattedUserData {
	sort.Slice(userData, func(i, j int) bool {
		return userData[i].Login < userData[j].Login
	})

	return userData
}

func formatUserData(u domain.UserData) domain.FormattedUserData {
	var formattedUserData domain.FormattedUserData
	formattedUserData.Login = u.Login
	formattedUserData.Name = u.Name
	formattedUserData.Company = u.Company
	formattedUserData.Followers = u.Followers
	formattedUserData.PublicRepos = u.PublicRepos

	if u.PublicRepos > 0 {
		formattedUserData.AverageFollowers = u.Followers / u.PublicRepos
	} else {
		formattedUserData.AverageFollowers = 0
	}

	return formattedUserData
}

func sanitizePayload(usernames []string) []string {
	uniqueMap := make(map[string]bool)
	result := []string{}

	for _, str := range usernames {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			result = append(result, str)
		}
	}

	return result
}
