package kgd
import (
	"testing"
	"fmt"
	"math"
	"strings"
	"reflect"
)

func ToBinaryString64(v interface{}) (string){
	s := ToBinaryString(v)
	s=strings.Replace(s,"[","",-1)
	s=strings.Replace(s,"]","",-1)
	s=strings.Replace(s," ","",-1)
	return s
}
func Test0001(t *testing.T) {
	//var X=[]float64{1,2,3,4,5,6,7,8}
	s := ToBinaryString64(uint64(866))
	fmt.Println(s,len(s))
	for n,x :=range s{
		if x=='1'{
			fmt.Println("find 1 in %d",n+1)
			fmt.Println(n,x,reflect.TypeOf(x),reflect.ValueOf(x))
		}
	}
	fmt.Println(s,s[63])
	//kgd(X)
}
type CodeMap struct{
	n int //n元可公度
	k int //n元可公度的运算编码[1,2^n-1]
	data []float64  //原始序列数据
	code []float64
	avg float64
	kgd bool
	next float64
	percent float64
}
func (cm CodeMap) calculate(){
	var (
		length = len(cm.data)
	    code = []float64{}
	    i=0
	    max = 0.00
	    sum=0.00
	    N=0.00
	)
	if cm.n<2 ||cm.n>16 || cm.n>length-2{
		return
	}
	s:=ToBinaryString64(uint64(cm.n))
	fmt.Println(s)
    s= strings.Replace(s,"[","",-1)
    s=strings.Replace(s,"]","",-1)
    s=strings.Replace(s," ","",-1)
    fmt.Println(s)
	for true{
		N=0.00
		for a,k :=range s[64-cm.n:]{
				if k=='0'{
					N -=cm.data[i+a]
				}else{
					N +=cm.data[i+a]
				}
		}
		i++
		sum +=N
		if max<N{max = N}
		code = append(code,N)
		if i>length{
			break
		}
	}
	cm.code=code
	avg:=sum/float64(len(code))
	cm.avg = avg
	cm.percent = (max-avg)/avg*100
	if (max-avg)/avg >(1-0.618){
		cm.kgd=false
	}else{
		cm.kgd=true
	}
	var (
		knext=-1.00
		NN=0.00
	)
	for a,k :=range s[64-cm.n:]{
		if a==cm.n{
			if k =='0'{
				knext=-1.00
			}else{
				knext=1.00
			}
		}else {
			if k == '0' {
				NN -= cm.data[length-cm.n+a+1]
			} else {
				NN += cm.data[length-cm.n+a+1]
			}
		}
	}
    cm.next = (avg-NN)/knext
}
//n=2~16
func kgd(X []float64) {
	var(
		length = len(X)
		codemap=CodeMap{}
	)
	for _,d:=range X{
		codemap.data= append(codemap.data,d)
	}

	for n:=2;n<length-2 && n<16;n++{
		for k:=0;k< int(math.Pow(float64(2),float64(n)));k++ {
			codemap.n=n
			codemap.k = k
			codemap.calculate()
			if codemap.kgd {
				fmt.Println(codemap.next,codemap.percent)
			}
		}
	}
}
