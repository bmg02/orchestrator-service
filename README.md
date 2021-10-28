# orchestrator-service

## Candidate details
**Name:** Bhuvan Gandhi

**Exercise:** Orchestrator Service

## Development environment
**Operating System:** Windows

**GO version:** 1.17.2

## File Structure
- The main branch has been merged with other branches in each part of the exercise, as per the specified requirements.
- In the end, I have restructurized a little bit so that it can become easy to navigate and understanding.
- Three folders - GetUserByName (:9000), GetUser (:9001) and Datamock (:10000).
- All those folders contain mainly two folders: server and proto<Port_Number>.
- GetUserByName folder contains the client folder - it contains the script for client.

## Things which are different from actual exercise
- With the approach of actual exercise, the :9000 and :9001 will be crashed on input of less than 6 characters. We need to manually activate them. To handle this, I haven't handled the exception in errors of Golang. Instead I have used `-1` status of `Roll` to know if the length of name is less than 6 or not.

## How To Use
1. Open 4 different terminals. And navigate to the root folder in all of them.
2. We will first activate the orchestrator services. While activating the server, you may get asked to allow the program to communicate over the network through firewall. Allow it.
3. So, let's first activate the :10000 service. Type `go run datamock/server/main.go` and hit enter.
4. Then, let's activate the :9001 service. Type `go run getUser/server/main.go` and hit enter.
5. Then, let's activate the :9000 service. Type `go run getUserByName/server/main.go` and hit enter.
5. Then, let's activate the client. Type `go run getUserByName/client/main.go` and hit enter. It will ask you to write the name of user. Then after providing the name, the requested will be propogated to :9000 to :9001 and :9001 to :10000. On the :10000 server, the actual logic is written. Then the response object will be propogated to :10000 to :9001, :9001 to :9000 and then finally to the client.

## Test Cases
### Case 1
**Input:**
> Enter a name of user: Bhuvan Gandhi

**Output:**
> 2021/10/29 03:40:39 Response: Name:"Bhuvan Gandhi" Class:"13" Roll:130

### Case 2
**Input:**
> Enter a name of user: BG

**Output:**
> 2021/10/29 03:41:15 Response: Name length must be greater than 6.