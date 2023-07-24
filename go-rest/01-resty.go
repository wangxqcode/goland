package main

import (
	"github.com/go-resty/resty/v2"
)

// 详情地址   https://segmentfault.com/a/1190000040247099
/*
func main() {
	client := resty.New()

	resp, err := client.R().Get("https://baidu.com")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Info:")
	fmt.Println("Status Code\n:", resp.StatusCode())
	fmt.Println("Status\n:", resp.Status())
	fmt.Println("Proto:\n", resp.Proto())
	fmt.Println("Time:\n", resp.Time())
	fmt.Println("Received At:\n", resp.ReceivedAt())
	fmt.Println("Size:\n", resp.Size())
	fmt.Println("Headers============:")
	for key, value := range resp.Header() {
		fmt.Println(key, "=", value)
	}
	fmt.Println("Cookies============:")
	for i, cookie := range resp.Cookies() {
		fmt.Printf("cookie%d: name:%s value:%s\n", i, cookie.Name, cookie.Value)
	}
}

 */

// resty拉取信息，自动 Unmarshal：


/*
type Library struct {
	Name   string
	Latest string
}

type Libraries struct {
	Results []*Library
}

//
func main() {

	client := resty.New()

	Libraries := &Libraries{}

	// setResult

	client.R().SetResult(Libraries).Get("https://api.cdnjs.com/libraries")

	fmt.Printf("%d libraries\n", len(Libraries.Results))

	for _, lib := range Libraries.Results {
		fmt.Println("first library:")
		fmt.Printf("name:%s latest:%s\n", lib.Name, lib.Latest)
		break
	}
}

 */


type Repository struct {
	ID              int        `json:"id"`
	NodeID          string     `json:"node_id"`
	Name            string     `json:"name"`
	FullName        string     `json:"full_name"`
	Owner           *Developer `json:"owner"`
	Private         bool       `json:"private"`
	Description     string     `json:"description"`
	Fork            bool       `json:"fork"`
	Language        string     `json:"language"`
	ForksCount      int        `json:"forks_count"`
	StargazersCount int        `json:"stargazers_count"`
	WatchersCount   int        `json:"watchers_count"`
	OpenIssuesCount int        `json:"open_issues_count"`
}

type Developer struct {
	Login      string `json:"login"`
	ID         int    `json:"id"`
	NodeID     string `json:"node_id"`
	AvatarURL  string `json:"avatar_url"`
	GravatarID string `json:"gravatar_id"`
	Type       string `json:"type"`
	SiteAdmin  bool   `json:"site_admin"`
}


// token 临时生成，关闭页面token消失,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

func main() {


	/*   https://zhuanlan.zhihu.com/p/585115845
	// header
	1. 定义：客户端和服务端通过http请求和相应发送附加信息
		- 请求标头：
			Accept:	作用 浏览器端可以接受的媒体类型, 例如： Accept: text/html 代表浏览器可以接受服务器回发的类型为 text/html
	也就是我们常说的html文档,如果服务器无法返回text/html类型的数据,服务器应该返回一个406错误(non acceptable)通配符 代表任意类型 例如 (Accept: )
	代表浏览器可以处理所有类型,(一般浏览器发给服务器都是发这个)
			Accept-Encoding: 作用： 浏览器申明自己接收的编码方法，通常指定压缩方法，是否支持压缩，支持什么压缩方法（gzip，deflate），
	（注意：这不是只字符编码）; 例如： Accept-Encoding: zh-CN,zh;q=0.8
			Accept-Language: 作用： 浏览器申明自己接收的语言。 语言跟字符集的区别：中文是语言，中文有多种字符集，比如big5，gb2312，gbk等等； 例如： Accept-Language: en-us。
			Connection: keep-alive 当一个网页打开完成后，客户端和服务器之间用于传输HTTP数据的TCP连接不会关闭，如果客户端再次访问这个服务器上的网页，会继续使用这一条已经建立的连接
						close 代表一个Request完成后，客户端和服务器之间用于传输HTTP数据的TCP连接会关闭， 当客户端再次发送Request，需要重新建立TCP连接。
			Host:（发送请求时，该报头域是必需的）请求报头域主要用于指定被请求资源的Internet主机和端口号，它通常从HTTP URL中提取出来的
	例如: 我们在浏览器中输入：http://www.hzau.edu.cn 浏览器发送的请求消息中，就会包含Host请求报头域，如下：Host：www.hzau.edu.cn 此处使用缺省端口号80，
	若指定了端口号，则变成：Host：指定端口号
			Referer: 当浏览器向web服务器发送请求的时候，一般会带上Referer，告诉服务器我是从哪个页面链接过来的，服务器籍此可以获得一些信息用于处理
	。比如从我主页上链接到一个朋友那里，他的服务器就能够从HTTP Referer中统计出每天有多少用户点击我主页上的链接访问他的网站
			User-Agent: 告诉 HTTP 服务器， 客户端使用的操作系统和浏览器的名称和版本

		- 响应标头
		- 表示标头
		- 有效负荷标头

	 */

	client :=resty.New()



	var result []*Repository

	client.R().
		SetAuthToken("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").
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetQueryParams(map[string]string{
			"per_page":  "3",
			"page":      "1",
			"sort":      "created",
			"direction": "asc",
		}).
		SetPathParams(map[string]string{
			"org": "golang",
		}).
		SetResult(&result).
		Get("https://api.github.com/orgs/{org}/repos")


	//for i, repo := range result {
	//	fmt.Printf("repo%d: name:%s stars:%d forks:%d\n", i+1, repo.Name, repo.StargazersCount, repo.ForksCount)
	//}




}

