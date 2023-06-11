package controller

import "strconv"

func GetWordDetail(params map[string]interface{}) map[string]interface{} {
	wordIdStr := params["id"]
	wordDict := params["dict"]
	wordId, err := strconv.Atoi(wordIdStr.(string))
	if err != nil {
		return nil
	}
	if wordDict != nil && len(wordDict.(string)) >= 1 {
		detail := GetWordDetailFromDict(wordId, wordDict.(string))
		return detail
	}
	return map[string]interface{}{
		"id": wordId,
	}
}
