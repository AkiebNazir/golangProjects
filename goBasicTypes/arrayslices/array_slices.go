package arrayslices

type ArrayNums struct {
	Arr [20]int
}

type SliceNums struct {
	Arr []int
}

func (a *ArrayNums) ArrayMonupulation() {
	a.Arr[2] = 20
	a.Arr[4] = 40
}

func (a *SliceNums) SliceEdit() {
	a.Arr[4] = 40
	a.Arr[2] = 20
}
