package leetcode

//
//import (
//	"reflect"
//	"testing"
//)
//
//func TestConstructor(t *testing.T) {
//	tests := []struct {
//		name string
//		want Trie
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Constructor() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestTrie_Insert(t *testing.T) {
//
//	trie := Constructor()
//	trie.Insert("app")
//	trie.Insert("apple")
//	trie.Insert("beer")
//	trie.Insert("add")
//
//	println(trie.StartsWith("ad"))
//	println(trie.StartsWith("app"))
//
//}
//
//func TestTrie_Search(t *testing.T) {
//	type fields struct {
//		root *Node
//	}
//	type args struct {
//		word string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			this := &Trie{
//				root: tt.fields.root,
//			}
//			if got := this.Search(tt.args.word); got != tt.want {
//				t.Errorf("Search() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestTrie_StartsWith(t *testing.T) {
//	type fields struct {
//		root *Node
//	}
//	type args struct {
//		prefix string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			this := &Trie{
//				root: tt.fields.root,
//			}
//			if got := this.StartsWith(tt.args.prefix); got != tt.want {
//				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
