package graph

import "card-service/graph/model"

var cards = []*model.Card{
	{ID: "1", CardMask: "**** **** **** 1234", UserID: "1"},
	{ID: "2", CardMask: "**** **** **** 4444", UserID: "1"},
	{ID: "3", CardMask: "**** **** **** 5678", UserID: "2"},
}
