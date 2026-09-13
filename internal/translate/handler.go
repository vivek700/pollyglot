package translate

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

type Service struct {
	Logger  *slog.Logger
	Client  openai.Client
	AIModel string
}

func (s *Service) Translate(w http.ResponseWriter, r *http.Request) {

	type reqBody struct {
		Text string `json:"text"`
		Lang string `json:"lang"`
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	req := reqBody{}
	err := decoder.Decode(&req)

	if err != nil {
		s.Logger.Error("Decoding failed", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(req.Lang, req.Text)

	_, _ = s.translateText(r.Context(), req.Text, req.Lang)

	w.WriteHeader(http.StatusNoContent)

}

func (s *Service) translateText(ctx context.Context, text, lang string) (string, error) {

	ques := "Write me a haiku about computers."

	resp, err := s.Client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(ques)},
		Model: s.AIModel,
	})

	if err != nil {
		s.Logger.Error("failed to get a response", "err", err)
		return "", nil
	}

	fmt.Println(resp.OutputText())

	return "", nil

}
