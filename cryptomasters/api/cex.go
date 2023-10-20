package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"shtb.dev/go/crypto/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

// *datatype.Rate returning the address, since the address can be nil
func GetRate(currency string) (*datatypes.Rate, error) {

	if len(currency) != 3 {
		return nil, fmt.Errorf("The currency should be 3 letters.")
	}

	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}

	return &rate, nil
}
