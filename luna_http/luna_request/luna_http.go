package luna_request

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"github.com/gamexg/proxyclient"
	mtls "github.com/refraction-networking/utls"
	//mtls "gitlab.com/yawning/utls.git"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
	"utls-master/http2"
)

func JhyGet(url string) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "Get--->注意查看错误")
		}
	}()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyRes := JhyGetResponse(resp, false)
	return bodyRes
	//fmt.Println(resp.StatusCode, bodyRes)
}


var pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 4096))
	},
}



func JhyGetResponse(resp *http.Response, isPint bool) (result string) {
	//start:=time.Now().UnixNano()/1e6
	var err error
	var reader io.ReadCloser
	if !resp.Uncompressed{
		reader, err = gzip.NewReader(resp.Body)
		fmt.Println(err)
		if err != nil {
			return ""
		}
	}else if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		fmt.Println(err)
		if err != nil {
			return ""
		}
	}else {
		reader = resp.Body
	}
	buffer := pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer pool.Put(buffer)
	io.Copy(buffer, reader)
	result=buffer.String()
	if isPint {
		fmt.Println(result)
	}
	//fmt.Println("读取耗时:", time.Now().UnixNano()/1e6-start)
	return result
}


func JhyGetClient(isproxy bool, username, password, proxyip, proxyport, servername string) (errCode int, client *http.Client, conn *mtls.UConn, transport *http2.Transport) {
	errCode = 0
	defer func() {
		if err := recover(); err != nil {
			errCode = 0
			fmt.Println("代理ip问题", err, "线程:")
		}
	}()
	var err error
	config := mtls.Config{ServerName: servername,
		InsecureSkipVerify: true,
	}
	var dialConn net.Conn
	var proxyClient proxyclient.ProxyClient

	if isproxy {
		if len(username) > 0 && len(password) > 0 {
			//proxyDialer, err = proxy.SOCKS5("tcp", proxyAddr, auth, proxy.Direct)
			proxyClient, err = proxyclient.NewProxyClient("http://" + username + ":" + password + "@" + proxyip + ":" + proxyport + "?insecureskipverify=true")
			if err != nil {
				fmt.Println("初始化连接出现错误:" + err.Error())
				return 0, nil, nil, nil
			}
			dialConn, err = proxyClient.Dial("tcp", servername+":443")
		} else {
			proxyClient, err = proxyclient.NewProxyClient("http://" + proxyip + ":" + proxyport + "?insecureskipverify=true")

			if err != nil {
				fmt.Println("初始化连接出现错误:" + err.Error())
				return 0, nil, nil, nil
			}
			dialConn, err = proxyClient.Dial("tcp", servername+":443")
		}
		//start := time.Now().UnixNano() / 1e6
		if err != nil {
			fmt.Println("初始化连接出现错误:" + err.Error())
			return 0, nil, nil, nil
		}
		//fmt.Println("探测proxy-连接到目标网络--速度--响应耗时-01:",servername, strconv.FormatInt(time.Now().UnixNano()/1e6-start, 10), "ip:", proxyip, "毫秒", "如果过慢请考虑更换proxy问题 or 网速")
		//wllog.InfoS("探测proxy-连接到目标网络--速度--响应耗时-01:",servername, strconv.FormatInt(time.Now().UnixNano()/1e6-start, 10), "ip:", proxyip, "毫秒", "如果过慢请考虑更换proxy问题 or 网速")
	} else {
		dialConn, err = net.DialTimeout("tcp", servername+":443", time.Duration(15*time.Minute))
		if err != nil {
			fmt.Println("初始化连接出现错误:" + err.Error())
			return 0, nil, nil, nil
		}
	}
	if err != nil {
		fmt.Println("初始化连接出现错误:" + err.Error())
		return 0, nil, nil, nil
	}

	var tlsConn *mtls.UConn
	//4539973;3485744 绑定83 稳稳的；
	tlsConn = mtls.UClient(dialConn, &config, mtls.HelloChrome_72)
	//start := time.Now().UnixNano() / 1e6
	err = tlsConn.Handshake()
	if err != nil {
		fmt.Println("初始化连接出现错误:" + err.Error())
		return 0, nil, nil, nil
	}
	//https://gitlab.com/yawning/utls
	//fmt.Println("建立连接02--响应耗时:", strconv.FormatInt(time.Now().UnixNano()/1e6-start, 10), "毫秒", "如果过慢请考虑更换proxy问题 or 网速")
	//wllog.InfoS("建立连接02--响应耗时:", strconv.FormatInt(time.Now().UnixNano()/1e6-start, 10), "毫秒", "如果过慢请考虑更换proxy问题 or 网速")
	//--
	transport = &http2.Transport{
		TLSClientConfig: &tls.Config{
			ServerName: servername, InsecureSkipVerify: true,
		},
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("代理ip问题>", "线程:")
				}
			}()
			return tlsConn, err
		},
	}
	if tlsConn == nil {
		return 0, nil, nil, transport
	}
	client = &http.Client{
		Timeout:   time.Minute * 10,
		Transport: transport,
	}
	return 1, client, tlsConn, transport
}


