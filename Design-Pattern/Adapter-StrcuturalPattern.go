package main

import "fmt"

// EnglishSpeaker represents the English-speaking interface.
type EnglishSpeaker interface {
	SpeakEnglish() string
}

// SpanishSpeaker represents the Spanish-speaking interface.
type SpanishSpeaker interface {
	SpeakSpanish() string
}

// SpanishSpeakerImpl implements the SpanishSpeaker interface.
type SpanishSpeakerImpl struct{}

func (s *SpanishSpeakerImpl) SpeakSpanish() string {
	return "Hola, ¿cómo está?"
}

// EnglishSpeakerImpl implements the EnglishSpeaker interface.
type EnglishSpeakerImpl struct{}

func (e *EnglishSpeakerImpl) SpeakEnglish() string {
	return "Hello, how are you?"
}

// SpanishToEnglishAdapter adapts SpanishSpeaker to EnglishSpeaker.
type SpanishToEnglishAdapter struct {
	SpanishSpeaker SpanishSpeaker
}

func (adapter *SpanishToEnglishAdapter) SpeakEnglish() string {
	// Perform translation from Spanish to English here
	spanishMessage := adapter.SpanishSpeaker.SpeakSpanish()
	// For demonstration purposes, let's assume a simple translation
	englishMessage := "Translated: " + spanishMessage
	return englishMessage
}

func main() {
	spanishSpeaker := &SpanishSpeakerImpl{}
	englishSpeaker := &EnglishSpeakerImpl{}
	adapter := &SpanishToEnglishAdapter{SpanishSpeaker: spanishSpeaker}

	fmt.Println("Spanish speaker says:", spanishSpeaker.SpeakSpanish())
	fmt.Println("English speaker says:", englishSpeaker.SpeakEnglish())

	// Now, using the adapter to make the Spanish speaker speak English
	fmt.Println("Adapter makes Spanish speaker speak English:", adapter.SpeakEnglish())
}
