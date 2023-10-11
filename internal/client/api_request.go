package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"machship/internal/core/constant"
	"machship/internal/core/domain"
	"machship/internal/core/port"
	"machship/internal/core/util"
	"net/http"
)

type apiRequestImpl struct {
	client *http.Client
}

func NewApiRequest() port.ApiRequest {
	apiRequest := new(apiRequestImpl)
	apiRequest.client = &http.Client{}

	return apiRequest
}

func (a *apiRequestImpl) GetUserInfo(ctx context.Context, username string) (domain.UserData, error) {
	util.Infoln(ctx, "### Start GetUserInfo() client")
	var userData domain.UserData

	url := fmt.Sprintf("%s%s", constant.GITHUB_ENDPOINT, username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		util.Errorf("Error creating request: %v", err.Error())
		return userData, err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		util.Errorf("Error sending request: %v", err.Error())
		return userData, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.Errorf("Error reading response body: %v", err.Error())
		return userData, err
	}

	if err := json.Unmarshal(body, &userData); err != nil {
		util.Errorf("Error Unmarshalling response body: %v", err.Error())
		return userData, err
	}

	util.Infoln(ctx, "### End GetUserInfo() client")
	return userData, nil
}
