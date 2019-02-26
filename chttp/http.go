// Create by Yale 2019/2/26 11:23
package chttp

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Get(url string) (string,error) {
	return GetByHeader(url,nil)
}
func DownloadFile(url , downPath string) error {

	f,err:=os.Create(downPath)
	if err!=nil {
		return err
	}

	client:=&http.Client{}

	req,err:= http.NewRequest("GET",url,nil)
	if err!=nil {
		return err
	}
	resp,err:=client.Do(req)
	if err!=nil {
		return err
	}


	defer resp.Body.Close()
	defer f.Close()

	if resp.StatusCode!=http.StatusOK {
		return errors.New(resp.Status)
	}

	_,err=io.Copy(f,resp.Body)
	if err!=nil {
		return err
	}

	return nil

}
func Post(u string,body io.Reader,headerMap map[string]string) (string,error) {

	client:=&http.Client{}

	req,err:= http.NewRequest("POST",u,body)
	if err!=nil {
		return "",err
	}

	for k,v:=range headerMap{
		req.Header.Add(k,v)
	}

	resp,err:=client.Do(req)
	if err!=nil {
		return "",err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(resBody),nil

}
func map2Reader(params map[string]string) io.Reader {
	reqParams:=make(url.Values)
	if params !=nil{
		for k,v:=range params{
			reqParams[k] =[]string{v}
		}
	}
	return strings.NewReader(reqParams.Encode())
}
func getHeaderMap(contentType string,headerMap map[string]string) map[string]string {
	header:=map[string]string{"Content-Type":contentType}
	for k,v:=range headerMap{
		if strings.ToLower(k)!="content-type" {
			header[k]=v
		}
	}
	return header
}
func PostFormByHeader(u string,params map[string]string,headerMap map[string]string) (string,error) {
	return Post(u,map2Reader(params),getHeaderMap("application/x-www-form-urlencoded",headerMap))
}
func PostForm(u string,params map[string]string) (string,error) {
	return Post(u,map2Reader(params),getHeaderMap("application/x-www-form-urlencoded",nil))
}
func PostJson(u ,jsonStr string) (string,error) {
	return Post(u,bytes.NewBuffer( []byte(jsonStr)),getHeaderMap("application/json",nil))
}
func GetByHeader(url string,headerMap map[string]string)(string,error)  {
	client:=&http.Client{}

	req,err:= http.NewRequest("GET",url,nil)
	if err!=nil {
		return "",err
	}

	for k,v:=range headerMap{
		req.Header.Add(k,v)
	}
	resp,err:=client.Do(req)
	if err!=nil {
		return "",err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(body),nil
}