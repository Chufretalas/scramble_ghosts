package utils

import f "golang.org/x/image/font"

func GetTextWidth(font f.Face, text string) int {
	_, adv := f.BoundString(font, text)
	return int(adv >> 6)
}

func GetTextHeight(font f.Face, text string) int {
	bounds, _ := f.BoundString(font, text)
	return int((bounds.Max.Y - bounds.Min.Y) >> 6)
}
