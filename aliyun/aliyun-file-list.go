package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	var data = strings.NewReader(`{"drive_id":"22467661","parent_file_id":"root","limit":20,"all":false,"url_expire_sec":14400,"image_thumbnail_process":"image/resize,w_256/format,jpeg","image_url_process":"image/resize,w_1920/format,jpeg/interlace,1","video_thumbnail_process":"video/snapshot,t_1000,f_jpg,ar_auto,w_256","fields":"*","order_by":"updated_at","order_direction":"DESC"}`)
	req, err := http.NewRequest("POST", "https://api.aliyundrive.com/adrive/v3/file/list?jsonmask=next_marker%2Citems(name%2Cfile_id%2Cdrive_id%2Ctype%2Csize%2Ccreated_at%2Cupdated_at%2Ccategory%2Cfile_extension%2Cparent_file_id%2Cmime_type%2Cstarred%2Cthumbnail%2Curl%2Cstreams_info%2Ccontent_hash%2Cuser_tags%2Cuser_meta%2Ctrashed%2Cvideo_media_metadata%2Cvideo_preview_metadata%2Csync_meta%2Csync_device_flag%2Csync_flag%2Cpunish_flag)", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.aliyundrive.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIyYjQzMjYwY2Y4YzA0NTc5YjEwNjNhNThmMmVhZDkwZiIsImN1c3RvbUpzb24iOiJ7XCJjbGllbnRJZFwiOlwiMjVkelgzdmJZcWt0Vnh5WFwiLFwiZG9tYWluSWRcIjpcImJqMjlcIixcInNjb3BlXCI6W1wiRFJJVkUuQUxMXCIsXCJTSEFSRS5BTExcIixcIkZJTEUuQUxMXCIsXCJVU0VSLkFMTFwiLFwiVklFVy5BTExcIixcIlNUT1JBR0UuQUxMXCIsXCJTVE9SQUdFRklMRS5MSVNUXCIsXCJCQVRDSFwiLFwiT0FVVEguQUxMXCIsXCJJTUFHRS5BTExcIixcIklOVklURS5BTExcIixcIkFDQ09VTlQuQUxMXCIsXCJTWU5DTUFQUElORy5MSVNUXCIsXCJTWU5DTUFQUElORy5ERUxFVEVcIl0sXCJyb2xlXCI6XCJ1c2VyXCIsXCJyZWZcIjpcImh0dHBzOi8vd3d3LmFsaXl1bmRyaXZlLmNvbS9cIixcImRldmljZV9pZFwiOlwiNDhkZDRlZDRkN2M5NGY1MGExMWU3ZjAwOWVhZTljMWZcIn0iLCJleHAiOjE2NzY0NzYyNTAsImlhdCI6MTY3NjQ2ODk5MH0.ds5WT6P8b4ZBw8s7-f8E06XHcBxzLYtzfECITrgC59M3vL8WqDdE6BOsGSz01hUleAgfNQlrbuAGAdyWKoBjQrvoMGD2nIXRZlVe4NVeuZRnIogOyRgPtgQ_7SO41tE6Kt0UlJfxEmbWFipSjAUXPW-py-2JjXYdcQznKLozqAM")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("origin", "https://www.aliyundrive.com")
	req.Header.Set("referer", "https://www.aliyundrive.com/")
	req.Header.Set("sec-ch-ua", `"Chromium";v="110", "Not A(Brand";v="24", "Google Chrome";v="110"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("x-canary", "client=web,app=adrive,version=v3.17.0")
	req.Header.Set("x-device-id", "bc0e824d-30f9-4f7b-8e32-ed5fd59b9356")
	req.Header.Set("x-signature", "93e2f8bf415adacb294ed08a16019dfadc3d9b993d725fe941e153a1655e91d153c02dbc32cecca96014d42de63ab831e7a0b4000b0fdf5821f452dbc2952fd101")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
