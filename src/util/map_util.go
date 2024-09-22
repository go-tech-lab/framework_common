package util

//FlatMap2Array 将map转化为key-value数组
func FlatMap2Array(params map[string]interface{}) []interface{} {
	length := len(params)
	if length <= 0 {
		return nil
	}
	flatParams := make([]interface{}, 0, length)
	for k, v := range params {
		flatParams = append(flatParams, k, v)
	}
	return flatParams
}
