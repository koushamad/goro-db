package appProvider

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-db/app/database/sql"
)

func Boot(egn *gin.Engine) {
	sql.Boot()
}

func Kill()  {

}
