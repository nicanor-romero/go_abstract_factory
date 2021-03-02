package main

type CDR struct {
	recordType    string
	originId      string
	destinationId string
	usage         int
}

type iRater interface {
	rate(cdr *CDR) float64
}

type iRaterFactory interface {
	getRater(cdr *CDR) iRater
}

func getRaterFactory(recordType string) (factory iRaterFactory) {
	if recordType == "voice" {
		factory = voiceRaterFactory{}
	} else if recordType == "sms" {
		factory = smsRaterFactory{}
	}
	return
}

func rateCDR(cdr *CDR) (cost float64) {
	factory := getRaterFactory(cdr.recordType)
	rater := factory.getRater(cdr)
	cost = rater.rate(cdr)
	return
}
