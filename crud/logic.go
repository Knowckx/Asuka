package crud

// 对比最新的数据newData，找出oldData里已经过期的数据
func DiffClientSecrets(oldData []*AzureClient, newData []*AzureClient) []*AzureClient {
	filterMap := make(map[string]int, len(newData))
	for _, da := range newData {
		filterMap[da.SubsName] = 1
	}
	unfound := []*AzureClient{}
	for _, da := range oldData {
		_, ok := filterMap[da.SubsName]
		if ok {
			continue
		}
		unfound = append(unfound, da)
	}
	return unfound
}
