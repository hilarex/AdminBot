package framework

/*
import (
)
*/

func IsInSlice(val string, slice []string) bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}

func MergeMap(maps ...map[string]int64) map[string]int64{
	// first map is the source, second one are new values
	
	res := map[string]int64{}

	for _, m := range maps{
		for k,v := range m {
			res[k] = v
		}
	}

	return res
}