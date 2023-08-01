package beans

import "sort"

//postStart
type postStart interface {
	StartOrder() int
	Start()
}
type postStarts []postStart

func (p postStarts) Len() int {
	return len(p)
}
func (p postStarts) Less(i, j int) bool {
	return p[i].StartOrder() < p[j].StartOrder()
}
func (p postStarts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p postStarts) start() {
	sort.Sort(p)
	for i := range p {
		p[i].Start()
	}
}

//preStop
type preStop interface {
	StopOrder() int
	Stop()
}

type preStops []preStop

func (p preStops) Len() int {
	return len(p)
}
func (p preStops) Less(i, j int) bool {
	return p[i].StopOrder() < p[j].StopOrder()
}
func (p preStops) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p preStops) stop() {
	sort.Sort(p)
	for i := range p {
		p[i].Stop()
	}
}
