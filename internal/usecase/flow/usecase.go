package flow

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"log"
	"review-chatbot/internal/domain/api"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"
	"review-chatbot/internal/util"
	"strconv"
	"strings"
	"sync"

	"github.com/jdkato/prose/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:embed files/stop_words.json
var stopWordsJson []byte

type Flow interface {
	ID() string
	FinalStep() int
	Name() string
	Ask(step int) string
	Answer(step int, userAnswer string) (int, string)
	SaveHistory(ctx context.Context, step int, customerID string, orderID string, history string) error
	GetHistory(ctx context.Context, customerID string, orderID string) (entity.Chat, error)
}

type useCase struct {
	stopWords map[string]struct{}
	flow      api.Flow
	chatRepo  repo.Repo[entity.Chat]
	timer     util.Time
}

func New(flowJson []byte, chatHistoryRepo repo.Repo[entity.Chat]) (Flow, error) {
	usecase := useCase{
		chatRepo: chatHistoryRepo,
		timer:    util.NewTimer(),
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

func (usecase useCase) Ask(step int) string {
	return usecase.flow.Steps[strconv.Itoa(step)].Question
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

func (usecase useCase) SaveHistory(ctx context.Context, step int, customerID string, orderID string, history string) error {
	chatHistory, err := usecase.GetHistory(ctx, customerID, orderID)
	if err != nil {
		return errors.New("error finding chat history: " + err.Error())
	}
	if chatHistory.ID.IsZero() {
		customerObjID, _ := primitive.ObjectIDFromHex(customerID)
		orderObjID, _ := primitive.ObjectIDFromHex(orderID)

		_, err = usecase.chatRepo.InsertOne(ctx, entity.Chat{
			CustomerID:  customerObjID,
			OrderID:     orderObjID,
			History:     history,
			Status:      entity.ChatStatusInProgress,
			Timestamp:   usecase.timer.Now(),
			CurrentStep: step,
		})
		if err != nil {
			return errors.New("error inserting chat history: " + err.Error())
		}
		return nil
	}

	chatHistory.History = history
	chatHistory.CurrentStep = step
	if usecase.flow.FinalStep == step {
		chatHistory.Status = entity.ChatStatusDone
	}
	err = usecase.chatRepo.UpdateOne(ctx, chatHistory)
	if err != nil {
		return errors.New("error updating chat history: " + err.Error())
	}
	return nil
}

func (usecase useCase) GetHistory(ctx context.Context, customerID string, orderID string) (entity.Chat, error) {
	customerObjID, _ := primitive.ObjectIDFromHex(customerID)
	orderObjID, _ := primitive.ObjectIDFromHex(orderID)

	chatHistory, err := usecase.chatRepo.FindOne(ctx,
		bson.D{
			{Key: "$and", Value: bson.A{
				bson.D{{Key: "customer_id", Value: customerObjID}},
				bson.D{{Key: "order_id", Value: orderObjID}}},
			},
		},
	)
	if err != nil {
		return entity.Chat{}, errors.New("error finding chat history: " + err.Error())
	}
	return chatHistory, nil
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
