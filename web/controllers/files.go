package controllers

import (
	"github.com/gin-gonic/gin"
	"cc-fileshare/web/forms"
	"cc-fileshare/web/models"
	"github.com/satori/go.uuid"
	"time"
	"log"
	"github.com/jackc/pgx/pgtype"
)

type FilesController struct{}

type File interface {
	CreateFile(c *gin.Context)
	GetFile(c *gin.Context)
	UpdateFile(c *gin.Context)
	RemoveFile(c *gin.Context)
	Compress(c *gin.Context)
}

func (ctrl FilesController) AddFile(c *gin.Context) {
	var form forms.VFSForm
	var model = new (models.VFSModel)
	if err := c.Bind(&form); err == nil {
		var vfs = new(models.VFS)
		vfs.Creation_date.Set(time.Now())
		vfs.Userid.Set("1")
		vfs.Attributes.Set(form.Attributes)
		vfs.Filename.Set(form.Filename)
		vfs.Fileid.Set(uuid.NewV4().Bytes())
		if err := model.Insert(vfs); err != nil{
			log.Fatal(err)
		}

	}else{
		log.Fatal(err)
	}
}

func (ctrl FilesController) SearchFile(c *gin.Context) {
	var form forms.VFSForm
	var model = new (models.VFSModel)
	if err := c.Bind(&form); err == nil {
		var jsonb pgtype.JSONB
		jsonb.Set(form.Attributes)
		if row, err := model.SearchByAttr(jsonb); err == nil{
			c.JSON(200, row)
		}else{
			log.Fatal(err)
		}

	}else{
		log.Fatal(err)
	}
}

func (ctrl FilesController) CacheSearch(c *gin.Context) {
	var form forms.VFSForm
	var model = new (models.VFSModel)
	if err := c.Bind(&form); err == nil {
		var jsonb pgtype.JSONB
		jsonb.Set(form.Attributes)
		if row, err := model.SearchByAttr(jsonb); err == nil{
			c.JSON(200, row)
		}else{
			log.Fatal(err)
		}

	}else{
		log.Fatal(err)
	}
}
