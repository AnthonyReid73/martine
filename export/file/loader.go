package file

import (
	"encoding/binary"
	"fmt"
	"github.com/jeromelesaux/m4client/cpc"
	"github.com/jeromelesaux/martine/constants"
	x "github.com/jeromelesaux/martine/export"
	"image/color"
	"os"
	"path/filepath"
	"strings"
)

// CPC plus loader nb colors *2 offset 0x9d

var (
	BasicLoaderBasic = []byte{
		0x36, 0x00, 0x05, 0x00, 0x8c, 0x20, 0x30, 0x30,
		0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30,
		0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c,
		0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30,
		0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30,
		0x30, 0x2c, 0x30, 0x30, 0x2c, 0x30, 0x30, 0x2c,
		0x30, 0x30, 0x2c, 0x30, 0x30, 0x00, 0x0e, 0x00,
		0x0a, 0x00, 0xaa, 0x20, 0x1c, 0x00, 0x40, 0x20,
		0xf5, 0x20, 0x0f, 0x00, 0x18, 0x00, 0x14, 0x00,
		0xa8, 0x22, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x2e, 0x70, 0x61, 0x6c, 0x22, 0x2c,
		0x1c, 0x00, 0x40, 0x00, 0x0e, 0x00, 0x1e, 0x00,
		0xad, 0x20, 0xff, 0x12, 0x28, 0x1c, 0x00, 0x40,
		0x29, 0x00, 0x13, 0x00, 0x28, 0x00, 0x9e, 0x20,
		0x0d, 0x00, 0x00, 0xf0, 0xef, 0x0e, 0x20, 0xec,
		0x20, 0x19, 0x0f, 0x20, 0x00, 0x0b, 0x00, 0x32,
		0x00, 0xc3, 0x20, 0x0d, 0x00, 0x00, 0xe3, 0x00,
		0x10, 0x00, 0x46, 0x00, 0xa2, 0x20, 0x0d, 0x00,
		0x00, 0xf0, 0x2c, 0x0d, 0x00, 0x00, 0xe3, 0x00,
		0x07, 0x00, 0x50, 0x00, 0xb0, 0x20, 0x00, 0x18,
		0x00, 0x5a, 0x00, 0xa8, 0x22, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x20, 0x2e, 0x73, 0x63,
		0x72, 0x22, 0x2c, 0x1c, 0x00, 0xc0, 0x00, 0x00,
		0x00, 0x0a}
	startPaletteValues   = 6
	startPaletteName     = 58 + 16
	startScreenName      = 149 + 16
	PaletteCPCPlusLoader = []byte{
		0x00, 0x50, 0x41, 0x4C, 0x50, 0x4C, 0x55, 0x53,
		0x20, 0x42, 0x49, 0x4E, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x3E, 0x00, 0x00, 0x30, 0x00,
		0x3E, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x3E, 0x00, 0x00, 0x36, 0x04, 0x00, 0x00, 0x50,
		0x41, 0x4C, 0x50, 0x4C, 0x55, 0x53, 0x20, 0x24,
		0x24, 0x24, 0xFF, 0x00, 0xFF, 0x00, 0x00, 0x02,
		0x20, 0x67, 0x65, 0x6E, 0x65, 0x72, 0x61, 0x74,
		0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x52, 0x41,
		0x53, 0x4D, 0x20, 0x76, 0x30, 0x2E, 0x31, 0x31,
		0x33, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xF3, 0x01, 0x00, 0xBC, 0x21, 0x2D, 0x30, 0x1E,
		0x11, 0x7E, 0xED, 0x79, 0x23, 0x1D, 0x20, 0xF9,
		0xFB, 0x01, 0xB8, 0x7F, 0xED, 0x49, 0x21, 0x3E,
		0x30, 0x11, 0x00, 0x64, 0x01, 0x20, 0x00, 0xED,
		0xB0, 0x21, 0xF9, 0xB7, 0xC3, 0xDD, 0xBC, 0x01,
		0xA0, 0x7F, 0xED, 0x49, 0xC9, 0xFF, 0x00, 0xFF,
		0x77, 0xB3, 0x51, 0xA8, 0xD4, 0x62, 0x39, 0x9C,
		0x46, 0x2B, 0x15, 0x8A, 0xCD, 0xEE}
	// offset file name 24
	startScreenPlusName     = 38
	BasicCPCPlusLoaderBasic = []byte{
		0x3f, 0x00, 0x0a, 0x00, 0xaa, 0x20, 0x1c, 0xff,
		0x2f, 0x01, 0x20, 0xad, 0x20, 0x10, 0x01, 0x20,
		0xa8, 0x22, 0x70, 0x61, 0x6c, 0x70, 0x6c, 0x75,
		0x73, 0x2e, 0x62, 0x69, 0x6e, 0x22, 0x2c, 0x1c,
		0x00, 0x30, 0x01, 0x20, 0xa8, 0x22, 0x74, 0x69,
		0x67, 0x72, 0x65, 0x20, 0x20, 0x20, 0x2e, 0x73,
		0x63, 0x72, 0x22, 0x2c, 0x1c, 0x00, 0xc0, 0x01,
		0x20, 0x83, 0x20, 0x1c, 0x00, 0x30, 0x00, 0x00,
		0x00, 0x1a, 0x00}

	FlashBasicLoader = []byte{
		0x08, 0x00, 0x0a, 0x00, 0xad, 0x20, 0x0f, 0x00,
		0x0e, 0x00, 0x14, 0x00, 0xaa, 0x20, 0x1c, 0x00,
		0x30, 0x20, 0xf5, 0x20, 0x0f, 0x00, 0x18, 0x00,
		0x1e, 0x00, 0xa8, 0x22, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x2e, 0x73, 0x63, 0x72,
		0x22, 0x2c, 0x1c, 0x00, 0x40, 0x00, 0x18, 0x00,
		0x28, 0x00, 0xa8, 0x22, 0x20, 0x20, 0x20, 0x20,
		0x20, 0x20, 0x20, 0x20, 0x2e, 0x73, 0x63, 0x72,
		0x22, 0x2c, 0x1c, 0x00, 0xc0, 0x00, 0x15, 0x00,
		0x32, 0x00, 0xa8, 0x22, 0x66, 0x6c, 0x61, 0x73,
		0x68, 0x2e, 0x62, 0x69, 0x6e, 0x22, 0x2c, 0x1c,
		0x00, 0x30, 0x00, 0x0a, 0x00, 0x3c, 0x00, 0x83,
		0x20, 0x1c, 0x00, 0x30, 0x00, 0x00, 0x00, 0x0a,
		0x1a, 0x20, 0x1c, 0x00, 0x30, 0x00, 0x00, 0x00,
		0x1A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}

	flashScreen1Offset = 28
	flashScreen2Offset = 52

	FlashBinaryLoader = []byte{
		0xF3, 0x2A, 0x38, 0x00, 0x22, 0x76, 0x30, 0x21,
		0xFB, 0xC9, 0x22, 0x38, 0x00, 0xD9, 0x08, 0xF5,
		0xC5, 0xD5, 0xE5, 0xFB, 0x06, 0xF5, 0xED, 0x78,
		0x1F, 0x30, 0xFB, 0x06, 0x7F, 0x3E, 0x8D, 0xEE,
		0x01, 0x32, 0x1E, 0x30, 0xED, 0x79, 0x01, 0x0C,
		0xBC, 0xED, 0x49, 0x3E, 0x10, 0xEE, 0x20, 0x32,
		0x2C, 0x30, 0x04, 0xED, 0x79, 0xCB, 0x6F, 0x28,
		0x0A, 0x21, 0xA6, 0x30, 0x3E, 0x10, 0xCD, 0x86,
		0x30, 0x18, 0x07, 0x3E, 0x04, 0x21, 0x96, 0x30,
		0x18, 0xF2, 0x76, 0x01, 0x0E, 0xF4, 0xED, 0x49,
		0x01, 0xC0, 0xF6, 0xED, 0x49, 0xAF, 0xED, 0x79,
		0x01, 0x92, 0xF7, 0xED, 0x49, 0x01, 0x45, 0xF6,
		0xED, 0x49, 0x06, 0xF4, 0xED, 0x78, 0x01, 0x82,
		0xF7, 0xED, 0x49, 0x01, 0x00, 0xF6, 0xED, 0x49,
		0x17, 0xDA, 0x14, 0x30, 0xF3, 0x21, 0x00, 0x00,
		0x22, 0x38, 0x00, 0xE1, 0xD1, 0xC1, 0xF1, 0x08,
		0xD9, 0xFB, 0xC3, 0xA7, 0xBC, 0xC9, 0x57, 0x01,
		0x00, 0x7F, 0xED, 0x49, 0x5E, 0xED, 0x59, 0x0C,
		0x2C, 0x3D, 0x20, 0xF6, 0x7A, 0xC9, 0x54, 0x5E,
		0x40, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4E, 0x4A,
		0x43, 0x5C, 0x54, 0x40, 0x5E, 0x40, 0x4B, 0x56,
		0x44, 0x59, 0x58, 0x46, 0x00, 0x00,
	}
	flashBinaryPalette1LenghtOffset = 68
	flashBinaryPalette2LenghtOffset = 61
	flashModeOffset                 = 30
	flashBinaryPalette1Offset       = 150
	flashBinaryPalette2Offset       = 166
)

