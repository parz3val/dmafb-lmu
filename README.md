# Do Me A Favor Buddy (DMAFB): Monorepo  

## Components of the Platform

1. Mobile Frontend
3. Web Frontend
2. Backend
3. Data Analysis and Machine Learning Server


## Architecture Diagram
```mermaid
architecture-beta
    group api(cloud)[DMAFB Platform]
    service server(server)[GoLang Rest API Server] in api
    service server2(server)[ML Server Python] in api
    service application(server)[Mobile Application] in api
    service web(cloud)[Admin Web App] in api

    server:L -- R:server2
    application:T -- T:server
    web:T -- T:server

```

## Tech Stack
1. Mobile Application : React Native, Expo.
2. Backend Application: Golang with Echo and Mongodb.
3. Admin Web Application: React and Typescript
4. Machine Learning Server: Python with Starlette, keras, and tensorflow for ML
