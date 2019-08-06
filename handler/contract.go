package handler

import (
	"fmt"
	"net/http"
        
        "log"     
        "encoding/json"
        "github.com/ethereum/go-ethereum/ethclient"    
        "github.com/bansatya/godapp/contracts"
)

type Message struct {
    question string
    answer string
}

func Pong(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World!")
	respondJSON(w, http.StatusOK, `name:satya`)
}
func GetAccount(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, nil)
}
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusCreated, nil)
}

func CreateContract(session quiz.QuizSession, client *ethclient.Client, w http.ResponseWriter, r *http.Request) {
	
     msg := ParseRequestBody(r)
     contractAddress, tx, instance, err := quiz.DeployQuiz(&session.TransactOpts, client, msg.question, stringToKeccak256(msg.answer))
    if err != nil {
        log.Fatalf("could not deploy contract: %v\n", err)
    }
    fmt.Printf("Contract deployed! Wait for tx %s to be confirmed.\n", tx.Hash().Hex())

    session.Contract = instance
    fmt.Fprintf(w,"tx hash: %v\n contact address: %v ",tx.Hash().Hex(),contractAddress.Hex())
    respondJSON(w, http.StatusCreated, nil)
}

func Execute(w http.ResponseWriter, r *http.Request) {
	
    respondJSON(w, http.StatusCreated, nil)
}

func Call(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, nil)
}

func ParseRequestBody(r *http.Request) (msg *Message){
   msg = &Message{}
   err := json.NewDecoder(r.Body).Decode(&msg)
   if err != nil {
      panic(err)
    }
   return msg;
}
