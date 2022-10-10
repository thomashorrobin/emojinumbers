package emojinumbers

import "testing"

func TestToEmoji(t *testing.T) {
	testCases := []struct {
		desc           string
		number         int
		expectedString string
	}{
		{
			desc:           "Basic Test",
			number:         3241,
			expectedString: "3️⃣2️⃣4️⃣1️⃣",
		},
		{
			desc:           "Specical case for 10",
			number:         10,
			expectedString: "🔟",
		},
		{
			desc:           "Specical case for 100",
			number:         100,
			expectedString: "💯",
		},
		{
			desc:           "Negitive number",
			number:         -7635,
			expectedString: "-7️⃣6️⃣3️⃣5️⃣",
		},
		{
			desc:           "Specical case for zero",
			number:         0,
			expectedString: "0️⃣",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			outputString := ToEmoji(tC.number)
			if tC.expectedString != outputString {
				t.Errorf("Expected %s but got %s", tC.expectedString, outputString)
			}
		})
	}
}
