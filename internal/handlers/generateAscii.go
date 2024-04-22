package handlers

import (
	"errors"

	"github.com/ediallocyf/asciiartweb/util"
)

// GenerateAsciiArt generates ASCII art based on the provided text and banner type.
func GenerateAsciiArt(text, bannerType string) (string, error) {
	var patternFileName string

	switch bannerType {
	case shadowBanner:
		patternFileName = shadowPatternFileAddress
	case standardBanner:
		patternFileName = standardPatternfileAdrress
	case thinkerBanner:
		patternFileName = thinkerToyPatternFileAddress
	default:
		logger.Error("Unknown Banner----<<--GenerateAsciiArt---", "type", bannerType)
	}

	patternContent, errPattern := util.ReadFileToStr(patternFileName)
	if errPattern != nil {
		errMsg = "----<-GenerateAsciiArt------<--patternContent----" + errPattern.Error()
		logMsg = "Error in loading file content"
		logger.Error(logMsg + errMsg)
		return "", errors.New(errMsg)
	} else {
		logger.Info("file content loaded successfullly---<---GenerateAsciiArt---<-patternContent--")
	}
	patternMap, errPatternMap := util.ContentToMap(patternContent, chLength)
	if errPatternMap != nil {
		errMsg = "----<-GenerateAsciiArt------<--patternMap----" + errPatternMap.Error()
		logMsg = "Errorinc converting pathern to map"
		logger.Error(logMsg + errMsg)
		return "", errors.New(errMsg)
	}
	return util.StrMaker(text, patternMap, chLength), nil
}
