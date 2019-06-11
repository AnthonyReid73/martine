package gfx

import (
	"encoding/binary"
	"fmt"
	"github.com/jeromelesaux/m4client/cpc"
	"github.com/jeromelesaux/martine/constants"
	"image/color"
	"os"
)

// CPC plus loader nb colors *2 offset 0x9d

var (
	BasicLoaderBasic = []byte{
		0x36, 0x00, 0x05, 0x00, 0x8c, 0x20, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30,
		0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30,
		0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c,
		0x30, 0x30, 0x2c, 0x30, 0x30, 0x00, 0x0e, 0x00, 0x0a, 0x00, 0xaa, 0x20, 0x1c, 0x00, 0x40, 0x20,
		0xf5, 0x20, 0x0f, 0x00, 0x18, 0x00, 0x14, 0x00, 0xa8, 0x22, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x2e, 0x70, 0x61, 0x6c, 0x22, 0x2c, 0x1c, 0x00, 0x40, 0x00, 0x0e, 0x00, 0x1e, 0x00,
		0xad, 0x20, 0xff, 0x12, 0x28, 0x1c, 0x00, 0x40, 0x29, 0x00, 0x13, 0x00, 0x28, 0x00, 0x9e, 0x20,
		0x0d, 0x00, 0x00, 0xf0, 0xef, 0x0e, 0x20, 0xec, 0x20, 0x19, 0x0f, 0x20, 0x00, 0x0b, 0x00, 0x32,
		0x00, 0xc3, 0x20, 0x0d, 0x00, 0x00, 0xe3, 0x00, 0x10, 0x00, 0x46, 0x00, 0xa2, 0x20, 0x0d, 0x00,
		0x00, 0xf0, 0x2c, 0x0d, 0x00, 0x00, 0xe3, 0x00, 0x07, 0x00, 0x50, 0x00, 0xb0, 0x20, 0x00, 0x18,
		0x00, 0x5a, 0x00, 0xa8, 0x22, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x73, 0x63,
		0x72, 0x22, 0x2c, 0x1c, 0x00, 0xc0, 0x00, 0x00, 0x00, 0x0a}
	startPaletteValues   = 6
	startPaletteName     = 58 + 16
	startScreenName      = 149 + 16
	PaletteCPCPlusLoader = []byte{
		0x00, 0x50, 0x41, 0x4C, 0x50, 0x4C, 0x55, 0x53, 0x20, 0x42, 0x49, 0x4E,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x3E, 0x00, 0x00, 0x30, 0x00,
		0x3E, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 
		0x00, 0x00, 0x00, 0x00, 0x3E, 0x00, 0x00, 0x36, 0x04, 0x00, 0x00, 0x50, 
		0x41, 0x4C, 0x50, 0x4C, 0x55, 0x53, 0x20, 0x24, 0x24, 0x24, 0xFF, 0x00, 
		0xFF, 0x00, 0x00, 0x02, 0x20, 0x67, 0x65, 0x6E, 0x65, 0x72, 0x61, 0x74, 
		0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x52, 0x41, 0x53, 0x4D, 0x20, 0x76, 
		0x30, 0x2E, 0x31, 0x31, 0x33, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xF3, 0x01, 0x00, 0xBC, 
		0x21, 0x2D, 0x30, 0x1E, 0x11, 0x7E, 0xED, 0x79, 0x23, 0x1D, 0x20, 0xF9,
		0xFB, 0x01, 0xB8, 0x7F, 0xED, 0x49, 0x21, 0x3E, 0x30, 0x11, 0x00, 0x64, 
		0x01, 0x20, 0x00, 0xED, 0xB0, 0x21, 0xF9, 0xB7, 0xC3, 0xDD, 0xBC, 0x01, 
		0xA0, 0x7F, 0xED, 0x49, 0xC9, 0xFF, 0x00, 0xFF, 0x77, 0xB3, 0x51, 0xA8, 
		0xD4, 0x62, 0x39, 0x9C,	0x46, 0x2B, 0x15, 0x8A, 0xCD, 0xEE}
	// offset file name 24
	startScreenPlusName     = 24
	BasicCPCPlusLoaderBasic = []byte{
		0x1a, 0x38, 0x00, 0x0a, 0x00, 0xaa, 0x20, 0x1c, 0xff, 0x2f, 0x01, 0xad, 0x20, 0x0e, 0x01, 0xa8,
		0x22, 0x70, 0x61, 0x6c, 0x70, 0x6c, 0x75, 0x73, 0x2e, 0x62, 0x69, 0x6e, 0x22, 0x2c, 0x1c, 0x00,
		0x30, 0x01, 0xa8, 0x22, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x73, 0x63, 0x72,
		0x22, 0x2c, 0x1c, 0x00, 0xc0, 0x01, 0x83, 0x20, 0x1c, 0x00, 0x30, 0x00, 0x00, 0x00, 0x1a}
)

