package handler

import (
	"encoding/json"
	"fmt"
	"os"
)

func Compare(followers, following *os.File) {
	var followersJSON map[string]interface{}
	var followingJSON map[string]interface{}

	json.NewDecoder(followers).Decode(&followersJSON)
	json.NewDecoder(following).Decode(&followingJSON)

	followersArray := followersJSON["relationships_followers"].([]interface{})
	followingArray := followingJSON["relationships_following"].([]interface{})

	var notFollowedBack []string
	for _, following := range followingArray {
		followingStringListData := following.(map[string]interface{})["string_list_data"].([]interface{})
		for _, followingStringListDataItem := range followingStringListData {
			followingStringListDataItemValue := followingStringListDataItem.(map[string]interface{})["value"].(string)
			var isFollowing bool
			for _, followers := range followersArray {
				followersStringListData := followers.(map[string]interface{})["string_list_data"].([]interface{})
				for _, followersStringListDataItem := range followersStringListData {
					followersStringListDataItemValue := followersStringListDataItem.(map[string]interface{})["value"].(string)
					if followingStringListDataItemValue == followersStringListDataItemValue {
						isFollowing = true
					}
				}
			}
			if !isFollowing {
				notFollowedBack = append(notFollowedBack, followingStringListDataItemValue)
			}
		}
	}

	for _, notFollowedBackItem := range notFollowedBack {
		fmt.Println(notFollowedBackItem)
	}
	fmt.Println("Not followed back:", len(notFollowedBack))
}
