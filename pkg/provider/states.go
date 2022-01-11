package provider

import (
	"github.com/koodeex/allure-testify/pkg/allure"
)

type stepState interface {
	addStep(step *allure.Step)
	addNested(step *allure.Step)
	finishNesting()
	addNestedParam(...allure.Parameter)
	addAttachment(attachment *allure.Attachment)
	addNestedAttachment(attachment *allure.Attachment)
}

/*
	isTestState
*/

type testState struct {
	t    *T
	name string
}

func (i *testState) addStep(step *allure.Step) {
	i.t.safely(func(result *allure.Result) {
		result.Steps = addStep(step, result.Steps, &result.StepsQueue)
	})
}

func (i *testState) addNested(step *allure.Step) {
	i.t.safely(func(result *allure.Result) {
		i.addStep(step)
		addNested(result.Steps, &result.StepsQueue)
	})
}

func (i *testState) finishNesting() {
	i.t.safely(func(result *allure.Result) {
		finishNesting(result.Steps, &result.StepsQueue)
	})
}

func (i *testState) addNestedParam(params ...allure.Parameter) {
	i.t.safely(func(result *allure.Result) {
		addNestingParameter(params, &result.StepsQueue)
	})
}

func (i *testState) addAttachment(attachment *allure.Attachment) {
	i.t.safely(func(result *allure.Result) {
		result.Attachments = append(result.Attachments, attachment)
	})
}

func (i *testState) addNestedAttachment(attach *allure.Attachment) {
	i.t.safely(func(result *allure.Result) {
		addNestingAttachment(attach, &result.StepsQueue)
	})
}

/*
	beforeTest state
*/

type beforeTest struct {
	t    *T
	name string
}

func (i *beforeTest) addStep(step *allure.Step) {
	i.t.safely(func(result *allure.Result) {
		container := result.Container
		container.Befores = addStep(step, container.Befores, &container.BeforesQueue)
	})
}

func (i *beforeTest) addNested(step *allure.Step) {
	i.t.safely(func(result *allure.Result) {
		i.addStep(step)
		container := result.Container
		addNested(container.Befores, &container.BeforesQueue)
	})
}

func (i *beforeTest) finishNesting() {
	i.t.safely(func(result *allure.Result) {
		container := result.Container
		finishNesting(container.Befores, &container.BeforesQueue)
	})
}

func (i *beforeTest) addNestedParam(params ...allure.Parameter) {
	i.t.safely(func(result *allure.Result) {
		container := result.Container
		addNestingParameter(params, &container.BeforesQueue)
	})
}

func (i *beforeTest) addAttachment(attachment *allure.Attachment) {
	i.t.safely(func(result *allure.Result) {
		addAttachmentToContainer(i, attachment)
	})
}

func (i *beforeTest) addNestedAttachment(attachment *allure.Attachment) {
	i.t.safely(func(result *allure.Result) {
		container := result.Container
		addNestingAttachment(attachment, &container.BeforesQueue)
	})
}

/*
	afterTest state
*/

type afterTest struct {
	t    *T
	name string
}

func (i afterTest) addStep(step *allure.Step) {
	i.t.safely(func(result *allure.Result) {
		container := result.Container
		container.Afters = addStep(step, container.Afters, &container.AftersQueue)
	})
}

func (i *afterTest) addNested(step *allure.Step) {
	i.t.safely(func(result *allure.Result) {
		i.addStep(step)
		container := result.Container
		addNested(container.Afters, &container.AftersQueue)
	})
}

func (i *afterTest) finishNesting() {
	i.t.safely(func(result *allure.Result) {
		container := result.Container
		finishNesting(container.Afters, &container.AftersQueue)
	})
}

func (i *afterTest) addNestedParam(params ...allure.Parameter) {
	i.t.safely(func(result *allure.Result) {
		container := result.Container
		addNestingParameter(params, &container.AftersQueue)
	})
}

func (i *afterTest) addAttachment(attachment *allure.Attachment) {
	i.t.safely(func(result *allure.Result) {
		addAttachmentToContainer(i, attachment)
	})
}

