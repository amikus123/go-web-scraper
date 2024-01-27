package headless

type HeadlessSelectors struct {
	Main string
	Href string
	Img  string
	Text string
}

var FrondaSelectors = HeadlessSelectors{
	Main: ".itemBox",
	Href: "a",
	Img:  "img",
	Text: ".itemBox-title",
}
