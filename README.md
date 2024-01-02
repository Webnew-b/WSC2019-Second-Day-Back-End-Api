# WSC
 这是针对2019年的世界技能大赛的项目第二天下午api项目的golang重构，提供大家学习参考使用

原本的要求是使用php进行编写，在这里我使用了golang重构了原本项目，以供学习与参考。

本项目技术栈为：
1. golang(echo+gorm+go-redis)
2. redis
3. mysql

运行之前从sql文件夹中手动导入sql文件到对应数据库，请检查配置文件的数据库名称与实际名称是否相符。

本项目需要使用到redis，请确保运行环境有redis服务



运行命令：
```shell
go run main.go
```

打包命令：
```shell
go bulid -trimpath -o ./bulid/wscmakebygo
```

以下是当时的题目

## Attendee API

Additional information for the REST-API tests/specification: the body contains some static example data, adjust it to fit the requirements. Placeholder parameters in the URL are marked with curly braces: {slug}
The tests are referenced by its function name [-> test<scenario-id><text>].
The order of properties in objects does not matter, but the order of items in an array does.

### Events overview

When an attendee lands on the index page of the application, he will get a list with all upcoming events across all organizers. Events should be ordered ascendilgly by their date.
Specification:

#### Feature: B1 – Read events: As a visitor I want to read the public data of all upcoming events so that I can display them in a list

Scenario: B1a – Read all upcoming events

Request

```api
URL: /api/v1/events
Method: GET
Header:
Body: -
Response
If success [-> testB1aGetEvents]
Header: Response code: 200
Body: {"events": [{"id": 1, "name": "someText", "slug": "some-text",  "date": "2019-08-15", "organizer": {"id": 1, "name": "someText", "slug": "some-text"}}]
```

### Event detail
When an event is selected, all information of the selected event should be returned in the same endpoint.
If a session has no cost or a ticket no description, NULL will be used.
The description of a ticket is generated as following:
-NULL if no special validity rule set
-"Available until {month} {day}, {year}" if special validity rule "date" is used
(example: "Available until September 1, 2019")
-"X tickets available" if special validity rule "amount" is used
(example: "30 tickets available" where 30 is the total amount of tickets and not the remaining ones)
#### Specification:
##### Feature: B2 - Read event: As a visitor I want to read all information about an event

Scenario: B2a – Read all detail information of a single event

```
Request
URL: /api/v1/organizers/{organizer-slug}/events/{event-slug}
Method: GET
Header:
Body: -

Response

If success [-> testB2aGetEventDetail | testB2aTicketValidity]
Header: Response code: 200
Body: {"id": 1, "name": "someText", "slug": "some-text",  "date": "2019-08-15", "channels": [{"id": 1, "name": "someText", "rooms": [{"id": 1, "name": "someText", "sessions": [{"id": 1, "title": "someText", "description": "someText", "speaker": "someText", "start": "2019-08-15 10:00:00", "end": "2019-08-15 10:45:00", "type": "workshop", "cost": 50|null}]}]}], "tickets": [{"id": 1, "name": "someText", "description": "Available until July 7, 2019"|null, "cost": 199.99, "available": true}]}

If event does not exist or was not created by the organizer [-> testB2aGetInvalidEvent]
Header: Response code: 404
Body: {"message": "Event not found"}

If organizer does not exist [-> testB2aGetInvalidOrganizer]
Header: Response code: 404
Body: {"message": "Organizer not found"}
```

### Attendee login and logout
After signup (which is not within the scope of this project), attendees get assigned a random 6 characters long registration code which they need when entering an event. To avoid requiring yet another password, this registration code is also used for logging the user in on the website, together with his lastname as his login credentials. The registration code is shared across all events and is not only valid for a single event.
For testing purposes, the registration code and last name can be seen in the database dump.
On a successful login, data of the loggedin attendee gets returned together with a token which can be sumitted with subsequent requests. For simplicity, the token is the username md5-hashed and is valid until the logout endpoint is called.
Specification:
#### Feature: B3 - Login and Logout as atteendee: As a visitor I want to login/logout so that I can attend an event and see my registrations

Scenario: B3a - Attendee login

```
Request
URL: /api/v1/login
Method: POST
Header:
Body: {"lastname": "someText", "registration_code": "someText"}

Response

If success [-> testB3aCorrectLogin|SameLastname|SameRegistrationCode]
Header: Response code: 200
Body: {"firstname: "someText", "lastname": "someText", "username": "someText", "email": "someText", "token": "AUTHORIZATION_TOKEN"} (md5 hashed username, valid until logout)

If user login info not correct [-> testB3aInvalidLastname|RegistrationCode|Request]
Header: Response code: 401
Body: {"message": "Invalid login"}
```

Scenario: B3b – Attendee logout
```
Request
URL: /api/v1/logout?token={AUTHORIZATION_TOKEN}
Method: POST
Header:
Body: -

Response
If success [-> testB3bLogout]
Header: Response code: 200
Body: {"message": "Logout success"}

If invalid token [-> testB3bInvalidToken | testB3bAlreadyLoggedOut]
Header: Response code: 401
Body: {"message": "Invalid token"}
```

### Event registration

Users will register and buy a ticket with this endpoint. The different ticket validity rules (date or amount) should be validated. Registrations should be sorted ascendingly by their id.
Specification:
#### Feature: B4 – Event registration: As a visitor I want to register for an event

Scenario: B4a – New event registration
```
Request
URL: /api/v1/organizers/{organizer-slug}/events/{event-slug}/registration?token={AUTHORIZATION_TOKEN}
Method: POST
Header:
Body: {"ticket_id": 1, "session_ids": [1, 2, 3]} (session_ids is optional)

Response

If success [-> testB4aCorrectRegistration]
Header: Response code: 200
Body: {"message": "Registration successful"}

If user is not logged in or token is invalid [-> testB4aLoggedOut]
Header: Response code: 401
Body: {"message": "User not logged in"}

If user already registered for this event [-> testB4aAlreadyRegistered]
Header: Response code: 401
Body: {"message": "User already registered"}

If ticket is not available anymore [-> testB4aInvalidTicket]
Header: Response code: 401
Body: {"message": "Ticket is no longer available"}
```



Scenario: B4b – Get registrations of a user
```
Request
URL: /api/v1/registrations?token={AUTHORIZATION_TOKEN}
Method: GET
Header:
Body: -

Response

If success [-> testB4bGetRegistrations | testB4bNewRegistration]
Header: Response code: 200
Body: {"registrations": [{"event": {"id": 1, "name": "someText", "slug": "some-text",  "date": "2019-08-15", "organizer": {"id": 1, "name": "someText", "slug": "some-text"}}, "session_ids": [1, 2, 3]}]}

If user is not logged in or token is invalid [-> testB4bLoggedOut]
Header: Response code: 401
Body: {"message": "User not logged in"}
```

