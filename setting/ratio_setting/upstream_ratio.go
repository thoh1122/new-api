package ratio_setting

import (
	"sync"
	"new-api/common"
)

var (
	upstreamModelRatioMap      = make(map[string]float64)
	upstreamModelRatioMapMutex = sync.RWMutex{}
	upstreamModelPriceMap      = make(map[string]float64)
	upstreamModelPriceMapMutex = sync.RWMutex{}
	upstreamCompletionRatioMap      = make(map[string]float64)
	upstreamCompletionRatioMapMutex = sync.RWMutex{}
	upstreamCacheRatioMap      = make(map[string]float64)
	upstreamCacheRatioMapMutex = sync.RWMutex{}
)

func UpdateUpstreamModelRatioByJSONString(jsonStr string) error {
	upstreamModelRatioMapMutex.Lock()
	defer upstreamModelRatioMapMutex.Unlock()
	upstreamModelRatioMap = make(map[string]float64)
	return common.Unmarshal([]byte(jsonStr), &upstreamModelRatioMap)
}

func UpstreamModelRatio2JSONString() string {
	upstreamModelRatioMapMutex.RLock()
	defer upstreamModelRatioMapMutex.RUnlock()
	jsonBytes, _ := common.Marshal(upstreamModelRatioMap)
	return string(jsonBytes)
}

func UpdateUpstreamModelPriceByJSONString(jsonStr string) error {
	upstreamModelPriceMapMutex.Lock()
	defer upstreamModelPriceMapMutex.Unlock()
	upstreamModelPriceMap = make(map[string]float64)
	return common.Unmarshal([]byte(jsonStr), &upstreamModelPriceMap)
}

func UpstreamModelPrice2JSONString() string {
	upstreamModelPriceMapMutex.RLock()
	defer upstreamModelPriceMapMutex.RUnlock()
	jsonBytes, _ := common.Marshal(upstreamModelPriceMap)
	return string(jsonBytes)
}

func UpdateUpstreamCompletionRatioByJSONString(jsonStr string) error {
	upstreamCompletionRatioMapMutex.Lock()
	defer upstreamCompletionRatioMapMutex.Unlock()
	upstreamCompletionRatioMap = make(map[string]float64)
	return common.Unmarshal([]byte(jsonStr), &upstreamCompletionRatioMap)
}

func UpstreamCompletionRatio2JSONString() string {
	upstreamCompletionRatioMapMutex.RLock()
	defer upstreamCompletionRatioMapMutex.RUnlock()
	jsonBytes, _ := common.Marshal(upstreamCompletionRatioMap)
	return string(jsonBytes)
}

func UpdateUpstreamCacheRatioByJSONString(jsonStr string) error {
	upstreamCacheRatioMapMutex.Lock()
	defer upstreamCacheRatioMapMutex.Unlock()
	upstreamCacheRatioMap = make(map[string]float64)
	return common.Unmarshal([]byte(jsonStr), &upstreamCacheRatioMap)
}

func UpstreamCacheRatio2JSONString() string {
	upstreamCacheRatioMapMutex.RLock()
	defer upstreamCacheRatioMapMutex.RUnlock()
	jsonBytes, _ := common.Marshal(upstreamCacheRatioMap)
	return string(jsonBytes)
}

func GetUpstreamModelRatioCopy() map[string]float64 {
	upstreamModelRatioMapMutex.RLock()
	defer upstreamModelRatioMapMutex.RUnlock()
	copyMap := make(map[string]float64, len(upstreamModelRatioMap))
	for k, v := range upstreamModelRatioMap {
		copyMap[k] = v
	}
	return copyMap
}

func GetUpstreamModelPriceCopy() map[string]float64 {
	upstreamModelPriceMapMutex.RLock()
	defer upstreamModelPriceMapMutex.RUnlock()
	copyMap := make(map[string]float64, len(upstreamModelPriceMap))
	for k, v := range upstreamModelPriceMap {
		copyMap[k] = v
	}
	return copyMap
}

func GetUpstreamCompletionRatioCopy() map[string]float64 {
	upstreamCompletionRatioMapMutex.RLock()
	defer upstreamCompletionRatioMapMutex.RUnlock()
	copyMap := make(map[string]float64, len(upstreamCompletionRatioMap))
	for k, v := range upstreamCompletionRatioMap {
		copyMap[k] = v
	}
	return copyMap
}

func GetUpstreamCacheRatioCopy() map[string]float64 {
	upstreamCacheRatioMapMutex.RLock()
	defer upstreamCacheRatioMapMutex.RUnlock()
	copyMap := make(map[string]float64, len(upstreamCacheRatioMap))
	for k, v := range upstreamCacheRatioMap {
		copyMap[k] = v
	}
	return copyMap
}
