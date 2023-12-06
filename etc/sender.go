package etc

import (
	"context"
	"time"

	gohttpclient "github.com/bozd4g/go-http-client"
	jsoniter "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	models "github.com/mehdi-shokohi/go-utils/config"
)

func Send(c *fiber.Ctx, Code int, Data interface{}, Error interface{}) error {
	return c.Status(Code).JSON(models.Response{Data: Data, Error: Error})
}

func SentryLog() {

}

func SendHttpPost(c context.Context, headers map[string]string, address string, data interface{}) (*gohttpclient.Response, error) {
	opts := []gohttpclient.ClientOption{
		gohttpclient.WithDefaultHeaders(),
		gohttpclient.WithTimeout(time.Second * 3),
	}
	client := gohttpclient.New(address, opts...)
	json, err := jsoniter.Marshal(data)
	if err != nil {
		return nil, err
	}
	reqOpts := []gohttpclient.Option{
		gohttpclient.WithBody(json),
		gohttpclient.WithHeader("Content-type", "application/json"),
	}
	if headers != nil {
		for hKey, hv := range headers {
			reqOpts = append(reqOpts, gohttpclient.WithHeader(hKey, hv))
		}
	}

	response, err := client.Post(c, "", reqOpts...)
	if err != nil {
		return nil, err
	}
	return response, err
}
