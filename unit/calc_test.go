package main

import "testing"

//func TestAdd(t *testing.T) {
//	res := Add(1, -1)
//	if res != 0 {
//		t.Errorf("测试失败")
//		return
//	}
//	t.Logf("测试通过")
//}

// 子测试
//
//	func TestAdd(t *testing.T) {
//		t.Run("add1", func(t *testing.T) {
//			if ans := Add(1, 2); ans != 3 {
//				t.Fatalf("1+2 expected be 3,but %d got", ans)
//			}
//		})
//		t.Run("add2", func(t *testing.T) {
//			if ans := Add(-10, -20); ans != -30 {
//				t.Fatalf("-10 + -20 expected be -30, but %d got", ans)
//			}
//		})
//	}
func TestAdd(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"a1", 2, 3, 5},
		{"a2", 2, -3, -1},
		{"a3", 2, 0, 2},
		{"a4", 2, 4, 6},
	}
	for _, s := range cases {
		t.Run(s.Name, func(t *testing.T) {
			if Add(s.A, s.B) != s.Expected {
				t.Fatalf("Add(%d, %d) != %d", s.A, s.B, s.Expected)

			}
		})

	}
}
