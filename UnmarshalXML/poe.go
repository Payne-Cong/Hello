package UnmarshalXML

import (
	es "Hello/error"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func readXMLFile(path string) (*os.File, error) {
	indexAbs, err := filepath.Abs(fmt.Sprintf("resource/%s", path))
	if err != nil {
		es.ErrorToString(err)
		return nil, err
	}

	//读取 XML 文件
	xmlFile, err := os.Open(indexAbs)
	if err != nil {
		es.ErrorToString(err)
		return nil, err
	}

	return xmlFile, nil
}

func createOrDeleteNewXML(newPath string) (*os.File, error) {
	newPath = fmt.Sprintf("resource/%s", newPath)
	if _, err := os.Stat(newPath); err == nil {
		// 如果存在，删除文件
		err = os.Remove(newPath)
		if err != nil {
			// 如果删除文件出错，输出错误信息
			log.Fatal(err)
			return nil, err
		}
	}

	// create new file xml
	newFile, err := os.Create(newPath)
	if err != nil {
		es.ErrorToString(err)
		return nil, err
	}

	return newFile, nil
}

func ParseXML() {

	xmlFile, err := readXMLFile("UISettings.xml")
	if err != nil {
		es.ErrorToString(err)
		return
	}
	defer xmlFile.Close()

	// 创建解码器
	//decoder := xml.NewDecoder(xmlFile)

	// 读取utf16-le格式的xml文件需要转码
	dec := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	utf8r := transform.NewReader(xmlFile, dec)

	decoder := xml.NewDecoder(utf8r)

	// 打开新文件，用于保存修改后的内容
	// 判断文件是否存在,存在就删除,否则新建
	newPath := "newUISettings.xml"
	newFile, err := createOrDeleteNewXML(newPath)
	if err != nil {
		es.ErrorToString(err)
		return
	}
	defer newFile.Close()

	//transform.NewWriter(newFile, )
	// 创建编码器
	utf16LEEncoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	utf16LeWriter := transform.NewWriter(newFile, utf16LEEncoder)
	encoder := xml.NewEncoder(utf16LeWriter)

	// 创建 XML 格式的头
	//header := xml.StartElement{
	//	Name: xml.Name{Local: "xml"},
	//	Attr: []xml.Attr{
	//		{Name: xml.Name{Local: "version"}, Value: "1.0"},
	//		//{Name: xml.Name{Local: "encoding"}, Value: "UTF-16 LE"},
	//	},
	//}
	// 使用编码器写入 XML 格式的头
	//encoder.EncodeToken(header)

	if err := encoder.EncodeToken(xml.ProcInst{
		Target: "xml",
		Inst:   []byte("version=\"1.0\""),
	}); err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	newline := xml.CharData("\n")
	encoder.EncodeToken(newline)

	// 设置缩进字符
	encoder.Indent("", "\t")

	for {
		// 读取下一个 XML 元素
		originToken, err := decoder.Token()
		if err == io.EOF {

			//// end
			//endElement := xml.EndElement{Name: xml.Name{
			//	Local: "xml",
			//}}
			//encoder.EncodeToken(endElement)

			// write
			encoder.Flush()
			break
		} else if err != nil {
			es.ErrorToString(err)
			panic(err)
		}

		// 处理 XML 元素
		token := xml.CopyToken(originToken)
		switch element := token.(type) {
		case xml.StartElement:
			// 元素开始，处理该元素
			fmt.Printf("Start element %s\n", element.Name.Local)
			var attributes []xml.Attr
			for _, attr := range element.Attr {
				if element.Name.Local == "Font" && attr.Name.Local == "typeface" {
					attr.Value = "Rainbow Party" // 你要修改的字体 Rainbow Party
				}
				attributes = append(attributes, attr)
				fmt.Printf("  Attribute %s = %s\n", attr.Name.Local, attr.Value)
			}

			startElement := xml.StartElement{
				Name: element.Name,
				Attr: attributes,
			}
			encoder.EncodeToken(startElement)

			//propsErr := decoder.DecodeElement(&props, &element)
			//if propsErr != nil {
			//	es.ErrorToString(propsErr)
			//}

		case xml.EndElement:
			// 元素结束，处理该元素
			//fmt.Printf("End element %s\n", element.Name.Local)
			endElement := xml.EndElement{Name: element.Name}
			encoder.EncodeToken(endElement)

		case xml.CharData:
			// 文本内容，处理该文本
			//fmt.Printf("  Character data: %s\n", string(element))
		}

	}

	//dec := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	//utf8r := transform.NewReader(xmlFile, dec)
}

// ReadFileUTF16 Similar to ioutil.ReadFile() but decodes UTF-16.  Useful when
// reading data from MS-Windows systems that generate UTF-16BE files,
// but will do the right thing if other BOMs are found.
func ReadFileUTF16(filename string) ([]byte, error) {

	// Read the file into a []byte:
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Make an tranformer that converts MS-Win default to UTF8:
	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)

	// decode and print:
	decoded, err := ioutil.ReadAll(unicodeReader)
	return decoded, err
}

// Deprecated
// 解析不支持utf8格式的xml, 暂不完善
func unmarshal(b []byte) (*Props, error) {
	var data Props
	decoder := xml.NewDecoder(bytes.NewBuffer(b))
	decoder.CharsetReader = func(charset string, reader io.Reader) (io.Reader, error) {
		enc, err := ianaindex.IANA.Encoding(charset)
		if err != nil {
			return nil, fmt.Errorf("charset %s: %s", charset, err.Error())
		}
		if enc == nil {
			// Assume it's compatible with (a subset of) UTF-8 encoding
			// Bug: https://github.com/golang/go/issues/19421
			return reader, nil
		}
		return enc.NewDecoder().Reader(reader), nil
	}

	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