func CheckConnIsCloseAndTrySleep(conn *mtls.UConn,client *http.Client) (isConn bool) {
	isConn=true
	defer func() {
		err := recover()
		if err != nil {
			isConn=true
			fmt.Println("[尝试连接错误!]")
			fmt.Println(err)
		}
	}()
	_, err := conn.Write(make([]byte, 0))
	if err==nil{
		fmt.Println("[尝试保持ip-活性]", `keep Alive`)
		req, _ := http.NewRequest(http.MethodGet, `https://www.ti.com.cn/AYNd9-lfjJ/WFZo7zlN/8V/DiEGmzhc/L3tjPw93/SyogdztF/YQY`, nil)
		req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.109 Safari/537.36`)
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
		req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		GetH2(NewCookie(),req,client,false)
		return false
	}
	return isConn
}

func CheckConnIsCloseAndTry(conn *mtls.UConn, username, password, proxyip, proxyport, servername string) (isConn bool) {
	isConn=true
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("[尝试连接错误!]")
			fmt.Println(err)
		}
	}()
	for i:=0;i<2;i++{
		_, err := conn.Write(make([]byte, 0))
		if err==nil{
			return false
		}else if i==0 {
			state, client, conn, transport:= JhyGetClient(true, username, password, proxyip, proxyport, servername)
			client=client
			conn=conn
			transport=transport
			if state==0{
				return true
			}
			fmt.Println("[jhyhttp-CheckConnIsCloseAndTry-判断连接已经关闭--尝试连接]",err,username, password, proxyip, proxyport, servername)
			continue
		}
	}
	return isConn
}

func CheckConnIsClose(conn *mtls.UConn) bool  {
	_, err := conn.Write(make([]byte, 0))
	if err==nil{
		return false
	}else {
		fmt.Println("[jhyhttp-CheckConnIsClose-判断连接已经关闭--开始换IP]",err)
		return true
	}
}



func GetH1(tlsConn *mtls.UConn,cookie *Cookie,req *http.Request,cookieListOrLine,isCloseReps bool) (state int,result string){
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "getStatus--->注意查看错误")
		}
	}()
	if cookieListOrLine{
		req.Header.Set("cookie",cookie.ToString())
	}else {
		cookie.WrriteCookie(req)
	}
	req.Write(tlsConn)
	resp, err := http.ReadResponse(bufio.NewReader(tlsConn), req)
	if err != nil {
		fmt.Println("Failed get: %s", err)
	}
	cookie.ReadCookie(resp)
	defer resp.Body.Close()
	bodyRes:=``
	bodyRes= JhyGetResponse(resp,false)
	return resp.StatusCode,bodyRes
}
func GetH2(cookie *Cookie, req *http.Request, client *http.Client, cookieListOrLine bool) (state int, result string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "getStatus--->注意查看错误")
		}
	}()
	if cookieListOrLine{
		req.Header.Set("jhy-cookie",cookie.ToString())
	}else {
		cookie.WrriteCookie(req)
	}
	resp, err := client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	cookie.ReadCookie(resp)
	bodyRes:=``
	bodyRes= JhyGetResponse(resp,false)
	return resp.StatusCode,bodyRes
}



func Geth2301(cookie *Cookie, req *http.Request, client *http.Client, cookieListOrLine bool) (state int, result, url string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "getStatus--->注意查看错误")
		}
	}()
	if cookieListOrLine{
		req.Header.Set("jhy-cookie",cookie.ToString())
	}else {
		cookie.WrriteCookie(req)
	}
	resp, err := client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	url=resp.Header.Get("Location")
	cookie.ReadCookie(resp)
	defer resp.Body.Close()
	bodyRes:=``
	bodyRes= JhyGetResponse(resp,false)
	return resp.StatusCode,bodyRes,url
}

func PostH1(tlsConn *mtls.UConn, cookie *Cookie, req *http.Request, cookieListOrLine bool) (state int, result string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "getStatus--->注意查看错误")
		}
	}()
	if cookieListOrLine{
		req.Header.Set("cookie",cookie.ToString())
	}else {
		cookie.WrriteCookie(req)
	}
	req.Write(tlsConn)
	resp, err := http.ReadResponse(bufio.NewReader(tlsConn), req)
	if err != nil {
		fmt.Println("Failed get: %s", err)
	}
	cookie.ReadCookie(resp)
	defer resp.Body.Close()
	bodyRes:=``
	bodyRes= JhyGetResponse(resp,false)
	return resp.StatusCode,bodyRes
}

func PostH2_301(cookie *Cookie, req *http.Request, client *http.Client, cookieListOrLine bool) (state int, result, url string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "getStatus--->注意查看错误")
		}
	}()
	if cookieListOrLine{
		req.Header.Set("jhy-cookie",cookie.ToString())
	}else {
		cookie.WrriteCookie(req)
	}
	resp, err := client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	url=resp.Header.Get("Location")
	defer resp.Body.Close()
	cookie.ReadCookie(resp)
	bodyRes:=``
	bodyRes= JhyGetResponse(resp,false)
	return resp.StatusCode,bodyRes,url
}
func PostH2(cookie *Cookie, req *http.Request, client *http.Client, cookieListOrLine bool) (state int, result string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "getStatus--->注意查看错误")
		}
	}()
	if cookieListOrLine{
		req.Header.Set("jhy-cookie",cookie.ToString())
	}else {
		cookie.WrriteCookie(req)
	}
	resp, err := client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	cookie.ReadCookie(resp)
	bodyRes:= JhyGetResponse(resp,false)
	return resp.StatusCode,bodyRes
}
