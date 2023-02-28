package main

//#include <../c/source.h>
//
// DX_FILE_HEAD Head;
// LORA_REC g_LoraRec;
// DXK_REC g_DxkRec;
// MQTTCLIENT_REC g_MqttClientRec;
// JDX_REC *g_JdxRecArray = NULL;
// int g_JdxRecArraySize = 0;
// DXZREC *g_DxzRecArray=NULL;
// int g_DxzRecArraySize = 0;
//void init_lora_para( )
// {
// 	 memset( &g_LoraRec,0,sizeof(LORA_REC) );
// 	 g_LoraRec.m_bsptype = 0;//bsp类型:0-A40I,1:ESM6800
// 	 g_LoraRec.m_ServerPort = 1883;//默认=1883
// 	 g_LoraRec.m_nRes1 = 0;
// 	 g_LoraRec.m_nRes2 = 0;
// 	 strcpy( g_LoraRec.m_ServerIP,"localhost");
// 	 strcpy( g_LoraRec.m_ApplicationID,"ApplicationID");//设备标识
// 	 strcpy( g_LoraRec.m_szRes1,"res1");//设备类型
// 	 strcpy( g_LoraRec.m_szRes2,"res2");
// 	 strcpy( g_LoraRec.m_szRes3,"res3");
// 	 strcpy( g_LoraRec.m_szRes4,"res4");
// }
//  int read_lora_para_file()
//  {
//          int nReadSize;
//          int nResult = 0;
//          char szFileName[256];
//          // sprintf(szFileName,"%s/dx/%s",A40I_CCU_CONF_HBDX, DX_LORA_FILENAME);
//          sprintf(szFileName, "%s", "/home/tiger/Desktop/code/vscode/demo/c/lorapara.dat");
//          char *szFile = szFileName;
//          FILE *fp = fopen(szFile, "rb");
//          if (!fp)
//          {
//                 return -1;
//          }
//          DX_FILE_HEAD hd;
//          memset(&hd, 0, sizeof(DX_FILE_HEAD));
//          nReadSize = fread(&hd, sizeof(unsigned char), sizeof(DX_FILE_HEAD), fp);
//          if (nReadSize == sizeof(DX_FILE_HEAD))
//          {
//                 int nCount = hd.m_PointsNum;
//				   Head = hd;
//	 			   //printf("%u\n", sizeof(hd));
//                 if (nCount == 1)
//                 {
//                                 LORA_REC st;
//                                 memset(&st, 0, sizeof(LORA_REC));
//								   //printf("%u\n", sizeof(st));
//                                 nReadSize = fread(&st, sizeof(unsigned char), sizeof(st), fp);
//                                 if (nReadSize == sizeof(st))
//                                         g_LoraRec = st;
//                                 else
//                                 {
//                                         nResult = -1;
//                                 }
//                 }
//          }
//          fclose(fp);
//          return nResult;
//  }
import "C"
import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"unsafe"
)

type DX_FILD_HEAD struct {
	StartCode   uint32
	DbFileVison uint32
	WYear       int16
	Month       uint8
	Day         uint8
	Hour        uint8
	Minute      uint8
	Second      uint8
	ByRes       uint8
	PointsNum   uint32
	DwRes5      uint32
	DwRes6      uint32
	DwRes7      uint32
}
type LORA_REC struct {
	Bsptype       int32
	ServerPort    int32
	NRes1         int32
	NRes2         int32
	ServerIP      [64]int8
	ApplicationID [64]int8
	SzRes1        [64]int8
	SzRes2        [64]int8
	SzRes3        [64]int8
	SzRes4        [64]int8
}

var (
	head DX_FILD_HEAD
	rec  LORA_REC
)

func lora_init() {
	C.init_lora_para()
	C.read_lora_para_file()

	head := C.Head
	// rec := C.g_LoraRec

	fmt.Printf("%X\n", head)
	// fmt.Printf("%X\n", rec)
}

func main() {
	lora_init()
	file, err := os.Open("/home/tiger/Desktop/code/vscode/demo/c/lorapara.dat")
	if err != nil {
		fmt.Println("open lora para file failed")
		return
	}
	buf := new(bytes.Buffer)
	data_head := make([]byte, unsafe.Sizeof(head))
	_, err = file.Read(data_head)
	if err != nil {
		fmt.Println("read lora para file failed")
		return
	}
	copy(buf.Bytes(), data_head)

	if err = gob.NewDecoder(buf).Decode(&head); err != nil {
		fmt.Println("failed")
	}
	fmt.Printf("%X\n", head)

	// data_rec := make([]byte, unsafe.Sizeof(rec))
	// _, err = file.ReadAt(data_rec, int64(num))

	// if err != nil {
	// 	fmt.Println("read lora para file failed")
	// 	return
	// }
	fmt.Printf("%X\n", data_head)

}