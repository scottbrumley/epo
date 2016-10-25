# EPO API Go Program

## Usage Example
```Go
package main

import (
	"github.com/scottbrumley/epo"
	"fmt"
)

func main() {
	var myParms epo.ParamStruct
    	
    	myParms.Cmd = "system.find"
    	myParms.Url = "https://epohost:8443/remote"
    	myParms.Parms = "searchText=."
    	myParms.Output = "json"
    	myParms.SslIgnore = true
    	myParms.UserName = "username"
    	myParms.UserPass = "password"
	
	myParms := epo.GetParams()
	jsonStr := epo.GetUrl(myParms)
	data := epo.DecodeJson(jsonStr)

	lineNum := 1
	for i := range data {
		fmt.Printf("Record # %v\n", lineNum)
		for key, value := range data[i] {
			fmt.Println("Key:", key, "Value:", value)
		}
		lineNum = lineNum + 1
		fmt.Println("")
	}

}

```

### Find All Systems
..-UserName=someuser 
..-UserPass=somepass 
..-SslIgnore=true 
..-Url=https://epohost:8443/remote 
..-Cmd=system.find 
..-Parms=searchText=. 
..-Output=json

### Find All Policies
..-UserName=someuser 
..-UserPass=somepass 
..-SslIgnore=true 
..-Url=https://epohost:8443/remote 
..-Cmd=policy.find 
..-Parms=searchText=. 
..-Output=json

### Find All Repositories
UserName=someuser 
UserPass=somepass 
SslIgnore=true 
Url=https://epohost:8443/remote 
Cmd=repository.find 
Parms=searchText=. 
Output=json


### Find All Queries
UserName=someuser 
UserPass=somepass 
SslIgnore=true 
Url=https://epohost:8443/remote 
Cmd=core.listQueries
Parms=searchText=. 
Output=json

### Find All Groups
UserName=someuser 
UserPass=somepass 
SslIgnore=true 
Url=https://epohost:8443/remote 
Cmd=systemFindGroups
Output=json