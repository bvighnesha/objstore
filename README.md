# objstore

###Technologies Used:

Language: Go
Database: Redis
Docker

How to use:

Start the application by running the following command:
`make start`

1. Which will do the following
2. Build the docker image
3. Start the docker container at port `6379`
4. Run the application

When application start it will ask for the following options:
Please provide a command:  `SET`, `GET`, `LIST` or `DEL`

To set object in the store:

`SET`

``` JSON 
{
"Name":      "test",
"LastName":  "Bojja",
"Birthday":  "04-11-1989",
"BirthDate": "2021-02-18T21:54:42.123Z"
}
```

`NOTE: name is unique`

`GET`

Do you want to get by name or ID?

Choose `ID` or `NAME`

Example: `ID`

And enter name like `test` 

`LIST`

Please provide a `kind` to list

`person` or `animal`

Example: `person`

`DEL`

Please provide a ID to delete

Example: `21082c45-36b8-4ed0-9131-34036a0dc8b6`