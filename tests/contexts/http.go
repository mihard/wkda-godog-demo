package contexts

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/cucumber/godog"
//	"io/ioutil"
//	"net/http"
//)
//
//type httpContext struct {
//	host string
//	resp *http.Response
//	body []byte
//}
//
//func InitHTTPFeature(serviceHost string, s *godog.ScenarioContext) *httpContext {
//	feature := &httpContext{host: serviceHost}
//
//	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, feature.iSendRequestTo)
//	s.Step(`^the response code should be (\d+)$`, feature.theResponseCodeShouldBe)
//	s.Step(`^the response should match json:$`, feature.theResponseShouldMatchJSON)
//
//	return feature
//}
//
//func (c *httpContext) iSendRequestTo(method, endpoint string) (err error) {
//	req, err := http.NewRequest(method, c.host+endpoint, nil)
//	if err != nil {
//		return
//	}
//
//	c.resp, err = http.DefaultClient.Do(req)
//	if err != nil {
//		return
//	}
//	defer c.resp.Body.Close()
//	c.body, err = ioutil.ReadAll(c.resp.Body)
//	if err != nil {
//		return
//	}
//
//	return
//}
//
//func (c *httpContext) theResponseCodeShouldBe(code int) error {
//	if c.resp == nil {
//		return fmt.Errorf("you need to make an http call first")
//	}
//	if code != c.resp.StatusCode {
//		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, c.resp.StatusCode)
//	}
//	return nil
//
//}
//
//func (c *httpContext) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
//	if c.body == nil {
//		return fmt.Errorf("you need to make an http call first")
//	}
//
//	var expected, actual []byte
//	var data interface{}
//	if err = json.Unmarshal([]byte(body.Content), &data); err != nil {
//		return
//	}
//	if expected, err = json.Marshal(data); err != nil {
//		return
//	}
//	actual = c.body
//	if !bytes.Equal(actual, expected) {
//		err = fmt.Errorf("expected json, does not match actual: %s", string(actual))
//	}
//	return
//}
