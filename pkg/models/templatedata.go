package models

//TemplateData is to hold data to be passed to templates
type TemplateData struct {
	StringMap    map[string]string
	IntMap       map[string]int
	FloatMap     map[string]float32
	Data         map[string]interface{}
	CSRFToken    string
	FlashMessage string
	Warning      string
	Error        string
}
