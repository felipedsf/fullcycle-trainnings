## Project Briefing: Temperature API
The Temperature API is a project designed to provide weather-related data through a simple API interface. It allows users to retrieve temperature information for various cities, making it a valuable tool for applications that require weather data integration.

### Key Features:
- **API Access**: The project includes examples of how to call the API both locally and when deployed on Google Cloud Run.
- **Environment Configuration**: Users can easily configure the API by setting the WEATHER_API_KEY environment variable, which is essential for authenticating requests to the weather data provider.
- **Local and Docker Support**: The API can be run locally using Go or deployed in a Docker container, providing flexibility for different development and production environments.
### Setup Instructions:
 - Firstly, you need to get a key on ``https://www.weatherapi.com/``
 - To run the API locally, set the ``WEATHER_API_KEY`` and specify the ``SERVICE_PORT`` before executing the main application.
> WEATHER_API_KEY={{ WEATHER_API_KEY_VALUE }} SERVICE_PORT=:8080 go run ./cmd/api/main.go

 - For Docker deployment, build the Docker image and run it with the necessary environment variables to expose the API on port 8080.

> docker build --tag city-temperature:v1 . 
>
> docker run -e WEATHER_API_KEY={{ WEATHER_API_KEY_VALUE }} -p 8080:8080 city-temperature:v1
