package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/novychok/niletrongrid/models"
)

const (
	METHOD_POST = "POST"
	METHOD_GET  = "GET"
)

var nileBase = "https://nile.trongrid.io/"

func ApiCall(ctx context.Context, url string, payload string, method string) ([]byte, error) {

	p := strings.NewReader(payload)

	req, err := http.NewRequestWithContext(ctx, method, nileBase+url, p)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func BroadcastHex(ctx context.Context, signedTx, broadcastHexUrl, method string) ([]byte, error) {

	broadcastHex := &models.BroadcastHex{
		Transaction: signedTx,
	}

	jsonBroadcastHex, err := json.Marshal(broadcastHex)
	if err != nil {
		return nil, err
	}

	broadcastHexResponse, err := ApiCall(ctx, broadcastHexUrl, string(jsonBroadcastHex), method)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(broadcastHexResponse))

	return broadcastHexResponse, nil
}
