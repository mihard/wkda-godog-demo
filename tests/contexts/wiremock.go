package contexts

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/cucumber/godog"
//	"net/http"
//)
//
//type wireMockContext struct {
//	url string
//}
//
//type Stub struct {
//	RequiredScenarioState string        `json:"requiredScenarioState,omitempty"`
//	Request               *StubRequest  `json:"request"`
//	Response              *StubResponse `json:"response"`
//}
//
//type BodyPatterns []map[string]string
//
//type StubRequest struct {
//	Method string       `json:"method"`
//	Url    string       `json:"url"`
//	Body   BodyPatterns `json:"bodyPatterns"`
//}
//
//type StubResponse struct {
//	Body    string            `json:"body"`
//	Headers map[string]string `json:"headers"`
//	Status  int               `json:"status"`
//}
//
//func InitWireMockFeature(url string, s *godog.ScenarioContext) *wireMockContext {
//	return &wireMockContext{
//		fmt.Sprintf("%s/__admin/mappings", url),
//	}
//}
//
//func (w *wireMockContext) RemoveAllStubs() (*http.Response, error) {
//	req, err := http.NewRequest(http.MethodDelete, w.url, nil)
//	if err != nil {
//		return nil, err
//	}
//	return http.DefaultClient.Do(req)
//}
//
//func (w *wireMockContext) AddStub(body *Stub) (*http.Response, error) {
//	data, err := json.Marshal(body)
//	if err != nil {
//		return nil, err
//	}
//	buffer := bytes.NewBuffer(data)
//
//	req, err := http.NewRequest(http.MethodPost, w.url, buffer)
//	if err != nil {
//		return nil, err
//	}
//	resp, err := http.DefaultClient.Do(req)
//	if err == nil {
//		defer resp.Body.Close()
//	}
//
//	return resp, err
//}
