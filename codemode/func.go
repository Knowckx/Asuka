// 检查redis的缓存是否有改动
func WatchRankCacheUpdate() {
	watchGap, err := time.ParseDuration("1m")
	if err != nil {
		fmlog.Err("time parse err:%s", err)
		return
	}

	for {
		rankCache.UpdateRankCacheFromRedis() //先执行一次，作为Init
		t := time.NewTimer(watchGap)
		<-t.C
	}
}