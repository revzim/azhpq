// Package azhpq --
// revzim <https://github.com/revzim>
// AZHPQ - HEAP PRIORITY QUEUE
package azhpq

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHPQ(t *testing.T) {

	hpq := New()

	qn1 := &QueueNode{
		Value:    "andy",
		Priority: 1,
	}

	qn2 := &QueueNode{
		Value:    "john",
		Priority: 2,
	}

	qn3 := &QueueNode{
		Value:    "jim",
		Priority: 5,
	}

	if !assert.Nil(t, hpq.Poll(), "hpq is empty!") {
		log.Println("FAILED EMPTY TEST!")
	}

	hpq.AddMany(qn1, qn2, qn3)

	hpq.ForEach(func(qn *QueueNode, i int) {
		log.Println(fmt.Sprintf("%d ==> %+v", i, qn))
		qn.Priority += 2 // rand.Intn(2)
	})

	hpq.ForEach(func(qn *QueueNode, i int) {
		log.Println(fmt.Sprintf("%d ==> %+v", i, qn))
	})

	if !assert.Equal(t, qn3, hpq.Peek(), "hpq peek values are equal!") {
		log.Println("FAILED PEEK TEST!")
	}

	qn4 := &QueueNode{
		Value:    "better andy",
		Priority: 99,
	}

	removedNode := hpq.Poll()

	if !assert.Equal(t, qn3, removedNode, "hpq poll values are equal!") {
		log.Println("FAILED REMOVED TEST!")
	}

	log.Println(fmt.Sprintf("removed highest prio node: %s", removedNode.Value))

	hpq.Trim()

	if !assert.Equal(t, qn2, hpq.Peek(), "post add better andy hpq peek values are equal!") {
		log.Println("FAILED TRIM & PEEK TEST!")
	}

	hpq.ForEach(func(qn *QueueNode, i int) {
		log.Println(fmt.Sprintf("%d ==> %+v", i, qn))
	})

	hpq.Add(qn4)

	if !assert.Equal(t, qn4, hpq.Peek(), "post add better andy hpq peek values are equal!") {
		log.Println("FAILED POST ADD NEW HIGH PRIO!")
	}

	hpq.RemoveMany(func(val *QueueNode) bool {
		if val.Priority >= 12 {
			return false
		}
		return true
	}, 10)

	hpq.ForEach(func(qn *QueueNode, i int) {
		log.Println(fmt.Sprintf("%d ==> %+v", i, qn))
	})

	if !assert.Equal(t, 1, hpq.Size, "hpq size correct!") {
		log.Println("FAILED HPQ SIZE TEST 1!")
	}

	if !assert.Equal(t, qn4, hpq.Peek(), "hpq peek is correct!") {
		log.Println("FAILED FINAL PEEK TEST!")
	}

	_ = hpq.Poll()

	_ = hpq.Poll()

	if !assert.Equal(t, 0, hpq.Size, "hpq size correct!") {
		log.Println("FAILED HPQ SIZE TEST 2!")
	}

}
