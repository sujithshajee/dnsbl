package worker

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gammazero/workerpool"
	"go.uber.org/zap"

	"github.com/sujithshajee/dnsbl/app/ent"
	"github.com/sujithshajee/dnsbl/app/ent/task"
)

const Workers = 100

type Worker interface {
	Start() error
	Stop()
	Register(task.Type, Job)
}
type worker struct {
	client *ent.Client
	pool   *workerpool.WorkerPool
	jobs   map[task.Type]Job
	log    *zap.SugaredLogger
}

type Job interface {
	Execute(ctx context.Context, state string) error
}

type State interface {
	Get() map[string]interface{}
}

func New(cl *ent.Client, l *zap.SugaredLogger) Worker {
	pool := workerpool.New(Workers)

	return &worker{
		jobs:   make(map[task.Type]Job),
		client: cl,
		pool:   pool,
		log:    l,
	}
}

func (w *worker) Start() error {
	for {
		w.log.Info("checking for jobs")
		ctx := context.Background()
		op, err := w.client.Task.Query().
			Where(
				task.StatusEQ(task.StatusWAITING),
			).First(ctx)
		if ent.IsNotFound(err) {
			w.log.Info("none found, sleeping")
			time.Sleep(500 * time.Millisecond)
			continue
		}
		if err != nil {
			w.log.Error(err.Error())
			return fmt.Errorf("finding tasks: %w", err)
		}

		if j, ok := w.jobs[op.Type]; ok {
			w.log.Infof("job:%s found, submitting", op.Type)

			pj, err := w.prepJob(ctx, j, op)
			if err != nil {
				return fmt.Errorf("preparing job: %w", err)
			}
			w.pool.Submit(pj)
		}
	}
}

func (w *worker) Stop() {
	w.pool.Stop()
}

func (w *worker) Register(opType task.Type, j Job) {
	w.jobs[opType] = j
}

func (w *worker) prepJob(ctx context.Context, j Job, op *ent.Task) (func(), error) {
	w.log.Infof("job:%s progressed", op.Type)
	deviation := time.Duration((rand.Intn(1000) + 300) % 1500)
	time.Sleep(deviation * time.Millisecond)

	_, err := op.Update().SetStatus(task.StatusIN_PROGRESS).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("progressing task: %w", err)
	}

	return func() {
		w.log.Infof("job:%s", op.Type)
		ctx := context.Background()
		err := j.Execute(ctx, op.Ipaddress)

		if err != nil {
			w.log.Errorf("job:%s state:%s failed: %s", op.Type, op.Ipaddress, err)
			_, _ = op.Update(). // TODO: we can't deal with the error
						SetStatus(task.StatusERROR).
						SetError(err.Error()).
						Save(ctx)
		}

		w.log.Info("job worked")
		_, _ = op.Update().
			SetStatus(task.StatusDONE).Save(ctx)
	}, nil
}
