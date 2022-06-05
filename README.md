# gorobots

#### 介绍 Introduction
This is a project with golang robots.txt parsing and robots checking for spider.

#### 安装教程 Installing

```bash
go get github.com/ocaraworks/gorobots
```

#### 使用方式 Using Method
```golang
    //get robots content
    r, _ := http.Get("http://www.163.com/robots.txt")
	
	//parse robots
	rules, sitemaps, err := Parse(r.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(rules)
	fmt.Println(sitemaps)

	//check url dir
	robotsAllowed, matchData, delaySeconds := CheckPath("http://www.163.com/bbs", "Googlebot", rules)

	if !robotsAllowed {
		fmt.Println(matchData)
		fmt.Println(delaySeconds)
		fmt.Println("It's not allowed.")
	}
```