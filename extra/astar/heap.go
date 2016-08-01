package astar

type memo struct {
	index  int
	link   int
	dist   uint
	weight uint
	off    int
}

type heap struct {
	core []*memo
}

func (hp *heap) Size() int {
	return len(hp.core)
}
func (hp *heap) IsEmpty() bool {
	return len(hp.core) == 0
}

func (hp *heap) ShiftUp(unit *memo, dist uint) {
	unit.weight = dist
	hp.shiftUp(unit.off)
}

func (hp *heap) shiftUp(pos int) {
	var unit = hp.core[pos]
	for pos > 0 {
		var parent = (pos - 1) / 2
		if hp.core[parent].weight <= unit.weight {
			break
		}
		hp.core[pos] = hp.core[parent]
		hp.core[pos].off = pos
		pos = parent
	}
	hp.core[pos], unit.off = unit, pos
}

func (hp *heap) Push(unit *memo) {
	var place = len(hp.core)
	hp.core = append(hp.core, unit)
	unit.off = place
	hp.shiftUp(place)
}

func (hp *heap) Pop() *memo {
	var size = hp.Size()
	if size == 0 {
		return nil
	}
	var result = hp.core[0]
	if size == 1 {
		hp.core = hp.core[:0]
		return result
	}

	var unit = hp.core[size-1]
	hp.core = hp.core[:size-1]

	var pos, kid, last = 0, 1, size - 2
	for kid < last {
		if hp.core[kid+1].weight < hp.core[kid].weight {
			kid++
		}
		if unit.weight <= hp.core[kid].weight {
			break
		}
		hp.core[pos] = hp.core[kid]
		hp.core[pos].off = pos
		pos, kid = kid, kid*2+1
	}
	if kid == last && unit.weight > hp.core[kid].weight {
		hp.core[pos] = hp.core[kid]
		hp.core[pos].off = pos
		pos = kid
	}
	hp.core[pos], unit.off = unit, pos

	return result
}
