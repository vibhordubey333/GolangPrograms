package main

import (
	"MoviesDB/constants"
	connect "MoviesDB/dbconnection"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type POC_Registry struct {
	ID        string     `json:"id"`
	MovieName string     `json:"moviename"`
	Comments  []Comments `json:"comments"`
}
type Comments struct {
	Username string `json:"username"`
	Comment  string `json:"comment"`
}

func main() {
	log.Println("Starting POC.")
	//Getting mongoDB object.
	mongoObj := connect.GetMongoObject()
	collection := mongoObj.Database(constants.DB_NAME).Collection("POC_Registry")

	commentsObj1 := Comments{"admin", "Nice Movie."}
	commentsObj2 := Comments{"chief", "Great Movie"}
	commentsObj := []Comments{commentsObj1, commentsObj2}

	commentsObj3 := Comments{"avds333", "Nice Movie."}
	commentsObj4 := Comments{"avds444", "Great Movie"}
	commentsObjCop := []Comments{commentsObj3, commentsObj4}

	pocObject1 := POC_Registry{"Movie_Hannibal_1", "Hannibal", commentsObj}
	fmt.Println("POCObjec1 :", pocObject1.Comments)
	pocObject2 := POC_Registry{"Movie_FightClub_1", "FightClub", commentsObjCop}
	projectObjectAll := []interface{}{pocObject1, pocObject2}
	collection.InsertMany(context.TODO(), projectObjectAll)

	fmt.Println("POC data inserted successfully.")

	fmt.Println("Displaying records: ")

	readObjects := ReturnAllRecords(mongoObj, bson.M{})

	// MongoDB args
	newUpdate := make(map[string]interface{})
	filter := make(map[string]interface{})
	//MongoDB args ends.
	movie
	for i, v := range readObjects {
		fmt.Println(i, " : ", v.ID, " , ", v.MovieName, " , ", v.Comments)
		for i1, v1 := range v.Comments {

			fmt.Println(i1, ": ", v1.Comment, ", ", v1.Username)
		}
		tempElement := Comments{"TestMovie", "TestComment, nice."}

		v.Comments = append(v.Comments, tempElement)

		fmt.Println("After updating comments:")
		for i1, v1 := range v.Comments {
			fmt.Println(i1, ": ", v1.Comment, ", ", v1.Username)
		}
		newUpdate["comments"] = v.Comments
		filter["ID"] = v.ID

	}
	//updateResponse, err := UpdateDocument(filter, "$set", newUpdate)
	//fmt.Println("Update Response Err: ", err)
	updateResponse := UpdateDocument(mongoObj, bson.M{"comments": newUpdate["comments"]}, bson.M{"id": "Movie_Hannibal_1"})
	//updateResponse := UpdateDocument(mongoObj, bson.M{"Comments": newUpdate["Comments"]}, bson.M{"id": "Movie_Hannibal_1"})
	fmt.Println("Updated Records: ", updateResponse)
	fmt.Println("Updating data: ")
	/*
		updatedResponse = UpdateDocument(mongoObj, bson.M{"MovieName": "HannibalUpdated"}, bson.M{"id": "Movie_Hannibal_1"})
		fmt.Println("Updated Records: ", updatedResponse)
		fmt.Println("**************After updating records: ")
		readObjects = ReturnAllRecords(mongoObj, bson.M{})
		for i, v := range readObjects {
			fmt.Println(i, " : ", v.ID, " , ", v.MovieName, " , ", len(v.Comments))

			for i1, v1 := range v.Comments {
				fmt.Println(i1, ": ", v1.Comment, ", ", v1.Username)
			}
		}
	*/

	/*
		fmt.Println("After updating struct slice.")

		copyCommentsSlice := make([]Comments, 0)

		updatedResponse := UpdateDocument(mongoObj, bson.M{"MovieName": "HannibalUpdated"}, bson.M{"id": "Movie_Hannibal_1"})
		for i, v := range readObjects {
			fmt.Println(i, " : ", v.ID, " , ", v.MovieName, " , ", len(v.Comments))

			for i1, v1 := range v.Comments {

			}
		}*/
	fmt.Println("Program finished.")
}

func UpdateDocument(filterObject interface{}, operation string, update map[string]interface{}) (int64, error) {

	mongoObj := connect.GetMongoObject()

	if mongoObj == nil {
		log.Fatalln(constants.ERR_DB)
		return 0, errors.New("Error while initializing DB connection UpdateDocument.") // TODO
	}
	collection := mongoObj.Database(constants.DB_NAME).Collection(constants.COLLECTION_NAME_USER)

	jsonData, err := json.Marshal(filterObject)
	if err != nil {
		log.Println("Error while Marshalling in UpdateDocument")
		return 0, err
	}
	fmt.Println("jsonData is : ", string(jsonData))
	var m interface{}
	json.Unmarshal(jsonData, &m)
	filter := m.(map[string]interface{})

	jsonDataUpdate, err := json.Marshal(update)
	if err != nil {
		return 0, err
	}
	var m1 interface{}
	json.Unmarshal(jsonDataUpdate, &m1)
	fmt.Println("jsonDataUpdate: ", string(jsonDataUpdate))
	fmt.Println("update string :", update)
	updateString := bson.M{operation: update}
	fmt.Println("Update String : ", updateString)
	result, err := collection.UpdateOne(context.TODO(), filter, updateString)
	log.Println("Result is ::: ", result, " and err : ", err)
	if err != nil {
		return 0, err
	} else {
		log.Println("Returning from here.")
		return result.ModifiedCount, nil
	}
}

/*
func UpdateDocument(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {

	collection := client.Database(constants.DB_NAME).Collection("POC_Registry")
	updateObject := bson.D{{Key: "$set", Value: updatedData}}
	updatedResult, err := collection.UpdateOne(context.TODO(), filter, updateObject)

	if err != nil {
		fmt.Println("Error on updating movie:", err)
	}

	return updatedResult.ModifiedCount
}
*/
func ReturnAllRecords(client *mongo.Client, filter bson.M) []*POC_Registry {
	//When filter is empty it returns all the objects
	var poc_object []*POC_Registry
	collection := client.Database(constants.DB_NAME).Collection("POC_Registry")

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error in ReturnAllRecords: ", err.Error())
	}

	for cur.Next(context.TODO()) {
		var records POC_Registry
		err = cur.Decode(&records)
		if err != nil {
			fmt.Println("Error while iterating records :")
		}
		poc_object = append(poc_object, &records)
	}
	return poc_object
}
