UCAS 2022春季”模式串匹配与信息过滤“课程作业

Implement Aho-Corasick Algorithm based golang, and tesing this in url data.

# Usage
```go
package main

import (
	ac "PatternCourse/ac"
)

func main() {
	dictionary := []string{"he", "hello", "hers"}
	newac := ac.New(dictionary)
	newac.ACMatch("hershe", dictionary, 0)

}
````

# Produce Test File 
``` go
package main

import (
 	"PatternCourse/FileSplit"
 )

func main() {
 	fileToBeChunked := "./data/rule/url_data.txt"
 	// FileSplit.Split_xMB(fileToBeChunked, 5)
 	FileSplit.Split_xLINE(fileToBeChunked, 1000)
 }
```