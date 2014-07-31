package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Doc struct {
	XMLName     xml.Name        `xml:"Comprobante"`
	Complemento CFDIComplemento `xml:"Complemento"`
}

type CFDIComplemento struct {
	XMLName             xml.Name               `xml:"Complemento"`
	TimbreFiscalDigital TFDTimbreFiscalDigital `xml:"TimbreFiscalDigital"`
}

type TFDTimbreFiscalDigital struct {
	XMLName           xml.Name `xml:"TimbreFiscalDigital"`
	NumeroCertificado string   `xml:"noCertificadoSAT,attr"`
	FechaTimbrado     string   `xml:"FechaTimbrado,attr"`
	UUID              string   `xml:"UUID,attr"`
}

func parseXml(path string) interface{} {
	var query Doc
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return query
	}
	defer xmlFile.Close()

	rawContent, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(rawContent, &query)
	return query
}
