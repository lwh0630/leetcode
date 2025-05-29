/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
package lc

import (
	"fmt"
)

type MaxHeap_215 struct {
	elements []int // 存储元素的切片
}

func (h *MaxHeap_215) heapUp(idx int) {
	for {
		parent := (idx - 1) / 2
		if parent < 0 || h.elements[parent] >= h.elements[idx] {
			break
		}
		h.elements[parent], h.elements[idx] = h.elements[idx], h.elements[parent]
		idx = parent
	}
}

func (h *MaxHeap_215) Push(val int) {
	h.elements = append(h.elements, val)
	h.heapUp(len(h.elements) - 1)
}

func (h *MaxHeap_215) Pop() (int, error) {
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	maxVal := h.elements[0]
	lastIdx := len(h.elements) - 1
	h.elements[0], h.elements[lastIdx] = h.elements[lastIdx], h.elements[0]
	h.elements = h.elements[:lastIdx]
	h.heapDown(0)
	return maxVal, nil
}

func (h *MaxHeap_215) heapDown(idx int) {
	n := len(h.elements)
	for {
		left := 2*idx + 1
		right := 2*idx + 2
		largest := idx
		if left < n && h.elements[left] > h.elements[largest] {
			largest = left
		}
		if right < n && h.elements[right] > h.elements[largest] {
			largest = right
		}
		if largest == idx {
			break
		}
		h.elements[idx], h.elements[largest] = h.elements[largest], h.elements[idx]
		idx = largest
	}
}

func (h *MaxHeap_215) init(data []int) {
	h.elements = data
	for i := len(data) / 1; i >= 0; i-- {
		h.heapDown(i)
	}
}

func findKthLargest(nums []int, k int) int {
	MaxHeap_215 := MaxHeap_215{}
	MaxHeap_215.init(nums)
	for i := 0; i < k-1; i++ {
		MaxHeap_215.Pop()
	}
	return MaxHeap_215.elements[0]
}

// @lc code=end
