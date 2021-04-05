package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Make chain of stages
	for _, stage := range stages {
		in = stage(in)
	}

	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case x, ok := <-in:
				if !ok {
					return
				}
				out <- x
			case <-done:
				return
			}
		}
	}()

	return out
}
