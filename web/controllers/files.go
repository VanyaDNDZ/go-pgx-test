package controllers

import (
	"github.com/gin-gonic/gin"
	"cc-fileshare/web/forms"
	"cc-fileshare/web/models"
	"github.com/satori/go.uuid"
	"time"
	"log"
	"github.com/jackc/pgx/pgtype"
	"fmt"
	"net/http"
	"os"
	"github.com/nfnt/resize"
	"image"
	"io/ioutil"
	"image/png"
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
	var model = new(models.VFSModel)
	if err := c.Bind(&form); err == nil {

		var vfs = new(models.VFS)
		vfs.Creation_date.Set(time.Now())
		vfs.Userid.Set("1")
		vfs.Attributes.Set(form.Attributes)
		vfs.Filename.Set(form.Filename)
		vfs.Fileid.Set(uuid.NewV4().Bytes())

		if err := model.Insert(vfs); err != nil {
			log.Fatal(err)
		}
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		filename := new(string)
		vfs.Fileid.AssignTo(filename)
		if err := c.SaveUploadedFile(file, *filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

	} else {
		log.Fatal(err)
	}
}

func (ctrl FilesController) SearchFile(c *gin.Context) {
	var form forms.VFSForm
	var model = new(models.VFSModel)
	if err := c.Bind(&form); err == nil {
		var jsonb pgtype.JSONB
		jsonb.Set(form.Attributes)
		if row, err := model.SearchByAttr(jsonb); err == nil {
			c.JSON(200, row)
		} else {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}
}

func (ctrl FilesController) UpdateAttributes(c *gin.Context) {
	var inputForm forms.VFSForm
	if err := c.Bind(&inputForm); err == nil {
		var model = new(models.VFSModel)
		var fileid = new(pgtype.UUID)
		fileid.Set(c.Params.ByName("fileid"))
		jsonb := new(pgtype.JSONB)
		jsonb.Set(inputForm.Attributes)
		if vfs, err := model.UpdateAttributes(fileid, jsonb); err == nil {
			c.JSON(200, vfs)
		} else {
			c.Error(err)
		}
	} else {
		fmt.Printf("Wrong parameters: %v", err)
		c.String(422, ("Wrong parameters"))
	}
}

func (ctrl FilesController) GetFileInfo(c *gin.Context) {
	var model = new(models.VFSModel)
	var fileid = new(pgtype.UUID)
	fileid.Set(c.Params.ByName("fileid"))
	if vfs, err := model.GetById(fileid); err == nil {
		c.JSON(200, vfs)
	} else {
		c.Error(err)
	}
}

func (ctrl FilesController) GetFile(c *gin.Context) {
	var model = new(models.VFSModel)
	var fileid = new(pgtype.UUID)
	fileid.Set(c.Params.ByName("fileid"))
	if vfs, err := model.GetById(fileid); err == nil {
		var resizeForm forms.VFSResizeForm
		if err := c.Bind(&resizeForm); err == nil {

			img := resizeImage(vfs.Fileid, resizeForm.Width, resizeForm.Height)
			tmpfile, err := ioutil.TempFile(".", "tmp-"+vfs.Fileid)
			if err != nil {
				log.Fatal(err)
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			err = png.Encode(tmpfile, img)
			tmpfile.Close()
			c.File(tmpfile.Name())
			defer os.Remove(tmpfile.Name())
		} else {
			c.File(vfs.Fileid)
		}
	} else {
		c.Error(err)
	}
}

func resizeImage(originFileName string, width, height uint) image.Image {
	file, err := os.Open(originFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return resize.Resize(width, height, img, resize.Lanczos3)
}
