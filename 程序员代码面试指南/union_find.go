package main

type UnionSet struct {
	rankMap    map[Element]int
	fatherMap  map[Element]Element
	elementMap map[interface{}]Element
}

type Element struct {
	value interface{}
}

func NewUnionSet(list []Element) *UnionSet {
	res := &UnionSet{
		rankMap:    make(map[Element]int, 0),
		fatherMap:  make(map[Element]Element, 0),
		elementMap: make(map[interface{}]Element, 0),
	}
	for _, v := range list {
		res.rankMap[v] = 1
		res.fatherMap[v] = v
		res.elementMap[v.value] = v
	}
	return res
}

func (u *UnionSet) findHead(e Element) Element {
	paths := make([]Element, 0)
	for e != u.fatherMap[e] {
		paths = append(paths, e)
		e = u.fatherMap[e]
	}
	for _, p := range paths {
		u.fatherMap[p] = e
	}
	return e
}

func (u *UnionSet) isSameSet(a, b interface{}) bool {
	_, aBool := u.elementMap[a]
	_, bBool := u.elementMap[b]
	if aBool && bBool {
		return u.findHead(u.elementMap[a]) == u.findHead(u.elementMap[b])
	}
	return false
}

func (u *UnionSet) union(a, b interface{}) {
	av, aBool := u.elementMap[a]
	bv, bBool := u.elementMap[b]
	if aBool && bBool {
		af := u.fatherMap[av]
		bf := u.fatherMap[bv]
		if af != bf {
			afrank := u.rankMap[af]
			bfrank := u.rankMap[bf]
			var big, small Element
			if afrank > bfrank {
				big = af
				small = bf
			} else {
				big = bf
				small = af
			}
			u.rankMap[big] = afrank + bfrank
			u.fatherMap[small] = big
			delete(u.rankMap, small)
		}
	}

}
