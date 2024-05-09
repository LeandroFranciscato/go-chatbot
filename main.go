package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jdkato/prose/v2"
)

type ReviewFlow struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	Greeting string          `json:"greeting"`
	Steps    map[string]Step `json:"steps"`
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

var stopWords = map[string]struct{}{
	"a": {}, "an": {}, "and": {}, "are": {}, "as": {}, "at": {}, "be": {}, "by": {}, "for": {}, "from": {},
	"has": {}, "he": {}, "in": {}, "is": {}, "it": {}, "its": {}, "of": {}, "on": {}, "that": {}, "the": {},
	"to": {}, "was": {}, "were": {}, "will": {}, "with": {},
}

//go:embed flow/review.json
var reviewFlowJson []byte
var reviewFlow ReviewFlow

const finalStep = 199

func main() {

	LoadReviewFlow()

	e := gin.Default()

	api := e.Group("/api")
	api.GET("customer/:customer/order/:order/delivery-confirmed", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/chat/review/1")
	})

	e.LoadHTMLFiles("form.html")

	chat := e.Group("/chat")
	chat.GET("/review", func(c *gin.Context) {
		step := 1
		c.HTML(http.StatusOK, "form.html", gin.H{
			"title":    "Review",
			"question": Ask(step),
			"step":     strconv.Itoa(step),
		})
	})

	chat.POST("/review", func(c *gin.Context) {
		step, _ := strconv.Atoi(c.Request.FormValue("step"))
		userAnswer := c.Request.FormValue("answer")
		nextStep, answer := Answer(step, userAnswer)
		c.HTML(http.StatusOK, "form.html", gin.H{
			"title":    "Review",
			"answer":   answer,
			"question": Ask(nextStep),
			"step":     strconv.Itoa(nextStep),
		})
	})

	e.Run()
}

func LoadReviewFlow() {
	err := json.Unmarshal(reviewFlowJson, &reviewFlow)
	if err != nil {
		panic(err)
	}
}

func Ask(step int) string {
	return reviewFlow.Steps[strconv.Itoa(step)].Question
}

func Answer(step int, userAnswer string) (int, string) {
	reviewFlowStep := reviewFlow.Steps[strconv.Itoa(step)]
	intent := identifyIntent(userAnswer, reviewFlowStep.IntentsDataset)
	intentAnswer, ok := reviewFlowStep.IntentConfig[intent]
	if !ok {
		return finalStep, ""
	}
	nextStep, _ := strconv.Atoi(intentAnswer.NextStep)
	return nextStep, intentAnswer.Answer
}

func identifyIntent(message string, intents map[string]string) string {
	doc, err := prose.NewDocument(message)
	if err != nil {
		log.Fatal(err)
	}
	messageTokens := removeStopWords(doc.Tokens())

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
			intentTokens := removeStopWords(doc.Tokens())

			match := calculateMatch(messageTokens, intentTokens)
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

func removeStopWords(tokens []prose.Token) []string {
	var result []string
	for _, token := range tokens {
		if _, ok := stopWords[strings.ToLower(token.Text)]; !ok {
			result = append(result, strings.ToLower(token.Text))
		}
	}
	return result
}

func calculateMatch(messageTokens, intentTokens []string) float64 {
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
