package task

import (
	"context"
	"github.com/pkg/errors"
	"sync"
	"sync/atomic"
)

type step struct {
	syncJobs  []SyncJob
	asyncJobs []AsyncJob
}

func (s *step) Run(ctx context.Context) error {
	err := s.runSyncJob(ctx)
	if err != nil {
		return err
	}

	_, err = s.runAsyncJobs(ctx)
	return err
}

func (s *step) runSyncJob(ctx context.Context) error {
	for _, syncJob := range s.syncJobs {
		err := syncJob(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *step) runAsyncJobs(ctx context.Context) ([]error, error) {
	count := len(s.asyncJobs)
	if count == 0 {
		return nil, nil
	}

	var wg sync.WaitGroup
	wg.Add(count)
	errs := make([]error, count)

	for i, asyncJob := range s.asyncJobs {
		go func(index int) {
			errs[index] = asyncJob(ctx)
			wg.Done()
		}(i)
	}

	wg.Wait()
	var err error
	for _, v := range errs {
		if err == nil {
			err = v
			continue
		}
		err = errors.Wrapf(err, v.Error())
	}
	return errs, err
}

func (s *step) asyncJobsRunErrorReturn(ctx context.Context) ([]error, error) {
	count := len(s.asyncJobs)
	if count == 0 {
		return nil, nil
	}

	errs := make([]error, count)
	errExit := make(chan error)

	var counter int32 = 0
	var all = int32(count)
	finish := make(chan struct{})

	var exitFlag int32
	for i, asyncJob := range s.asyncJobs {
		go func(index int) {
			err := asyncJob(ctx)
			errs[i] = err
			if err != nil {
				if atomic.AddInt32(&exitFlag, 1) != 1 {
					return
				}

				errExit <- err
				return
			}
			if atomic.AddInt32(&counter, 1) == all {
				close(finish)
			}
		}(i)
	}
	select {
	case err := <-errExit:
		close(errExit)
		return errs, err
	case <-finish:
		return errs, nil
	}
}
