package utils

import "strings"

func FindMatchString(src, skip, start, end string) (string, int) {
	var index = strings.Index(src, skip)
	if index != -1 {
		var position = index + len(skip)
		if position < len(src) {
			//has more char after string skip
			src = src[position:]
			index = strings.Index(src, start)
			if index != -1 {
				var newPos = position  //record how many char has skipped since origin string
				position = index + len(start)
				newPos += position
				if position < len(src) {
					//more char after string start
					src = src[position:]
					index = strings.Index(src, end)
					if index != -1 {
						return src[:index], newPos + index + len(end)
					}
				}
			}
		}
	}
	return "", -1
}

/*
var index, position int
	if len(skip) != 0 {
		index = strings.Index(src, skip)
		if index != -1 {
			position = index + len(skip)
			if position < len(src) {
				//has more char after string skip
				src = src[position:]
			}
		}
	}
	index = strings.Index(src, start)
	if index != -1 {
		var newPos = position  //record how many char has skipped since origin string
		position = index + len(start)
		newPos += position
		if position < len(src) {
			//more char after string start
			src = src[position:]
			index = strings.Index(src, end)
			if index != -1 {
				return src[:index], newPos + index + len(end)
			}
		}
	}
	return "", -1
	*/
