package gosqlite3

import (
	"database/sql"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"
	"net/url"
	"os"
)

func getMagicHeader() [16]byte {
	return [...]byte{'S', 'Q', 'L', 'i', 't', 'e', ' ', 'f', 'o', 'r', 'm', 'a', 't', ' ', '3', 0}
}

type gosqlite3Driver struct {
	magic                   [16]byte
	pagesize                uint16
	fileformatwrite         byte
	fileformatread          byte
	bytesofunused           byte
	maximumembeddedpayload  byte
	minimumembeddedpayload  byte
	leafpayload             byte
	filechangecounter       uint32
	sizeofdatabaseinpages   uint32
	pagenumberoffirst       uint32
	totalnumberoffreelist   uint32
	schemacookie            uint32
	schemaformat            uint32
	defaultpagecachesize    uint32
	pagenumberoflargestroot uint32
	databasetextencoding    uint32
	userversion             uint32
	flagincremental         uint32
	applicationID           uint32
	reserved                [20]byte
	versionvalidfornumber   uint32
	SQLITEVERSIONNUMBER     uint32
}

func (my *gosqlite3Driver) toByte() []byte {
	ret := make([]byte, 0)
	ret = append(ret, my.magic[:]...)
	ret = binary.LittleEndian.AppendUint16(ret, my.pagesize)
	ret = append(ret, my.fileformatwrite)
	ret = append(ret, my.fileformatread)
	ret = append(ret, my.bytesofunused)
	ret = append(ret, my.maximumembeddedpayload)
	ret = append(ret, my.minimumembeddedpayload)
	ret = append(ret, my.leafpayload)
	ret = binary.LittleEndian.AppendUint32(ret, my.filechangecounter)
	ret = binary.LittleEndian.AppendUint32(ret, my.sizeofdatabaseinpages)
	ret = binary.LittleEndian.AppendUint32(ret, my.pagenumberoffirst)
	ret = binary.LittleEndian.AppendUint32(ret, my.totalnumberoffreelist)
	ret = binary.LittleEndian.AppendUint32(ret, my.schemacookie)
	ret = binary.LittleEndian.AppendUint32(ret, my.schemaformat)
	ret = binary.LittleEndian.AppendUint32(ret, my.defaultpagecachesize)
	ret = binary.LittleEndian.AppendUint32(ret, my.pagenumberoflargestroot)
	ret = binary.LittleEndian.AppendUint32(ret, my.databasetextencoding)
	ret = binary.LittleEndian.AppendUint32(ret, my.userversion)
	ret = binary.LittleEndian.AppendUint32(ret, my.flagincremental)
	ret = binary.LittleEndian.AppendUint32(ret, my.applicationID)
	ret = append(ret, my.reserved[:]...)
	ret = binary.LittleEndian.AppendUint32(ret, my.versionvalidfornumber)
	ret = binary.LittleEndian.AppendUint32(ret, my.SQLITEVERSIONNUMBER)
	return ret
}

func (my *gosqlite3Driver) defaultHeader() {
	my.magic = getMagicHeader()
	my.pagesize = 4096
	my.fileformatwrite = 1
	my.fileformatread = 1
	my.bytesofunused = 0
	my.maximumembeddedpayload = 64
	my.minimumembeddedpayload = 32
	my.leafpayload = 32
	my.filechangecounter = 1
	my.sizeofdatabaseinpages = 0
	my.pagenumberoffirst = 0
	my.totalnumberoffreelist = 0
	my.schemacookie = 0
	my.schemaformat = 0
	my.defaultpagecachesize = 0
	my.pagenumberoflargestroot = 0
	my.databasetextencoding = 1
	my.userversion = 0
	my.flagincremental = 0
	my.applicationID = 0
	for i := range my.reserved {
		my.reserved[i] = 0
	}
	my.versionvalidfornumber = 1
	my.SQLITEVERSIONNUMBER = 3037000
}

// Open implements driver.Driver
func (my *gosqlite3Driver) createDB(filename string) (*os.File, error) {
	f, err := os.Create(filename)
	defer f.Sync()
	if err != nil {
		return nil, err
	}
	my.defaultHeader()
	_, err = f.Write(my.toByte())
	if err != nil {
		return nil, err
	}

	return f, err
}

func (d *gosqlite3Driver) Open(name string) (driver.Conn, error) {
	conn := new(gosqlite3Connection)
	URL, err := url.Parse(name)
	if err != nil {
		return nil, err
	}
	if URL.Scheme != "sqlite" {
		return nil, fmt.Errorf("this no sqlite URI")
	}
	fstat, err := os.Stat(URL.Host)
	if errors.Is(err, os.ErrNotExist) {
		d.createDB(URL.Host)
	}
	if err != nil {
		return nil, err
	}
	if fstat.IsDir() {
		return nil, fmt.Errorf("database is dir")
	}
	hFile, err := os.OpenFile(URL.Host, os.O_RDWR|os.O_EXCL, 0755)
	if err != nil {
		return conn, err
	}
	conn.hFile = hFile
	return conn, err
}

func init() {
	//log.Print("init gosqlite3Driver\n")
	sql.Register("gosqlite3", &gosqlite3Driver{})
}
