package main

// var nileBaseUrl = "https://nile.trongrid.io/"

// var (
// 	getAccountResourceUrl = nileBaseUrl + "wallet/getaccountresource"
// 	delegateResourceUrl   = nileBaseUrl + "wallet/delegateresource"
// 	freezeBalanceV2Url    = nileBaseUrl + "wallet/freezebalancev2"
// )

// type api struct{}

// func (w *wallet) GetAccountResource() (*AccountResourceResponse, error) {

// 	payload := strings.NewReader("{\"owner_address\":\"TZ4UXDV5ZhNW7fb2AMSbgfAEZ7hWsnYS2g\",\"frozen_balance\":10000000,\"frozen_duration\":3,\"resource\":\"ENERGY\",\"visible\":true}")

// 	accountResourceBytes, err := json.Marshal(accountResource)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to marshal the request %s", err)
// 	}

// 	payload := strings.NewReader(string(accountResourceBytes))
// 	req, err := http.NewRequest("POST", getAccountResourceUrl, payload)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to prepare the request %s", err)
// 	}

// 	req.Header.Add("accept", "application/json")
// 	req.Header.Add("content-type", "application/json")

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("error to Do the request %s", err)
// 	}
// 	defer res.Body.Close()

// 	var accountResourceResponse AccountResourceResponse
// 	if err := json.NewDecoder(res.Body).Decode(&accountResourceResponse); err != nil {
// 		return nil, fmt.Errorf("error to decode the response body %s", err)
// 	}

// 	return &accountResourceResponse, nil
// }

// func NewApi() *api {
// 	return &api{}
// }
