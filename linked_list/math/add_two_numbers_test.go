package linkedlist_test

import (
	linkedlist "go-programming-techniques/linkedlist/math"
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		name string
		l1   *linkedlist.ListNode
		l2   *linkedlist.ListNode
		want *linkedlist.ListNode
	}{
		// l1 = [2,4,3]
		// l2 = [5,6,4]
		// want = [7,0,8]
		{
			name: "test case 1",
			l1: &linkedlist.ListNode{Val: 2,
				Next: &linkedlist.ListNode{Val: 4,
					Next: &linkedlist.ListNode{Val: 3, Next: nil},
				},
			},
			l2: &linkedlist.ListNode{Val: 5,
				Next: &linkedlist.ListNode{Val: 6,
					Next: &linkedlist.ListNode{Val: 4, Next: nil},
				},
			},
			want: &linkedlist.ListNode{Val: 7,
				Next: &linkedlist.ListNode{Val: 0,
					Next: &linkedlist.ListNode{Val: 8, Next: nil},
				},
			},
		},
		// l1 = [0]
		// l2 = [0]
		// want = [0]
		{
			name: "test case 2",
			l1:   &linkedlist.ListNode{Val: 0, Next: nil},
			l2:   &linkedlist.ListNode{Val: 0, Next: nil},
			want: &linkedlist.ListNode{Val: 0, Next: nil},
		},
		// l1 = [9,9,9,9,9,9,9]
		// l2 = [9,9,9,9]
		// want = [8,9,9,9,0,0,0,1]
		{
			name: "test case 3",
			l1: &linkedlist.ListNode{Val: 9,
				Next: &linkedlist.ListNode{Val: 9,
					Next: &linkedlist.ListNode{Val: 9,
						Next: &linkedlist.ListNode{Val: 9,
							Next: &linkedlist.ListNode{Val: 9,
								Next: &linkedlist.ListNode{Val: 9,
									Next: &linkedlist.ListNode{Val: 9, Next: nil},
								},
							},
						},
					},
				},
			},
			l2: &linkedlist.ListNode{Val: 9,
				Next: &linkedlist.ListNode{Val: 9,
					Next: &linkedlist.ListNode{Val: 9,
						Next: &linkedlist.ListNode{Val: 9, Next: nil},
					},
				},
			},
			want: &linkedlist.ListNode{Val: 8,
				Next: &linkedlist.ListNode{Val: 9,
					Next: &linkedlist.ListNode{Val: 9,
						Next: &linkedlist.ListNode{Val: 9,
							Next: &linkedlist.ListNode{Val: 0,
								Next: &linkedlist.ListNode{Val: 0,
									Next: &linkedlist.ListNode{Val: 0,
										Next: &linkedlist.ListNode{Val: 1, Next: nil},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := linkedlist.AddTwoNumbers(tt.l1, tt.l2)
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
