package sendcloud

import (
    "encoding/json"
    "testing"
)

func TestSend(t *testing.T) {
    client := &Client{
        APIUser: "imxfly",
        APIKey:  "DJYppkdfM50tBhO1",
    }

    params := map[string]string{
        "from":     "support@blog.imxfly.com",
        "to":       "stephenfxl@gmail.com",
        "subject":  "测试",
        "html":     "<h1>仅测试</h1>",
        "fromName": "X.FLY",
    }
    _, err := client.Send(params)
    if err != nil {
        t.Errorf("普通发送失败，错误信息：%s", err.Error())
    }
}

func TestClient_SendTemplate(t *testing.T) {
    client := &Client{
       APIUser: "imxfly",
       APIKey:  "DJYppkdfM50tBhO1",
    }

    xsmtpapi := map[string]interface{}{
        "to": []string{"stephenfxl@gmail.com"},
        "sub": map[string][]string{
            "%title%":    {"测试"},
            "%time%":     {"2020"},
            "%nickname%": {"stephen"},
            "%content%":  {"加油！"},
            "%url%":      {"https://blog.imxfly.com"},
        },
    }

    xsmtpapiStr, _ := json.Marshal(xsmtpapi)

    params := map[string]string{
        "from":               "support@blog.imxfly.com",
        "xsmtpapi":           string(xsmtpapiStr),
        "templateInvokeName": "new_comment",
        "fromName":           "X.FLY",
    }
    _, err := client.SendTemplate(params)
    if err != nil {
       t.Errorf("普通发送失败，错误信息：%s", err.Error())
    }
}
