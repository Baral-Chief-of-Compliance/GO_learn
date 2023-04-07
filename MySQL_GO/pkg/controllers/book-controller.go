package controllers


import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"book_managment_system/pkg/utils"
	"book_managment_system/pkg/models"
)

var NewBook models.Book

