package pipline

import "sync"

type stepFn func(chan any) chan any

type Pipeline struct {
	Steps []func(chan any) chan any
	In    chan any
	Out   chan any
	wgIn  sync.WaitGroup
}

// Создает новый pipline
func NewPipline() *Pipeline {
	return &Pipeline{
		In: make(chan any, 3),
	}
}

// Добавляет заказ в pipline
func (p *Pipeline) PushOrder(do DeviceOrder) {
	p.wgIn.Add(1)
	go func() {
		defer p.wgIn.Done()
		p.In <- do
	}()
}

// Добавляет этап в pipline
func (p *Pipeline) AddStep(step stepFn) {
	p.Steps = append(p.Steps, step)
}

// Запускает pipline
func (p *Pipeline) Run() chan any {
	out := p.In
	for _, step := range p.Steps {
		out = step(out)
	}

	return out
}

// Закрывает pipline
func (p *Pipeline) Close() {
	p.wgIn.Wait()
	close(p.In)
}
