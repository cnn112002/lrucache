package cache

import (
	"container/list"
	"sync"
)
type LruCache struct {

	maxLength int

	elements *list.List
	index map[string]*list.Element
    mutex sync.RWMutex


}


type CacheElement struct {
	key string
	value interface{}
    count int

}

func (c *LruCache) Set(key string,value interface{})  {

	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.elements==nil {
		c.elements=list.New()
		c.index=make(map[string]*list.Element)
	}

	if elem, ok:=c.index[key];ok{

		cacheEle:=elem.Value.(CacheElement)
		cacheEle.value=value
		cacheEle.count++
		c.elements.MoveToFront(elem)
	}else {
		if c.elements.Len()+1>c.maxLength{
			c.removeOldest()
		}
		cacheEle:=CacheElement{key:key,value:value,count:1}
		listEle:=c.elements.PushFront(cacheEle)
		c.index[key]=listEle

	}


}

func (c *LruCache) removeOldest()  {
	tail:=c.elements.Back()
	c.elements.Remove(tail)
	cacheEle:=tail.Value.(CacheElement)
	delete(c.index,cacheEle.key)

}

func (c *LruCache) Get(key string)  interface{}{

	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if elem, ok:=c.index[key];ok{

		cacheEle:=elem.Value.(CacheElement)
		cacheEle.count++
		c.elements.MoveToFront(elem)
		return cacheEle.value
	}
	return nil

}
