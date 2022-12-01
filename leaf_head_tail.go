/**
 * @Author: Log1c
 * @Description:
 * @File:  leaf_head_tail
 * @Version: 1.0.0
 * @Date: 2022/12/1 16:52
 */

package art



func (this *tree) AddLast(x *leaf) {
	x.pre = this.tail.pre
	x.next = this.tail
	this.tail.pre.next = x
	this.tail.pre = x
	this.size++
}

func (this *tree) Remove(t Tree, x *leaf) Key {
	key := x.key
	x.pre.next = x.next
	x.next.pre = x.pre
	x.pre = nil
	x.next = nil

	this.size--

	t.Delete(key)

	return key
}

func (this *tree) RemoveFirst(t Tree) *leaf {
	if this.head.next == this.tail {
		return nil
	}
	first := this.head.next
	this.Remove(t, first)
	return first
}