func Loader(filePath string, p color.Palette, mode uint8, exportType *ExportType) error {
	if exportType.CpcPlus {
		return BasicLoaderCPCPlus(filePath, p, mode, exportType)
	}
	return BasicLoader(filePath, p, exportType)
}

func BasicLoaderCPCPlus(filePath string, p color.Palette, mode uint8, exportType *ExportType) error {
	// export de la palette assemblée
	loader := PaletteCPCPlusLoader

	for i := 0; i < len(p); i++ {
		cp := NewCpcPlusColor(p[i])
		b := cp.Bytes()
		loader = append(loader, b[1])
		loader = append(loader, b[0])
	}
 
	loader[0x9d] = uint8(len(p) * 2)
	
	paletteHeader,err := cpc.BytesCpcHeader(loader)
	if err != nil {
		return err
	}
	paletteHeader.Size = uint16(binary.Size(loader))
	paletteHeader.Size2 = uint16(binary.Size(loader))
	paletteHeader.LogicalSize = uint16(binary.Size(loader))
	paletteHeader.Checksum = uint16(paletteHeader.ComputedChecksum16())
	data, err := paletteHeader.Bytes()
	if err != nil {
		return err
	}
	copy(loader[0:128],data)


	osFilepath := exportType.AmsdosFullPath("PALPLUS", ".BIN")
	// modifier checksum amsdos header
	fw, err := os.Create(osFilepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating file (%s) error :%s\n", osFilepath, err)
		return err
	}
	binary.Write(fw, binary.LittleEndian, loader)
	fw.Close()

	exportType.AddFile(osFilepath)

	// export fichier basic loader
	loader = BasicCPCPlusLoaderBasic
	filename := exportType.AmsdosFilename()
	copy(loader[startScreenPlusName:], filename[:])
	loader[0x1d] = mode
	fmt.Println(loader)
	header := cpc.CpcHead{Type: 0, User: 0, Address: 0x170, Exec: 0x0,
		Size:        uint16(binary.Size(loader)),
		Size2:       uint16(binary.Size(loader)),
		LogicalSize: uint16(binary.Size(loader))}
	file := string(filename) + ".BAS"
	copy(header.Filename[:], file)
	header.Checksum = uint16(header.ComputedChecksum16())
	osFilepath = exportType.AmsdosFullPath(filePath, ".BAS")
	fw, err = os.Create(osFilepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating file (%s) error :%s\n", osFilepath, err)
		return err
	}
	if !exportType.NoAmsdosHeader {
		binary.Write(fw, binary.LittleEndian, header)
	}
	binary.Write(fw, binary.LittleEndian, loader)
	fw.Close()

	exportType.AddFile(osFilepath)

	return nil
}

func BasicLoader(filePath string, p color.Palette, exportType *ExportType) error {
	var out string
	for i := 0; i < len(p); i++ {
		v, err := constants.FirmwareNumber(p[i])
		if err == nil {
			out += fmt.Sprintf("%0.2d", v)
		} else {
			fmt.Fprintf(os.Stderr, "Error while getting the hardware values for color %v, error :%v\n", p[0], err)
		}
		if i+1 < len(p) {
			out += ","
		}

	}

	var loader []byte
	loader = BasicLoaderBasic
	copy(loader[startPaletteValues:], out[0:len(out)])
	filename := exportType.AmsdosFilename()
	copy(loader[startPaletteName:], filename[:])
	copy(loader[startScreenName:], filename[:])
	fmt.Println(loader)
	header := cpc.CpcHead{Type: 0, User: 0, Address: 0x170, Exec: 0x0,
		Size:        uint16(binary.Size(loader)),
		Size2:       uint16(binary.Size(loader)),
		LogicalSize: uint16(binary.Size(loader))}
	file := string(filename) + ".BAS"
	copy(header.Filename[:], file)
	header.Checksum = uint16(header.ComputedChecksum16())
	osFilepath := exportType.AmsdosFullPath(filePath, ".BAS")
	fw, err := os.Create(osFilepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating file (%s) error :%s\n", osFilepath, err)
		return err
	}
	if !exportType.NoAmsdosHeader {
		binary.Write(fw, binary.LittleEndian, header)
	}
	binary.Write(fw, binary.LittleEndian, loader)
	fw.Close()

	exportType.AddFile(osFilepath)
	return nil
}
