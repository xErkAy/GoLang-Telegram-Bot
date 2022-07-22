## How to run the project
```sh
cd docker/
docker-compose up -d --build
```

## Getting access to the database
```sh
pgadmin4 >>> localhost:5051
Copy and paste the sql code from 'docker' folder to create a table.
```

## Available functions
```sh
SendMessage()
ReplyToMessageID()
SendPhoto()
ReplyWithPhoto()
SendAudioMessage()
ReplyWithAudioMessage()
SendVideo()
ReplyWithVideo()
SendVideoURL()
ReplyWithVideoURL()
SendDocument()
ReplyWithDocument()
SendSticker()
SendLocation()
ReplyWithLocation()
CreatePoll()
```
