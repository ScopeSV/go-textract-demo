package main

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

var tSession *textract.Textract

func init() {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	}))
	tSession = textract.New(mySession)
}

func detectDocumentText(file []byte) {
	doc, err := tSession.DetectDocumentText(&textract.DetectDocumentTextInput{
		Document: &textract.Document{
			Bytes: file,
		},
	})
	if err != nil {
		panic(err)

	}

	for _, d := range doc.Blocks {
		if *d.BlockType == "WORD" {
			fmt.Println(*d.Text)
		}

	}

}

func createFeatureType(t string) *string {
	return &t
}

func askQuestion(question string) *textract.Query {
	return &textract.Query{
		Text: &question,
	}
}

func analyzeDocument(file []byte) {
	var featureTypes []*string
	var questions []*textract.Query

	featureTypes = append(featureTypes,
		createFeatureType("FORMS"),
		createFeatureType("TABLES"),
		createFeatureType("QUERIES"),
	)

	questions = append(questions,
		askQuestion("What is the product?"),
		askQuestion("Hva er kontonummer"),
	)

	doc, err := tSession.AnalyzeDocument(&textract.AnalyzeDocumentInput{
		Document: &textract.Document{
			Bytes: file,
		},
		FeatureTypes: featureTypes,
		QueriesConfig: &textract.QueriesConfig{
			Queries: questions,
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(doc)
}

func main() {
	file, err := ioutil.ReadFile("invoice.jpeg")
	if err != nil {
		panic(err)
	}

	//detectDocumentText(file)
	analyzeDocument(file)
}
