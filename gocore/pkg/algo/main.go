package main

import "fmt"

/*
https://www.hungarianalgorithm.com/solve.php?c=725-660-188-806-354-78-418-501-867-675--135-842-429-853-200-923-904-274-457-665--35-134-347-102-871-863-277-356-827-568--710-318-311-427-51-564-93-582-574-825--562-529-211-486-259-983-880-822-154-788--179-987-458-623-244-101-250-464-66-825--106-332-77-785-354-499-884-260-60-637--673-256-454-32-530-219-431-168-527-837--257-220-388-662-596-172-910-172-989-108--927-406-184-346-591-395-232-457-50-706
*/
type Bipartite struct {
	left  []int
	right []int
	p     [][]int
	q     [][]int
}

func (b *Bipartite) kuhnstep(v int, mem map[int]bool, mt *Matching) bool {
	if mem[v] {
		return false
	}

	mem[v] = true

	for _, u := range b.p[v] {
		if mt.ji[u] == -1 || b.kuhnstep(mt.ji[u], mem, mt) {
			mt.ji[u] = v
			mt.ij[v] = u
			return true
		}
	}

	return false
}

func (b *Bipartite) dfs(v int, meml map[int]bool, memr map[int]bool, ij []int) [][]int {
	meml[v] = true
	path := [][]int{[]int{v}, []int{}}
	for _, u := range b.p[v] {
		if memr[u] {
			continue
		}
		memr[u] = true
		path[1] = append(path[1], u)
		for _, vv := range b.q[u] {
			if ij[vv] != u {
				meml[vv] = true
				continue
			}
			if !meml[vv] {
				path2 := b.dfs(vv, meml, memr, ij)
				path[0] = append(path[0], path2[0]...)
				path[1] = append(path[1], path2[1]...)
				break
			}
		}
	}

	return path
}

type Matching struct {
	ij []int
	ji []int
}

func (m *Matching) Size() int {
	res := 0
	for _, el := range m.ij {
		if el != -1 {
			res += 1
		}
	}
	return res
}

func (m *Matching) Print() {
	for i, j := range m.ij {
		if j != -1 {
			fmt.Printf("%d -> %d\n", i, j)
		}
	}
}

func Algorithm(c [][]int) []int {
	n := len(c)

	mt := Matching{ij: make([]int, n), ji: make([]int, n)}
	for i := range mt.ij {
		mt.ij[i] = -1
		mt.ji[i] = -1
	}

	for i, row := range c {
		delta := -1
		for _, el := range row {
			if delta == -1 || delta > el {
				delta = el
			}
		}

		for j := range row {
			c[i][j] -= delta
		}
	}

	for j := 0; j < n; j++ {
		delta := -1
		for i := 0; i < n; i++ {
			if delta == -1 || delta > c[i][j] {
				delta = c[i][j]
			}
		}

		for i := 0; i < n; i++ {
			c[i][j] -= delta
		}
	}

	for mt.Size() < n {
		// build Bipartite
		b := Bipartite{left: make([]int, n), right: make([]int, n), p: make([][]int, n), q: make([][]int, n)}
		for i := range b.left {
			b.left[i] = -1
		}
		for i := range b.right {
			b.right[i] = -1
		}

		for i, row := range c {
			for j, el := range row {
				if el == 0 {
					b.left[i] = i
					b.p[i] = append(b.p[i], j)
					b.right[j] = j
					b.q[j] = append(b.q[j], i)
				}
			}
		}

		for i, el := range b.left {
			if el == -1 {
				if len(b.left) == 1 {
					b.left = []int{}
				} else {
					b.left = append(b.left[:i], b.left[i+1:]...)
				}
			}
		}
		for j, el := range b.right {
			if el == -1 {
				if len(b.right) == 1 {
					b.right = []int{}
				} else {
					b.right = append(b.right[:j], b.right[j+1:]...)
				}
				//				b.q[j] = []int{}
			}
		}

		mt = Matching{ij: make([]int, n), ji: make([]int, n)}
		for i := range mt.ij {
			mt.ij[i] = -1
			mt.ji[i] = -1
		}
		seen := make(map[int]bool)
		for i := range b.left {
			mem := make(map[int]bool)
			b.kuhnstep(i, mem, &mt)
			for k, val := range mem {
				if !seen[k] {
					seen[k] = val
				}
			}
		}

		if mt.Size() == n {
			continue
		}

		// run dfs from first free vertex of a left part
		// check what we mark
		// update matrix based on marted left and right marked vertex ( include rows of a left part, exclude columns of a right part)

		// get free vertex from L+ ( marked by Kuhn, but not in matching )

		meml := make(map[int]bool)
		memr := make(map[int]bool)
		path := make([][]int, 2)
		for i, j := range mt.ij {
			if j == -1 {
				p := b.dfs(i, meml, memr, mt.ij)
				path[0] = append(path[0], p[0]...)
				path[1] = append(path[1], p[1]...)
			}
		}

		// meml := make(map[int]bool)
		// memr := make(map[int]bool)

		// path := b.dfs(freev, meml, memr)

		pathl := make(map[int]bool)
		pathr := make(map[int]bool)

		for _, el := range path[0] {
			pathl[el] = true
		}
		for _, el := range path[1] {
			pathr[el] = true
		}

		delta := -1
		for _, i := range path[0] {
			for j, el := range c[i] {
				if pathr[j] {
					continue
				}
				if delta == -1 || delta > el {
					delta = el
				}
			}
		}

		for i, row := range c {
			for j := range row {
				if pathl[i] {
					c[i][j] -= delta
				}
				if pathr[j] {
					c[i][j] += delta
				}

			}
		}
	}

	//	mt.Print()

	return mt.ij
}
