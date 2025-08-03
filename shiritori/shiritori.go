package shiritori

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type State struct {
	LastWord string
	UsedWords []string
}

type Entry struct {
	Words []string
	Readings []string
	Definitions []string
}
var Dict = map[string]jmdictEntry{}

type k_ele struct {
	Keb string `xml:"keb"`
	Ke_inf []string `xml:"ke_inf"`
	Ke_pri []string `xml:"ke_pri"`
}

type r_ele struct {
	Reb string `xml:"reb"`
	Re_nokanji string `xml:"re_nokanji"`
	Re_restr string `xml:"re_restr"`
	Re_inf string `xml:"re_inf"`
	Re_pri string `xml:"re_pri"`
}

type sense struct {
	Stagk []string `xml:"stagk"`
	Stagr []string `xml:"stagr"`
	Pos []string `xml:"pos"`
	Xref []string `xml:"xref"`
	Ant []string `xml:"ant"`
	Field []string `xml:"field"`
	Misc []string `xml:"misc"`
	S_inf []string `xml:"s_inf"`
	Lsource []string `xml:"lsource"`
	Dial []string `xml:"dial"`
	Gloss []string `xml:"gloss"`
}

type jmdictEntry struct {
	XMLName xml.Name `xml:"entry"`
	Ent_seq int64 `xml:"ent_seq"`
	K_ele []*k_ele `xml:"k_ele"`
	R_ele []*r_ele `xml:"r_ele"`
	Sense []*sense `xml:"sense"`
}
type xmlStruct struct {
	XMLName xml.Name `xml:"JMdict"`
	Entries []*jmdictEntry `xml:"entry"`
}

func Read_dict(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	d := xml.NewDecoder(bytes.NewReader(data))
	d.Strict = false
	var xmlData xmlStruct
	d.Decode(&xmlData)
	fmt.Println(xmlData.Entries[0].K_ele)
	for _, entry := range xmlData.Entries {
		for _, kele := range entry.K_ele {
			Dict[kele.Keb] = *entry
		}
	}
}

func Get_entry(word string) Entry {
	var ent = Dict[word]
	var retVal = Entry{}
	for _, k_ele_p := range ent.K_ele {
		retVal.Words = append(retVal.Words, k_ele_p.Keb)
	}
	for _, r_ele_p := range ent.R_ele {
		retVal.Readings = append(retVal.Readings, r_ele_p.Reb)
	}
	for _, sense := range ent.Sense {
		for _, gloss := range sense.Gloss {
			retVal.Definitions = append(retVal.Definitions, gloss)
		}
	}

	return retVal
}
