package split

import (
	"strings"
)

//切割字符串

func Split(str string,sep string) ([]string ){
	ret := make([]string,0,len(str))
	if sep == ""{
		for i := 0; i < len(str);i++{
			ret = append(ret, str[i:i+1])
		}
		return ret
	}
	for{
		index := strings.Index(str,sep)
		if index == -1{
			ret = append(ret, str)
			return ret
		}else {
			ret = append(ret,str[:index])
			str = str[index+len(sep):]
		}
	}

}
