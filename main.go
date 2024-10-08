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
					<CPF></CPF>
					<xNome></xNome>
					<enderDest>
						<xLgr></xLgr>
						<nro></nro>
						<xBairro></xBairro>
						<cMun></cMun>
						<xMun></xMun>
						<UF></UF>
						<CEP></CEP>
						<cPais></cPais>
						<xPais></xPais>
						<fone></fone>
					</enderDest>
					<indIEDest></indIEDest>
					<email></email>
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
