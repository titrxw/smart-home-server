package http

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/titrxw/smart-home-server/app/pkg/helper"
	"github.com/titrxw/smart-home-server/app/pkg/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetSign(secret string, params map[string]string) string {
	_, ok := params["sign"]
	if ok {
		delete(params, "sign")
	}

	keys := make([]string, len(params))
	numFieldCount := 0
	for k, _ := range params {
		keys[numFieldCount] = k
		numFieldCount++
	}
	sort.Strings(keys)

	numFieldCount = 0
	paramList := make([]string, len(params))
	for _, k := range keys {
		paramList[numFieldCount] = fmt.Sprintf("%s=%v", k, params[k])
		numFieldCount++
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strings.Join(paramList, "&") + secret))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func PostWithAppSignByWWWForm(ctx context.Context, appid string, secret string, url string, body map[string]interface{}, headers map[string][]string, options ...http.Option) (map[string]interface{}, error) {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	params := make(map[string]string)
	params["appid"] = appid
	params["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["nonce"] = helper.RandomStr(16)
	params["body"] = string(bodyStr)
	params["sign"] = GetSign(secret, params)

	response, err := http.PostWithWWWForm(ctx, url, params, headers, options...)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(response.Body(), &data)
	if err != nil {
		return nil, err
	}
	msg, exists := data["error"]
	if exists {
		return nil, errors.New(msg.(string))
	}

	return data, nil
}
