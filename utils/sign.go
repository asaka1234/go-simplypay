package utils

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parseExtra(ext map[string]any) string {
	if ext == nil {
		return ""
	}
	keys := make([]string, 0, len(ext))
	for key := range ext {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(ext))
	for _, key := range keys {
		value := ext[key]
		if value == nil || value == "" {
			continue
		}
		switch v := value.(type) {
		case float64:
			parts = append(parts, fmt.Sprintf("%s=%s", key, strconv.FormatFloat(v, 'f', -1, 64)))
		case int64:
			parts = append(parts, fmt.Sprintf("%s=%d", key, v))
		case string:
			parts = append(parts, fmt.Sprintf("%s=%s", key, v))
		default:
			parts = append(parts, key+"="+fmt.Sprintf("%v", v))
		}
	}
	return strings.Join(parts, "&")
}

func Sign(params map[string]any, appSecret string) string {
	// 1. 获取所有请求参数，不包括字节类型参数，如文件、字节流，剔除 sign 字段，
	keys := make([]string, 0, len(params))
	for key, value := range params {
		if key == "sign" || value == nil || value == "" {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 2. 按照字符的键值 ASCII 码递增排序
	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		value := params[key]
		if value == nil {
			continue
		}
		switch v := value.(type) {
		case float64:
			parts = append(parts, fmt.Sprintf("%s=%s", key, strconv.FormatFloat(v, 'f', -1, 64)))
		case int64:
			parts = append(parts, fmt.Sprintf("%s=%d", key, v))
		case string:
			parts = append(parts, fmt.Sprintf("%s=%s", key, v))
		case map[string]any:
			parts = append(parts, fmt.Sprintf("%s=%s", key, parseExtra(v)))
		default:
			parts = append(parts, key+"="+fmt.Sprintf("%v", v))
		}
	}

	// 3. 组合成“参数=参数值”的格式，并且把这些参数用 & 字符连接起来
	paramsString := strings.Join(parts, "&")

	// 4. 追加私钥
	if appSecret != "" {
		paramsString += "&key=" + appSecret
	}
	fmt.Println("sign: =====> ", paramsString)

	// 5. 计算签名
	h := sha256.New()
	h.Write([]byte(paramsString))
	sign := fmt.Sprintf("%x", h.Sum(nil))
	return sign
}

// 验证签名
func VerifySign(params map[string]interface{}, signKey string) bool {
	// Check if sign exists in params
	signValue, exists := params["sign"]
	if !exists {
		return false
	}

	// Get the sign value and remove it from params
	key := fmt.Sprintf("%v", signValue)
	delete(params, "sign")

	// Generate current signature
	currentKey := Sign(params, signKey)

	// Compare the signatures
	return key == currentKey
}
