// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/brianykl/customer-feedback/api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// Map the route for creating feedback
	beego.Router("/feedback", &controllers.FeedbackController{}, "post:CreateFeedback")

	// Add more routes as needed
	// For example, a route to retrieve feedback might look like this:
	// beego.Router("/feedback/:id", &controllers.FeedbackController{}, "get:GetFeedback")
}
