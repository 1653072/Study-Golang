package api

// Configurations of router
const (
	Host = "localhost"
	Port = "8080"
)

// Api/Url path
const WebUrl = "https://localhost:8080"
const DomainUrl = Host + ":" + Port
const apiPoint = "/api"
const (
	// Single urls
	LoginUrl = apiPoint + "/login"

	// Single urls
	RegisterAccountUrl = apiPoint + "/accounts/new"

	// Game url group
	// "BaseUrl" stands for a group of api
	GameBaseUrl = apiPoint + "/games"
	PlayGameUrl = "/play"
	SeeGameHistoryUrl = "/history"
	SeeTopPlayersUrl = "/top"
	SeeGameTurnsUrl = "/turns" // "/turns?gameid=XXX"
)
