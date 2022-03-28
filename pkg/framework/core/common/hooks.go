package common

import (
	"runtime/debug"
)

func BeforeAllHook(t InternalT, provider HookProvider) {
	t.WG().Add(1)
	defer t.WG().Done()
	if provider.GetSuiteMeta().GetBeforeAll() != nil {
		provider.BeforeAllContext()
		defer func() {
			r := recover()
			if r != nil {
				t.Errorf("BeforeAll hook panicked:%v\n%s", r, debug.Stack())
				t.FailNow()
			}
		}()
		provider.GetSuiteMeta().GetBeforeAll()(t)
	}
}

func AfterAllHook(t InternalT, provider HookProvider) {
	t.WG().Add(1)
	defer t.WG().Done()
	if provider.GetSuiteMeta().GetAfterAll() != nil {
		provider.AfterAllContext()
		defer func() {
			r := recover()
			if r != nil {
				t.Errorf("AfterAll hook panicked:%v\n%s", r, debug.Stack())
				t.FailNow()
			}
		}()
		provider.GetSuiteMeta().GetAfterAll()(t)
	}
}

func BeforeEachHook(t InternalT, provider HookProvider) {
	if provider.GetTestMeta().GetBeforeEach() != nil {
		provider.BeforeEachContext()
		defer func() {
			r := recover()
			if r != nil {
				t.Errorf("BeforeEach hook panicked:%v\n%s", r, debug.Stack())
				t.FailNow()
			}
		}()
		provider.GetTestMeta().GetBeforeEach()(t)
	}
}

func AfterEachHook(t InternalT, provider HookProvider) {
	if provider.GetTestMeta().GetAfterEach() != nil {
		provider.AfterEachContext()
		defer func() {
			r := recover()
			if r != nil {
				t.Errorf("AfterEach hook panicked:%v\n%s", r, debug.Stack())
				t.FailNow()
			}
		}()
		provider.GetTestMeta().GetAfterEach()(t)
	}
}
