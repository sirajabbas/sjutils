package sjutils

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/nu7hatch/gouuid"
	"strconv"
	"strings"
)

func PrintPretty(obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(b))
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func ContainsObjectId(s []primitive.ObjectID, e primitive.ObjectID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ToJson(ip interface{}) string {
	out, err := json.Marshal(ip)
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

//this method generate unique UUID
func UUID() string {
	u, _ := uuid.NewV4()
	return u.String()
}

func ExtractSecondsFromTimeString(t string) (op int, err error) {
	if strings.Contains(t, ":") {
		list := strings.Split(t, ":")
		//reversing
		for i := len(list)/2 - 1; i >= 0; i-- {
			opp := len(list) - 1 - i
			list[i], list[opp] = list[opp], list[i]
		}
		for k, v := range list {
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			} else {
				if k == 0 {
					op = val
				}
				if k == 1 {
					op = op + (val * 60)
				}
				if k == 2 {
					op = op + (val * 60 * 60)
				}
			}
		}
		return
	} else {
		op, err = strconv.Atoi(t)
		return
	}
}
