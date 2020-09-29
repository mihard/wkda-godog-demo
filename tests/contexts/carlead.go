package contexts

//
//import (
//	"fmt"
//	"github.com/cucumber/godog"
//	"github.com/cucumber/messages-go/v10"
//	"net/http"
//)
//
//type carLeadContext struct {
//	wm *wireMockContext
//}
//
//func InitCarLead(wm *wireMockContext, s *godog.ScenarioContext) *carLeadContext {
//	cl := &carLeadContext{wm}
//
//	s.Step(`^I expect to receive following response from carlead-service with carId (\d+)$`, cl.IExpectToHaveACarLead)
//
//	return cl
//}
//
//func (b *carLeadContext) IExpectToHaveACarLead(id int64, expectedResponse *messages.PickleStepArgument_PickleDocString) error {
//	path := fmt.Sprintf("/v1/carlead/%d", id)
//
//	_, err := b.wm.AddStub(&Stub{
//		Request: &StubRequest{
//			Method: http.MethodGet,
//			Url:    path,
//		},
//		Response: &StubResponse{
//			Body: expectedResponse.Content,
//			Headers: map[string]string{
//				"Content-type": "application/json",
//			},
//			Status: 200,
//		},
//	})
//
//	return err
//}
//
//func (b *carLeadContext) IExpectNotFound(id int64) error {
//	path := fmt.Sprintf("/v1/carlead/%d", id)
//
//	_, err := b.wm.AddStub(&Stub{
//		Request: &StubRequest{
//			Method: http.MethodGet,
//			Url:    path,
//		},
//		Response: &StubResponse{
//			Body: "",
//			Headers: map[string]string{
//				"Content-type": "application/json",
//			},
//			Status: 404,
//		},
//	})
//
//	return err
//}
