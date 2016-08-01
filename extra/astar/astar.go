package astar

type Path struct {
	Next int
	Dist uint
}

const MaxDistance = ^uint(0)

//这是个启发式算法，有猜测成分，不一定能得到最优解
func AStar(roads [][]Path, start int, end int,
	evaluate func(int) uint) []int {
	var size = len(roads)
	if start < 0 || end < 0 || start >= size || end >= size {
		return nil
	}
	if start == end {
		return []int{start}
	}

	const FAKE = -1
	var list = make([]memo, size)
	for i := 0; i < size; i++ {
		list[i].link = FAKE
	}
	var trace = func() []int {
		var path []int
		for idx := end; idx != start; idx = list[idx].link {
			path = append(path, idx)
		}
		path = append(path, start)
		reverse(path)
		return path
	}

	list[start].index, list[start].link = start, start
	list[start].dist = 0
	var hp = heap{core: make([]*memo, 0, size)}

	hp.Push(&list[start])
	for !hp.IsEmpty() {
		var curr = hp.Pop()
		if curr.index == end {
			return trace()
		}
		var index = curr.index
		curr.index = FAKE //入围
		for _, path := range roads[index] {
			var peer = &list[path.Next]
			if peer.link == FAKE { //未涉及点
				peer.index, peer.link = path.Next, index
				peer.dist = curr.dist + path.Dist
				//dist记录了起点到当前点的距离，evaluate评估当前点到终点的距离
				peer.weight = peer.dist + evaluate(peer.index)
				hp.Push(peer)
			} else if peer.index != FAKE { //外围点
				var distance = curr.dist + path.Dist
				if distance < peer.dist {
					hp.ShiftUp(peer, distance)
					peer.link, peer.dist = index, distance
				}
			}
		}
	}
	return nil
}
func reverse(list []int) {
	for left, right := 0, len(list)-1; left < right; {
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
}
