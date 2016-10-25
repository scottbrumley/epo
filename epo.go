package epo

import (
	"os"
	"fmt"
	"flag"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"encoding/json"
	"log"
	"strings"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
	"bufio"
)

// Parameters from command line
type ParamStruct struct{
	UserName string
	UserPass string
	Test bool
	SslIgnore bool
	Url string
	Cmd string
	Parms string
	Output string
}

func showSyntax(){
	fmt.Println("For Help " + os.Args[0] + " -h")
}

func getPasswd()(string){
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err == nil {
		fmt.Println("\nPassword typed: " + string(bytePassword))
	}
	password := string(bytePassword)

	return strings.TrimSpace(password)
}

func getUser()(string){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')

	return strings.TrimSpace(username)
}

// Collect parameters from the command line
func GetParams()(retParams ParamStruct){
	var userFlag = flag.String("user","","ePO Username")
	var passFlag = flag.String("password","","ePO Password")
	var cmdFlag = flag.String("cmd","core.help","Command String")
	var parmsFlag = flag.String("parms","","Parameter String")
	var sslIgnoreFlag = flag.Bool("ignoressl",false,"Ignore Insecure SSL")
	var outputFlag = flag.String("output","json","Enable output: verbose, terse, xml, json")
	var testFlag = flag.Bool("test",false,"Testing Mode")
	var urlFlag = flag.String("url","","URL for API")
	flag.Parse()

	retParams.UserName = *userFlag
	retParams.UserPass = *passFlag
	retParams.Cmd = *cmdFlag
	retParams.Parms = *parmsFlag
	retParams.SslIgnore = *sslIgnoreFlag
	retParams.Output = *outputFlag
	retParams.Url = *urlFlag
	retParams.Test = *testFlag

	// Test Params
	if retParams.UserName == "" {
		retParams.UserName = getUser()
	}
	if retParams.UserPass == "" {
		retParams.UserPass = getPasswd()
	}

	if retParams.Url == "" {
		showSyntax()
		fmt.Println("url requires a value")
		os.Exit(0)
	} else {
		if (retParams.Cmd != "") {
			if (retParams.Parms != "") {
				retParams.Url = retParams.Url + "/" + retParams.Cmd + "?" + retParams.Parms + "&:output=" + retParams.Output
			} else {
				retParams.Url = retParams.Url + "/" + retParams.Cmd + "?:output=" + retParams.Output
			}
		}
	}

	return retParams
}

// System Property Type
type SystemProperties []map[string]interface{}

// Decode Json to array of SystemProperty Type
func DecodeJson(jsonStr string)(res SystemProperties){
	jsonStr = strings.Replace(jsonStr,"OK:","",-1)
	jsonStr = strings.Replace(jsonStr,"\n","",-1)
	if err := json.Unmarshal([]byte(jsonStr), &res); err != nil {
		fmt.Println("Error:")
		log.Fatal(err)
	}
	return res
}

// Authenticate to ePO API and Get HTTP response
func GetUrl(myParms ParamStruct)(retVal string){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: myParms.SslIgnore},
	}

	client := &http.Client{
		Transport: tr,
	}

	fmt.Println(myParms.Url)
	req, err := http.NewRequest("GET", myParms.Url, nil)
	req.SetBasicAuth(myParms.UserName, myParms.UserPass)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		retVal = string(body)
		//fmt.Println("response Body:", string(body))
	}
	return retVal
}
