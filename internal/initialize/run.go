package initialize

func Run() {
	// Initialize configuration
	LoadConfig()

	// Initialize the logger
	InitLogger()

	// Initialize MySQL connection
	InitMysql()

	// Initialize Redis connection
	InitRedis()

	// Initialize the router
	r := InitRouter()

	r.Run(":8080")
}
