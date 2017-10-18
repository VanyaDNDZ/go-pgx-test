package models

import (
	"github.com/jackc/pgx/pgtype"
	"cc-fileshare/web/db"
	"cc-fileshare/web/forms"
	"github.com/jackc/pgx"
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

func (m *VFSModel) Scan(row *pgx.Row) (*VFS, error)  {
	var vfs = new(VFS)
	if err:= row.Scan(&vfs.Userid, &vfs.Fileid, &vfs.Filename, &vfs.Creation_date, &vfs.Attributes);  err == nil{
		return vfs, nil
	}else{
		return nil, err
	}

}

func (m VFSModel) GetById(fileId *pgtype.UUID) (*forms.VFSResponse, error) {

	row := db.GetPool().QueryRow("select userid, fileId, filename, creation_date, attributes from vfs where fileId=$1", fileId)
	vfs, err := m.Scan(row)
	if err != nil{
		return nil, err
	} else {
		form := new(forms.VFSResponse)
		vfs.AssignTo(form)
		return form, nil
	}
}

func (m VFSModel) SearchByAttr(jsonb pgtype.JSONB) (*forms.VFSResponse, error) {
	row := db.GetPool().QueryRow("select userid, fileId, filename, creation_date, attributes from vfs where attributes @> $1", jsonb.Bytes)
	vfs, err := m.Scan(row)
	if err != nil{
		return nil, err
	} else {
		form := new(forms.VFSResponse)
		vfs.AssignTo(form)
		return form, nil
	}
}

func (m VFSModel) UpdateAttributes(fileId *pgtype.UUID, attributes *pgtype.JSONB) (*forms.VFSResponse, error) {
	if _, err := db.GetPool().Exec(`
	update vfs
	set attributes = attributes || $2
	where fileid = $1
	`, fileId, attributes); err == nil{
		return m.GetById(fileId)
	} else{
		return nil, err
	}
}

func (vfs *VFS)  AssignTo(form *forms.VFSResponse) {
	vfs.Attributes.AssignTo(&form.Attributes)
	vfs.Creation_date.AssignTo(&form.Creation_date)
	vfs.Filename.AssignTo(&form.Filename)
	vfs.Fileid.AssignTo(&form.Fileid)
	vfs.Userid.AssignTo(&form.Userid)
}