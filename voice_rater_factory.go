package main

type voiceRaterFactory struct {
}

func (f voiceRaterFactory) getRater(cdr *CDR) (rater iRater) {
	if cdr.originId == "SPAIN" && cdr.destinationId == "GERMANY" {
		rater = voiceRater{
			connectionFee: 0.15,
			rateValue:     0.02,
		}
	} else if cdr.originId == "SPAIN" && cdr.destinationId == "ARGENTINA" {
		rater = voiceRater{
			connectionFee: 0.20,
			rateValue:     0.05,
		}
	}
	return
}

type voiceRater struct {
	connectionFee float64
	rateValue     float64
}

func (r voiceRater) rate(cdr *CDR) (cost float64) {
	cost = r.connectionFee + r.rateValue * float64(cdr.usage)
	return
}
