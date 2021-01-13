package utilsgo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"testing"
	"time"
)

var m map[string]int

type Model struct {
	Name string
	Age  int
}

func TestPrintPretty(t *testing.T) {
	m = map[string]int{
		"key1": 1,
		"key2": 2,
	}
	PrintPretty(m)
}

func TestContainsString(t *testing.T) {
	var source []string
	for i := 0; i < 10; i++ {
		source = append(source, strconv.Itoa(i))
	}
	target := "5"
	found := ContainsString(source, target)

	if found {
		t.Log("Found Item:")
	} else {
		t.Fatal("Test failed item not found")
	}
}

func TestContainsObjectId(t *testing.T) {
	var source []primitive.ObjectID
	for i := 0; i < 10; i++ {
		source = append(source, primitive.NewObjectID())
	}
	target := source[0]
	found := ContainsObjectId(source, target)

	if found {
		t.Log("Found Item:")
	} else {
		t.Fatal("Test failed item not found")
	}
}

func TestToJson(t *testing.T) {
	m := Model{
		Name: "Joy",
		Age:  52,
	}
	json := ToJson(m)
	t.Log(json)
	if json == "" {
		t.Fail()
	}
}

func TestUUID(t *testing.T) {
	id := UUID()
	if id == "" {
		t.Fail()
	}
}

func TestExtractSecondsFromTimeString(t *testing.T) {
	timStr := time.Now().String()
	s, err := ExtractSecondsFromTimeString(timStr)
	if err != nil || s == 0 {
		t.Fail()
	} else {
		t.Log("seconds: ", s)
	}
}

func TestFindCommonInSlice(t *testing.T) {
	set := []string{"a", "b", "c", "d"}
	subset := []string{"b", "d"}
	common := FindCommonInSlice(set, subset)
	fmt.Println("common:", common)

	set1 := []int{1, 2, 4, 5, 6}
	subset1 := []int{2, 4, 6}
	common1 := FindCommonInSlice(set1, subset1)
	fmt.Println("common:", common1)
}

func TestFindUnCommonInSlice(t *testing.T) {
	set := []string{"a", "b", "c", "d"}
	subset := []string{"b", "d"}
	uncommon := FindUnCommonInSlice(set, subset)
	fmt.Println("uncommon:", uncommon)
}