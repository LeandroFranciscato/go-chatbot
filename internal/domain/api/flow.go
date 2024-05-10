package api

type Flow struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Greeting  string          `json:"greeting"`
	FinalStep int             `json:"finalStep"`
	Steps     map[string]Step `json:"steps"`
}

type Step struct {
	Question       string            `json:"question"`
	IntentsDataset map[string]string `json:"intents_dataset"`
	IntentConfig   map[string]Intent `json:"intent_config"`
}

type Intent struct {
	Answer   string `json:"answer"`
	NextStep string `json:"next_step"`
}
