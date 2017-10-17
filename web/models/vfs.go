package models

import (
	"github.com/jackc/pgx/pgtype"
	"cc-fileshare/web/db"
	"cc-fileshare/web/forms"
)

type VFS struct {
	Userid pgtype.Text
	Fileid pgtype.UUID
	Filename pgtype.Text
	Creation_date pgtype.Date
	Attributes pgtype.JSONB
}

type VFSModel struct {}


func (m VFSModel) Insert(vfs *VFS) error {
	if _, err := db.GetPool().Exec(`
	insert into vfs (userid, fileid, filename, creation_date, attributes) values($1, $2, $3, $4, $5)
	`, vfs.Userid.String, vfs.Fileid.Bytes, vfs.Filename.String, vfs.Creation_date.Time, vfs.Attributes.Bytes); err == nil{
		return nil
	} else{
		return err
	}
}

func (m VFSModel) GetById(fileId pgtype.UUID) (*VFS, error) {
	var vfs = new(VFS)
	err := db.GetPool().QueryRow("select userid, fileId, filename, creation_date, attributes from vfs where fileId=$1", fileId).Scan(*vfs)
	if err != nil{
		return nil, err
	} else {
		return vfs, nil
	}
}

func (m VFSModel) SearchByAttr(jsonb pgtype.JSONB) (*forms.VFSResponse, error) {
	var vfs = new(VFS)
	err := db.GetPool().QueryRow("select userid, fileId, filename, creation_date, attributes from vfs where attributes @> $1", jsonb.Bytes).Scan(&vfs.Userid, &vfs.Fileid, &vfs.Filename, &vfs.Creation_date, &vfs.Attributes)
	if err != nil{
		return nil, err
	} else {
		form := new(forms.VFSResponse)
		vfs.AssignTo(form)
		return form, nil
	}
}

func (vfs *VFS)  AssignTo(form *forms.VFSResponse) {
	vfs.Attributes.AssignTo(&form.Attributes)
	vfs.Creation_date.AssignTo(&form.Creation_date)
	vfs.Filename.AssignTo(&form.Filename)
	vfs.Fileid.AssignTo(&form.Fileid)
	vfs.Userid.AssignTo(&form.Userid)
}