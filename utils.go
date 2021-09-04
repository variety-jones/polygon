package polygon

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// unixTimeNow returns the current unix time in string format
func unixTimeNow() string {
	currentTime := time.Now()
	seconds := strconv.FormatInt(currentTime.Unix(), 10)

	return seconds
}

// createSHA512Hash creates the SHA512 hash of the the input
func createSHA512Hash(toHash string) string {
	hasher := sha512.New()
	hasher.Write([]byte(toHash))
	hashedString := hex.EncodeToString(hasher.Sum(nil))

	return hashedString
}

// generateRandomPrefix creates a random string of length 6
func generateRandomPrefix(stringLength int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, stringLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// prepareURL prepares the URL according to polygon's API criteria
func (api *PolygonApi) prepareURL(parametersOriginal map[string]string, methodName string) string {
	sec := unixTimeNow()

	// First, copy the map, so that you do not modify user's map
	parameters := make(map[string]string)
	for key, value := range parametersOriginal {
		parameters[key] = value
	}

	// Add "time" and "apiKey" parameter
	parameters["time"] = sec
	parameters["apiKey"] = api.ApiKey

	// Add "problemId" parameter only if it is a prob
	if !(methodName == "problems.list" || methodName == "contest.problems") {
		parameters["problemId"] = api.ProblemId
	}

	// Extract all keys of the map
	keys := make([]string, len(parameters))
	index := 0
	for k := range parameters {
		keys[index] = k
		index++
	}

	// Sort all the keys
	sort.Strings(keys)

	// Create the parameters encoding
	commonPart := methodName + "?"
	commonPartEscaped := commonPart

	// Concatenate the parameters in sorted order
	first := true
	for _, key := range keys {
		if !first {
			commonPartEscaped += "&"
			commonPart += "&"
		}

		commonPartEscaped += key + "=" + url.QueryEscape(parameters[key])
		commonPart += key + "=" + parameters[key]
		first = false
	}

	randPrefix := generateRandomPrefix(6)
	toHash := randPrefix + "/" + commonPart + "#" + api.Secret
	hashedString := createSHA512Hash(toHash)
	requestURL := base + commonPartEscaped + "&apiSig=" + randPrefix + hashedString
	return requestURL
}

// processRequest makes API calls and reports if there was any error
func (api *PolygonApi) processRequest(parameters map[string]string, methodName string) (body []byte, err error) {
	URL := api.prepareURL(parameters, methodName)
	resp, err := http.Get(URL)
	if err != nil {
		return body, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	// TODO : Revisit this section later
	if resp.StatusCode != http.StatusOK {
		mp := make(map[string]string)
		json.Unmarshal(body, &mp)
		return body, errors.New("the http request returned a FAILED status with comment: " + mp["comment"])
	}
	return body, err
}

// extractView is a utility function for all methods that return a view
func (api *PolygonApi) extractView(parameters map[string]string, methodName string) (viewName string, err error) {
	body, err := api.processRequest(parameters, methodName)
	if err != nil {
		return viewName, err
	}
	return string(body), err
}

// extractName is a utility function for all methods that return a name of the resource
func (api *PolygonApi) extractName(parameters map[string]string, methodName string) (name string, err error) {
	body, err := api.processRequest(parameters, methodName)
	if err != nil {
		return name, err
	}

	wrapper := wrapperString{}
	err = json.Unmarshal(body, &wrapper)
	if err != nil {
		return name, err
	}
	return wrapper.Result, err
}

// checkForErrors is a utility function.
// It checks whether an API call which returns nothing succeeded or not
func (api *PolygonApi) checkForErrors(parameters map[string]string, methodName string) (err error) {
	_, err = api.processRequest(parameters, methodName)
	return err
}
