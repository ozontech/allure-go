package provider

type AllureContext interface {
	setState(state stepState)
	getState() stepState
	beforeTestState()
	runningTestState()
	afterTestState()
	beforeSuiteState()
	afterSuiteState()
}

//contextController allows control allure context (test/setup setup/tear down)
type contextController struct {
	test        stepState
	beforeTest  stepState
	afterTest   stepState
	beforeSuite stepState
	afterSuite  stepState

	currentState stepState
}

func newStateMachine(t *T) *contextController {
	test := testState{t: t, name: "Test"}
	bt := beforeTest{t: t, name: "BeforeTest"}
	at := afterTest{t: t, name: "AfterTest"}
	bs := beforeSuite{t: t, name: "BeforeSuite"}
	as := afterSuite{t: t, name: "AfterSuite"}

	return &contextController{
		test:         &test,
		beforeTest:   &bt,
		afterTest:    &at,
		beforeSuite:  &bs,
		afterSuite:   &as,
		currentState: &test,
	}
}

func (s *contextController) setState(state stepState) {
	s.currentState = state
}

func (s *contextController) getState() stepState {
	return s.currentState
}

func (s *contextController) beforeTestState() {
	s.setState(s.beforeTest)
}

func (s *contextController) afterTestState() {
	s.setState(s.afterTest)
}

func (s *contextController) runningTestState() {
	s.setState(s.test)
}

func (s *contextController) beforeSuiteState() {
	s.setState(s.beforeSuite)
}

func (s *contextController) afterSuiteState() {
	s.setState(s.afterSuite)
}
