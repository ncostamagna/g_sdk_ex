package user

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/ncostamagna/g_ms_domain_ex/domain"
	c "github.com/ncostamagna/go_http_client/client"
)

type (
	DataResponse struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}

	Transport interface {
		Get(id string) (*domain.User, error)
	}

	clientHTTP struct {
		client c.Transport
	}
)

func NewHttpClient(baseURLSGE, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &clientHTTP{
		client: c.New(header, baseURLSGE, 5000*time.Millisecond, true),
	}
}

// Get entities by ids
func (c *clientHTTP) Get(id string) (*domain.User, error) {

	dataResponse := DataResponse{Data: &domain.User{}}

	u := url.URL{}
	u.Path += fmt.Sprintf("/users/%s", id)
	reps := c.client.Get(u.String())
	if reps.Err != nil {
		return nil, reps.Err
	}

	if reps.StatusCode == 404 {
		return nil, ErrNotFound{fmt.Sprintf("%s", reps)}
	}

	if reps.StatusCode > 299 {
		return nil, fmt.Errorf("%s", reps)
	}

	if err := reps.FillUp(&dataResponse); err != nil {
		return nil, err
	}

	return dataResponse.Data.(*domain.User), nil
}
