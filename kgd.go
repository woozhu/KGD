package kgd
“””
二元
三元
四元
五元
六元
七元
八元
九元
十元
十一元
十二元
“””
import (
	"testing"
	"fmt"
	"math"
	"strings"
	//"reflect"
)

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
var e=0.001
var a=0.618
//n元，那么k中有n个1。
func TwoYuanKGD(data []float64) (float64,float64){
    var (
        T= 1.00
        length=len(data)
        kgd= []float64{}
    )
    for i,k:=range data{
        for j,d:= range data{
            if j<=i{return}
            if math.Abs(data[j]-data[i])<e { 
                kgd= append(kgd,j-i)
            }
        }
        
    }
    
}
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
func Test0001(t *testing.T) {
	var X=[]float64{1,2,3,4,5,6,7,8}
   // var kkl=GetKLofN(2.00,4)
	//fmt.Println(len(kkl))
    //fmt.Println(kkl)
	kgd(X)
}
type CodeMap struct{
	n int //n元可公度
	k int //n元可公度的运算编码[1,2^n-1]
	data []float64  //原始序列数据
	code [][]float64
	avg float64
	kgd bool
	next float64
	percent []float64
}
func (cm *CodeMap) calculate(n float64,k float64){
	var (
		length = len(cm.data)
	    code = []float64{}
	    //i=0
	    max = 0.00
	    min = 0.00

	    //N=0.00
	)
	if n<2 ||n>16 || n>float64(length-2){
		return
	}

   // kl:=GetKLfloat64(uint64(cm.n))   //编码数组
    //k:=math.Pow(2,float64(cm.n))
    //cm.k=int(k)
	for  i:=1;i<int(k)-1;i++{
		kl:=GetKLofNK(float64(n),float64(i))
		fmt.Println(len(kl),kl)
		sum:=0.00
		code=[]float64{}
        for j:=0;float64(j)<float64(length)-n-1;j++{
        	N:=0.00
			for a,k :=range kl {
				N += k*cm.data[int(j)+int(a)]
				//fmt.Println(N)
			}
			code=append(code,N)
			if max < N { max = N }
			if min > N { min = N}
			sum +=N
		}
		cm.code =append(cm.code,code)
		fmt.Println(code)
		var(
		    avg =float64(float64(sum)/float64(len(code)))
		    percentmax = (1- (math.Abs(max)-math.Abs(avg))/math.Abs(avg))*100
			percentmin = (1- (math.Abs(avg)-math.Abs(min))/math.Abs(avg))*100
		)
		fmt.Println(sum,len(cm.code),avg,percentmax,percentmin)
		if percentmax>99&&percentmin>99{
			cm.kgd=true
		}
		cm.percent=[]float64{percentmax,percentmin}
		cm.avg=avg
		if cm.kgd{
			var (
				knext=-1.00
				NN=0.00
			)
			for a,k :=range GetKLofNK(float64(cm.n),float64(i)){
				if a==cm.n{
					knext=k
				}else {
					NN += k*cm.data[length-cm.n+a]
				}
			}
			cm.next = (avg-NN)/knext
		}else{
			cm.next=0.00
		}
	}

		//kl:= GetKLfloat64(uint64(i))
		//fmt.Println(kl,len(kl),kl[64-int(n):])
		//KN=append(KN,kl[])
		//}
		//return KN



}
//n=2~16
func kgd(X []float64) {
	var(
		length = len(X)
	)

    //fmt.Println(codemap.data)
	for n:=2;n<length-2 && n<16;n++{
		for k:=1;k< int(math.Pow(float64(2),float64(n)))-1; {
			var codemap = new(CodeMap)
			for _,d:=range X{
				codemap.data= append(codemap.data,d)
			}
			fmt.Println(n,k)
			codemap.n=n
			codemap.k = k
			//fmt.Println(codemap.n,codemap.k)
			codemap.calculate(float64(n),float64(k))
			if codemap.kgd {
				fmt.Println(codemap.next,codemap.percent)
			}
			k++
		}
	}
}
