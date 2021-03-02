package main

type smsRaterFactory struct {
}

func (f smsRaterFactory) getRater(cdr *CDR) (rater iRater) {
	if cdr.originId == "SPAIN" && cdr.destinationId == "GERMANY" {
		rater = smsRater{
			rateValue: 0.10,
		}
	} else if cdr.originId == "SPAIN" && cdr.destinationId == "ARGENTINA" {
		rater = smsRater{
			rateValue: 0.25,
		}
	}
	return
}

type smsRater struct {
	rateValue float64
}

func (r smsRater) rate(cdr *CDR) (cost float64) {
	cost = r.rateValue * float64(cdr.usage)
	return
}