package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type ActionExecution struct {
	ID                primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Action            map[string]interface{} `json:"action,omitempty" bson:"action,omitempty"`
	Parameters        map[string]interface{} `json:"parameters,omitempty" bson:"parameters,omitempty"`
	Context           map[string]interface{} `json:"context,omitempty" bson:"context,omitempty"`
	Result            map[string]interface{} `json:"result,omitempty" bson:"result,omitempty"`
	Trigger           map[string]interface{} `json:"trigger,omitempty" bson:"trigger,omitempty"`
	TriggerType       map[string]interface{} `json:"trigger_type,omitempty" bson:"trigger_type,omitempty"`
	TriggerInstance   map[string]interface{} `json:"trigger_instance,omitempty" bson:"trigger_instance,omitempty"`
	Rule              map[string]interface{} `json:"rule,omitempty" bson:"rule,omitempty"`
	Runner            map[string]interface{} `json:"runner,omitempty" bson:"runner,omitempty"`
	LiveAction        map[string]interface{} `json:"liveaction,omitempty" bson:"liveaction,omitempty"`
	TaskExecution     string                 `json:"task_execution,omitempty" bson:"task_execution,omitempty"`
	WorkflowExecution string                 `json:"workflow_execution,omitempty" bson:"workflow_execution,omitempty"`
	Status            string                 `json:"status,omitempty" bson:"status,omitempty"`
	StartTimestamp    string                 `json:"start_timestamp,omitempty" bson:"start_timestamp,omitempty"`
	EndTimestamp      string                 `json:"end_timestamp,omitempty" bson:"end_timestamp,omitempty"`
	ElapsedSeconds    int                    `json:"elapsed_seconds,omitempty" bson:"elapsed_seconds,omitempty"`
	WebUrl            string                 `json:"web_url,omitempty" bson:"web_url,omitempty"`
	ResultSize        int                    `json:"result_size,omitempty" bson:"result_size,omitempty"`
	Parent            string                 `json:"parent,omitempty" bson:"parent,omitempty"`
	Children          []string               `json:"children,omitempty" bson:"children,omitempty"`
	Log               []interface{}          `json:"log,omitempty" bson:"log,omitempty"`
	Delay             int                    `json:"delay,omitempty" bson:"delay,omitempty"`
	Delay1            int                    `bson:"-"`
}

type Test struct {
	fi int `bson:"-"`
}

func main() {
	actionExecutionAttrs := make(map[string]bool)
	st := reflect.TypeOf(Test{})
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if tag, ok := bsoncodec.DefaultStructTagParser(field); ok == nil {
			if !tag.Skip {
				actionExecutionAttrs[tag.Name] = true
			}
		} else {
			fmt.Println("(not specified)")
		}
	}
	fmt.Println(actionExecutionAttrs)

	x := struct {
		x int
	}{1}

	fmt.Println(reflect.TypeOf(x).Kind().String() == "struct")
}
