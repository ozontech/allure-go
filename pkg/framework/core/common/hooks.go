package common

import (
	"fmt"
	"runtime/debug"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type HookFunc func(t InternalT, provider HookProvider) (bool, error)

type HookType string

// HookType constants
const (
	BeforeAll  HookType = "BeforeAll"
	AfterAll   HookType = "AfterAll"
	BeforeEach HookType = "BeforeEach"
	AfterEach  HookType = "AfterEach"
)

func CarriedHook(hook HookType, getHookBody func() func(t provider.T)) HookFunc {
	return func(t InternalT, provider HookProvider) (result bool, err error) {
		result = true
		if hookBody := getHookBody(); hookBody != nil {
			t.WG().Add(1)
			defer t.WG().Wait()

			// for correct logs
			oldT := t.RealT()
			defer t.SetRealT(oldT)

			// VERY dirt hack.
			// That allows let testing library control routines to avoid deadlocks and appropriate waiting
			result = t.RealT().Run(string(hook), func(realT *testing.T) {
				defer t.WG().Done()
				switch hook {
				case BeforeAll:
					provider.BeforeAllContext()
				case AfterAll:
					provider.AfterAllContext()
				case BeforeEach:
					provider.BeforeEachContext()
				case AfterEach:
					provider.AfterEachContext()
				}
				defer func() {
					r := recover()
					if r != nil {
						err = fmt.Errorf("%s hook panicked:%v\n%s", hook, r, debug.Stack())
						t.Errorf("%s hook panicked:%v\n%s", hook, r, debug.Stack())
						t.FailNow()
					}
				}()
				t.SetRealT(realT)
				hookBody(t)
			})
		}
		return
	}
}
