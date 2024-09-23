package comsoc

import "errors"

func TieBreakFactory(orderedAlts []Alternative) func([]Alternative) (Alternative, error) {
	return func(alts []Alternative) (winner Alternative, err error) {
		for _, alt := range orderedAlts {
			for _, possibleWinner := range alts {
				if possibleWinner == alt {
					return possibleWinner, nil
				}
			}
		}
		return alts[0], errors.New("Erreur, pas de gagant dans la liste")
	}
}

func SWFFactory(swf func(p Profile) (Count, error), tb func([]Alternative) (Alternative, error)) func(Profile) ([]Alternative, error) {

}

func SCFFactory(scf func(p Profile) ([]Alternative, error), tb func([]Alternative) (Alternative, error)) func(Profile) (Alternative, error) {

}
