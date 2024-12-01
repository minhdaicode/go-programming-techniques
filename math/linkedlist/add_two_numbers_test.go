package linkedlist_test

import (
	"go-programming-techniques/math/linkedlist"
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
		{
			name: "test case 2",
			l1:   &linkedlist.ListNode{Val: 0, Next: nil},
			l2:   &linkedlist.ListNode{Val: 0, Next: nil},
			want: &linkedlist.ListNode{Val: 0, Next: nil},
		},
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
