package service

import (
	"webmis/app/config"
	"webmis/app/util"
	"webmis/core"
)

/* 数据类 */
type Data struct {
	core.Base
}

/* 分区时间 */
func (*Data) Partition() map[string]int {
	partition := map[string]int{
		"p2601": 1769875200,
		"p2602": 1772294400,
		"p2603": 1774972800,
		"p2604": 1777564800,
		"p2605": 1780243200,
		"p2606": 1782835200,
		"p2607": 1785513600,
		"p2608": 1788192000,
		"p2609": 1790784000,
		"p2610": 1793462400,
		"p2611": 1796054400,
		"p2612": 1798732800,
		"plast": 1798732800,
	}
	return partition
}

/* 图片地址 */
func (*Data) Img(img string, isTmp bool) string {
	if img == "" {
		return ""
	}
	if isTmp {
		return config.Env().Img_url + img
	} else {
		return config.Env().Img_url + img + "?" + util.Str(util.Time())
	}
}

/* 图片地址-商品 */
func (d *Data) ImgGoods(sku_id string, isTmp bool) string {
	if sku_id == "" {
		return ""
	}
	return d.Img("img/sku/"+sku_id+".jpg", isTmp)
}
