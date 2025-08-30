package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// LogOutgoingRequest mencetak request yang dikirim ke Shopee Xpress
func LogOutgoingRequest(method, url string, headers map[string]string, body []byte) {
	fmt.Println("===== [SPX REQUEST] =====")
	fmt.Printf("TIME      : %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("METHOD    : %s\n", method)
	fmt.Printf("URL       : %s\n", url)
	fmt.Printf("HEADERS   :\n")
	for k, v := range headers {
		fmt.Printf("  %s: %s\n", k, v)
	}
	if len(body) > 0 {
		fmt.Println("BODY      :")
		fmt.Println(prettyJSON(body))
	} else {
		fmt.Println("BODY      : {}")
	}
	fmt.Println("=========================")
}

// LogIncomingResponse mencetak respons mentah dari Shopee Xpress
func LogIncomingResponse(resp *http.Response, respBody []byte) {
	fmt.Println("===== [SPX RESPONSE] =====")
	fmt.Printf("TIME      : %s\n", time.Now().Format(time.RFC3339))
	fmt.Printf("STATUS    : %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("BODY      :")
	fmt.Println(prettyJSON(respBody))
	fmt.Println("==========================")
}

// prettyJSON menata format JSON agar rapi
func prettyJSON(b []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		return string(b)
	}
	return out.String()
}
