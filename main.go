package main

import (
	"co.edu.meli/luisillera/prueba-tecnica/application/entrypoint"
	_ "co.edu.meli/luisillera/prueba-tecnica/docs"
	"co.edu.meli/luisillera/prueba-tecnica/domain/dto"
	"co.edu.meli/luisillera/prueba-tecnica/domain/model"
	"co.edu.meli/luisillera/prueba-tecnica/domain/usecase"
	"co.edu.meli/luisillera/prueba-tecnica/infrastructure"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
)

type Application struct {
	router             *mux.Router
	messageProvider    infrastructure.MessageProvider
	satelliteProvider  infrastructure.SatelliteProvider
	calculatePosition  usecase.CalculatePositionUsecase
	extractInformation usecase.ExtractInformationUsecase
	messageBuilder     usecase.MessageBuilderUsecase
	requestExtractor   usecase.RequestExtractorUsecase
}

func main() {
	a := Application{}
	a.initializeSatelliteProvider()
	a.initializeMessageProvider()
	a.initializeRoutes()
	a.initializeUseCases()
	a.start()
}
func (a *Application) start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":8080", a.router))
}

// @title Prueba Tecnica MELI - Golang
// @version 1.0
// @description Extract secret message from Spaceship
// @termsOfService http://swagger.io/terms/
// @contact.name Luis Fernando Illera Sanmartin
// @contact.email luisfernando.illera@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func (a *Application) initializeRoutes() {
	a.router = mux.NewRouter()
	a.router.HandleFunc("/topsecret/", a.TopSecretRoute).Methods(http.MethodPost)
	a.router.HandleFunc("/topsecret_split/", a.TopSecretSplitRoute).Methods(http.MethodGet)
	a.router.HandleFunc("/topsecret_split/{satellite_name}", a.TopSecretRouteSingleLoad).Methods(http.MethodPost)
	a.router.PathPrefix("/swagger-ui/").Handler(httpSwagger.WrapHandler)
}

// TopSecret_Split godoc
// @Summary Get position and message from splitted message
// @Description Find the position respect all satellites of a spaceship and the message sended in various signals
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.Response
// @Router /topsecret_split/{satellite_name} [post]
// @param satellite_name path string true "satellite name to register"
func (a *Application) TopSecretRouteSingleLoad(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["satellite_name"]
	if len(name) == 0 {
		errorResponse(w, 404, "la variable [satellite_name] no fue recibida.")
		return
	}
	var body dto.Satellite
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		errorResponse(w, 404, err.Error())
		return
	}
	tssr := entrypoint.TopSecretSplitResource{}
	tssr.Initialize(a.messageProvider)
	body.Name = name
	response, err := tssr.LoadMessage(body)
	if err != nil {
		errorResponse(w, 200, "No hay suficiente informacion.")
		return
	}
	jsonResponse(w, 200, response)
}

// TopSecret godoc
// @Summary Get position and message from splitted message
// @Description Find the position respect all satellites of a spaceship and the message sended in various signals
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.Response
// @Failure 400,404 {string} Not Found
// @Router /topsecret/ [post]
func (a *Application) TopSecretRoute(w http.ResponseWriter, r *http.Request) {
	var body dto.SatelliteRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		errorResponse(w, 404, err.Error())
		return
	}
	defer r.Body.Close()
	tsr := entrypoint.TopSecretResource{
		ExtractInformation: a.extractInformation,
		RequestExtractor:   a.requestExtractor,
	}
	response, err := tsr.HandleRequest(body)
	if err != nil {
		errorResponse(w, 404, err.Error())
		return
	}
	jsonResponse(w, 200, response)
}

// TopSecret_Split godoc
// @Summary Get position and message from splitted message
// @Description Find the position respect all satellites of a spaceship and the message sended in various signals
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.Response
// @Router /topsecret_split/ [get]
func (a *Application) TopSecretSplitRoute(w http.ResponseWriter, r *http.Request) {
	tssr := entrypoint.TopSecretSplitResource{}
	tssr.Initialize(a.messageProvider)
	response, err := tssr.LocateMessage()
	if err != nil {
		fmt.Printf("%s", err.Error())
		errorResponse(w, 200, "No hay suficiente informacion.")
		return
	}
	jsonResponse(w, 200, response)
}

func (a *Application) initializeSatelliteProvider() {
	kenobi := model.Point{
		East:  -500,
		North: -200,
	}
	skywalker := model.Point{
		East:  100,
		North: -100,
	}
	sato := model.Point{
		East:  500,
		North: 100,
	}
	a.satelliteProvider = infrastructure.SatelliteProvider{}
	a.satelliteProvider.Initialize(kenobi, skywalker, sato)
}

func (a *Application) initializeMessageProvider() {
	a.messageProvider = infrastructure.MessageProvider{}
	a.messageProvider.Initialize()
}

func (a *Application) initializeUseCases() {
	a.requestExtractor = usecase.RequestExtractorUsecase{}
	a.messageBuilder = usecase.MessageBuilderUsecase{}

	a.calculatePosition = usecase.CalculatePositionUsecase{
		SatelliteProvider: a.satelliteProvider,
	}
	a.extractInformation = usecase.ExtractInformationUsecase{
		CalculatePosition: a.calculatePosition,
		MessageBuilder:    a.messageBuilder,
	}
}

func errorResponse(w http.ResponseWriter, code int, message string) {
	jsonResponse(w, code, map[string]string{"error": message})
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
