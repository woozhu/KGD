package kgd

import(
    “strings”
)

//获得一个长度为64的list，十进制v的二进制编码字符串。
func ToBinaryString64(v interface{}) (string){
	s := ToBinaryString(v)
	s=strings.Replace(s,"[","",-1)
	s=strings.Replace(s,"]","",-1)
	s=strings.Replace(s," ","",-1)
	return s
}
//把0变-1，1变1
func GetSubPlusList(v interface{}) []float64{
	s:=ToBinaryString64(v)
	var kl=[]float64{}
	for _,x :=range s{
		if x=='1'{
			kl=append(kl,1)
		}else{
			kl=append(kl,-1)
		}
	}
	return kl
}
 
//得到一个0到0，1到1的list
func GetZeroOneList(v interface{}) []float64{
    s:=ToBinaryString64(v)
    var kl=[]float64{}
    for _,x := range s{
        if x==‘1’{
            kl=append(kl,1)
        }else{
            kl=append(kl,0)
        }
    }
    return kl
}
//返回十进制v的二进制字符串中1的个数。
func GetOneNum(v interface{}) float64{
    var num=0.00
    kl:= GetZeroOneList(v)
    for _,k:=range kl{
        num+=k
    }
    return num
}
//过滤编码
//n次元，数据长度l，n次元的数据组合
func BinaryCode(n,l int,c int,data []float64)[][]float64{
    var (
        result=[][]float64{}
        length=len(data)
   
    )
    for c:= 2;c<math.pow2(length)-2;c++{
        r,f:=math.modf(math.log2(float64(c)))
        r1,f1:=math.modf(math.log2(float64(c+1)))
    //把2的n次方和2的n次方-1过滤掉
        if f==float64(0)||f1==float64(0){
            return
        }
    //把kl变短
        kl:=GetZeroOneList(uint64(c))
        kl= kl[binary_length-length:]
    //如果kl中的1为n个，返回kl，否则过滤掉
        if GetOneNum(uint64(c))==float64(n){
           result=append(result,kl)
        }
    }
    return result
}
//n 的最大值建议在20以内，以免内存泄露。
func GetKLofNK(n float64,k float64) []float64{
	kl:= GetSubPlusList(uint64(k))
	//fmt.Println(n,k,kl[64-int(n):])
	return kl[64-int(n):]
	//for i:=1;i<int(k)-1;i++{
		//kl:= GetKLfloat64(uint64(i))
		//fmt.Println(kl,len(kl),kl[64-int(n):])
		//KN=append(KN,kl[])
	//}
    //return KN
}
//n元，那么k中有n个1。
func GetSubPlusResult(n,k float64,v []float64)float64{
    var (
        length=len(v)
        sp =GetSubPlusList(uint64(k))
        nsp=sp[64-int(n):]
     )
    for i:=0;i<length-int(n)+1;i +=int(n){
        for j,k:=range nsp{
            v[i+j]*k
        ｝
    }
}