package SystemCode

import (
  "errors"
)

// System error code
var (
  ErrRunServer                        = errors.New("ERROR_RUN_SERVER")
  ErrConnectToDbFail                  = errors.New("ERROR_CONNECT_TO_DB_FAIL")
  ErrLoadFileFail                     = errors.New("ERROR_LOAD_FILE_FAIL")
  ErrServiceUnavailable               = errors.New("ERROR_SERVICE_UNAVAILABLE")
  ErrGenerateJwtFail                  = errors.New("ERROR_GENERATE_JWT_FAIL")
)

// Response message from server to client
const (
  InappropriateAccountInfo            = "INAPPROPRIATE_ACCOUNT_INFO"
  AccountExisted                      = "ACCOUNT_EXISTED"
  RegisterAccountSuccessfully         = "REGISTER_ACCOUNT_SUCCESSFULLY"
  JwtTokenNotFound                    = "JWT_TOKEN_NOT_FOUND"
  InappropriateJwtToken               = "INAPPROPRIATE_JWT_TOKEN"
  ExpiredJwtToken                     = "EXPIRED_JWT_TOKEN"
  LoginInSuccessfully                 = "LOGIN_SUCCESSFULLY"
  GetGameHistorySuccessfully          = "GET_GAME_HISTORY_SUCCESSFULLY"
  GetGameTurnHistorySuccessfully      = "GET_GAME_TURN_HISTORY_SUCCESSFULLY"
  GetGameTopPlayersSuccessfully       = "GET_GAME_TOP_PLAYERS_SUCCESSFULLY"
  InappropriateRequestBody            = "INAPPROPRIATE_REQUEST_BODY"
  LackOfFieldInRequestBody            = "LACK_OF_FIELD_IN_REQUEST_BODY"
  LackOfFieldInRequestParam           = "LACK_OF_FIELD_IN_REQUEST_PARAM"
  SomethingWentWrong                  = "SOMETHING_WENT_WRONG"

)