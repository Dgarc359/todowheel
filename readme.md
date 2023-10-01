# Why
For when someone has a lot of things they need to do but don't know where to start

# Installation
Make sure you have docker desktop installed and running.

# Usage
To start up the backend:
```bash
# build a docker container that runs the backend
./dockerfile.sh
```

# API specification
All requests should be routed to localhost.
The backend accepts proto payloads OR JSON payloads.
Specify which you want with `application/proto` or `application/json`
respectively in your `Content-Type` header. For the latest proto schemas, please check the `proto` folder and read the `.proto` files. Below is the outline for the required JSON payloads.


`/get-health`
**GET**
Returns health of application.


`/create-task`
**POST**
Create a task.

Request
```ts
{
    task_length: number,
    task_name: string,
    task_id?: number,
}
```

Response
```ts
{
    status: string,
    statusCode: number,
}
```

`/get-task`
**POST**
Get a task by task name

Request
```ts
{
    task_name: string,
}
```

Response
```ts
{
    status: string,
    statusCode: number,
}
```


`/get-tasks`
Get a list of tasks. If min / max task length are not provided, then return all tasks.

Request
```ts
{
    min_task_length: number, // optional
    max_task_length: number, // optional
}
```

Response
```
{
    status: string,
    statusCode: number,
    Data: { id: string, task_name: string, task_length: number }[]
}
```

`/spin`
returns a single todo based on optional min / max task length. If no min / max task length provided, then use it picks a random todo from all available todos

Request
```ts
{
    min_task_length: number,
    max_task_length: number,
}
```

Response
```
{
    status: string,
    statusCode: number,
    Data: { id: string, task_name: string, task_length: number }
}
```

# Features
* Set an estimated length for task
* Spin a wheel to pick a task
* You can set a minimum / maximum task length so that you can say, 'I wanna do something that will take me 5 minutes'

# Technical Requirements for MVP #1
* Docker
* Go
* Protobuf
* Roulette app is dockerized and exposes either an API or something else to interface with a frontend
