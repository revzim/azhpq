# azhpq

[![PkgGoDev](https://pkg.go.dev/badge/github.com/revzim/azhpq)](https://pkg.go.dev/github.com/revzim/azhpq)

## golang heap priority queue
- simple, quick, & lightweight
- zero dependencies

## USE:
```
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
  
  hpq.Add(qn1)
  hpq.Add(qn2) 
  hpq.Add(qn3)
  // OR
  hpq.AddMany(qn1, qn2, qn3)

  hpq.ForEach(func (qn *QueueNode, i int) {
    log.Println(fmt.Sprintf("%d ==> %+v", i, qn))
  })

  // OUT: 
  // 0 ==> &{Value:jim Priority:5}
  // 1 ==> &{Value:john Priority:2}
  // 2 ==> &{Value:andy Priority:1}

  highestPrioNode := hpq.Poll()
  log.Println(fmt.Sprintf("removed highest prio node: %s", highestPrioNode.Value))

  // OUT: 
  // removed highest prio node: jim

  hpq.ForEach(func (qn *QueueNode, i int) {
    log.Println(fmt.Sprintf("%d ==> %+v", i, qn))
  })

  // OUT: 
  // 0 ==> &{Value:john Priority:2}
  // 1 ==> &{Value:andy Priority:1}
```

## TEST:
- `go test -v -run TestHPQ azhpq.go azhpq_test.go`
```
  === RUN   TestHPQ
  2021/08/17 10:54:59 0 ==> &{Value:jim Priority:5}
  2021/08/17 10:54:59 1 ==> &{Value:john Priority:2}
  2021/08/17 10:54:59 2 ==> &{Value:andy Priority:1}
  2021/08/17 10:54:59 0 ==> &{Value:jim Priority:7}
  2021/08/17 10:54:59 1 ==> &{Value:john Priority:4}
  2021/08/17 10:54:59 2 ==> &{Value:andy Priority:3}
  2021/08/17 10:54:59 removed highest prio node: jim
  2021/08/17 10:54:59 0 ==> &{Value:john Priority:4}
  2021/08/17 10:54:59 1 ==> &{Value:andy Priority:3}
  --- PASS: TestHPQ (0.00s)
  ok      command-line-arguments  0.138s
```

## AUTHOR:
- [revzim](https://github.com/revzim)
