package service

import (
	"context"
	"errors"
	"machship/internal/core/domain"
	"machship/internal/core/mock"
	"testing"
)

func TestGetUsers(t *testing.T) {
	scenarios := []struct {
		testcase             string
		request              domain.RetrieveUsersRequest
		usernameExpectations map[string]struct {
			returnValue domain.UserData
			returnError error
		}
		totalDataReturned    int
		getUserResponseError error
	}{
		{
			testcase: "Normal test",
			request:  domain.RetrieveUsersRequest{Usernames: []string{"testuser", "anotheruser"}},
			usernameExpectations: map[string]struct {
				returnValue domain.UserData
				returnError error
			}{
				"testuser": {
					returnValue: domain.UserData{Login: "testuser", Followers: 0, PublicRepos: 0},
					returnError: nil,
				},
				"anotheruser": {
					returnValue: domain.UserData{Login: "anotheruser", Followers: 0, PublicRepos: 0},
					returnError: nil,
				},
			},
			totalDataReturned:    2,
			getUserResponseError: nil,
		},
		{
			testcase: "Non existing username removal test",
			request:  domain.RetrieveUsersRequest{Usernames: []string{"testuser", "anotheruser", "nonexistinguser", "anothernonexistinguser"}},
			usernameExpectations: map[string]struct {
				returnValue domain.UserData
				returnError error
			}{
				"testuser": {
					returnValue: domain.UserData{Login: "testuser", Followers: 0, PublicRepos: 0},
					returnError: nil,
				},
				"anotheruser": {
					returnValue: domain.UserData{Login: "anotheruser", Followers: 0, PublicRepos: 0},
					returnError: nil,
				},
				"nonexistinguser": {
					returnValue: domain.UserData{},
					returnError: nil,
				},
				"anothernonexistinguser": {
					returnValue: domain.UserData{},
					returnError: nil,
				},
			},
			totalDataReturned:    2,
			getUserResponseError: nil,
		},
		{
			testcase: "Api Client error test",
			request:  domain.RetrieveUsersRequest{Usernames: []string{"testuser"}},
			usernameExpectations: map[string]struct {
				returnValue domain.UserData
				returnError error
			}{
				"testuser": {
					returnValue: domain.UserData{},
					returnError: errors.New("API client error"),
				},
			},
			totalDataReturned:    0,
			getUserResponseError: errors.New("API client error"),
		},
	}

	for _, tc := range scenarios {
		mockApiClient := mock.NewMockApiRequest()

		mockApiClient.GetUserInfoFunc = func(ctx context.Context, username string) (domain.UserData, error) {
			exp, exists := tc.usernameExpectations[username]
			if exists {
				return exp.returnValue, exp.returnError
			}

			return domain.UserData{}, nil
		}

		userService := NewUserServiceImpl(mockApiClient)
		result, err := userService.GetUsers(context.Background(), tc.request)
		if err != nil && tc.getUserResponseError != nil && err.Error() != tc.getUserResponseError.Error() {
			t.Errorf("%s: expected status %v, got %v", tc.testcase, err, tc.getUserResponseError)
		} else if (err == nil && tc.getUserResponseError != nil) || (err != nil && tc.getUserResponseError == nil) {
			t.Errorf("%s: one of the errors is nil, but the other is not", tc.testcase)
		}

		if len(result.UserData) != tc.totalDataReturned {
			t.Errorf("%s: expected status %d, got %d", tc.testcase, len(result.UserData), tc.totalDataReturned)
		}
	}
}
