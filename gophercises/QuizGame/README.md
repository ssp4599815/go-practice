# Quiz Game

## part 1 


## part 2


## 主要学习
- csv
- os
- time
- select 
- chan

### go 使用 csv
打开文件并且创建一个`csv.Reader`对象
```go
csvFile,err := os.Open("./foo.csv")
if err != nil {
	panic(err)
}
defer csvFile.Close()
csvReader := csv.NewReader(csvFile)

// 读取单行
row, err := csvReader.Read()

// 读取所有
rows,err := csvReader.ReadAll() 
type [][]string
if err != nil {
	panic(err)
}
for i,row := range rows {
	// process the row here
}
```

### go 创建定时器
初始化一个到期时间据此时的间隔为3小时30分的定时器
```go
t := time.Newtimer(3*time.Hour + 30*time.Minute)
```



#### 参考： 
- https://gophercises.com/exercises/quiz
- [使用Go语言标准库对CSV文件进行读写](https://xiequan.info/%E4%BD%BF%E7%94%A8go%E8%AF%AD%E8%A8%80%E6%A0%87%E5%87%86%E5%BA%93%E5%AF%B9csv%E6%96%87%E4%BB%B6%E8%BF%9B%E8%A1%8C%E8%AF%BB%E5%86%99/)