package gongdu
//string 中的0对应-1,1对应1.
func GetKByString(s string) []int{
    var r=[]int
    for _, k :=range s{
        if k =="0" {
            append(r,-1)
        }else{
            append(r,1)
        }
    }
    return r
}
//n元可公度
func KGD(X [...]float64) float64{
    if x.len()<=8 {return 0.000} //数据太少
    for n:=2,n<X.len()-1,{
        
    }
}

func getN(n int,X [n]float64) float64{
    
}

// append bytes of string in binary format.
func appendBinaryString(bs []byte, b byte) []byte {
	var a byte
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			bs = append(bs, zero)
		default:
			bs = append(bs, one)
		}
		b <<= 1
	}
	return bs
}


//
import "encoding/binary"

// Uint16ToBinaryString get the string of a uint16 number in binary format.
func Uint16ToBinaryString(i uint16) string {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, i)
	return BytesToBinaryString(bs)
}

// Uint32ToBinaryString get the string of a uint32 number in binary format.
func Uint32ToBinaryString(i uint32) string {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, i)
	return BytesToBinaryString(bs)
}

// Uint64ToBinaryString get the string of a uint64 number in binary format.
func Uint64ToBinaryString(i uint64) string {
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, i)
	return BytesToBinaryString(bs)
}
