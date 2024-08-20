package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	dir := "./"
	novoDest := `<dest>
<CPF>12345678910</CPF>
<xNome>PRIQUITO</xNome>
<enderDest>
<xLgr>Rua 31 de Maio</xLgr>
<nro>180</nro>
<xBairro>Sitio Grande</xBairro>
<cMun>2928901</cMun>
<xMun>Sao Desiderio</xMun>
<UF>BA</UF>
<CEP>47825000</CEP>
<cPais>1058</cPais>
<xPais>Brasil</xPais>
<fone>77999930046</fone>
</enderDest>
<indIEDest>9</indIEDest>
<email>kauek78942@gmail.com</email>
</dest>`

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".xml" {
			xmlData, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			xmlStr := string(xmlData)

			xmlStr = regexp.MustCompile(`(?s)<dest>.*?</dest>`).ReplaceAllString(xmlStr, novoDest)

			xmlStr = regexp.MustCompile(`(?s)<nfeProc([^>]*)>`).ReplaceAllString(xmlStr, `<nfeProc$1 xmlns="http://www.portalfiscal.inf.br/nfe">`)
			xmlStr = regexp.MustCompile(`(?s)<NFe([^>]*)>`).ReplaceAllString(xmlStr, `<NFe$1 xmlns="http://www.portalfiscal.inf.br/nfe">`)

			xmlStr = regexp.MustCompile(`(?s)<(/?\w+)([^>]*)\s+xmlns="http://www.portalfiscal.inf.br/nfe"([^>]*)>`).ReplaceAllString(xmlStr, `<$1$2$3>`)

			err = ioutil.WriteFile(path, []byte(xmlStr), 0644)
			if err != nil {
				return err
			}

			fmt.Printf("Arquivo %s processado e substituído\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Erro ao processar arquivos: %v\n", err)
	}
}
