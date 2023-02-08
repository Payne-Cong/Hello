package UnmarshalXML

import "encoding/xml"

type Size struct {
	XMLName xml.Name `xml:"Size"`
	ID      string   `xml:"id,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

type Point struct {
	XMLName xml.Name `xml:"Point"`
	ID      string   `xml:"id,attr"`
	X       int      `xml:"x,attr"`
	Y       int      `xml:"y,attr"`
}

type InstalledFont struct {
	XMLName  xml.Name `xml:"InstalledFont"`
	ID       string   `xml:"id,attr"`
	Typeface string   `xml:"typeface,attr"`
	Style    string   `xml:"style,attr"`
	Value    string   `xml:"value,attr"`
}

type FallbackFont struct {
	XMLName xml.Name `xml:"FallbackFont"`
	ID      string   `xml:"id,attr"`
	Ranges  string   `xml:"ranges,attr"`
	Fonts   string   `xml:"fonts,attr"`
}

type ImageDefinitionList struct {
	XMLName xml.Name `xml:"ImageDefinitionList"`
	ID      string   `xml:"id,attr"`
	Value   string   `xml:"value,attr"`
}

type Font struct {
	XMLName  xml.Name `xml:"Font"`
	ID       string   `xml:"id,attr"`
	Typeface string   `xml:"typeface,attr"`
	Size     string   `xml:"size,attr"`
	Inherits string   `xml:"inherits,attr"`
	Bold     string   `xml:"bold,attr"`
	Italic   string   `xml:"italic,attr"`
}

type Colour struct {
	XMLName xml.Name `xml:"Colour"`
	ID      string   `xml:"id,attr"`
	Value   string   `xml:"value,attr"`
}

type Number struct {
	XMLName xml.Name `xml:"Number"`
	ID      string   `xml:"id,attr"`
	Value   string   `xml:"value,attr"`
}

type String struct {
	XMLName xml.Name `xml:"String"`
	ID      string   `xml:"id,attr"`
	Value   string   `xml:"value,attr"`
}

type ImageIdList struct {
	XMLName xml.Name  `xml:"ImageIdList"`
	Value   string    `xml:"value,attr"`
	ImageId []ImageId `xml:"ImageId"`
}

type ImageId struct {
	XMLName xml.Name `xml:"ImageId"`
	Value   string   `xml:"value,attr"`
}

type ImageIncludePath struct {
	XMLName xml.Name `xml:"ImageIncludePath"`
	ID      string   `xml:"id,attr"`
	Value   string   `xml:"value,attr"`
}

type Props struct {
	XMLName       xml.Name        `xml:"Props"`
	ID            string          `xml:"id,attr"`
	Size          Size            `xml:"Size"`
	Point         Point           `xml:"Point"`
	InstalledFont []InstalledFont `xml:"InstalledFont"`
	//ImageIncludePath    ImageIncludePath      `xml:"ImageIncludePath"`
	//FallbackFont        []FallbackFont        `xml:"FallbackFont"`
	//ImageDefinitionList []ImageDefinitionList `xml:"ImageDefinitionList"`
	//Font                []Font                `xml:"Font"`
	//Colour              []Colour              `xml:"Colour"`
	//Number              []Number              `xml:"Number"`
	//Props               []struct {
	//	ID     string   `xml:"id,attr"`
	//	Colour []Colour `xml:"Colour"`
	//	String []String `xml:"String"`
	//	Font   []Font   `xml:"Font"`
	//	Size   Size     `xml:"Size"`
	//	Number []Number `xml:"Number"`
	//	Point  Point    `xml:"Point"`
	//} `xml:"Props"`
	//ImageIdList []ImageIdList `xml:"ImageIdList"`
}
