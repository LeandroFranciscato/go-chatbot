package flow

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"review-chatbot/internal/domain/api"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"
	"strconv"
	"strings"
	"sync"

	"github.com/jdkato/prose/v2"
)

//go:embed files/stop_words.json
var stopWordsJson []byte

type Flow interface {
	Ask(order entity.Order, step int) (template.HTML, error)
	Answer(step int, userAnswer string) (int, string)
	ID() string
	FinalStep() int
	Name() string
}

type useCase struct {
	stopWords map[string]struct{}
	flow      api.Flow
	orderRepo repo.Repo[entity.Order]
}

func New(flowJson []byte, orderRepo repo.Repo[entity.Order]) (Flow, error) {
	usecase := useCase{
		orderRepo: orderRepo,
	}

	err := json.Unmarshal(flowJson, &usecase.flow)
	if err != nil {
		return usecase, errors.New("error parsing flow: " + err.Error())
	}

	err = json.Unmarshal(stopWordsJson, &usecase.stopWords)
	if err != nil {
		return usecase, errors.New("error parsing stopWords: " + err.Error())
	}
	return usecase, nil
}

func (usecase useCase) ID() string {
	return usecase.flow.ID
}

func (usecase useCase) FinalStep() int {
	return usecase.flow.FinalStep
}

func (usecase useCase) Name() string {
	return usecase.flow.Name
}

func (usecase useCase) Ask(order entity.Order, step int) (template.HTML, error) {

	tmpl, err := template.New("question").Parse(usecase.flow.Steps[strconv.Itoa(step)].Question)
	if err != nil {
		return "", errors.New("error creating template: " + err.Error())
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, order)
	if err != nil {
		return "", errors.New("error executing template: " + err.Error())
	}

	return template.HTML(buf.String()), nil
}

func (usecase useCase) Answer(step int, userAnswer string) (int, string) {
	reviewFlowStep := usecase.flow.Steps[strconv.Itoa(step)]
	intent := usecase.identifyIntent(userAnswer, reviewFlowStep.IntentsDataset)
	intentAnswer, ok := reviewFlowStep.IntentConfig[intent]
	if !ok {
		return usecase.flow.FinalStep, ""
	}
	nextStep, _ := strconv.Atoi(intentAnswer.NextStep)
	return nextStep, intentAnswer.Answer
}

func (usecase useCase) identifyIntent(message string, intents map[string]string) string {
	doc, err := prose.NewDocument(message)
	if err != nil {
		log.Fatal(err)
	}
	messageTokens := usecase.removeStopWords(doc.Tokens())

	maxMatch := 0.0
	bestIntent := "unknown"

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(intents))
	for k, v := range intents {
		go func(k, v string) {
			defer wg.Done()

			doc, err := prose.NewDocument(k)
			if err != nil {
				log.Fatal(err)
			}
			intentTokens := usecase.removeStopWords(doc.Tokens())

			match := usecase.calculateMatch(messageTokens, intentTokens)
			mu.Lock()
			if match > maxMatch {
				maxMatch = match
				bestIntent = v
			}
			mu.Unlock()
		}(k, v)
	}
	wg.Wait()

	return bestIntent
}

func (usecase useCase) removeStopWords(tokens []prose.Token) []string {
	var result []string
	for _, token := range tokens {
		if _, ok := usecase.stopWords[strings.ToLower(token.Text)]; !ok {
			result = append(result, strings.ToLower(token.Text))
		}
	}
	return result
}

func (usecase useCase) calculateMatch(messageTokens, intentTokens []string) float64 {
	matchCount := 0
	for _, mt := range messageTokens {
		for _, it := range intentTokens {
			if mt == it {
				matchCount++
				break
			}
		}
	}
	return float64(matchCount) / float64(len(messageTokens))
}
