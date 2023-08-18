// Author: yangzq80@gmail.com
// Date: 2023/6/14
package flow

import (
	"context"
	"io"
	"os"
)

type Executor interface {
	SetStdout(out io.Writer)
	SetStderr(out io.Writer)
	Kill(sig os.Signal) error
	Run() error
}

type Creator func(ctx context.Context, step *Step) (Executor, error)

var executors = make(map[string]Creator)

func Register(name string, register Creator) {
	executors[name] = register
}

//func CreateExecutor(ctx context.Context, step *Step) (Executor, error) {
//	f, ok := executors[step.ExecutorConfig.Type]
//	if ok {
//		return f(ctx, step)
//	}
//	return nil, fmt.Errorf("invalid executor: %s", step.ExecutorConfig)
//}

func ExecutorIsValid(name string) bool {
	_, ok := executors[name]
	return ok
}
