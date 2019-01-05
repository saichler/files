package model

import (
	"errors"
	. "github.com/saichler/utils/golang"
	"io/ioutil"
	"strconv"
)

type FileInfo struct {
	Name string
	Size int64
	IsDir bool
	Files []*FileInfo
}

func NewDirInfo(path string) (*FileInfo,error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil,errors.New("No Such Path:"+path)
	}

	fi:=&FileInfo{}
	fi.Name = path
	fi.IsDir = true

	if files!=nil && len(files)>0 {
		fi.Files=make([]*FileInfo,0)
		for _, f := range files {
			if f.IsDir() {
				file,err:=NewDirInfo(fi.Name+"/"+f.Name())
				if err==nil {
					fi.Files = append(fi.Files, file)
					fi.IsDir = true
				}
			} else {
				file :=&FileInfo{}
				file.Name = f.Name()
				file.Size = f.Size()
				file.IsDir = false
				fi.Files = append(fi.Files,file)
			}
		}
	}

	return fi,nil
}

func (fi *FileInfo) Print() {
	fi.print(0)
}

func (fi *FileInfo) print(lvl int) {
	if fi==nil {
		return
	}
	space:=""
	for i:=0;i<lvl;i++{
		space+="   "
	}
	if fi.IsDir {
		Info(space + fi.Name)
	} else {
		Info(space + fi.getSize())
	}
	if fi.IsDir {
		for _,f:=range fi.Files {
			f.print(lvl+1)
		}
	}
}

func (fi *FileInfo) getSize() string {
	kb:=fi.Size/1024
	if int(kb)==0 {
		return fi.Name+" "+strconv.Itoa(int(fi.Size))+" by"
	}
	mb:=kb/1024
	if int(mb)==0 {
		return fi.Name + " " + strconv.Itoa(int(kb)) + " kb"
	}
	return fi.Name + " " + strconv.Itoa(int(mb)) + " mb"
}

func (fi *FileInfo) Bytes() []byte {
	ba := NewByteSlice()
	fi.bytes(ba)
	return ba.Data()
}

func (fi *FileInfo) bytes(ba *ByteSlice) {
	ba.AddString(fi.Name)
	ba.AddInt64(fi.Size)
	ba.AddBool(fi.IsDir)
	if fi.IsDir {
		if fi.Files == nil {
			ba.AddInt(0)
		} else {
			ba.AddInt(len(fi.Files))
			for _, file := range fi.Files {
				file.bytes(ba)
			}
		}
	}
}

func FromBytes(ba *ByteSlice) *FileInfo {
	fi:=&FileInfo{}
	fi.Name=ba.GetString()
	fi.Size=ba.GetInt64()
	fi.IsDir=ba.GetBool()
	if fi.IsDir {
		size:=ba.GetInt()
		if size>0 {
			fi.Files=make([]*FileInfo,size)
			for i:=0;i<size;i++ {
				fi.Files[i]=FromBytes(ba)
			}
		}
	}
	return fi
}