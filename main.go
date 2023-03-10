package main

import (
	"Enigma/attack"
	"Enigma/machine"
	"fmt"
	"strings"
)

var plainText string = `DURING the whole of a dull dark and soundless day in the autumn of the year when the clouds hung oppressively low in the heavens I had been passing alone on horseback through a singularly dreary tract of country and at length found myself as the shades of the evening drew on within view of the melancholy House of Usher I know not how it was but with the first glimpse of the building a sense of insufferable gloom pervaded my spirit I say insufferable for the feeling was unrelieved by any of that halfpleasurable because poetic sentiment with which the mind usually receives even the sternest natural images of the desolate or terrible I looked upon the scene before me upon the mere house and the simple landscape features of the domain upon the bleak walls upon the vacant eyelike windows upon a few rank sedges and upon a few white trunks of decayed trees with an utter depression of soul which I can compare to no earthly sensation more properly than to the afterdream of the reveller upon opium the bitter lapse into everyday life the hideous dropping off of the veil There was an iciness a sinking a sickening of the heart an unredeemed dreariness of thought which no goading of the imagination could torture into aught of the sublime What was it I paused to think what was it that so unnerved me in the contemplation of the House of Usher It was a mystery all insoluble nor could I grapple with the shadowy fancies that crowded upon me as I pondered I was forced to fall back upon the unsatisfactory conclusion that while beyond doubt there are combinations of very simple natural objects which have the power of thus affecting us still the analysis of this power lies among considerations beyond our depth It was possible I reflected that a mere different arrangement of the particulars of the scene of the details of the picture would be sufficient to modify or perhaps to annihilate its capacity for sorrowful impression and acting upon this idea I reined my horse to the precipitous brink of a black and lurid tarn that lay in unruffled lustre by the dwelling and gazed down but with a shudder even more thrilling than before upon the remodelled and inverted images of the gray sedge and the ghastly treestems and the vacant and eyelike windows`

func main() {
	fmt.Println("Running...")
	encoderEnigma := machine.NewEnigmaMachineTest()
	encodedText := encoderEnigma.EncodeDecodeText(enigmafyText(plainText))
	fmt.Println(encodedText)
	attack.Attack(encodedText)
}

func enigmafyText(text string) string {
	return strings.ToUpper(strings.ReplaceAll(text, " ", ""))
}
