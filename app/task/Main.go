package task

import "webmis/core"

type Main struct{ core.Base }

/* 首页 */
func (r Main) Index() {
	r.Print("Cli")
}
