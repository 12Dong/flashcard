package routers

import (
	"flashcard/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/flashcard", &controllers.FlashcardController{})
}
