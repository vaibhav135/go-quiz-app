package quiz

const (
	Json = "json"
	Csv  = "csv"
)

type QuizContent struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
