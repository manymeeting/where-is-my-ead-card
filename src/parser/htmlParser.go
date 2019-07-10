package htmlparser

import (
	"fmt"
	"strings"
)

/**

Result Data Srruct

[
	{
		
		Date: "2019-07-01",
		Status: [
			{
				CaseRange: "YSCXXX0",
				Stats: {
					Received: 123,
					CardBeingProduced: 234,
					CardWasMailed: 345
				}
			},
			{
				CaseRange: "YSCXXX1",
				Stats: {
					Received: 321,
					CardBeingProduced: 432,
					CardWasMailed: 543
				}
			}
		]
	}
]
*/

const (

	CARD_DELIVERED_SLN = 296
	CARD_DELIVERED_ELN = 496

	CASE_RECEIVED_SLN = 1216
	CASE_RECEIVED_ELN = 1416

	CASE_ID_SLN = 1494
	CASE_ID_ELN = 1694

)

/**
 * @param {string}
 */
func ParseCheckUSCIS(rawHTML string) {
	caseIDs := strings.Split(rawHTML, "\n")[CASE_ID_SLN:CASE_ID_ELN]
	cardDeliveredCnt := strings.Split(rawHTML, "\n")[CARD_DELIVERED_SLN:CARD_DELIVERED_ELN]
	caseReceivedCnt := strings.Split(rawHTML, "\n")[CASE_RECEIVED_SLN:CASE_RECEIVED_ELN]

	// Stripe spaces and comma
	fmtJSONLineToArr(caseIDs)
	fmtJSONLineToArr(cardDeliveredCnt)
	fmtJSONLineToArr(caseReceivedCnt)
	
	fmt.Println(caseIDs)
	fmt.Println(cardDeliveredCnt)
	fmt.Println(caseReceivedCnt)


}

/**
 * @param  {[[]string]} Lines in string fromat, contains comma and/or spaces
 */
func fmtJSONLineToArr(dataLines []string){
	for idx, ln := range dataLines {
		s := strings.TrimSpace(ln)
		s = strings.Replace(s, ",", "", -1)
		dataLines[idx] = s
	}
}