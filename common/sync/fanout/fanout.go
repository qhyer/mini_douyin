package fanout

import (
	"context"
	"douyin/common/ecode"
	"github.com/go-kratos/kratos/v2/log"
	"runtime"
	"sync"
)

type options struct {
	worker int
	buffer int
}

type Option func(*options)

type item struct {
	f   func(c context.Context)
	ctx context.Context
}

type Fanout struct {
	ch      chan item
	options *options
	waiter  sync.WaitGroup
	ctx     context.Context
	cancel  func()
}

func Worker(num int) Option {
	if num <= 0 {
		panic("worker num must grater than 0")
	}
	return func(o *options) {
		o.worker = num
	}
}

func Buffer(num int) Option {
	if num <= 0 {
		panic("buffer num must grater than 0")
	}
	return func(o *options) {
		o.buffer = num
	}
}

func New(opts ...Option) *Fanout {
	opt := &options{
		worker: 1,
		buffer: 1024,
	}
	for _, op := range opts {
		op(opt)
	}
	f := &Fanout{
		ch:      make(chan item, opt.buffer),
		options: opt,
	}
	f.ctx, f.cancel = context.WithCancel(context.Background())
	f.waiter.Add(opt.worker)
	for i := 0; i < opt.worker; i++ {
		go f.proc()
	}
	return f
}

func (f *Fanout) proc() {
	defer f.waiter.Done()
	for {
		select {
		case t := <-f.ch:
			wrapFunc(t.f)(t.ctx)
		case <-f.ctx.Done():
			return
		}
	}
}

func wrapFunc(f func(c context.Context)) (res func(context.Context)) {
	res = func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				buf := make([]byte, 64*1024)
				buf = buf[:runtime.Stack(buf, false)]
				log.Errorf("Panic in fanout proc, err: %s, stack: %s", r, buf)
			}
		}()
		f(ctx)
	}
	return
}

func (f *Fanout) Do(ctx context.Context, fun func(ctx context.Context)) (err error) {
	if fun == nil || f.ctx.Err() != nil {
		return f.ctx.Err()
	}
	select {
	case f.ch <- item{f: fun, ctx: ctx}:
	default:
		err = ecode.ChanFullErr
	}
	return
}

func (f *Fanout) Close() error {
	if err := f.ctx.Err(); err != nil {
		return err
	}
	f.cancel()
	f.waiter.Wait()
	return nil
}
