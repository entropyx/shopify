package shopify

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {

	Convey("Given a product", t, func() {
		product := &Product{
			Title:       "test product",
			BodyHTML:    "test body",
			Vendor:      "Nike",
			ProductType: "test product type",
			Tags:        `Barnes & Noble, John's Fav, "Big Air"`,
		}
		auth := Authenticate()
		store := auth.NewStore("linduritas")
		err := store.Create(product)
		So(err, ShouldBeNil)
	})
}
