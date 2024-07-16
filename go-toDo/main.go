package main

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"strings"
// 	"time"

// 	"github.com/globalsign/mgo"
// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// 	"github.com/thedevsaddam/renderer"
// 	mgo "gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bison"
// 	"gopkg.in/mgo.v2/bson"
// )

// var rnd * renderer.Render
// var db *mgo.Database

// const (
// 	hostName string = "localhost:27017"
// 	dbName string = "demo_todo"
// 	collectionName string = "todo"
// 	port 			string = ":9000"
// )

// type (
// 	todoModel struct{
// 		ID		bson.ObjectId`bson:"id,omitempty"`
// 		Title string `bson:"`
// 	}
// )