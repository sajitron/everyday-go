package astros

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type person struct {
	Name string `json:"name"`
}

type people struct {
	Number int      `json:"number"`
	Person []person `json:"people"`
}

func GetAstros(apiURL string) (people, error) {
	p := people{}

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return p, err
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return p, err
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}
