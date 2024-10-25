package vndb

var path = struct {
	stats       string
	user        string
	authInfo    string
	vn          string
	release     string
	producer    string
	character   string
	staff       string
	tag         string
	trait       string
	uList       string
	rList       string
	uListLabels string
}{
	stats:       "/stats",
	user:        "/user",
	authInfo:    "/authinfo",
	vn:          "/vn",
	release:     "/release",
	producer:    "/producer",
	character:   "/character",
	staff:       "/staff",
	tag:         "/tag",
	trait:       "/trait",
	uList:       "/ulist",
	rList:       "/rlist",
	uListLabels: "/ulist_labels",
}
