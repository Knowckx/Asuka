// 真实启动参数
func MakeSrvInitArgs(env string) {
	newArgs := []string{}
	if env == "dev" || env == "" {
		newArgs = []string{"--registry=consul",
			"--registry_address=192.168.8.6:8500",
			"--server", "grpc",
			"--client", "grpc",
		}
	// os.Args = append(os.Args, newArgs...)
	os.Args = newArgs
}


// 随机浮点数 2位小数
func Float64() float64 {
	out := rand.Float64()
	sout := fmt.Sprintf("%.2f", out)
	out, _ = strconv.ParseFloat(sout, 64)
	return out
}

func RandomTime(in time.Time) time.Time {
	ranDay := rand.Int31n(20) - 10
	ss := fmt.Sprintf("%dh", ranDay*24)
	d, _ := time.ParseDuration(ss)
	new := in.Add(d)
	return new
}