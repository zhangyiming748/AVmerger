package AVmerger

import "github.com/zhangyiming748/AVmerger/merge"

func Single(src, dst string) {
	merge.Single(src, dst)
}

func Multi(src, dst string) {
	merge.Multi(src, dst)
}
