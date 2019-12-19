package task

import "context"

type task struct {
	ss []*step
}

func NewTask(taskCount ...int) *task {
	t := new(task)
	count := 0
	if len(taskCount) > 0 {
		count = taskCount[0]
	}
	t.ss = make([]*step, 0, count)
	return t
}

func (t *task) Run(ctx context.Context) error {
	return t.run(ctx)
}

func (t *task) run(ctx context.Context) error {
	for _, step := range t.ss {
		err := step.Run(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *task) AddStep(s *step) *task {
	t.ss = append(t.ss, s)
	return t
}

func (t *task) AddSteps(ss []*step) *task {
	t.ss = append(t.ss, ss...)
	return t
}
