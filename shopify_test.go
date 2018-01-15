package shopify

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {

	Convey("Given a product", t, func() {
		product := &Product{}
		auth := Authenticate()
		store := auth.NewStore("ropa-chida")
		err := store.Create(product)
		So(err, ShouldBeNil)
	})
}
