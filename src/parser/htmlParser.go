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
func ParseCheckUSCIS(rawHTML string) {
	cardWasMailedCnt := strings.Split(rawHTML, "\n")[296:496]
	// Stripe spaces and comma
	for idx, ln := range cardWasMailedCnt {
		s := strings.TrimSpace(ln)
		s = strings.Replace(s, ",", "", -1)
		cardWasMailedCnt[idx] = s
	} 
	fmt.Println(cardWasMailedCnt)

}