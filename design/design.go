package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("test", func() {
	Title("test")
	Description("This is to show the issue on goa new gzip middleware")
	Scheme("http")
	BasePath("/")
})

var _ = Resource("good", func() {
	Files("/good/*filepath", "./swagger")
})

var _ = Resource("bad", func() {
	Files("/bad/*filepath", "./swagger")
})
