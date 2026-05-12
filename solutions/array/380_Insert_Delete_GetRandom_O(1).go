package array

import "math/rand"

type RandomizedSet struct {
	values []int
	index  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values: []int{},
		index:  make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exist := this.index[val]
	if exist {
		return false
	}
	this.values = append(this.values, val)
	lastIndex := len(this.values) - 1
	this.index[val] = lastIndex
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	valIndex, exist := this.index[val]
	if !exist {
		return false
	}
	lastIndex := len(this.values) - 1

	lastValue := this.values[lastIndex]
	this.values[valIndex] = lastValue
	this.index[lastValue] = valIndex

	this.values = this.values[:lastIndex]
	delete(this.index, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(this.values))
	return this.values[random]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
