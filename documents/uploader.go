package documents

import (
	"github.com/ryotarai/paramedic/awsclient"
)

type Uploader struct {
	ssm awsclient.SSM
}

func NewUploader(ssm awsclient.SSM) *Uploader {
	return &Uploader{
		ssm: ssm,
	}
}

// func (u *Uploader) UploadFile(path string) error {
// 	log.Printf("DEBUG: loading %s\n", path)
// 	def, err := newDefinitionFromFile(path)
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("DEBUG: %+v\n", def)
// 	log.Printf("INFO: creating a new document\n")

// 	return u.uploadDefinition(def)
// }

// func (u *Uploader) uploadDefinition(d *definition) error {
// 	input := &ssm.CreateDocumentInput{
// 		Content:      aws.String(""),
// 		DocumentType: aws.String("Command"),
// 		Name:         aws.String(""),
// 	}
// 	u.ssm.CreateDocument(input)
// 	return nil
// }
