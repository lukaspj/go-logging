package main

import "./logging"

var logger = logging.GetLogger()

func main() {
	logger.Info("Unseen")
	logger.AddLogOutput(logging.GetFileOutput("logTest.log"))
	logger.Info("ToFile")
	logger.AddLogOutput(new(logging.StdoutOutput))
	logger.Info("ToFileAndPrint")
}
