package cache

import (
	"strings"
	"testing"
)

func TestLruCache_Set(t *testing.T) {

	cache:=LruCache{maxLength:3}
	cache.Set("1","test1")
	value:=cache.Get("1").(string)
	if strings.Compare(value,"test1")!=0{
		t.Fail()
	}

}

func TestLruCache_Set2(t *testing.T) {
	cache:=LruCache{maxLength:3}
	cache.Set("1","test1")
	cache.Set("2","test2")
	cache.Set("3","test3")
	cache.Set("4","test4")
	value:=cache.Get("1")
	if value!=nil {
		t.Fail()
	}
}


func BenchmarkLruCache_Set(b *testing.B) {
	b.ResetTimer()
	cache:=LruCache{maxLength:100}
	for i:=0; i<b.N;i++  {
		cache.Set("1",b.N)
	}
}