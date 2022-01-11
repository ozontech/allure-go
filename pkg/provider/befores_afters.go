package provider

//WithBeforeTest wraps f func() with Before Test context of allure (test setup)
func (t *T) WithBeforeTest(f func(t *T)) {
	t.beforeEach = func(_t *T) {
		_t.beforeTestState()
		f(_t)
		defer _t.runningTestState()
	}
}

//WithBeforeSuite wraps f func() with Before Suite context of allure (suite setup)
func (t *T) WithBeforeSuite(f func()) {
	t.beforeSuiteState()
	f()
	t.runningTestState()
}

//WithAfterTest wraps f func() with After Test context of allure (test tear down)
func (t *T) WithAfterTest(f func(t *T)) {
	t.afterEach = func(_t *T) {
		_t.afterTestState()
		f(_t)
		_t.runningTestState()
	}
}

//WithAfterSuite wraps f func() with After Suite context of allure (suite tear down)
func (t *T) WithAfterSuite(f func()) {
	t.afterSuiteState()
	f()
	t.runningTestState()
}

func ExecuteBefore(t *T, f func()) {
	t.beforeTestState()
	f()
	t.runningTestState()
}

func ExecuteAfter(t *T, f func()) {
	t.afterTestState()
	f()
	t.runningTestState()
}
