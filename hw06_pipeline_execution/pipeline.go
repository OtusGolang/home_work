package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		ch := make(chan interface{})
		close(ch)
		return ch
	}

	job := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan interface{} {
		jobStream := make(chan interface{})

		go func() {
			defer close(jobStream)

			for {
				select {
				case <-done:
					return
				default:
				}

				select {
				case <-done:
					return
				case val, ok := <-valueStream:
					// Не обрабатываем пустой канал
					if !ok {
						return
					}

					select {
					case <-done:
						return
					case jobStream <- val:
					}
				}
			}
		}()
		return jobStream
	}

	for _, stage := range stages {
		in = stage(job(done, in))
	}

	return in
}
