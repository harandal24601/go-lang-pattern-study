package main

import (
	"errors"
	"log"
)

var emptyS1 error = errors.New("s1 is empty")
var emptyS2 error = errors.New("s2 is empty")

func main() {
	log.Println(join1("", "s2", 2))
	log.Println(join2("s1", "", 2))
}

/*
* join 1
(1) Specific tasks are handled by the concatenate function, but errors may be returned.

The join function concatenates two strings and returns if the length is longer than max.
If we just look at the features, it was implemented accurately.
However, it is not intuitive to manage a mindfulness model that encompasses all cases.
This is because there are too many overlapping steps.
*/
func join1(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", emptyS1
	} else {
		if s2 == "" {
			return "", emptyS2
		} else {
			concat, err := concatenate(s1, s2) // (1)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

/*
* join 2

The code functions the same as join1, but requires much less cognitive effort to maintain the mental model.
This is because there are only two levels of superposition.
*/
func join2(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", emptyS1
	}
	if s2 == "" {
		return "", emptyS2
	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}
	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return "", nil
}
