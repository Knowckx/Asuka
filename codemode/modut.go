// 代表一个键值对,例如 "UID"：715432
type KVPair struct {
	Key   string
	Value string
}

func NewKVPair(in *pb.KVPair) *KVPair {
	out := &KVPair{}
	if in == nil {
		return out
	}
	key := strings.ToLower(in.Key) //默认小写
	if key != "" {
		out.Key = in.Key
		out.Value = in.Value
	}
	return out
}


// 表示一个double的范围值
type DoubleRange struct {
	Min float64
	Max float64
}

func NewDoubleRangePB(pb *pb.DoubleRange) *DoubleRange {
	out := &DoubleRange{}
	if pb == nil {
		return nil
	}
	out.Min = pb.Min
	out.Max = pb.Max
	return out
}

func NewDoubleRange(max, min float64) *DoubleRange {
	out := &DoubleRange{}
	out.Min = min
	out.Max = max
	return out
}

// [min,max)
func (in *DoubleRange) InRange(v float64) bool {
	if in == nil {
		return true
	}
	if in.Min != 0 && v < in.Min {
		return false
	}

	if in.Max != 0 && v >= in.Max {
		return false
	}
	return true
}

func (in *DoubleRange) ToProto() *pb.DoubleRange {
	if in == nil {
		return nil
	}
	out := &pb.DoubleRange{
		Max: in.Max,
		Min: in.Min,
	}
	return out
}




// Decimal to binary  10 -> 1010
func DecBin(n int) string {
	return fmt.Sprintf("%b", n)
}

// binary to Decimal  1010 -> 10
func BinDec(b string) (int, error) {
	ss := strings.Split(b, "")
	l := len(ss)
	d := float64(0)
	for i := 0; i < l; i++ {
		f, err := strconv.ParseFloat(ss[i], 10)
		if err != nil {
			return -1, fmt.Errorf("Binary to decimal error:%s", err.Error())
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int(d), nil
}



// 一串int存sql 经纪商ID条件
func SliceInt32ToString(ins []int32) string {
	ss := ""
	if len(ins) == 0 {
		return ss
	}
	for _, in := range ins {
		if ss == "" {
			ss = fmt.Sprintf("%d", in)
			continue
		}
		ss = fmt.Sprintf("%s_%d", ss, in)
	}
	return ss
}

func StringToSliceInt32(ss string) []int32 {
	strs := StringToSlice(ss)
	outs := []int32{}
	for _, str := range strs {
		out, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		outs = append(outs, int32(out))
	}
	return outs
}