func Loader(filePath string, p color.Palette, mode uint8, exportType *x.ExportType) error {
	if exportType.CpcPlus {
		return BasicLoaderCPCPlus(filePath, p, mode, exportType)
	}
	return BasicLoader(filePath, p, exportType)
}

func BasicLoaderCPCPlus(filePath string, p color.Palette, mode uint8, exportType *x.ExportType) error {
	// export de la palette assemblée
	loader := PaletteCPCPlusLoader

	for i := 0; i < len(p); i++ {
		cp := constants.NewCpcPlusColor(p[i])
		b := cp.Bytes()
		loader = append(loader, b[0])
		loader = append(loader, b[1])
	}

	loader[0x9d] = uint8(len(p) * 2)

	paletteHeader, err := cpc.BytesCpcHeader(loader)
	if err != nil {
		return err
	}
	paletteHeader.Size = uint16(binary.Size(loader)) - 128
	paletteHeader.Size2 = uint16(binary.Size(loader)) - 128
	paletteHeader.LogicalSize = uint16(binary.Size(loader)) - 128
	paletteHeader.Checksum = uint16(paletteHeader.ComputedChecksum16())
	data, err := paletteHeader.Bytes()
	if err != nil {
		return err
	}
	copy(loader[0:128], data)

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
	switch mode {
	case 0:
		loader[13] = 0x0e
	case 1:
		loader[13] = 0x0f
	case 2:
		loader[13] = 0x10
	}

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

func BasicLoader(filePath string, p color.Palette, exportType *x.ExportType) error {
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
	filename := exportType.GetAmsdosFilename(filePath, "")
	copy(loader[startPaletteName:], filename[:])
	copy(loader[startScreenName:], filename[:])
	fmt.Println(loader)
	header := cpc.CpcHead{Type: 0, User: 0, Address: 0x170, Exec: 0x0,
		Size:        uint16(binary.Size(loader)),
		Size2:       uint16(binary.Size(loader)),
		LogicalSize: uint16(binary.Size(loader))}
	file := exportType.GetAmsdosFilename(string(filename), ".BAS")
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

func FlashLoader(screenFilename1, screenFilename2 string, p1, p2 color.Palette, m1, m2 uint8, exportType *x.ExportType) error {
	// modification du binaire flash
	pal1 := make([]byte, 16)
	pal2 := make([]byte, 16)
	mode1 := m2
	mode2 := m1
	pl1 := p2
	pl2 := p1

	if m1 > m2 {
		pl1 = p1
		pl2 = p2
		mode1 = m1
		mode2 = m2
	}
	var l1, l2 uint8
	for i := 0; i < len(pl1); i++ {
		v, err := constants.HardwareValues(pl1[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while getting the hardware values for color %v, error :%v\n", pl1[i], err)
		} else {
			l1++
		}
		pal1[i] = v[0]
	}
	for i := 0; i < len(pl2); i++ {
		v, err := constants.HardwareValues(pl2[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while getting the hardware values for color %v, error :%v\n", pl2[i], err)
		} else {
			l2++
		}
		pal2[i] = v[0]
	}

	var flashLoader []byte
	flashLoader = FlashBinaryLoader
	copy(flashLoader[flashBinaryPalette1Offset:], pal1)
	copy(flashLoader[flashBinaryPalette2Offset:], pal2)

	switch mode2 {
	case 0:
		flashLoader[flashModeOffset] = 0x8d
		flashLoader[flashBinaryPalette2LenghtOffset] = 0x10
	case 1:
		flashLoader[flashModeOffset] = 0x8c
		flashLoader[flashBinaryPalette2LenghtOffset] = 0x04
	case 2:
		flashLoader[flashModeOffset] = 0x8e
		flashLoader[flashBinaryPalette2LenghtOffset] = 0x02
	}

	switch mode1 {
	case 0:
		flashLoader[flashBinaryPalette1LenghtOffset] = 0x10
	case 1:
		flashLoader[flashBinaryPalette1LenghtOffset] = 0x04
	case 2:
		flashLoader[flashBinaryPalette1LenghtOffset] = 0x02
	}

	binaryHeader := cpc.CpcHead{Type: 2, User: 0, Address: 0x3000, Exec: 0x3000,
		Size:        uint16(binary.Size(flashLoader)),
		Size2:       uint16(binary.Size(flashLoader)),
		LogicalSize: uint16(binary.Size(flashLoader))}
	copy(binaryHeader.Filename[:], "FLASH   BIN")
	binaryHeader.Checksum = uint16(binaryHeader.ComputedChecksum16())
	flashBinPath := exportType.OutputPath + string(filepath.Separator) + "FLASH.BIN"
	fw, err := os.Create(flashBinPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating file (%s) error :%s\n", flashBinPath, err)
		return err
	}
	if !exportType.NoAmsdosHeader {
		binary.Write(fw, binary.LittleEndian, binaryHeader)
	}
	binary.Write(fw, binary.LittleEndian, flashLoader)
	fw.Close()

	exportType.AddFile(flashBinPath)

	// modification du flash loader en basic
	var basicLoader []byte
	basicLoader = FlashBasicLoader
	copy(basicLoader[flashScreen1Offset:], strings.ToUpper(screenFilename1[0:len(screenFilename1)-4]))
	copy(basicLoader[flashScreen2Offset:], strings.ToUpper(screenFilename2[0:len(screenFilename2)-4]))
	basicHeader := cpc.CpcHead{Type: 0, User: 0, Address: 0x170, Exec: 0x0,
		Size:        uint16(binary.Size(basicLoader)),
		Size2:       uint16(binary.Size(basicLoader)),
		LogicalSize: uint16(binary.Size(basicLoader))}
	copy(binaryHeader.Filename[:], "-SWITCH.BAS")
	basicHeader.Checksum = uint16(basicHeader.ComputedChecksum16())
	basicPath := exportType.OutputPath + string(filepath.Separator) + "-SWITCH.BAS"
	fw2, err := os.Create(basicPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating file (%s) error :%s\n", basicPath, err)
		return err
	}
	if !exportType.NoAmsdosHeader {
		binary.Write(fw2, binary.LittleEndian, basicHeader)
	}
	binary.Write(fw2, binary.LittleEndian, basicLoader)
	fw2.Close()

	exportType.AddFile(basicPath)
	return nil
}
