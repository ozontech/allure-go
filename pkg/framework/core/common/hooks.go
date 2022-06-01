package common

import (
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"runtime/debug"
	"sync"
	"testing"
)

type HookFunc func(t InternalT, provider HookProvider, wg *sync.WaitGroup) (bool, error)

type HookType string

const (
	BeforeAll  HookType = "BeforeAll"
	AfterAll   HookType = "AfterAll"
	BeforeEach HookType = "BeforeEach"
	AfterEach  HookType = "AfterEach"
)

func CarriedHook(hook HookType, getHookBody func() func(t provider.T)) HookFunc {
	return func(t InternalT, provider HookProvider, wg *sync.WaitGroup) (result bool, err error) {
		result = true
		if hookBody := getHookBody(); hookBody != nil {
			t.WG().Add(1)
			oldT := t.RealT()
			defer t.SetRealT(oldT)
			// VERY dirt hack.
			// That allows let testing library control routines to avoid deadlocks and appropriate waiting
			go func() {
				result = t.RealT().Run(string(hook), func(realT *testing.T) {
					defer t.WG().Done()
					switch hook {
					case BeforeAll:
						provider.BeforeAllContext()
					case AfterAll:
						provider.AfterAllContext()
						//realT.Parallel()
						wg.Wait()
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
			}()
		}
		t.WG().Wait()
		return
	}
}

//
//func BeforeAllHook(t InternalT, provider HookProvider, wg *sync.WaitGroup) (err error) {
//	t.WG().Add(1)
//	defer t.WG().Done()
//
//	if provider.GetSuiteMeta().GetBeforeAll() != nil {
//		provider.BeforeAllContext()
//		oldT := t.RealT()
//		defer t.SetRealT(oldT)
//
//		t.RealT().Run("BeforeAll", func(realT *testing.T) {
//			t.SetRealT(realT)
//			defer func() {
//				r := recover()
//				if r != nil {
//					err = fmt.Errorf("BeforeAll hook panicked:%v\n%s", r, debug.Stack())
//					t.Errorf("BeforeAll hook panicked:%v\n%s", r, debug.Stack())
//					t.FailNow()
//				}
//			}()
//			provider.GetSuiteMeta().GetBeforeAll()(t)
//		})
//	}
//	return
//}
//
//func AfterAllHook(t InternalT, provider HookProvider, wg *sync.WaitGroup) (err error) {
//	t.WG().Add(1)
//	defer t.WG().Done()
//	if provider.GetSuiteMeta().GetAfterAll() != nil {
//		provider.AfterAllContext()
//		oldT := t.RealT()
//		defer t.SetRealT(oldT)
//
//		// VERY dirt hack
//		t.RealT().Run("AfterAll", func(realT *testing.T) {
//			realT.Parallel()
//			wg.Wait()
//			t.SetRealT(realT)
//			defer func() {
//				r := recover()
//				if r != nil {
//					err = fmt.Errorf("AfterAll hook panicked:%v\n%s", r, debug.Stack())
//					t.Errorf("AfterAll hook panicked:%v\n%s", r, debug.Stack())
//					t.FailNow()
//				}
//			}()
//			provider.GetSuiteMeta().GetAfterAll()(t)
//		})
//	}
//	return
//}
//
//func BeforeEachHook(t InternalT, provider HookProvider, wg *sync.WaitGroup) (err error) {
//	if provider.GetTestMeta().GetBeforeEach() != nil {
//		provider.BeforeEachContext()
//		oldT := t.RealT()
//		defer t.SetRealT(oldT)
//		t.RealT().Run("BeforeEach", func(realT *testing.T) {
//			t.SetRealT(realT)
//			defer func() {
//				r := recover()
//				if r != nil {
//					err = fmt.Errorf("BeforeEach hook panicked:%v\n%s", r, debug.Stack())
//					t.Errorf("BeforeEach hook panicked:%v\n%s", r, debug.Stack())
//					t.FailNow()
//				}
//			}()
//			provider.GetTestMeta().GetBeforeEach()(t)
//		})
//	}
//	return
//}
//
//func AfterEachHook(t InternalT, provider HookProvider, wg *sync.WaitGroup) (err error) {
//	if provider.GetTestMeta().GetAfterEach() != nil {
//		provider.AfterEachContext()
//		oldT := t.RealT()
//		defer t.SetRealT(oldT)
//
//		hookResult := t.RealT().Run("AfterEach", func(realT *testing.T) {
//			t.SetRealT(realT)
//			defer func() {
//				r := recover()
//				if r != nil {
//					err = fmt.Errorf("AfterEach hook panicked:%v\n%s", r, debug.Stack())
//					t.Errorf("AfterEach hook panicked:%v\n%s", r, debug.Stack())
//					t.FailNow()
//				}
//			}()
//			provider.GetTestMeta().GetAfterEach()(t)
//		})
//	}
//	return
//}
