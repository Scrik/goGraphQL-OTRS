package api

import "time"

type TypeXmlStorage struct {
	XmlType         string  `db:"xml_type"`
	XmlKey          *string `db:"xml_key"`
	XmlContentKey   *string `db:"xml_content_key"`
	XmlContentValue *string `db:"xml_content_value"`
}
