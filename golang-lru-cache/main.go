package main 
import (
	"fmt"
)

type Node struct{
	Val string 
	Left *Node
	Right *Node
}

type Hash map[string] *Node
type Cache struct {
	Queue *Queue 
	HashMap Hash
}


type Queue struct{
	Head *Node 
	Tail *Node
	Current uint8
	Capacity uint8

}

func  CreateNewCache(capacity uint8) (*Cache){
   queue:=CreateNewQueue(capacity)
   return &Cache{
	Queue: queue,
	HashMap: Hash{},
   }
}


func CreateNewQueue(capacity uint8) (*Queue){
	head:=Node{}
	tail:=Node{}
	//DUMMY NODE
	head.Right=&tail;
	tail.Left=&head
	return &Queue{
		Head:&head,
		Tail:&tail,
		Capacity: capacity,
		Current: 0,
	}
}


func (cache *Cache) CheckAndAdd(val string){
	//check capacity and remove 
    hashmap:=cache.HashMap;
	node,ok:=hashmap[val];
	if ok{
		cache.RemoveFromCurrentPos(node)
		cache.InsertAtHead(node)
	}else{
		cache.CheckCapacityAndRemove()
		newNode:=&Node{
			Val:val,
		}
		cache.InsertAtHead(newNode)
		cache.HashMap[val]=newNode;
		cache.Queue.Current++
	}
}

func (cache *Cache) InsertAtHead(node *Node){
	queue:=cache.Queue
	//first insertion
	if queue.Head.Right == queue.Tail{
		queue.Head.Right=node
		node.Right=queue.Tail;
		queue.Tail.Left=node
		node.Left=queue.Head;
	}else{
		currentHead:=queue.Head.Right;
		queue.Head.Right=node;
		node.Left=queue.Head
		node.Right=currentHead;
		currentHead.Left=node;
	}
	
}

func (cache *Cache) RemoveFromCurrentPos(node *Node){
	left:=node.Left
	right:=node.Right

	//make mapping
	left.Right=right;
	right.Left=left

	node.Left = nil
	node.Right = nil
}

func (cache *Cache) CheckCapacityAndRemove()  {
	queue:=cache.Queue;
	if queue.Capacity==queue.Current{
		val:=queue.Tail.Left.Val;
		secondLast:=queue.Tail.Left.Left;
		secondLast.Right=queue.Tail
		queue.Tail.Left=secondLast;
		queue.Current--
		delete(cache.HashMap, val)
		
		
	}
}

func (cache *Cache) GetAllValue(){
	tmp:=cache.Queue.Head.Right;
	for (tmp.Right!=nil){
		fmt.Println(tmp.Val)
		tmp=tmp.Right
	}

}

func (cache *Cache) MostRecentlyAdded() (string){
  return cache.Queue.Head.Right.Val;
}

func main(){
   cache:=CreateNewCache(3)
   cache.CheckAndAdd("100")
   cache.CheckAndAdd("200")
   cache.CheckAndAdd("300")
	cache.CheckAndAdd("400")
   cache.CheckAndAdd("500")
   cache.CheckAndAdd("600")
   	cache.CheckAndAdd("400")
   cache.GetAllValue()
//    fmt.Println("Most Recently Added",recent);
}