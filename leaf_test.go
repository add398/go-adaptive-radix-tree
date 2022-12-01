/**
 * @Author: Log1c
 * @Description:
 * @File:  leaf_test
 * @Version: 1.0.0
 * @Date: 2022/12/1 17:16
 */

package art

import (
	"fmt"
	"testing"
)

func Test_Tree_ForEachByLeaf(t *testing.T) {
	tree := New()
	terms := []string{"a", "A", "aa", "a"}

	for _, term := range terms {
		tree.Insert(Key(term), term)
	}
	tree.ForEachByLeaf()
	fmt.Println(tree)
}
