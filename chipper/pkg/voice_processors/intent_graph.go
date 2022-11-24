package wirepod

import (
	"strconv"

	"github.com/digital-dream-labs/chipper/pkg/vtt"
)

func (s *Server) ProcessIntentGraph(req *vtt.IntentGraphRequest) (*vtt.IntentGraphResponse, error) {
	botNum = botNum + 1
	var successMatched bool
	speechReq := reqToSpeechRequest(req)
	transcribedText, err := sttHandler(speechReq)
	if err != nil {
		botNum = botNum - 1
		IntentPass(req, "intent_system_noaudio", "voice processing error", map[string]string{"error": err.Error()}, true, speechReq.BotNum)
		return nil, nil
	}
	successMatched = processTextAll(req, transcribedText, matchListList, intentsList, speechReq.IsOpus, speechReq.BotNum)
	if !successMatched {
		logger("No intent was matched.")
		botNum = botNum - 1
		IntentPass(req, "intent_system_noaudio", transcribedText, map[string]string{"": ""}, false, speechReq.BotNum)
	}
	botNum = botNum - 1
	logger("Bot " + strconv.Itoa(speechReq.BotNum) + " request served.")
	return nil, nil
}