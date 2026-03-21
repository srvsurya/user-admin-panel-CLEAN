package logger

import ("go.uber.org/zap"
		"os"
		"Week_12/config")

var Log *zap.Logger

func Init(){
	config.LoadEnv()
	var err error
	mode:=os.Getenv("LOGGING_MODE")
	switch mode{ // Mode switch
	case "production":
		Log,err = zap.NewProduction()
	case "development":
		Log,err = zap.NewDevelopment()
	default:
		Log,err = zap.NewProduction()
	}
	if err!=nil{
		panic("Failed to load logger:" + err.Error())
	}
	
}