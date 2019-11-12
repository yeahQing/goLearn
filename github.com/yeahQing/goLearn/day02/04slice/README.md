# Go语言学习笔记
## Day02

### 切片(Slice)

```go

	a2 := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a2:%v type: %T len:%d cap:%d\n", a2, a2, len(a2), cap(a2))
	b2 := a2[0:3] // 切片
	fmt.Printf("b2:%v type: %T len:%d cap:%d\n", b2, b2, len(b2), cap(b2))
	c2 := b2[0:5] // 切片再切片
	fmt.Printf("c2:%v type: %T len:%d cap:%d\n", c2, c2, len(c2), cap(c2))

```

> cap($0) 是求切片的容量，切片的容量是指从切割的位置开始到原数组的长度。比如b2为a2的切片，a2的容量为6,b2 := a2[1:3],b2从下标为1的位置开始切割，那么b2的容量则为5。