func (i *afterTest) addNestedAttachment(attachment *allure.Attachment) {
	i.t.safely(func(result *allure.Result) {
		container := result.Container
		addNestingAttachment(attachment, &container.AftersQueue)
	})
}

/*
	beforeSuite state
*/

type beforeSuite struct {
	t    *T
	name string
}

func (i beforeSuite) addStep(step *allure.Step) {
	container := i.t.GetContainer()
	container.Befores = addStep(step, container.Befores, &container.BeforesQueue)
}

func (i *beforeSuite) addNested(step *allure.Step) {
	i.addStep(step)
	container := i.t.GetContainer()
	addNested(container.Befores, &container.BeforesQueue)
}

func (i *beforeSuite) finishNesting() {
	container := i.t.GetContainer()
	finishNesting(container.Befores, &container.BeforesQueue)
}

func (i *beforeSuite) addNestedParam(params ...allure.Parameter) {
	container := i.t.GetContainer()
	addNestingParameter(params, &container.BeforesQueue)
}

func (i *beforeSuite) addAttachment(attachment *allure.Attachment) {
	addAttachmentToContainer(i, attachment)
}

func (i *beforeSuite) addNestedAttachment(attachment *allure.Attachment) {
	container := i.t.GetContainer()
	addNestingAttachment(attachment, &container.BeforesQueue)
}

/*
	afterSuite state
*/

type afterSuite struct {
	t    *T
	name string
}

func (i *afterSuite) addStep(step *allure.Step) {
	container := i.t.GetContainer()
	container.Afters = addStep(step, container.Afters, &container.AftersQueue)
}

func (i *afterSuite) addNested(step *allure.Step) {
	i.addStep(step)
	container := i.t.GetContainer()
	addNested(container.Afters, &container.AftersQueue)
}

func (i *afterSuite) finishNesting() {
	container := i.t.GetContainer()
	finishNesting(container.Afters, &container.AftersQueue)
}

func (i *afterSuite) addNestedParam(params ...allure.Parameter) {
	container := i.t.GetContainer()
	addNestingParameter(params, &container.AftersQueue)
}

func (i *afterSuite) addAttachment(attachment *allure.Attachment) {
	addAttachmentToContainer(i, attachment)
}

func (i *afterSuite) addNestedAttachment(attachment *allure.Attachment) {
	container := i.t.GetContainer()
	addNestingAttachment(attachment, &container.AftersQueue)
}

func addStep(step *allure.Step, steps []*allure.Step, queue *allure.NestingQueue) []*allure.Step {
	last := queue.Last()
	if last != nil && step.Parent == "" {
		step.Parent = last.GetUUID()
	}
	return append(steps, step)
}

func addNested(steps []*allure.Step, queue *allure.NestingQueue) {
	queue.Push(steps[len(steps)-1])
}

func finishNesting(steps []*allure.Step, queue *allure.NestingQueue) {
	if queue.Last() == nil {
		return
	}
	step := queue.Pop()
	step.Status = checkAllInnerStatuses(*step, steps)
	step.Stop = allure.GetNow()
}

func addNestingParameter(parameters []allure.Parameter, queue *allure.NestingQueue) {
	step := queue.Last()
	if step == nil {
		return
	}
	step.AddParameters(parameters...)
}

func addNestingAttachment(attachment *allure.Attachment, queue *allure.NestingQueue) {
	step := queue.Last()
	if step == nil {
		return
	}
	step.Attachment(attachment)
}

func addAttachmentToContainer(ctx stepState, attachment *allure.Attachment) {
	step := allure.NewSimpleStep(attachment.Name)
	step.Attachment(attachment)
	ctx.addStep(step)
}

func checkAllInnerStatuses(parent allure.Step, steps []*allure.Step) allure.Status {
	if parent.Status == allure.Skipped {
		return parent.Status
	}
	for _, step := range steps {
		if step.Parent == parent.GetUUID() && step.Status != allure.Passed {
			return step.Status
		}
	}
	return parent.Status
}
