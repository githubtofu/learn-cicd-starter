package auth

import (
    "errors"
	"testing"
    "net/http"
)

var malformed = errors.New("malformed authorization header")

func TestGetApi(t *testing.T) {
    tests := []struct {
        input_key string
        input_value string
        f_resp string
        f_err error
    }{
        { input_key:"", input_value:"", f_resp:"", f_err: ErrNoAuthHeaderIncluded},
        { input_key:"Authorization", input_value:"KEYEX", f_resp:"", f_err:malformed},
        { input_key:"Authorization", input_value:"ApiKey KEYEX", f_resp:"KEYEX", f_err:nil},
    }
    for i, test := range tests{
        t.Log("TEST====================")
        req, _ := http.NewRequest("GET", "someURL", nil)
        req.Header.Set(test.input_key, test.input_value)
        api_func_resp, api_func_err := GetAPIKey(req.Header)
        if api_func_resp != test.f_resp {
            t.Fatalf("Test %d: expected: %v, got: %v", i+1, test.f_resp, api_func_resp)
        }
        if api_func_err == nil || test.f_err == nil {
            if api_func_err != nil || test.f_err != nil {
                t.Fatalf("Test %d: ERR expected: %v, got: %v", i+1, test.f_err, api_func_err)
            }
        }else if api_func_err.Error() != test.f_err.Error() {
            t.Fatalf("Test %d: ERR expected: %v, got: %v", i+1, test.f_err, api_func_err)
        }
        t.Logf("Test succeeded: %d", i + 1)
    }
}
