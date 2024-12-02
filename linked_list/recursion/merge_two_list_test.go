package linkedlist_test

import (
	linkedlist "go-programming-techniques/linked_list/math"
	list "go-programming-techniques/linked_list/recursion"
	"reflect"
	"testing"
)

func TestMergeListNode(t *testing.T) {
	var tests = []struct {
		name string
		l1   *linkedlist.ListNode
		l2   *linkedlist.ListNode
		want *linkedlist.ListNode
	}{
		// l1 = [1,2,4]
		// l2 = [1,3,4]
		// want = [1,1,2,3,4,4]
		{
			name: "test case 1",
			l1: &linkedlist.ListNode{Val: 1,
				Next: &linkedlist.ListNode{Val: 2,
					Next: &linkedlist.ListNode{Val: 4, Next: nil},
				},
			},
			l2: &linkedlist.ListNode{Val: 1,
				Next: &linkedlist.ListNode{Val: 3,
					Next: &linkedlist.ListNode{Val: 4, Next: nil},
				},
			},
			want: &linkedlist.ListNode{Val: 1,
				Next: &linkedlist.ListNode{Val: 1,
					Next: &linkedlist.ListNode{Val: 2,
						Next: &linkedlist.ListNode{Val: 3,
							Next: &linkedlist.ListNode{Val: 4,
								Next: &linkedlist.ListNode{Val: 4, Next: nil},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := list.MergeTwoLists(tt.l1, tt.l2)
			cur := tt.want
			for ans != nil {
				if !reflect.DeepEqual(ans.Val, cur.Val) {
					t.Errorf("got %d ,want %d", ans.Val, cur.Val)
				}
				ans = ans.Next
				cur = cur.Next
			}
		})
	}
}
